# CHAINCODE [Link](asset-transfer-chaincode/atcc.go)
- Can be coded in GO
- Javascript
- Java
- It needs a Interface
- Runs in a separate process from the peer and inits and manages the ledger state
- If we try to run a chaincode from another chaincode, in case they are in different channels, we will only be able to perform querys
- If we have external dependencies in go, we whould use "go mod vendor", before packaging the chaincode

## Fabric Contract API
- The contract interface, a high level API for app developers 
- Documentation for go is at "https://pkg.go.dev/github.com/hyperledger/fabric-contract-api-go/contractapi"
- When using this contract API, all the functions are passed in a transaction context "ctx". From this you can get the chaincode stub (GetStub()). By accessing it we have function to access the ledger and requests to update the ledger (PutState())
- We can have access control over the chaincode using ctx.GetStub().GetCreator(). This uses the client certificate to make some sort of operations, we can store the certificate as a key value pair and then use it to compare with the given certificate