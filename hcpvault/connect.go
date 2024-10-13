package hcpvault

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func getHCPAPIToken(clientID, clientSecret string) (string, error) {
	url := "https://auth.idp.hashicorp.com/oauth2/token"

	data := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"grant_type":    "client_credentials",
		"audience":      "https://api.hashicorp.cloud",
	}

	formData := []byte(fmt.Sprintf(
		"client_id=%s&client_secret=%s&grant_type=%s&audience=%s",
		data["client_id"],
		data["client_secret"],
		data["grant_type"],
		data["audience"],
	))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(formData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request failed with status %d: %s", resp.StatusCode, body)
	}

	var tokenResponse TokenResponse
	if err := json.Unmarshal(body, &tokenResponse); err != nil {
		return "", fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return tokenResponse.AccessToken, nil
}

func Connect() string {

	clientID := os.Getenv("HCP_CLIENT_ID")
	clientSecret := os.Getenv("HCP_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {

		return "Please set the following ENV variable HCP_CLIENT_ID and HCP_CLIENT_SECRET values"
	}

	token, err := getHCPAPIToken(clientID, clientSecret)

	if err != nil {

		return "Please check your ENV variables! Error retrieving HCP API token"
	}

	os.Setenv("HCP_API_TOKEN", token)

	return "Successfully Connected To Target Hashicorp Vault"

}
