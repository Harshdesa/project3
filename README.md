Drop in Remix<br />

Plugins to enable:<br />
  Solidity compiler<br />
  Deploy and Run Transactions<br />

Steps:<br /> 


  As server:<br />
    setNumberOfWorkers = 3 (or any number)<br />

  As worker for each worker:<br />
    sendLocalUpdateByWorker = 0x0000000000000000000000000000000000000000000000000000000000000001 (or any bytes32)<br />

  As server:<br />
    bitPackLocalUpdates<br />


Future Work:<br />
To add in sendLocalUpdateByWorker<br />

require(updates[msg.sender][globalMomenta.length] == "0x0", "Multiple operations performed by worker node. Please wait till global momenta is sent out.");<br />
require(numberOfUpdates++ < numberOfWorkers, "Not enough worker nodes have sent local updates");<br />
