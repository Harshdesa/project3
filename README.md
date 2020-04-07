Drop in Remix

Plugins to enable: 
  Solidity compiler
  Deploy and Run Transactions

Steps: 


  As server: 
    setNumberOfWorkers = 3 (or any number)

  As worker for each worker:
    sendLocalUpdateByWorker = 0x0000000000000000000000000000000000000000000000000000000000000001 (or any bytes32)

  As server:
    bitPackLocalUpdates
