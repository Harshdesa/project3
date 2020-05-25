/*
 * Copyright IBM Corp All Rights Reserved
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"
        "errors"
        "encoding/json"

        "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SimpleAsset implements a simple chaincode to manage an asset
type SimpleAsset struct {
          contractapi.Contract
}

type Payload struct {
	Parameters []uint64
}

// Init is called during chaincode instantiation to initialize any
// data. Note that chaincode upgrade also calls this function to reset
// or to migrate data.
func (t *SimpleAsset) Init(ctx contractapi.TransactionContextInterface, A string, Aval [] uint64, B string, Bval [] uint64) error {

        fmt.Println("ABstore Init")
	var err error
	// Initialize the chaincode
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)
	// Write the state to the ledger

        Atest := Payload{
		Parameters: Aval,
	}
	a, err := json.Marshal(Atest)
	if err != nil {
		fmt.Println("error:", err)
	}
	err = ctx.GetStub().PutState(A, a)
	if err != nil {
		return err
	}

        Btest := Payload{
                Parameters: Bval,
        }
        b, err := json.Marshal(Btest)
        if err != nil {
                fmt.Println("error:", err)
        }   
        err = ctx.GetStub().PutState(B, b) 

	return nil

}

func (t *SimpleAsset) Set(ctx contractapi.TransactionContextInterface, A string, Aval [] uint64, B string, Bval [] uint64) error {

        fmt.Println("ABstore Init")
	var err error
	// Initialize the chaincode
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)
	// Write the state to the ledger

        Atest := Payload{
		Parameters: Aval,
	}
	a, err := json.Marshal(Atest)
	if err != nil {
		fmt.Println("error:", err)
	}
	err = ctx.GetStub().PutState(A, a)
	if err != nil {
		return err
	}

        Btest := Payload{
                Parameters: Bval,
        }
        b, err := json.Marshal(Btest)
        if err != nil {
                fmt.Println("error:", err)
        }   
        err = ctx.GetStub().PutState(B, b) 

	return nil

}



// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The Set
// method may create a new asset by specifying a new key-value pair.
func (t *SimpleAsset) Invoke(ctx contractapi.TransactionContextInterface, A, B, C string) error {
	// Extract the function and args from the transaction proposal
        var err error
        var Cval []uint64
        Cval = make([]uint64, 4687500, 4687500)
	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := ctx.GetStub().GetState(A)
	if err != nil {
		return fmt.Errorf("Failed to get state")
	}
	if Avalbytes == nil {
		return fmt.Errorf("Entity not found")
	}

        testA := new(Payload)
	_ = json.Unmarshal(Avalbytes, testA)


	Bvalbytes, err := ctx.GetStub().GetState(B)
	if err != nil {
		return fmt.Errorf("Failed to get state")
	}
	if Bvalbytes == nil {
		return fmt.Errorf("Entity not found")
	}
        
        testB := new(Payload)
        _ = json.Unmarshal(Bvalbytes, testB) 

        fmt.Printf("Reached Here: ")
        fmt.Printf("length of testA.Parameters : %d", len(testA.Parameters))
        fmt.Printf("length of testB.Parameters : %d", len(testA.Parameters))
        fmt.Printf("Reached Here: ")

        if len(testA.Parameters) == len(testB.Parameters) {
          for i := 0; i < len(testA.Parameters) ; i++ {
              Cval[i] = testA.Parameters[i] + testB.Parameters[i]
              fmt.Printf("length of Cval[%d] : %d",i,Cval[i])
          }
        }

        Ctest := Payload{
                Parameters: Cval,
        }
        c, err := json.Marshal(Ctest)
        if err != nil {
                fmt.Println("error:", err)
        }   
        err = ctx.GetStub().PutState(C, c)


        return nil

}

func (t *SimpleAsset) Delete(ctx contractapi.TransactionContextInterface, A string) error {

	// Delete the key from the state in ledger
	err := ctx.GetStub().DelState(A)
	if err != nil {
		return fmt.Errorf("Failed to delete state")
	}

	return nil
}

func (t *SimpleAsset) Query(ctx contractapi.TransactionContextInterface, A string) (string, error) {
	var err error
	// Get the state from the ledger
	Avalbytes, err := ctx.GetStub().GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return "", errors.New(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return "", errors.New(jsonResp)
	}

        testA := new(Payload)
        _ = json.Unmarshal(Avalbytes, testA)

	fmt.Printf("Query Response:%d\n", testA.Parameters[0])
	fmt.Printf("Query Response:%d\n", testA.Parameters[1])
	fmt.Printf("Query Response:%d\n", testA.Parameters[2])
        
	return "success", nil
}

// main function starts up the chaincode in the container during instantiate
func main() {


        cc, err := contractapi.NewChaincode(new(SimpleAsset))
        if err != nil {
		panic(err.Error())
	}
	if err := cc.Start(); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}
