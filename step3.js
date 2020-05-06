const Web3 = require('web3');
const configuration = require('./configuration');
console.log(configuration.serverAddress);

/* Connect to ethereum node */
const etherUrl = configuration.etherUrl;

let web3 = new Web3();
web3.setProvider(new web3.providers.HttpProvider(etherUrl));

var contractInstance = new web3.eth.Contract(configuration.contractAbi, configuration.contractAddress);


contractInstance.methods.bitPackLocalUpdatesLoad().send({from: configuration.serverAddress, gas: 200000000000}).on('receipt', function(receipt){
    console.log(receipt); 
}); 

