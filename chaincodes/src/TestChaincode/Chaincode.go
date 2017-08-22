package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

var logger = shim.NewLogger("TestChaincode")

func (t *TestChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	if len(stub.GetStringArgs()) < 2 {
		logger.Errorf("UserRegistry init: Not enough arguments, expected 2")
		return shim.Error("UserRegistry init: Not enough arguments, expected 2")
	}

	jsonInitParams := stub.GetStringArgs()[1]

	logger.Info(jsonInitParams)

	user := new(User)

	if err := jsonpb.UnmarshalString(jsonInitParams, user); err != nil {
		logger.Errorf("Error unmarshaling user: %s", err.Error())
		return shim.Error("Error unmarshaling user:" + err.Error())
	}
	if err := t.init(stub, user); err != nil {
		logger.Errorf("Error initing user: %s", err.Error())
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *TestChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	args := stub.GetArgs()
	if len(args) <= 1 {
		logger.Errorf("Insufficient arguments found\n")
		return shim.Error(fmt.Sprintf("Insufficient arguments found\n"))
	}
	function := string(args[0][:])
	argsBytes := args[1]

	logger.Info("Function = ", function)

	switch function {
	case "add":
		parameters := User{}
		if err := proto.Unmarshal(argsBytes, &parameters); err != nil {
			logger.Errorf("Invalid argument expected User protocol buffer %s\n", err.Error())
			return shim.Error(fmt.Sprintf("Invalid argument expected User protocol buffer %s\n", err.Error()))
		}
		if err := t.CreateUser(stub, &parameters); err != nil {
			logger.Error("Error creating user: ", err.Error())
			return shim.Error(err.Error())
		}
	case "find":
		parameters := UpdateUserCert{}
		if err := proto.Unmarshal(argsBytes, &parameters); err != nil {
			logger.Errorf("Invalid argument expected User protocol buffer %s\n", err.Error())
			return shim.Error(fmt.Sprintf("Invalid argument expected User protocol buffer %s\n", err.Error()))
		}
		if err := t.UpdateUserPubKey(stub, &parameters); err != nil {
			logger.Error(err.Error())
			return shim.Error(err.Error())
		}
		return shim.Success(pbmessage)
	}
	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(TestChaincode))
	if err != nil {
		logger.Error("Error starting chaincode: ", err)
	}
}
