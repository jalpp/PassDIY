package hcpvault

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Location struct {
	Provider string `json:"provider"`
	Region   string `json:"region"`
}

type SecretBody struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func createAppKVSecret(orgID, projectID, appName, name, value, token string) error {

	url := fmt.Sprintf("https://api.cloud.hashicorp.com/secrets/2023-06-13/organizations/%s/projects/%s/apps/%s/kv", orgID, projectID, appName)

	body := SecretBody{
		Name:  name,
		Value: value,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(responseBody))
	}

	return nil
}

func Create(name, value string) string {
	orgID := os.Getenv("HCP_ORG_ID")
	projectID := os.Getenv("HCP_PROJECT_ID")
	appName := os.Getenv("HCP_APP_NAME")
	token := os.Getenv("HCP_API_TOKEN")

	if orgID == "" || projectID == "" || appName == "" || token == "" {
		return "Please set HCP_ORG_ID HCP_PROJECT_ID HCP_APP_NAME HCP_API_TOKEN"
	}

	if err := createAppKVSecret(orgID, projectID, appName, name, value, token); err != nil {
		return err.Error()
	}

	return "Successfully Created Secret In The Configured Hashicorp Vault"

}
