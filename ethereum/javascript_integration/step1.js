const Web3 = require('web3');
const configuration = require('./configuration');
console.log(configuration.serverAddress);

/* Connect to ethereum node */
const etherUrl = configuration.etherUrl;

let web3 = new Web3();
web3.setProvider(new web3.providers.HttpProvider(etherUrl));

var contractInstance = new web3.eth.Contract(configuration.contractAbi, configuration.contractAddress);

//Set number of workers by server
contractInstance.methods.setNumberOfWorkers(3).send({from: configuration.serverAddress}).on('receipt', function(receipt){
    console.log(receipt);});
