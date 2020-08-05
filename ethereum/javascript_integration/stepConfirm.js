const Web3 = require('web3');
const configuration = require('./configuration');

console.log("Server address is this: ")
console.log(configuration.serverAddress);

/* Connect to ethereum node */
const etherUrl = configuration.etherUrl;

let web3 = new Web3();
web3.setProvider(new web3.providers.HttpProvider(etherUrl));

var contractInstance = new web3.eth.Contract(configuration.contractAbi, configuration.contractAddress);


console.log("Worker Address A is this: ")
contractInstance.methods.workerNodes(0).call().then(console.log);

console.log("Worker Address B is this: ")
contractInstance.methods.workerNodes(1).call().then(console.log);

console.log("Worker Address C is this: ")
contractInstance.methods.workerNodes(2).call().then(console.log);

console.log("These are the number of workers: ")
contractInstance.methods.numberOfWorkers().call().then(console.log);

console.log("These are the number of updates that have been made by the workers: ")
contractInstance.methods.numberOfUpdatesLoad().call().then(console.log);

console.log("Debug number is this: ")
contractInstance.methods.debugNumber().call().then(console.log);

console.log("Sample update is this: ")
contractInstance.methods.updates_load( configuration.workerAddressA, 0, 1499).call().then(console.log);

console.log("Global momenta is this: ")
contractInstance.methods._globalMomentaretvalue( 6000 ).call().then(console.log);
