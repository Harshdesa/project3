module.exports = { contractAddress: "0xdDB6351AE2683bF5Ac4615933Efa2D7C107F76a5",
                   serverAddress:   "0xA49dAC1E098775A2E5Ef1E0D7B31FA29178f80e8",
                   workerAddressA:  "0xc362EbB1aC74d23C1F5871016A9e47aCb80B9019",
                   workerAddressB:  "0xd4dd600f790A33faF6265745f15cc547076791CC",
                   workerAddressC:  "0xDc2ECA72F09837560b4aa2C23BbcEbf061A91544",
                   etherUrl: "http://localhost:8545"
 };
module.exports.contractAbi = [ { "constant": true, "inputs": [ { "name": "", "type": "uint256" } ], "name": "_globalMomentaretvalue", "outputs": [ { "name": "", "type": "bytes32" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "uint256" } ], "name": "workerNodes", "outputs": [ { "name": "", "type": "address" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "address" }, { "name": "", "type": "uint256" }, { "name": "", "type": "uint256" } ], "name": "updates_load", "outputs": [ { "name": "", "type": "bytes32" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "uint256" } ], "name": "globalMomenta", "outputs": [ { "name": "", "type": "bytes32" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "uint256" }, { "name": "", "type": "uint256" } ], "name": "globalMomenta_load", "outputs": [ { "name": "", "type": "bytes32" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "uint256" } ], "name": "_globalMomenta", "outputs": [ { "name": "", "type": "bytes32" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "numberOfWorkers", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [ { "name": "", "type": "uint256" } ], "name": "debugArray", "outputs": [ { "name": "", "type": "bytes32" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "debugNumber", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "serverNode", "outputs": [ { "name": "", "type": "address" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": true, "inputs": [], "name": "numberOfUpdatesLoad", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "inputs": [], "payable": false, "stateMutability": "nonpayable", "type": "constructor" }, { "anonymous": false, "inputs": [ { "indexed": false, "name": "globalMomenta", "type": "bytes32" } ], "name": "SendUpdate", "type": "event" }, { "constant": false, "inputs": [ { "name": "_numberOfWorkers", "type": "uint256" } ], "name": "setNumberOfWorkers", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [], "name": "getNumberOfWorkers", "outputs": [ { "name": "", "type": "uint256" } ], "payable": false, "stateMutability": "view", "type": "function" }, { "constant": false, "inputs": [ { "name": "_update", "type": "bytes32[]" } ], "name": "sendLocalUpdateByWorkerLoad", "outputs": [], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": false, "inputs": [], "name": "bitPackLocalUpdatesLoad", "outputs": [ { "name": "", "type": "bytes32[]" } ], "payable": false, "stateMutability": "nonpayable", "type": "function" }, { "constant": true, "inputs": [], "name": "getGlobalMomenta", "outputs": [ { "name": "", "type": "bytes32[]" } ], "payable": false, "stateMutability": "view", "type": "function" } ];


/*
// RUN BELOW COMMAND TO START LOCAL BLOCKCHAIN IN GANACHE
// WHY ? -> So that, accounts are created with above addresses
 
ganache-cli -g 1 -l 2000000000000 --allowUnlimitedContractSize --account 0x02927dfe6b8eec0e3e8a0ad6a78be2516e8ae3d6768ed166881094856f0217,100000000000000000000 --account 0x02927dfe6b8eec0e3e8a0ad6a78be2516e8ae3d6768ed16688601094856f0218,100000000000000000000 --account 0x02927dfe6b8eec0e3e8a0ad6a78be2516e8ae3d6768ed16688601094856f0219,100000000000000000000 --account 0x02927dfe6b8eec0e3e8a0ad6a78be2516e8ae3d6768ed16688601094856f0214,100000000000000000000
*/