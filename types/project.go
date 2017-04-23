package types

// Project bladi bladi
type Project struct {
	GitURL      string `json:"git-url,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ReadMeURL   string `json:"readme-url,omitempty"`
	HTMLUrl     string `json:"html-url,omitempty"`
}
