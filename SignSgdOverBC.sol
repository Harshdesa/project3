pragma solidity >=0.4.0 <0.7.0;

contract SignSgdOverBC {
    uint public numberOfWorkers = 0;
    uint public numberOfUpdates = 0;

    //array for the global momenta over time
    bytes32[] public globalMomenta;



    //updates per address
    mapping(address => bytes32[]) public updates;

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

    function sendLocalUpdateByWorker(bytes32 _update) public {
        require(msg.sender != serverNode, "Illegal operation performed by serverNode");
        require(numberOfUpdates++ < numberOfWorkers, "Multiple operations performed by worker nodes");
        updates[msg.sender].push(_update);
        if(globalMomenta.length == 0) {
          workerNodes.push(msg.sender);
        }
    }

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

    function getGlobalMomenta() public view returns (bytes32[] memory) {
        return globalMomenta;
    }
}
