package onepassword

import (
	"context"
	"errors"
	"fmt"

	"github.com/1password/onepassword-sdk-go"
)

func GetVaultId(client *onepassword.Client) string {
	vaults, err := client.Vaults.ListAll(context.Background())
	var valId string
	if err != nil {
		panic(err)
	}
	for {
		vault, err := vaults.Next()
		if errors.Is(err, onepassword.ErrorIteratorDone) {
			break
		} else if err != nil {
			panic(err)
		}

		valId = vault.ID
	}
	return valId
}

func createItem(client *onepassword.Client, vaultID string, usernameVal string, passVal string, URL string) string {
	sectionID := "passDetails"
	itemParams := onepassword.ItemCreateParams{
		Title:    fmt.Sprintf("PassDIY Gen: %s", URL),
		Category: onepassword.ItemCategoryLogin,
		VaultID:  vaultID,
		Fields: []onepassword.ItemField{
			{
				ID:        "username",
				Title:     "username",
				Value:     usernameVal,
				FieldType: onepassword.ItemFieldTypeText,
			},
			{
				ID:        "password",
				Title:     "password",
				Value:     passVal,
				FieldType: onepassword.ItemFieldTypeConcealed,
			},
			{
				ID:        "onetimepassword",
				Title:     "one-time password",
				Value:     "otpauth://totp/my-example-otp?secret=jncrjgbdjnrncbjsr&issuer=1Password",
				SectionID: &sectionID,
				FieldType: onepassword.ItemFieldTypeTOTP,
			},
		},
		Sections: []onepassword.ItemSection{
			{
				ID:    sectionID,
				Title: "Extra Details",
			},
		},
		Tags: []string{"passdiy generated", fmt.Sprintf("web: %s", URL)},
		Websites: []onepassword.Website{
			{
				URL:              URL,
				AutofillBehavior: onepassword.AutofillBehaviorAnywhereOnWebsite,
				Label:            URL,
			},
		},
	}

	createdItem, err := client.Items.Create(context.Background(), itemParams)
	if err != nil {
		return fmt.Sprintf("Error %s", err.Error())
	}

	return fmt.Sprintf("Successfully created password item ID %s in 1Password Vault", createdItem.Title)
}

func Create(user string, pass string, URL string) string {

	var client = Connect()

	if client == nil {
		return "Error! Please check 1password config"
	}

	var valId = GetVaultId(client)

	return createItem(client, valId, user, pass, URL)

}
