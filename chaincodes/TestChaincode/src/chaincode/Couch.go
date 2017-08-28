package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/golang/protobuf/jsonpb"
	"encoding/json"
	"fmt"
)

// this method is required only when using CouchDB as peer database cause it stores data as JSON
func (t *TestChaincode) getEntityFromDB(stub shim.ChaincodeStubInterface, entityName *GetEntity) (*Entity, error) {

	entity := new(Entity)

	// getting real file descriptor by key
	jsonEntity, err := stub.GetState(entityName.Name)
	if err != nil {
		return entity, fmt.Errorf("Error getting entity from db: " + err.Error())
	}

	if jsonEntity == nil {
		return entity, fmt.Errorf("Entity not found for key: " + entityName.Name)
	}

	if len(jsonEntity) == 0 {
		return nil, fmt.Errorf("Empty value found for key:  " + entity.Name)
	}

	if err := json.Unmarshal(jsonEntity, &entity); err != nil {
		return entity, fmt.Errorf("Error parsing json: " + err.Error())
	}

	return entity, nil
}

// this method is required only when using CouchDB as peer database cause it stores data as JSON
func (t *TestChaincode) putEntityToDB(stub shim.ChaincodeStubInterface, entity *Entity) (*GetEntity, error) {

	key := entity.Name

	entityJSONasBytes, err := json.Marshal(entity)
	if err != nil {
		return nil, fmt.Errorf("Failed to create json for entity <%s> with error: %s" , key, err)
	}

	if err := stub.PutState(key, entityJSONasBytes); err != nil {
		return nil, fmt.Errorf("Failed to store entity <%s> with error: %s" , key, err)
	}
	return &GetEntity{Name:key}, nil
}



// this method is required only when using CouchDB as peer database cause it stores data as JSON
func (t *TestChaincode) getHistoryFromDB(stub shim.ChaincodeStubInterface, entityName *GetEntity) (*History, error) {

	key := entityName.Name

	history := new(History)

	// getting real file descriptor by key
	historyIterator, err := stub.GetHistoryForKey(key)
	if err != nil {
		return nil, fmt.Errorf("Error getting history from db: " + err.Error())
	}

	if historyIterator == nil {
		return nil, fmt.Errorf("Entity not found for key: " + key)
	}

	if !historyIterator.HasNext() {
		return nil, fmt.Errorf("Empty value found for key:  " + key)
	}

	for historyIterator.HasNext() {
		km, err := historyIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("Error iterating over history: " + err.Error())
		}

		item := new(HistoryItem)

		item.TxId = km.TxId
		item.Timestamp = km.Timestamp.Seconds * 1000 + int64(km.Timestamp.Nanos) / 1000
		item.IsDelete = km.IsDelete

		value := new(Entity)
		if err := jsonpb.UnmarshalString(string(km.Value), value); err != nil {
			return nil, fmt.Errorf("Error unmarshaling balanve value: " + err.Error())
		}

		item.Value = value

		history.History = append(history.History, item)
	}

	history.Key = key

	return history, nil
}