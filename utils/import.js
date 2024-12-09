const getFunctionLocation = require('get-function-location');

const {
  chainTypeToValidator,
  getAll,
} = require('multicoin-address-validator/src/currencies');

console.log('getting currencies');
const allCurrencies = getAll();

console.log('getting currencies');


const pp = [];

allCurrencies.forEach(cc => {
  pp.push( getFunctionLocation(сс[0].validator.isValidAddress))
  cc.p = p;
});

console.log('done', allCurrencies);
