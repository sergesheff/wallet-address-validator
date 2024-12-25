const getFunctionLocation = require('get-function-location');
const fs = require('node:fs');




const crypto = require('multicoin-address-validator/src/crypto/utils');

const cr = require('multicoin-address-validator/src/ethereum_validator');
 

cr.isValidAddress('0xE37c0D48d68da5c5b14E5c1a9f1CFE802776D9FF');

// console.log( crypto.sha256("12345688", 'B64'));
// console.log( crypto.sha256('12345688', 'B64'));

console.log( crypto.hexStr2byteArray('12345688'));
console.log( crypto.hexStr2byteArray('12345688'));

return




const packageString = 'package base';
const currencyTypesFileName = 'base/lst_currency_types.go';
const currenciesFileName = 'base/lst_currencies.go';

const fnParseAddressType = (arr) => {
  if (typeof arr === 'array') {
    return arr.prod.map((l) => `"${l}"`).join(',');
  }

  return undefined;
};

// function for formatting Address Types
const fnAddressType = (obj) => {
  let content = '';

  if (obj) {
    if (obj['prod']) {
      const data = fnParseAddressType(obj['prod']);

      if (data) {
        content += `Prod: [${data.map((l) => `"${l}"`).join(',')}],`;
      }
    }

    if (obj['testNet']) {
      const data = fnParseAddressType(obj['testNet']);

      if (data) {
        content += `TestNet: [${data.map((l) => `"${l}"`).join(',')}],`;
      }
    }

    if (obj['stageNet']) {
      const data = fnParseAddressType(obj['stageNet']);

      if (data) {
        content += `StageNet: [${data.map((l) => `"${l}"`).join(',')}],`;
      }
    }
  }

  return content;
};

// function which cleaning up the currency name
const fnCleanup = (name) => {
  let newName = undefined;

  // removing non letter or non digit symbols
  const arr = (name || '').split(/[^\w\d]/g);

  arr.forEach((t) => {
    t = (t + '').trim();
    if (t) {
      newName = (newName || '') + t.charAt(0).toUpperCase() + t.slice(1);
    }
  });

  return newName;
};

const fnInitCurrency = async (mCurrencies) => {
  let content = packageString + '\n';
  content +='import . "github.com/sergesheff/ref"\n';
  content += 'var Currencies = map[string]*Currency{\n';

  mCurrencies.forEach((v, k) => {
    content += `"${k}": {`;
    content += `Name: ${v.goType},`;
    content += `Symbol: "${v.symbol}",`;

    if (v.minLength) {
      content += `MinLength: Ref(${v.minLength}),`;
    }

    if (v.maxLength) {
      content += `MaxLength: Ref(${v.maxLength}),`;
    }

    if (v.expectedLength) {
      content += `ExpectedLength: Ref(${v.expectedLength}),`;
    }

    if (v.hashFunction) {
      content += `HashFunction: Ref("${v.hashFunction}"),`;
    }

    if (v.regex) {
      content += `Regex: Ref("${v.regex}"),`;
    }

    const bech32Hrp = fnAddressType(v.bech32Hrp);
    if (bech32Hrp) {
      content += `Bech32Hrp: &AddressType{${bech32Hrp}}`;
    }

    const addressTypes = fnAddressType(v.addressTypes);
    if (addressTypes) {
      content += `AddressTypes: &AddressType{${addressTypes}}`;
    }

    const iAddressTypes = fnAddressType(v.iAddressTypes);
    if (iAddressTypes) {
      content += `IAddressTypes: &AddressType{${iAddressTypes}}`;
    }

    content += '},\n';
  });

  content += '}';

  try {
    await fs.writeFile(currenciesFileName, content, 'utf-8', (err) => {
      if (err) {
        console.error(err);
      }
    });
  } catch (ex) {
    if (ex) {
      throw err;
    }
  }
};

// function for writing a currency types
const fnInitCurrencyType = async (mCurrencies) => {
  let content = packageString + '\n';
  content += 'const (\n';

  mCurrencies.forEach((v) => {
    content += `${v.goType} CurrencyTypes = "${v.symbol}" \n`;
  });

  content += ')';

  // writing the files
  try {
    await fs.writeFile(currencyTypesFileName, content, 'utf-8', (err) => {
      if (err) {
        throw err;
      }
    });
  } catch (ex) {
    console.error(ex);
    mCurrencies = undefined;
  }

  return mCurrencies;
};

const {
  chainTypeToValidator,
  getAll,
} = require('multicoin-address-validator/src/currencies');
const { log } = require('node:console');

console.log('getting currencies');
const allCurrencies = getAll();

console.log('getting currencies');

const mCurrencies = new Map();

allCurrencies.forEach(async (cc) => {
  // getting non empty entities only
  if (cc && cc.name) {
    // cleaning up the currency name
    const newName = fnCleanup(cc.name);

    if (newName) {
      if (!mCurrencies.has(newName)) {
        newCurrency = Object.assign({}, cc);
        newCurrency.goType = `CurrencyType${newName}`;

        // adding currency to the map
        mCurrencies.set(newName, newCurrency);
      }
    }
  }

  const fn = await getFunctionLocation(cc.validator.isValidAddress);
  cc.fn = fn.source;
});

// init Currency Types
fnInitCurrencyType(mCurrencies);

// init currencies
fnInitCurrency(mCurrencies);

console.log('done', allCurrencies);
