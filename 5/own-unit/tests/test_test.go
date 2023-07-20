package tests

import (
	"testing"

	"unit/chaincode"

	"github.com/hyperledger/fabric-samples/asset-transfer-basic/chaincode-go/chaincode/mocks"
	"github.com/stretchr/testify/require"
)

func TestInitLedger(t *testing.T) {
	//? GET THE STUB FOR THE CHAINCODE
	chaincodeSub := &mocks.ChaincodeStub{}
	//? GRAB THE TRANSACTION CONTEXT
	transactionContext := &mocks.TransactionContext{}
	//? Use the given chaincodeStub
	transactionContext.GetStubReturns(chaincodeSub)

	//? GRAB THE SMART CONTRACT
	assetTransfer := chaincode.SmartContract{}
	err := assetTransfer.InitLedger(transactionContext)
	require.NoError(t, err)
}
