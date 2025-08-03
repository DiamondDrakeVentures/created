package modrinth

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
)

const endpoint = "/project"

type Project struct {
	ID          string  `json:"id"`
	Slug        string  `json:"slug"`
	Title       string  `json:"title"`
	ProjectType string  `json:"project_type"`
	Description string  `json:"description"`
	License     License `json:"license"`

	Categories           []string `json:"categories"`
	AdditionalCategories []string `json:"additional_categories"`
	ClientSide           string   `json:"client_side"`
	ServerSide           string   `json:"server_side"`

	Body   string `json:"body"`
	Status string `json:"status"`

	IssuesURL    string `json:"issues_url"`
	SourceURL    string `json:"source_url"`
	WikiURL      string `json:"wiki_url"`
	DiscordURL   string `json:"discord_url"`
	DonationURLs []any  `json:"donation_urls"`
}

type DonationURL struct {
	ID       string `json:"id"`
	Platform string `json:"platform"`
	URL      string `json:"url"`
}

type License struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

func (api *API) Get(idOrSlug string) (Project, error) {
	url := api.v2() + endpoint + "/" + idOrSlug

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Project{}, err
	}

	resp, err := api.client.Do(req)
	if err != nil {
		return Project{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		contentType, _, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
		if err != nil {
			return Project{}, err
		}

		if contentType != "application/json" {
			return Project{}, fmt.Errorf("expected Content-Type to be %q, got %q instead",
				"application/json", resp.Header.Get("Content-Type"))
		}

		var project Project
		err = json.NewDecoder(resp.Body).Decode(&project)
		return project, err
	}

	return Project{}, fmt.Errorf("HTTP %d Error", resp.StatusCode)
}
