pragma solidity >=0.4.0 <0.7.0;

contract SignSgdOverBC {
    uint public numberOfWorkers = 0;

    uint public debugNumber = 0;
    mapping(uint =>  bytes32) public debugArray;

    //number of Updates 
    //uint public numberOfUpdates = 0;
    //number of Updates (Load)
    uint public numberOfUpdatesLoad = 0;

    //array for the global momenta over time
    bytes32[] public globalMomenta;
    //array for the global momenta (load test) over time
    bytes32[][] public globalMomenta_load;
    mapping(uint => bytes32) public _globalMomenta;
    bytes32[] public _globalMomentaretvalue;


    //updates per address
    //mapping(address => bytes32[]) public updates;
    //updates (load test) address
    mapping(address => bytes32[][]) public updates_load;


    //list of worker Nodes
    address[] public workerNodes;

    address public serverNode;

    event SendUpdate(
      bytes32 globalMomenta
    );  

    constructor() public {
      serverNode = msg.sender;
    }   

    function setNumberOfWorkers(uint _numberOfWorkers) public {
      require(serverNode == msg.sender, "Illegal operation performed by worker node.");
      require(numberOfWorkers == 0, "Cannot set number of workers more than once");
      numberOfWorkers = _numberOfWorkers;
    }   

    function getNumberOfWorkers() view public returns (uint) {
      return numberOfWorkers;
    }   


    // send local updates per worker (load test)
    function sendLocalUpdateByWorkerLoad(bytes32[] memory _update) public {
        require(msg.sender != serverNode, "Illegal operation performed by serverNode");
        require(updates_load[msg.sender].length == globalMomenta_load.length, "Multiple operations performed by worker node. Please wait till global momenta is sent out.");
        require(numberOfUpdatesLoad++ < numberOfWorkers, "Worker node limit has reached. Cannot send local updates");
        
        updates_load[msg.sender].push(_update);

        if(globalMomenta_load.length == 0) {
          workerNodes.push(msg.sender);
        }
    }

/*
    function sendLocalUpdateByWorker(bytes32 _update) public {
        require(msg.sender != serverNode, "Illegal operation performed by serverNode");
        require(updates[msg.sender].length == globalMomenta.length, "Multiple operations performed by worker node. Please wait till global momenta is sent out.");
        require(numberOfUpdates++ < numberOfWorkers, "Worker node limit has reached. Cannot send local updates");
        updates[msg.sender].push(_update);
        if(globalMomenta.length == 0) {
          workerNodes.push(msg.sender);
        }
    }
*/
    function bitPackLocalUpdatesLoad() public returns (bytes32[] memory) {
        require(msg.sender == serverNode, "Illegal operation performed by workerNodes");
        require(numberOfUpdatesLoad == numberOfWorkers, "Every worker node has committed");

        //bytes32[] memory _globalMomenta;
        //mapping(uint =>  bytes32) storage _globalMomenta = debugArray;


        for (uint indexOfLoadArray = 0; indexOfLoadArray < updates_load[workerNodes[0]][0].length ; indexOfLoadArray++) { //for each array
          bytes32 zerovalue = 0x0000000000000000000000000000000000000000000000000000000000000000;
          _globalMomenta[indexOfLoadArray] = zerovalue;
          for (uint i = 0; i < workerNodes.length; i++) { // for each worker
            //debugArray[debugNumber] = updates_load[workerNodes[i]][globalMomenta_load.length][indexOfLoadArray];
            debugNumber++;

            _globalMomenta[indexOfLoadArray] = _globalMomenta[indexOfLoadArray] ^ updates_load[workerNodes[i]][globalMomenta_load.length][indexOfLoadArray];
          }
          _globalMomentaretvalue.push(_globalMomenta[indexOfLoadArray]);
        }
         //globalMomenta_load.push(_globalMomenta);
         //for (uint i = 0; i < updates_load[workerNodes[0]][0].length; i++) {
         //      _globalMomentaretvalue.push(_globalMomenta[i]);
         //}
         numberOfUpdatesLoad = 0;
         //emit SendUpdate(_globalMomenta);
         return _globalMomentaretvalue;
    }
    
/*
    function bitPackLocalUpdates() public returns (bytes32) {
        require(msg.sender == serverNode, "Illegal operation performed by workerNodes");
        require(numberOfUpdates == numberOfWorkers, "Every worker node has committed");

        bytes32 _globalMomenta;

        for (uint i = 0; i < workerNodes.length; i++) {
            _globalMomenta = _globalMomenta ^ updates[workerNodes[i]][globalMomenta.length];
        }
         globalMomenta.push(_globalMomenta);
         numberOfUpdates = 0;
         emit SendUpdate(_globalMomenta);
         return _globalMomenta;
    }
*/
    function getGlobalMomenta() public view returns (bytes32[] memory) {
        return globalMomenta;
    }
}
