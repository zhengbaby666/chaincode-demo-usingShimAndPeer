package main

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type CatChaincode struct{}

func (cc *CatChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	args := stub.GetArgs()
	if len(args) != 2 {
		return shim.Error("Incorrect arguments, expecting a key and a value")
	}
	err := stub.PutState(string(args[0]), args[1])
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create asset : %s", args[0]))
	}
	return shim.Success(nil)
}

func (cc *CatChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fn, args := stub.GetFunctionAndParameters()

	var result string
	var err error

	if fn == "set" {
		result, err = set(stub, args)
	} else if fn == "get" {
		result, err = get(stub, args)
	}

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(result))
}

func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments, expecting a key and a value")
	}
	err := stub.PutState(args[0], []byte(args[1]))
	if err != nil {
		return "", err
	}
	return args[1], nil
}

func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments, expecting a key")
	}
	value, err := stub.GetState(args[0])
	if err != nil {
		return "", err
	}
	if value == nil {
		return "", fmt.Errorf("Asset %s does not exist.", args[0])
	}
	return string(value), nil
}
func main() {
	fmt.Println("开始启动链码...")
	if err := shim.Start(new(CatChaincode)); err != nil {
		fmt.Printf("Starting chaincode error! %s", err)
	}
}
