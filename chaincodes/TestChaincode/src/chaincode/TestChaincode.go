
package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"fmt"
	"encoding/json"
	"github.com/golang/protobuf/proto"
)

var logger = shim.NewLogger("TestChaincode")


type TestChaincode struct {
}


func (t *TestChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	logger.Info("Init")

	return shim.Success(nil)
}


func (t *TestChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("Invoke")

	function, args := stub.GetFunctionAndParameters()

	switch function {

	case "invoke":

		return t.invoke(stub, args)

	case "query":

		if len(args) < 1 {
			return loggedShimError(fmt.Sprintf("Insufficient arguments found\n"))
		}

		entityRequest := new(GetEntity)
		if err := proto.Unmarshal([]byte(args[0]), entityRequest); err != nil {
			return loggedShimError(fmt.Sprintf("Invalid argument expected User protocol buffer %s\n", err.Error()))
		}

		entity, err := t.query(stub, entityRequest)

		if err != nil {
			return loggedShimError(fmt.Sprintf("Error getting entity: %s\n", err.Error()))
		}

		pbmessage, err := proto.Marshal(entity)
		if err != nil {
			return loggedShimError(fmt.Sprintf("Failed to marshal Allowed protobuf (%s)", err.Error()))
		}

		return shim.Success(pbmessage)

	}

	return loggedShimError("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}


func (t *TestChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	return shim.Success(nil)
}


// query callback representing the query of a chaincode
func (t *TestChaincode) query(stub shim.ChaincodeStubInterface, ref *GetEntity) (*Entity, error) {

	if err := checkPermissions(stub); err != nil {
		return nil, err
	}

	entity, err := getEntityFromDB(*ref, stub)

	if err != nil {
		return nil, fmt.Errorf("Error getting file from db: " + err.Error())
	}

	return entity, nil
}

// this method is required only when using CouchDB as peer database cause it stores data as JSON
func getEntityFromDB(entityName GetEntity, stub shim.ChaincodeStubInterface) (*Entity, error) {

	entity := new(Entity)

	// getting real file descriptor by key
	jsonEntity, err := stub.GetState(entityName.Name)
	if err != nil {
		return entity, fmt.Errorf("Error getting entity from db: " + err.Error())
	}

	if jsonEntity == nil {
		return entity, fmt.Errorf("Entity not found for key: " + entityName.Name)
	}

	if err := json.Unmarshal(jsonEntity, &entity); err != nil {
		return entity, fmt.Errorf("Error parsing json: " + err.Error())
	}

	return entity, nil
}

// this method is required only when using CouchDB as peer database cause it stores data as JSON
func putEntityToDB( entity Entity, stub shim.ChaincodeStubInterface) (string, error) {

	key := entity.Name

	entityJSONasBytes, err := json.Marshal(entity)
	if err != nil {
		return "", fmt.Errorf("Failed to create json for entity <%s> with error: %s" , key, err)
	}

	if err := stub.PutState(key, entityJSONasBytes); err != nil {
		return "", fmt.Errorf("Failed to store entity <%s> with error: %s" , key, err)
	}
	return key, nil
}

// Sample code to call separate Auth chaincode for permissions check
func checkPermissions(stub shim.ChaincodeStubInterface) error {
	//args := [][]byte{[]byte(permissions)}
	//
	//if subject != nil {
	//	args = append(args, subject)
	//}
	//
	//resp := stub.InvokeChaincode(authServiceChaincodeId, args, "")
	//
	//if resp.Status != 200 {
	//	logger.Info("Response status:", resp.Status)
	//	logger.Info("Response message:", resp.Message)
	//	logger.Info("Response payload:", resp.Payload)
	//
	//	return errors.New(fmt.Sprintf("Error invoking %s chaincode: (%s)", authServiceChaincodeId, resp.Message))
	//}
	//if resp.Payload == nil {
	//	return errors.New("AuthService return empty permissions")
	//}
	//
	//allowed := new(Allowed)
	//if err := proto.Unmarshal(resp.Payload, allowed); err != nil {
	//	return errors.New("Cannot parse permission protobuf")
	//}
	//
	//if !allowed.Allowed {
	//	return errors.New("Action not allowed")
	//} else {
	//	return nil
	//}
	return nil
}

// Convenience method to make sure all errors are logged and to decrease code lines number
func loggedShimError(message string) pb.Response {
	logger.Error(message)
	return shim.Error(message)
}

func main() {
	err := shim.Start(new(TestChaincode))
	if err != nil {
		logger.Errorf("Error starting chaincode: %s", err)
	}
}
