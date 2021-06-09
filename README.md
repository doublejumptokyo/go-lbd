# go-lbd
Library for Golang to call [LINE Blockchain Developers API](https://docs-blockchain.line.biz/api-guide/). When you use Golang to develop a service, it allows you to easily send a request without separately setting HTTP connection or [signature](https://docs-blockchain.line.biz/api-guide/Authentication?id=generating-a-signature).

## For testing

### Environment values
Set environment values

```sh
export API_KEY=
export API_SECRET=
export SERVICE_ID=

export OWNER_ADDR=
export OWNER_SECRET=

export SERVICETOKEN_CONTRACT_ID=
export ITEMTOKEN_CONTRACT_ID=

export USER_ID=
```


### Transactions

Set environment value `TX=1`.
