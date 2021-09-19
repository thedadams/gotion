package notion

import "net/http"

const version = "v0.0.0"

// Versions holds a list of all valid versions of the Notion API.
var Versions = []string{
	"2021-05-13",
	"2021-08-16",
}

// Settings represents the settings needed to use the Notion API.
type Settings struct {
	APIKey, Version, UserAgent string
}

// LatestWithAPIKey returns a Settings struct using the provided API key and the latest version of the Notion API.
func LatestWithAPIKey(apiKey string) *Settings {
	return &Settings{APIKey: apiKey, Version: Versions[len(Versions)-1], UserAgent: "gotion/" + version}
}

// ToHeaders attaches the appropriate header information to the request.
func (s *Settings) ToHeaders(req *http.Request) {
	if s.Version == "" {
		s.Version = LatestWithAPIKey("").Version
	}
	req.Header.Set("Authorization", "Bearer "+s.APIKey)
	req.Header.Set("Notion-Version", s.Version)
	if s.UserAgent != "" {
		req.Header.Set("User-Agent", s.UserAgent)
	}

	if req.Method != http.MethodGet {
		req.Header.Set("Content-Type", "application/json")
	}
}
