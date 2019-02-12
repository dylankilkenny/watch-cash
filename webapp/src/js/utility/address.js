let BITBOXSDK = require('bitbox-sdk/lib/bitbox-sdk').default;
let BITBOX = new BITBOXSDK();

export default {
    ValidAddress: address => {
        try {
            BITBOX.Address.isCashAddress(address);
            BITBOX.Address.isLegacyAddress(address);
            return true;
        } catch (error) {
            return false;
        }
    },

    CashAddress: address => {
        if (BITBOX.Address.isCashAddress(address)) {
            console.log(1);
            return address;
        } else {
            console.log(2);

            return BITBOX.Address.toCashAddress(address, false);
        }
    }
};
