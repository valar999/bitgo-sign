# BitGo Sign Transaction example

This example shows how to sign a transaction with your private key.

```golang
hash, signature, err := SignTx(key, address, amountStr, expire, seqId)
```

Params:
- `key` - ECDSA private key in hex. See below
- `address` - recipient address in `0x1234...` format
- `amountStr` - transaction amount, as integer
- `expire` - timestamp when transaction should expire, see BigGo
  documentation
- `seqId` - internal transaction sequence number, received from `tx/build`

Result:
- `hash` - operationHash, hash of recipient, amount, expire and seqId
- `signature` - signature of hash with the key

## Key

BitGo gives a key for Ethereum in "wallet" format. It is encrypted with
password and looks like this:

```
{"iv":"XJSKS29eLUoF5FDCdg2g2Q==","v":1,"iter":10000,"ks":256,"ts":64,"mode"
:"ccm","adata":"","cipher":"aes","salt":"m2F+tC4jtTs=","ct":"XXfuUJD4XOvQTz
x0tJGye3T7TtNbiqFTwaLBNInukGdEoXeLZSpIcWWfhLg6LVmTdwz3tk41aqW9fbX2zh8/l/B1h
zXngBVf4QWM5Q2ZQHswva6LBhBzpGGcy6auRVDyxNcXm59db/oHUvw+59pdV2ZKpwogRh0="}
```

First we need to convert it to "xprv" format. We can do it with following JS
code:

```js
import { BitGoAPI } from '@bitgo/sdk-api';

const bitgo = new BitGoAPI({});

const key = '{"iv":"XJSKS29eLUoF5FDCdg2g2Q==","v":1,...';
const password = "secret-password";

const prv = bitgo.decrypt({ input: key, password });
```
