package chaintests

import (
	"testing"
	"unit-test/chaincode"
	"unit-test/mocks"

	"github.com/stretchr/testify/require"
)

func TestInitLedger(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContext := &mocks.TransactionContext{}
	transactionContext.GetStubReturns(chaincodeStub)

	assetTransfer := chaincode.SmartContract{}
	err := assetTransfer.InitLedger(transactionContext)
	require.NoError(t, err)
}

func TestCreateAsset(t *testing.T) {
	chaincodeStub := &mocks.ChaincodeStub{}
	transactionContex := &mocks.TransactionContext{}
	transactionContex.GetStubReturns(chaincodeStub)

	assetTransfer := chaincode.SmartContract{}

	err := assetTransfer.CreateAsset(transactionContex, "asset1", "", 2, "", 1)
	require.NoError(t, err)
	err2 := assetTransfer.UpdateAsset(transactionContex, "asset1", "", 3, "", 4)
	require.NoError(t, err2)
}
