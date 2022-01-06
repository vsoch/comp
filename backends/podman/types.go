package podman

import (
	"time"
)

// Manifests is a list of container manifests
type Manifests []Manifest

// Config is an image config
type Config struct {
	Env        []string          `json:"Env"`
	Entrypoint []string          `json:"Entrypoint"`
	Cmd        []string          `json:"Cmd"`
	Labels     map[string]string `json:"Labels"`
}

// Manifest is a container manifest, output of inspect
type Manifest struct {
	ID           string        `json:"Id"`
	Digest       string        `json:"Digest"`
	RepoTags     []interface{} `json:"RepoTags"`
	RepoDigests  []string      `json:"RepoDigests"`
	Parent       string        `json:"Parent"`
	Comment      string        `json:"Comment"`
	Created      time.Time     `json:"Created"`
	Config       Config        `json:"Config"`
	Version      string        `json:"Version"`
	Author       string        `json:"Author"`
	Architecture string        `json:"Architecture"`
	Os           string        `json:"Os"`
	Size         int           `json:"Size"`
	VirtualSize  int           `json:"VirtualSize"`
	GraphDriver  struct {
		Name string `json:"Name"`
		Data struct {
			LowerDir string `json:"LowerDir"`
			UpperDir string `json:"UpperDir"`
			WorkDir  string `json:"WorkDir"`
		} `json:"Data"`
	} `json:"GraphDriver"`
	RootFS struct {
		Type   string   `json:"Type"`
		Layers []string `json:"Layers"`
	} `json:"RootFS"`
	Labels      map[string]string `json:"Labels"`
	Annotations struct {
	} `json:"Annotations"`
	ManifestType string `json:"ManifestType"`
	User         string `json:"User"`
	History      []struct {
		Created    time.Time `json:"created"`
		CreatedBy  string    `json:"created_by"`
		EmptyLayer bool      `json:"empty_layer,omitempty"`
	} `json:"History"`
	NamesHistory []string `json:"NamesHistory"`
}
