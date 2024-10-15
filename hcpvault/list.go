package hcpvault

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type SecretsResponse struct {
	Secrets []Secret `json:"secrets"`
}

type Secret struct {
	Name          string     `json:"name"`
	Version       Version    `json:"version"`
	CreatedAt     string     `json:"created_at"`
	CreatedBy     CreatedBy  `json:"created_by"`
	LatestVersion string     `json:"latest_version"`
	SyncStatus    SyncStatus `json:"sync_status"`
}

type Version struct {
	Version string `json:"version"`
	Type    string `json:"type"`
}

type CreatedBy struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Email string `json:"email"`
}

type SyncStatus struct {
	VersionCount string `json:"version_count"`
}

func ListAppSec(orgID, projID, appName, token string) string {

	url := fmt.Sprintf("https://api.cloud.hashicorp.com/secrets/2023-06-13/organizations/%s/projects/%s/apps/%s/secrets", orgID, projID, appName)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "Error reading response body:"
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "Error reading response body"
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("Request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Error reading response body"
	}

	var secrets SecretsResponse
	var buffer string = "Hashicorp Secret Log Info \n"
	if err := json.Unmarshal(body, &secrets); err != nil {
		return fmt.Sprintf("Error parsing JSON: %v", err)
	}

	for _, secret := range secrets.Secrets {

		buffer += fmt.Sprintf("Name: %s Version: %s Type: %s Created_at %s Created_by %s\n", secret.Name, secret.Version.Version, secret.Version.Type, secret.CreatedAt, secret.CreatedBy.Type)
	}

	return buffer
}

func List() string {
	orgID := os.Getenv("HCP_ORG_ID")
	projectID := os.Getenv("HCP_PROJECT_ID")
	appName := os.Getenv("HCP_APP_NAME")
	token := os.Getenv("HCP_API_TOKEN")

	if orgID == "" || projectID == "" || appName == "" || token == "" {
		return "Please set the following ENV variables HCP_ORG_ID HCP_PROJECT_ID HCP_APP_NAME HCP_API_TOKEN values"
	}

	return ListAppSec(orgID, projectID, appName, token)

}
