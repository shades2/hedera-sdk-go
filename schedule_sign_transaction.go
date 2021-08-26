package hedera

import (
	"github.com/pkg/errors"
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2/proto"
)

type ScheduleSignTransaction struct {
	Transaction
	scheduleID ScheduleID
}

func NewScheduleSignTransaction() *ScheduleSignTransaction {
	transaction := ScheduleSignTransaction{
		Transaction: newTransaction(),
	}

	transaction.SetMaxTransactionFee(NewHbar(5))

	return &transaction
}

func scheduleSignTransactionFromProtobuf(transaction Transaction, pb *proto.TransactionBody) ScheduleSignTransaction {
	return ScheduleSignTransaction{
		Transaction: transaction,
		scheduleID:  scheduleIDFromProtobuf(pb.GetScheduleSign().GetScheduleID()),
	}
}

func (transaction *ScheduleSignTransaction) SetScheduleID(id ScheduleID) *ScheduleSignTransaction {
	transaction.requireNotFrozen()
	transaction.scheduleID = id

	return transaction
}

func (transaction *ScheduleSignTransaction) GetScheduleID() ScheduleID {
	return transaction.scheduleID
}

func (transaction *ScheduleSignTransaction) validateNetworkOnIDs(client *Client) error {
	if client == nil || !client.autoValidateChecksums {
		return nil
	}
	var err error
	err = transaction.scheduleID.Validate(client)
	if err != nil {
		return err
	}

	return nil
}

func (transaction *ScheduleSignTransaction) build() *proto.TransactionBody {
	body := &proto.ScheduleSignTransactionBody{}
	if !transaction.scheduleID.isZero() {
		body.ScheduleID = transaction.scheduleID.toProtobuf()
	}

	return &proto.TransactionBody{
		TransactionFee:           transaction.transactionFee,
		Memo:                     transaction.Transaction.memo,
		TransactionValidDuration: durationToProtobuf(transaction.GetTransactionValidDuration()),
		TransactionID:            transaction.transactionID.toProtobuf(),
		Data: &proto.TransactionBody_ScheduleSign{
			ScheduleSign: body,
		},
	}
}

func (transaction *ScheduleSignTransaction) constructScheduleProtobuf() (*proto.SchedulableTransactionBody, error) {
	return nil, errors.New("cannot schedule `ScheduleSignTransaction")
}

//
// The following methods must be copy-pasted/overriden at the bottom of **every** _transaction.go file
// We override the embedded fluent setter methods to return the outer type
//

func scheduleSignTransaction_getMethod(request request, channel *channel) method {
	return method{
		transaction: channel.getSchedule().SignSchedule,
	}
}

func (transaction *ScheduleSignTransaction) IsFrozen() bool {
	return transaction.isFrozen()
}

// Sign uses the provided privateKey to sign the transaction.
func (transaction *ScheduleSignTransaction) Sign(
	privateKey PrivateKey,
) *ScheduleSignTransaction {
	return transaction.SignWith(privateKey.PublicKey(), privateKey.Sign)
}

func (transaction *ScheduleSignTransaction) SignWithOperator(
	client *Client,
) (*ScheduleSignTransaction, error) {
	// If the transaction is not signed by the operator, we need
	// to sign the transaction with the operator

	if client == nil {
		return nil, errNoClientProvided
	} else if client.operator == nil {
		return nil, errClientOperatorSigning
	}
	if !transaction.IsFrozen() {
		_, err := transaction.FreezeWith(client)
		if err != nil {
			return transaction, err
		}
	}
	return transaction.SignWith(client.operator.publicKey, client.operator.signer), nil
}

// SignWith executes the TransactionSigner and adds the resulting signature data to the Transaction's signature map
// with the publicKey as the map key.
func (transaction *ScheduleSignTransaction) SignWith(
	publicKey PublicKey,
	signer TransactionSigner,
) *ScheduleSignTransaction {
	if !transaction.IsFrozen() {
		_, _ = transaction.Freeze()
	}

	if !transaction.keyAlreadySigned(publicKey) {
		transaction.signWith(publicKey, signer)
	}

	return transaction
}

// Execute executes the Transaction with the provided client
func (transaction *ScheduleSignTransaction) Execute(
	client *Client,
) (TransactionResponse, error) {
	if client == nil {
		return TransactionResponse{}, errNoClientProvided
	}

	if transaction.freezeError != nil {
		return TransactionResponse{}, transaction.freezeError
	}

	if !transaction.IsFrozen() {
		_, err := transaction.FreezeWith(client)
		if err != nil {
			return TransactionResponse{}, err
		}
	}

	transactionID := transaction.GetTransactionID()

	if !client.GetOperatorAccountID().isZero() && client.GetOperatorAccountID().equals(*transactionID.AccountID) {
		transaction.SignWith(
			client.GetOperatorPublicKey(),
			client.operator.signer,
		)
	}

	resp, err := execute(
		client,
		request{
			transaction: &transaction.Transaction,
		},
		transaction_shouldRetry,
		transaction_makeRequest(request{
			transaction: &transaction.Transaction,
		}),
		transaction_advanceRequest,
		transaction_getNodeAccountID,
		scheduleSignTransaction_getMethod,
		transaction_mapStatusError,
		transaction_mapResponse,
	)

	if err != nil {
		return TransactionResponse{
			TransactionID: transaction.GetTransactionID(),
			NodeID:        resp.transaction.NodeID,
		}, err
	}

	hash, err := transaction.GetTransactionHash()

	return TransactionResponse{
		TransactionID: transaction.GetTransactionID(),
		NodeID:        resp.transaction.NodeID,
		Hash:          hash,
	}, nil
}

func (transaction *ScheduleSignTransaction) Freeze() (*ScheduleSignTransaction, error) {
	return transaction.FreezeWith(nil)
}

func (transaction *ScheduleSignTransaction) FreezeWith(client *Client) (*ScheduleSignTransaction, error) {
	if transaction.IsFrozen() {
		return transaction, nil
	}
	transaction.initFee(client)
	err := transaction.validateNetworkOnIDs(client)
	if err != nil {
		return &ScheduleSignTransaction{}, err
	}
	if err := transaction.initTransactionID(client); err != nil {
		return transaction, err
	}
	body := transaction.build()

	return transaction, transaction_freezeWith(&transaction.Transaction, client, body)
}

func (transaction *ScheduleSignTransaction) GetMaxTransactionFee() Hbar {
	return transaction.Transaction.GetMaxTransactionFee()
}

// SetMaxTransactionFee sets the max transaction fee for this ScheduleSignTransaction.
func (transaction *ScheduleSignTransaction) SetMaxTransactionFee(fee Hbar) *ScheduleSignTransaction {
	transaction.requireNotFrozen()
	transaction.Transaction.SetMaxTransactionFee(fee)
	return transaction
}

func (transaction *ScheduleSignTransaction) GetTransactionMemo() string {
	return transaction.Transaction.GetTransactionMemo()
}

// SetTransactionMemo sets the memo for this ScheduleSignTransaction.
func (transaction *ScheduleSignTransaction) SetTransactionMemo(memo string) *ScheduleSignTransaction {
	transaction.requireNotFrozen()
	transaction.Transaction.SetTransactionMemo(memo)
	return transaction
}

func (transaction *ScheduleSignTransaction) GetTransactionValidDuration() time.Duration {
	return transaction.Transaction.GetTransactionValidDuration()
}

// SetTransactionValidDuration sets the valid duration for this ScheduleSignTransaction.
func (transaction *ScheduleSignTransaction) SetTransactionValidDuration(duration time.Duration) *ScheduleSignTransaction {
	transaction.requireNotFrozen()
	transaction.Transaction.SetTransactionValidDuration(duration)
	return transaction
}

func (transaction *ScheduleSignTransaction) GetTransactionID() TransactionID {
	return transaction.Transaction.GetTransactionID()
}

// SetTransactionID sets the TransactionID for this ScheduleSignTransaction.
func (transaction *ScheduleSignTransaction) SetTransactionID(transactionID TransactionID) *ScheduleSignTransaction {
	transaction.requireNotFrozen()

	transaction.Transaction.SetTransactionID(transactionID)
	return transaction
}

// SetNodeAccountID sets the node AccountID for this ScheduleSignTransaction.
func (transaction *ScheduleSignTransaction) SetNodeAccountIDs(nodeID []AccountID) *ScheduleSignTransaction {
	transaction.requireNotFrozen()
	transaction.Transaction.SetNodeAccountIDs(nodeID)
	return transaction
}

func (transaction *ScheduleSignTransaction) SetMaxRetry(count int) *ScheduleSignTransaction {
	transaction.Transaction.SetMaxRetry(count)
	return transaction
}

func (transaction *ScheduleSignTransaction) SetMaxBackoff(max time.Duration) *ScheduleSignTransaction {
	if max.Nanoseconds() < 0 {
		panic("maxBackoff must be a positive duration")
	} else if max.Nanoseconds() < transaction.minBackoff.Nanoseconds() {
		panic("maxBackoff must be greater than or equal to minBackoff")
	}
	transaction.maxBackoff = &max
	return transaction
}

func (transaction *ScheduleSignTransaction) GetMaxBackoff() time.Duration {
	if transaction.maxBackoff != nil {
		return *transaction.maxBackoff
	}

	return 8 * time.Second
}

func (transaction *ScheduleSignTransaction) SetMinBackoff(min time.Duration) *ScheduleSignTransaction {
	if min.Nanoseconds() < 0 {
		panic("minBackoff must be a positive duration")
	} else if transaction.maxBackoff.Nanoseconds() < min.Nanoseconds() {
		panic("minBackoff must be less than or equal to maxBackoff")
	}
	transaction.minBackoff = &min
	return transaction
}

func (transaction *ScheduleSignTransaction) GetMinBackoff() time.Duration {
	if transaction.minBackoff != nil {
		return *transaction.minBackoff
	}

	return 250 * time.Millisecond
}
