package main

import (
	"fmt"
        "errors"
        "encoding/json"
        "strconv"
        "math"

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
func (t *SimpleAsset) Init(ctx contractapi.TransactionContextInterface, NodeA string, NodeAParameters [] uint64, NodeB string, NodeBParameters [] uint64) error {

        fmt.Println("ABstore Init")
	var err error
	// Initialize the chaincode
	fmt.Printf("NodeAParameters = %d, NodeBParameters = %d\n", NodeAParameters, NodeBParameters)
	// Write the state to the ledger

        NodeAPayloadJson := Payload{
		Parameters: NodeAParameters,
	}

        NodeASubparty := []string {NodeA+"0"}
        NodeAMetadataJson := Metadata{
                Subparty: NodeASubparty,
        }
	NodeAPayload, err := json.Marshal(NodeAPayloadJson)
        NodeAMetadata, err := json.Marshal(NodeAMetadataJson)
	err = ctx.GetStub().PutState(NodeAMetadataJson.Subparty[0], NodeAPayload)
        err = ctx.GetStub().PutState(NodeA, NodeAMetadata)

        NodeBPayloadJson := Payload{
                Parameters: NodeBParameters,
        }

        NodeBSubparty := []string {NodeB+"0"}
        NodeBMetadataJson := Metadata{
                Subparty: NodeBSubparty, 
        }
        NodeBPayload, err := json.Marshal(NodeBPayloadJson)
        NodeBMetadata, err := json.Marshal(NodeBMetadataJson)
        if err != nil {
                fmt.Println("error:", err)
        }   
        err = ctx.GetStub().PutState(NodeBMetadataJson.Subparty[0], NodeBPayload) 
        err = ctx.GetStub().PutState(NodeB, NodeBMetadata)

	return nil

}

func (t *SimpleAsset) Set(ctx contractapi.TransactionContextInterface, NodeName string, NodeParameters [] uint64) error {


        // Get the state from the ledger
        NodeMetadataBytes, err := ctx.GetStub().GetState(NodeName)
        if err != nil {
          return err
        }

        if NodeMetadataBytes == nil {
                NodePayloadJson := Payload{
                       Parameters: NodeParameters,
                }

                NodeBaseSubparty := []string {NodeName+"0"}
                NodeMetadataJson := Metadata{
                       Subparty: NodeBaseSubparty, 
                }
                NodePayload, err := json.Marshal(NodePayloadJson)
                NodeMetadata, err := json.Marshal(NodeMetadataJson)

                err = ctx.GetStub().PutState(NodeMetadataJson.Subparty[0], NodePayload)
                err = ctx.GetStub().PutState(NodeName, NodeMetadata)

                if err != nil {
                        return err
                }
        } else {
            CurrentNodeMetadata := new(Metadata)
            _ = json.Unmarshal(NodeMetadataBytes, CurrentNodeMetadata)

            var Subparties []string
            Subparties = make([]string, 4687500, 4687500)
            Subparties = CurrentNodeMetadata.Subparty
            SubpartyIndex := len(Subparties)
            CurrentSubparty := []string {NodeName + strconv.Itoa(SubpartyIndex)}
            Subparties = append(Subparties, CurrentSubparty...)

	    // Write the state to the ledger

            NodeMetadataJson := Metadata{
	    	Subparty: Subparties,
	    }
            NodePayloadJson := Payload{
                       Parameters: NodeParameters,
            }
	    NodeMetadata, err := json.Marshal(NodeMetadataJson)
            NodePayload, err := json.Marshal(NodePayloadJson)

	    err = ctx.GetStub().PutState(NodeName, NodeMetadata)
            err = ctx.GetStub().PutState(CurrentSubparty[0], NodePayload)
	    if err != nil {
		return err
	    }
       }

	return nil

}


// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The Set
// method may create a new asset by specifying a new key-value pair.
func (t *SimpleAsset) Invoke(ctx contractapi.TransactionContextInterface, NodeNames [] string, ResultNodeName string, beg int, end int) error {
	// Extract the function and args from the transaction proposal
        //var err error
        var ResultParameters []uint64
        ResultParameters = make([]uint64, 4687500, 4687500)
        ResultParameters = nil
        
        allNodesMetadata := []Metadata{}
        allNodesMetadata = make([]Metadata, 10, 10)
        allNodesMetadata = nil

        //STORE METADATA
        for i := 0; i < len(NodeNames) ; i++ {
          NodeMetadataBytes, err := ctx.GetStub().GetState(NodeNames[i])
          fmt.Printf("NodeMetadataBytes: %b\n\n", NodeMetadataBytes )
          if err != nil {
                return fmt.Errorf("Failed to get state")
          }
          CurrentNodeMetadata := new(Metadata)
          _ = json.Unmarshal(NodeMetadataBytes, CurrentNodeMetadata)
          allNodesMetadata = append(allNodesMetadata, *CurrentNodeMetadata)
        }

        //GET NUMBER OF WORKER NODES
        numberOfNodes := len(allNodesMetadata)
        fmt.Printf("Number of nodes: %d\n\n", numberOfNodes)

        if len(allNodesMetadata[0].Subparty) == len(allNodesMetadata[1].Subparty) {
 
          var ResultSubparties []string
          ResultSubparties = make([]string, 4687500, 4687500)
          ResultSubparties = nil
          if (beg != 0) {
            ResultNodeMetadataBytes, err := ctx.GetStub().GetState(ResultNodeName)
            if err != nil {
                return fmt.Errorf("Failed to get state")
            }
            CurrentResultNodeMetadata := new(Metadata) 
            _ = json.Unmarshal(ResultNodeMetadataBytes, CurrentResultNodeMetadata)
            ResultSubparties = append(ResultSubparties, CurrentResultNodeMetadata.Subparty...)
          }

          //FOR EACH BATCH OF INPUTS
          for batchNumber := beg; batchNumber < end ; batchNumber++ {
            var batch []Payload
            batch = make([]Payload, 10, 10)
            batch = nil

            for nodeNumber := 0; nodeNumber < numberOfNodes; nodeNumber++ {
              Avalbytes, err := ctx.GetStub().GetState(allNodesMetadata[nodeNumber].Subparty[batchNumber])
              if err != nil {
                return fmt.Errorf("Failed to get state")
              }
              tempBatchbytes := new(Payload)
              _ = json.Unmarshal(Avalbytes, tempBatchbytes)
              batch = append(batch, *tempBatchbytes)
            }

            //For each 64 bit Parameter in a Batch
            for i := 0; i < len(batch[0].Parameters) ; i++ {
              var finalMajorityUINT64Parameter uint64 = 0

              //For each Bit of the 64 bit Parameter
              for x:= 0 ; x < 64 ; x++ {
                var mask uint64 = uint64(math.Pow(2, float64(x)))
                //TRACE LOG
                //fmt.Printf("Masker: %d\n", mask)                
                var aggregateParameter uint64 = 0

                //For each worker node
                for nodeNumber := 0; nodeNumber < numberOfNodes; nodeNumber++ {
                  aggregateParameter = aggregateParameter +  batch[nodeNumber].Parameters[i] & mask
                  //TRACE LOG
                  //fmt.Printf("aggregateParameter after & with node %d : %d\n", nodeNumber, aggregateParameter)
                }
                // END of for each worker node || aggregateParameter is addition of all the bits
                if (aggregateParameter > (mask*uint64(numberOfNodes/2)))  {
                   finalMajorityUINT64Parameter = finalMajorityUINT64Parameter + mask
                   //TRACE LOG
                   //fmt.Printf("Final Number : %d\n", finalMajorityUINT64Parameter)
                }
              }
              //END of For each Bit of the 64 bit Parameter || finalMajorityUINT64Parameter is the majority
              fmt.Printf("Final Number : %d\n", finalMajorityUINT64Parameter)
              ResultParameters = append(ResultParameters, finalMajorityUINT64Parameter)
              fmt.Printf("length of ResultParameters[%d] : %d",i,ResultParameters[i])
              fmt.Printf("Subparty: %d",batchNumber)
            }
            //END of For each 64 bit Parameter in a Batch || finalMajorityUINT64Parameter is loaded into ResultParameters

            CurrentResultSubparty := []string {ResultNodeName + strconv.Itoa(batchNumber)}
            ResultNodePayloadJson := Payload{
                Parameters: ResultParameters,
            }
            ResultNodePayload, err := json.Marshal(ResultNodePayloadJson)
            if err != nil {
                 fmt.Println("error:", err)
            }
            ResultSubparties = append(ResultSubparties, CurrentResultSubparty...)
            err = ctx.GetStub().PutState(CurrentResultSubparty[0], ResultNodePayload)
            ResultParameters = nil
          }
          //END of for each Batch of Inputs || Batches of Parameters are loaded

          ResultNodeMetadataJson := Metadata{
                Subparty: ResultSubparties,
          }
          ResultNodeMetadata, err := json.Marshal(ResultNodeMetadataJson)
          err = ctx.GetStub().PutState(ResultNodeName, ResultNodeMetadata)
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
        for i := 0; i < index ; i++ {
           fmt.Printf("Query Response Subparty:%s\n", testAstr.Subparty[i])
           Avalbytes, err := ctx.GetStub().GetState(testAstr.Subparty[i])
           if err != nil {
                jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
                return "", errors.New(jsonResp)
           }
           testA := new(Payload)
           _ = json.Unmarshal(Avalbytes, testA)
       
           fmt.Printf("Number of Parameters committed: %d\n", len(testA.Parameters))
           indext := len(testA.Parameters)
           fmt.Printf("Query Response Last Parameter:%d\n", testA.Parameters[indext-1])
        }
	//fmt.Printf("Query Response Last Subparty:%s\n", testAstr.Subparty[index-1])
        //Avalbytes, err := ctx.GetStub().GetState(testAstr.Subparty[index-1])
        //testA := new(Payload)
        //_ = json.Unmarshal(Avalbytes, testA)
       
        //fmt.Printf("Number of Parameters committed: %d\n", len(testA.Parameters))
        //indext := len(testA.Parameters)
        //fmt.Printf("Query Response Last Parameter:%d\n", testA.Parameters[indext-1])

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
