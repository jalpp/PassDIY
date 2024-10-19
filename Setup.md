## Hashicorp Setup

To allow PassDIY to store and connect to your Hashicorp vault you must create a [service principle](https://developer.hashicorp.com/hcp/docs/hcp/iam/service-principal) with ```Vault Secrets App Manager``` permission. You would also need to set up following Env variables to successfully connect and store the secrets to specific app in your vault. If you don't care about storing the values you can ignore the below variables and directly run 
`go run main.go`

`HCP_CLIENT_ID` your sp client ID

`HCP_CLIENT_SECRET` your sp secret

`HCP_ORG_ID` your HCP org ID

`HCP_PROJECT_ID` your HCP project ID

`HCP_APP_NAME` your HCP app name

`HCP_API_TOKEN` your HCP API token generated from `HCP_CLIENT_ID` and `HCP_CLIENT_SECRET` you don't need generate `HCP_API_TOKEN` every time it expires the PassDIY's `hcpvaultconnect` command will handle it by automatically connecting for you

you can now clone the repo and run from source

`go run main.go`

can also build PassDIY with build command and run passdiy

`go build`

`./passdiy`

## 1Password Setup

To allow PassDIY to connect to your 1Password Vault you would need to set [service principle](https://developer.1password.com/docs/sdks)

`export OP_SERVICE_ACCOUNT_TOKEN=<your-service-account-token>`