package onepassword

import (
	"context"
	"os"

	"github.com/1password/onepassword-sdk-go"
)

func Connect() *onepassword.Client {

	token := os.Getenv("OP_SERVICE_ACCOUNT_TOKEN")

	client, err := onepassword.NewClient(context.Background(),
		onepassword.WithServiceAccountToken(token),

		onepassword.WithIntegrationInfo("passdiy", "v1.0.0"),
	)
	if err != nil {
		return nil
	}

	return client
}
