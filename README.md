
# PassDIY

A personal password/token manager TUI for developers to generate various types of hash/salted secrets and store them in their cloud-based Hashicorp Vault

## Why PassDIY?

Because managing tokens, pins used in various dummy/dev apps require them to be generated first, and store them somewhere, I personally used 3 sites to generate random API dummy tokens, store them in other site. It become a big mess, and I thought there has to be a simple way where I can generate **pass**words and **D**o **I**t m**Y**self hence PassDIY.

## Features

- Generation of strong secrets like pins, passwords, API tokens, passphrases 
- Generate X multiple secrets at once and pick X and X password generation algorithms
- Hash tokens/passwords with Argon2id 
- Salt tokens/passwords
- Copy passwords to clipboard 
- Hashicorp Vault integration to connect to secure vault and store generated secrets on cloud

## Hashicorp Vault Commands
- hcpvaultconnect automatically connect to hcp vault via service principle
- hcpvaultstore store secrets into the vault via name=value format
- hcpvaultlist list log details about token created at, created by details

## Demo

![passdiydemo3](https://github.com/user-attachments/assets/b17a0164-9dcd-491f-9a6d-29c48f6d8e82)


## Setup

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

## Config

you can config PassDIY's password/token/pin char lengths additional confiurations in `config/config.go` by changing below values

```
const (
	PIN_DIGIT_LENGTH      int = 6   // number of ints in pin digit
	API_TOKEN_CHAR_LENGTH int = 60  // number of chars in a API token
	PASWORD_CHAR_LENGTH   int = 40  // number of chars in a password
	PASSPHRASE_COUNT_NUM  int = 5   // number of words in passphrase
	MULTIPLE_VALUE_COUNT  int = 5   // how many password/tokens you want to output
	LOTTERY_WHEEL_COUNT   int = 100 // how many times you want to generate token/password/pins to randomly pick one (pass100, pass10000)
	SALT_EXTRA_LENGTH     int = 10  // how many extra chars you want to add to a password/token
)

```
## Roadmap

- dynamically change config
- add more vaults possibly vercel/Azure key vault
- add more hashing algos


## Acknowledgements

 - [BubbleTea](https://github.com/charmbracelet/bubbletea)
 - [Lipgloss](github.com/charmbracelet/lipgloss)
 - [Argon2id](https://github.com/alexedwards/argon2id)
 - [Hashicorp Vault](https://developer.hashicorp.com/hcp/api-docs/vault-secrets#overview)
 - [English](github.com/gregoryv/english)
 - [Clipboard](https://github.com/atotto/clipboard)

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Authors

- [@jalpp](https://www.github.com/jalpp)

