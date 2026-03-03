package plugins

var manifestFile = "manifest.json"
var orgName = "summink"
var pluginPrefix = "sumi-plugin-"

type PluginPlatform struct {
	OS   string `json:"os"`
	Arch string `json:"arch"`
}

type PluginCommandData struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type PluginCommandInputs struct {
	PluginCommandData
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

type PluginManifest struct {
	Name        string                `json:"name"`
	Version     string                `json:"version"`
	Description string                `json:"description"`
	Author      string                `json:"author"`
	License     string                `json:"license"`
	Repo        string                `json:"repo"`
	Doc         string                `json:"doc"`
	Platforms   []PluginPlatform      `json:"platforms"`
	Tags        []string              `json:"tags"`
	Args        []PluginCommandInputs `json:"args"`
	Flags       []PluginCommandInputs `json:"flags"`
}

type PluginDetails struct {
	Name     string         `json:"name"`
	Manifest PluginManifest `json:"manifest"`
	Execute  string         `json:"execute"`
}

type PluginConfig struct {
	PluginDir string `json:"pluginDir"`
}
