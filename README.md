
# PassDIY

A personal password/token manager TUI for developers to generate various types of hash/salted secrets and store them in their cloud-based Hashicorp Vault

## Why PassDIY?

Because managing tokens, pins used in various dummy/dev apps require them to be generated first, and store them somewhere, I personally used 3 sites to generate random API dummy tokens, store them in other site. It become a big mess, and I thought there has to be a simple way where I can generate **pass**words and **D**o **I**t m**Y**self hence PassDIY.

## Features

- Generation of strong secrets like pins, passwords, API tokens, passphrases 
- Ability to generate 5 multiple secrets at once and pick 100 and 10000 password generation algorithms
- Ability to hash tokens/passwords with Argon2id 
- Ability to salt tokens/passwords
- Ability to copy passwords to clipboard 
- Hashicorp Vault integration to connect to secure vault and store generated secrets on cloud
- Simple to use TUI so you don't forget sub commands and what password you were trying to store

## Demo

![passdiydemo](https://github.com/user-attachments/assets/79e1d0ce-614f-45dd-a5d9-6143ffac259f)

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

can also build PassDIY with build command

`go build`


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

