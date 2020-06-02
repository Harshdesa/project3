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
        "strconv"

        "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SimpleAsset implements a simple chaincode to manage an asset
type SimpleAsset struct {
          contractapi.Contract
}

type Payload struct {
	Parameters []uint64
}

type Metadata struct {
        Subparty []string
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

        Astr := []string {A+"0"}
        Ateststr := Metadata{
                Subparty: Astr,
        }
	a, err := json.Marshal(Atest)
        astr, err := json.Marshal(Ateststr)
	if err != nil {
		fmt.Println("error:", err)
	}
	err = ctx.GetStub().PutState(Ateststr.Subparty[0], a)
        err = ctx.GetStub().PutState(A, astr)
	if err != nil {
		return err
	}

        Btest := Payload{
                Parameters: Bval,
        }

        Bstr := []string {B+"0"}
        Bteststr := Metadata{
                Subparty: Bstr, 
        }
        b, err := json.Marshal(Btest)
        bstr, err := json.Marshal(Bteststr)
        if err != nil {
                fmt.Println("error:", err)
        }   
        err = ctx.GetStub().PutState(Bteststr.Subparty[0], b) 
        err = ctx.GetStub().PutState(B, bstr)

	return nil

}

func (t *SimpleAsset) Set(ctx contractapi.TransactionContextInterface, A string, Aval [] uint64) error {


        var err error
        // Get the state from the ledger
        Astrbytes, err := ctx.GetStub().GetState(A)
        if err != nil {
                return err
        }   

        if Astrbytes == nil {
                Atest := Payload{
                       Parameters: Aval,
                }

                Astr := []string {A+"0"}
                Ateststr := Metadata{
                       Subparty: Astr, 
                }
                a, err := json.Marshal(Atest)
                astr, err := json.Marshal(Ateststr)

                if err != nil {
                       fmt.Println("error:", err)
                }
                err = ctx.GetStub().PutState(Ateststr.Subparty[0], a)
                err = ctx.GetStub().PutState(A, astr)

                if err != nil {
                        return err
                }
        } else {
            testAstr := new(Metadata)
            _ = json.Unmarshal(Astrbytes, testAstr)

            var tempAstr []string
            tempAstr = make([]string, 4687500, 4687500)
            tempAstr = testAstr.Subparty
            Astriter := len(tempAstr)
            Astr := []string {A + strconv.Itoa(Astriter)}
            tempAstr = append(tempAstr, Astr...)

	    // Write the state to the ledger

            Ateststr := Metadata{
	    	Subparty: tempAstr,
	    }
	    astr, err := json.Marshal(Ateststr)
            if err != nil {
		fmt.Println("error:", err)
	    }

            Atest := Payload{
                       Parameters: Aval,
            }
            a, err := json.Marshal(Atest)

	    err = ctx.GetStub().PutState(A, astr)
            err = ctx.GetStub().PutState(Astr[0], a)
	    if err != nil {
		return err
	    }
       }

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
        Cval = nil
	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Astrbytes, err := ctx.GetStub().GetState(A)
	if err != nil {
		return fmt.Errorf("Failed to get state")
	}
	if Astrbytes == nil {
		return fmt.Errorf("Entity not found")
	}

        testAstr := new(Metadata)
	_ = json.Unmarshal(Astrbytes, testAstr)


	Bstrbytes, err := ctx.GetStub().GetState(B)
	if err != nil {
		return fmt.Errorf("Failed to get state")
	}
	if Bstrbytes == nil {
		return fmt.Errorf("Entity not found")
	}
        
        testBstr := new(Metadata)
        _ = json.Unmarshal(Bstrbytes, testBstr) 

        fmt.Printf("Reached Here: ")
        fmt.Printf("length of testAstr.Subparty : %d", len(testAstr.Subparty))
        fmt.Printf("length of testBstr.Subparty : %d", len(testBstr.Subparty))
        fmt.Printf("Reached Here: ")

        if len(testAstr.Subparty) == len(testBstr.Subparty) {
 
          var tempCstr []string
          tempCstr = make([]string, 4687500, 4687500)
          tempCstr = nil


          for bigI := 0; bigI < len(testAstr.Subparty) ; bigI++ {
            Avalbytes, err := ctx.GetStub().GetState(testAstr.Subparty[bigI])
            testA := new(Payload)
            _ = json.Unmarshal(Avalbytes, testA)
 
            Bvalbytes, err := ctx.GetStub().GetState(testBstr.Subparty[bigI])
            testB := new(Payload)
            _ = json.Unmarshal(Bvalbytes, testB)


            fmt.Printf("Reached Here: ")
            fmt.Printf("length of testA.Parameters : %d", len(testA.Parameters))
            fmt.Printf("length of testB.Parameters : %d", len(testB.Parameters))
            fmt.Printf("Reached Here: ")

            for i := 0; i < len(testA.Parameters) ; i++ {
              tempCval := testA.Parameters[i] + testB.Parameters[i]
              Cval = append(Cval, tempCval)
              //Cval[i] = testA.Parameters[i] + testB.Parameters[i]
              fmt.Printf("length of Cval[%d] : %d",i,Cval[i])
              fmt.Printf("Subparty: %d",bigI)
            }

            Cstr := []string {C + strconv.Itoa(bigI)}

            Ctest := Payload{
                Parameters: Cval,
            }
            c, err := json.Marshal(Ctest)
            if err != nil {
                 fmt.Println("error:", err)
            }

            tempCstr = append(tempCstr, Cstr...)

            err = ctx.GetStub().PutState(Cstr[0], c)
            Cval = nil
          }


          Cteststr := Metadata{
                Subparty: tempCstr,
          }
          cstr, err := json.Marshal(Cteststr)
          err = ctx.GetStub().PutState(C, cstr)
          if err != nil {
                 fmt.Println("error:", err)
          }
        }

        return nil

}

func (t *SimpleAsset) Delete(ctx contractapi.TransactionContextInterface, A string) error {

	// Delete the key from the state in ledger
        Astrbytes, err := ctx.GetStub().GetState(A)
        testAstr := new(Metadata)
        _ = json.Unmarshal(Astrbytes, testAstr)
        for i := 0; i < len(testAstr.Subparty) ; i++ {
            err := ctx.GetStub().DelState(testAstr.Subparty[i])
            if err != nil {
                return fmt.Errorf("Failed to delete state")
            }
        }
	err = ctx.GetStub().DelState(A)
	if err != nil {
		return fmt.Errorf("Failed to delete state")
	}

	return nil
}

func (t *SimpleAsset) Query(ctx contractapi.TransactionContextInterface, A string) (string, error) {
	var err error
	// Get the state from the ledger
	Astrbytes, err := ctx.GetStub().GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return "", errors.New(jsonResp)
	}

	if Astrbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return "", errors.New(jsonResp)
	}

        testAstr := new(Metadata)
        _ = json.Unmarshal(Astrbytes, testAstr)

        fmt.Printf("Number of Subparties committed: %d\n", len(testAstr.Subparty))
        index := len(testAstr.Subparty)
	fmt.Printf("Query Response Last Subparty:%s\n", testAstr.Subparty[index-1])
        Avalbytes, err := ctx.GetStub().GetState(testAstr.Subparty[index-1])
        testA := new(Payload)
        _ = json.Unmarshal(Avalbytes, testA)
       
        fmt.Printf("Number of Parameters committed: %d\n", len(testA.Parameters))
        indext := len(testA.Parameters)
        fmt.Printf("Query Response Last Parameter:%d\n", testA.Parameters[indext-1])

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
