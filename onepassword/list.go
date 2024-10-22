package onepassword

import (
	"context"
	"errors"
	"fmt"

	"github.com/1password/onepassword-sdk-go"
)

func listItems(client *onepassword.Client, vaultID string) string {

	var buffer string = "1Password Vault Item Info \n"
	items, err := client.Items.ListAll(context.Background(), vaultID)
	if err != nil {
		panic(err)
	}
	for {
		item, err := items.Next()
		if errors.Is(err, onepassword.ErrorIteratorDone) {
			break
		} else if err != nil {
			panic(err)
		}
		buffer += fmt.Sprintf("%s %s\n", "Item Title:", item.Title)
	}

	return buffer
}

func List() string {

	var client = Connect()

	if client == nil {
		return "Error! Please check 1password config"
	}

	return listItems(client, GetVaultId(client))

}
