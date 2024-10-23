
<p align="center">
  <img src="https://github.com/jalpp/PassDIY/blob/dev/style/icon.png?raw=true" alt="Image Description" width=20% height=20% />
</p>

![MIT License](https://img.shields.io/badge/License-MIT-green.svg)

# PassDIY

A personal password/token manager TUI for developers to generate various types of hash/salted secrets and store them in different cloud based vaults

## Why PassDIY?

Because managing tokens, pins used in various dummy/dev apps require them to be generated first, and store them somewhere, I personally used 3 sites to generate random API dummy tokens, store them in other site. It become a big mess, and I thought there has to be a simple way where I can generate **pass**words and **D**o **I**t m**Y**self hence PassDIY.

## Features

- Generation of strong secrets like pins, passwords, API tokens, passphrases 
- Generate X multiple secrets at once and pick X and X password generation algorithms
- Hash tokens/passwords with Argon2Id and Bcrypt
- Salt tokens/passwords
- Copy passwords to clipboard 
- Automatically config password/token lengths and other settings
- Hashicorp Vault integration to connect to secure vault and store generated secrets on cloud
- 1Password integration to connect to secure vault and store generated secrets on cloud

## Hashicorp Vault Commands
- hcpvaultconnect automatically connect to hcp vault via service principle
- hcpvaultstore store secrets into the vault via name=value format
- hcpvaultlist list log details about token created at, created by details


## 1Password Commands
- 1passstore store secrets into the vault via name|password|url format
- 1passwordlist list secret names for connected vault


## Demo

![passdiydemo4](https://github.com/user-attachments/assets/b26dabb9-96ad-4a68-a742-dfeb4fc604d7)


## Hashicorp Setup

To allow PassDIY to store and connect to your Hashicorp vault you must create a [service principle](https://developer.hashicorp.com/hcp/docs/hcp/iam/service-principal) with ```Vault Secrets App Manager``` permission. Also would need set below envs

`export HCP_CLIENT_ID=<your-hcp-client-id>`
`export HCP_CLIENT_SECRET=<your-hcp-client-secret>`

more detailed in `./Setup.md`

## 1Password Setup

To allow PassDIY to connect to your 1Password Vault you would need to set [service principle](https://developer.1password.com/docs/sdks) anf the service account token

`export OP_SERVICE_ACCOUNT_TOKEN=<your-service-account-token>`

## Config

you can config PassDIY's password/token/pin char lengths additional confiurations by running config command



## Installation

If you have `make` installed, follow these steps to build, run, and install **passdiy**:

1. **Build the project**:
   ```
   make build
   ```

2. **Run the application**:
   ```
   make run
   ```

3. **Install globally** (optional):
   ```
   sudo make install
   ```

You can then run it from anywhere with: `passdiy`

If you do not have `make` you can build and run it traditionally with:

```
go run .
```

## Uninstall

You can uninstall passdiy with:

```
sudo make uninstall
```
## Roadmap

- dynamically change config
- add more vaults possibly vercel/Azure key vault
- add more hashing algos


## Acknowledgements

 - [BubbleTea](https://github.com/charmbracelet/bubbletea)
 - [Lipgloss](github.com/charmbracelet/lipgloss)
 - [Argon2id](https://github.com/alexedwards/argon2id)
 - [Bcrypt](https://golang.org/x/crypto/bcrypt)
 - [Hashicorp Vault](https://developer.hashicorp.com/hcp/api-docs/vault-secrets#overview)
 - [English](github.com/gregoryv/english)
 - [Clipboard](https://github.com/atotto/clipboard)
 - [1Password SDK](https://github.com/1Password/onepassword-sdk-go)

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Discord

[Invite](https://discord.gg/FU6DMKZuZY)

## Authors

- [@jalpp](https://www.github.com/jalpp)

## Contributor

- [@cursedcadaver](https://github.com/cursedcadaver)

