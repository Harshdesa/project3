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

1 byte = 8 bits<br />
bytes32 = 256 bits<br />
# of parameters = # of bits<br />
300 million parameters = 1.2 million bytes??<br />
find out if using -1 / +1 / 0 or using -1 / +1<br />
load test... try sending all parameters in an array of bytes32 and try to reach 300 million<br />
figure out the number of contracts we need<br />
make sure the number of calls to contracts are less<br />
