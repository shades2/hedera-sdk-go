package hedera

import (
	"time"

	"github.com/hashgraph/hedera-sdk-go/v2/proto"
)

// Revokes KYC to the account for the given token. Must be signed by the Token's kycKey.
// If the provided account is not found, the transaction will resolve to INVALID_ACCOUNT_ID.
// If the provided account has been deleted, the transaction will resolve to ACCOUNT_DELETED.
// If the provided token is not found, the transaction will resolve to INVALID_TOKEN_ID.
// If the provided token has been deleted, the transaction will resolve to TOKEN_WAS_DELETED.
// If an Association between the provided token and account is not found, the transaction will
// resolve to TOKEN_NOT_ASSOCIATED_TO_ACCOUNT.
// If no KYC Key is defined, the transaction will resolve to TOKEN_HAS_NO_KYC_KEY.
// Once executed the Account is marked as KYC Revoked
type TokenRevokeKycTransaction struct {
	Transaction
	tokenID   TokenID
	accountID AccountID
}

func NewTokenRevokeKycTransaction() *TokenRevokeKycTransaction {
	transaction := TokenRevokeKycTransaction{
		Transaction: newTransaction(),
	}
	transaction.SetMaxTransactionFee(NewHbar(30))

	return &transaction
}

func tokenRevokeKycTransactionFromProtobuf(transaction Transaction, pb *proto.TransactionBody) TokenRevokeKycTransaction {
	return TokenRevokeKycTransaction{
		Transaction: transaction,
		tokenID:     tokenIDFromProtobuf(pb.GetTokenRevokeKyc().GetToken()),
		accountID:   accountIDFromProtobuf(pb.GetTokenRevokeKyc().GetAccount()),
	}
}

// The token for which this account will get his KYC revoked. If token does not exist, transaction results in INVALID_TOKEN_ID
func (transaction *TokenRevokeKycTransaction) SetTokenID(id TokenID) *TokenRevokeKycTransaction {
	transaction.requireNotFrozen()
	transaction.tokenID = id
	return transaction
}

func (transaction *TokenRevokeKycTransaction) GetTokenID() TokenID {
	return transaction.tokenID
}

// The account to be KYC Revoked
func (transaction *TokenRevokeKycTransaction) SetAccountID(id AccountID) *TokenRevokeKycTransaction {
	transaction.requireNotFrozen()
	transaction.accountID = id
	return transaction
}

func (transaction *TokenRevokeKycTransaction) GetAccountID() AccountID {
	return transaction.accountID
}

func (transaction *TokenRevokeKycTransaction) validateNetworkOnIDs(client *Client) error {
	if !client.autoValidateChecksums {
		return nil
	}
	var err error
	err = transaction.tokenID.Validate(client)
	if err != nil {
		return err
	}
	err = transaction.accountID.Validate(client)
	if err != nil {
		return err
	}

	return nil
}

func (transaction *TokenRevokeKycTransaction) build() *proto.TransactionBody {
	body := &proto.TokenRevokeKycTransactionBody{}
	if !transaction.tokenID.isZero() {
		body.Token = transaction.tokenID.toProtobuf()
	}

	if !transaction.accountID.isZero() {
		body.Account = transaction.accountID.toProtobuf()
	}

	return &proto.TransactionBody{
		TransactionFee:           transaction.transactionFee,
		Memo:                     transaction.Transaction.memo,
		TransactionValidDuration: durationToProtobuf(transaction.GetTransactionValidDuration()),
		TransactionID:            transaction.transactionID.toProtobuf(),
		Data: &proto.TransactionBody_TokenRevokeKyc{
			TokenRevokeKyc: body,
		},
	}
}

func (transaction *TokenRevokeKycTransaction) Schedule() (*ScheduleCreateTransaction, error) {
	transaction.requireNotFrozen()

	scheduled, err := transaction.constructScheduleProtobuf()
	if err != nil {
		return nil, err
	}

	return NewScheduleCreateTransaction().setSchedulableTransactionBody(scheduled), nil
}

func (transaction *TokenRevokeKycTransaction) constructScheduleProtobuf() (*proto.SchedulableTransactionBody, error) {
	body := &proto.TokenRevokeKycTransactionBody{}
	if !transaction.tokenID.isZero() {
		body.Token = transaction.tokenID.toProtobuf()
	}

	if !transaction.accountID.isZero() {
		body.Account = transaction.accountID.toProtobuf()
	}

	return &proto.SchedulableTransactionBody{
		TransactionFee: transaction.transactionFee,
		Memo:           transaction.Transaction.memo,
		Data: &proto.SchedulableTransactionBody_TokenRevokeKyc{
			TokenRevokeKyc: body,
		},
	}, nil
}

//
// The following methods must be copy-pasted/overriden at the bottom of **every** _transaction.go file
// We override the embedded fluent setter methods to return the outer type
//

func tokenRevokeKycTransaction_getMethod(request request, channel *channel) method {
	return method{
		transaction: channel.getToken().RevokeKycFromTokenAccount,
	}
}

func (transaction *TokenRevokeKycTransaction) IsFrozen() bool {
	return transaction.isFrozen()
}

// Sign uses the provided privateKey to sign the transaction.
func (transaction *TokenRevokeKycTransaction) Sign(
	privateKey PrivateKey,
) *TokenRevokeKycTransaction {
	return transaction.SignWith(privateKey.PublicKey(), privateKey.Sign)
}

func (transaction *TokenRevokeKycTransaction) SignWithOperator(
	client *Client,
) (*TokenRevokeKycTransaction, error) {
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
func (transaction *TokenRevokeKycTransaction) SignWith(
	publicKey PublicKey,
	signer TransactionSigner,
) *TokenRevokeKycTransaction {
	if !transaction.IsFrozen() {
		_, _ = transaction.Freeze()
	}

	if !transaction.keyAlreadySigned(publicKey) {
		transaction.signWith(publicKey, signer)
	}

	return transaction
}

// Execute executes the Transaction with the provided client
func (transaction *TokenRevokeKycTransaction) Execute(
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
		tokenRevokeKycTransaction_getMethod,
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

func (transaction *TokenRevokeKycTransaction) Freeze() (*TokenRevokeKycTransaction, error) {
	return transaction.FreezeWith(nil)
}

func (transaction *TokenRevokeKycTransaction) FreezeWith(client *Client) (*TokenRevokeKycTransaction, error) {
	if transaction.IsFrozen() {
		return transaction, nil
	}
	transaction.initFee(client)
	err := transaction.validateNetworkOnIDs(client)
	if err != nil {
		return &TokenRevokeKycTransaction{}, err
	}
	if err := transaction.initTransactionID(client); err != nil {
		return transaction, err
	}
	body := transaction.build()

	return transaction, transaction_freezeWith(&transaction.Transaction, client, body)
}

func (transaction *TokenRevokeKycTransaction) GetMaxTransactionFee() Hbar {
	return transaction.Transaction.GetMaxTransactionFee()
}

// SetMaxTransactionFee sets the max transaction fee for this TokenRevokeKycTransaction.
func (transaction *TokenRevokeKycTransaction) SetMaxTransactionFee(fee Hbar) *TokenRevokeKycTransaction {
	transaction.requireNotFrozen()
	transaction.Transaction.SetMaxTransactionFee(fee)
	return transaction
}

func (transaction *TokenRevokeKycTransaction) GetTransactionMemo() string {
	return transaction.Transaction.GetTransactionMemo()
}

// SetTransactionMemo sets the memo for this TokenRevokeKycTransaction.
func (transaction *TokenRevokeKycTransaction) SetTransactionMemo(memo string) *TokenRevokeKycTransaction {
	transaction.requireNotFrozen()
	transaction.Transaction.SetTransactionMemo(memo)
	return transaction
}

func (transaction *TokenRevokeKycTransaction) GetTransactionValidDuration() time.Duration {
	return transaction.Transaction.GetTransactionValidDuration()
}

// SetTransactionValidDuration sets the valid duration for this TokenRevokeKycTransaction.
func (transaction *TokenRevokeKycTransaction) SetTransactionValidDuration(duration time.Duration) *TokenRevokeKycTransaction {
	transaction.requireNotFrozen()
	transaction.Transaction.SetTransactionValidDuration(duration)
	return transaction
}

func (transaction *TokenRevokeKycTransaction) GetTransactionID() TransactionID {
	return transaction.Transaction.GetTransactionID()
}

// SetTransactionID sets the TransactionID for this TokenRevokeKycTransaction.
func (transaction *TokenRevokeKycTransaction) SetTransactionID(transactionID TransactionID) *TokenRevokeKycTransaction {
	transaction.requireNotFrozen()

	transaction.Transaction.SetTransactionID(transactionID)
	return transaction
}

// SetNodeTokenID sets the node TokenID for this TokenRevokeKycTransaction.
func (transaction *TokenRevokeKycTransaction) SetNodeAccountIDs(nodeID []AccountID) *TokenRevokeKycTransaction {
	transaction.requireNotFrozen()
	transaction.Transaction.SetNodeAccountIDs(nodeID)
	return transaction
}

func (transaction *TokenRevokeKycTransaction) SetMaxRetry(count int) *TokenRevokeKycTransaction {
	transaction.Transaction.SetMaxRetry(count)
	return transaction
}

func (transaction *TokenRevokeKycTransaction) AddSignature(publicKey PublicKey, signature []byte) *TokenRevokeKycTransaction {
	transaction.requireOneNodeAccountID()

	if !transaction.isFrozen() {
		transaction.Freeze()
	}

	if transaction.keyAlreadySigned(publicKey) {
		return transaction
	}

	if len(transaction.signedTransactions) == 0 {
		return transaction
	}

	transaction.transactions = make([]*proto.Transaction, 0)
	transaction.publicKeys = append(transaction.publicKeys, publicKey)
	transaction.transactionSigners = append(transaction.transactionSigners, nil)

	for index := 0; index < len(transaction.signedTransactions); index++ {
		transaction.signedTransactions[index].SigMap.SigPair = append(
			transaction.signedTransactions[index].SigMap.SigPair,
			publicKey.toSignaturePairProtobuf(signature),
		)
	}

	//transaction.signedTransactions[0].SigMap.SigPair = append(transaction.signedTransactions[0].SigMap.SigPair, publicKey.toSignaturePairProtobuf(signature))
	return transaction
}
