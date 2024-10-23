package hcpvault

import "strings"

func ConnectUI() string {
	return Connect()
}

func ListUI() string {
	var list string = List()
	if strings.Contains(list, "Unauthorized") {
		return "Please connect to Hashicorp vault via hcpvaultconnect"
	}
	return list
}

func StoreUI(userInput string) string {
	parts := strings.SplitN(userInput, "=", 2)
	if len(parts) == 2 {
		name := parts[0]
		value := parts[1]
		return Create(name, value)
	}
	return "Invalid format. Use 'name=value'."
}
