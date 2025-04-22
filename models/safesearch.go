package models

// SafeSearchConfig - Safe search settings.
type SafeSearchConfig struct {
	Enabled    bool `json:"enabled"`
	Bing       bool `json:"bing"`
	Duckduckgo bool `json:"duckduckgo"`
	Ecosia     bool `json:"ecosia"`
	Google     bool `json:"google"`
	Pixabay    bool `json:"pixabay"`
	Yandex     bool `json:"yandex"`
	Youtube    bool `json:"youtube"`
}
