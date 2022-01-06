package docker

import (
	"time"
)

type Manifests []Manifest
type Manifest struct {
	ID              string    `json:"Id"`
	RepoTags        []string  `json:"RepoTags"`
	RepoDigests     []string  `json:"RepoDigests"`
	Parent          string    `json:"Parent"`
	Comment         string    `json:"Comment"`
	Created         time.Time `json:"Created"`
	Container       string    `json:"Container"`
	ContainerConfig struct {
		Hostname     string      `json:"Hostname"`
		Domainname   string      `json:"Domainname"`
		User         string      `json:"User"`
		AttachStdin  bool        `json:"AttachStdin"`
		AttachStdout bool        `json:"AttachStdout"`
		AttachStderr bool        `json:"AttachStderr"`
		Tty          bool        `json:"Tty"`
		OpenStdin    bool        `json:"OpenStdin"`
		StdinOnce    bool        `json:"StdinOnce"`
		Env          []string    `json:"Env"`
		Cmd          []string    `json:"Cmd"`
		ArgsEscaped  bool        `json:"ArgsEscaped"`
		Image        string      `json:"Image"`
		Volumes      interface{} `json:"Volumes"`
		WorkingDir   string      `json:"WorkingDir"`
		Entrypoint   []string    `json:"Entrypoint"`
		OnBuild      interface{} `json:"OnBuild"`
		Labels       struct {
			Maintainer string `json:"Maintainer"`
		} `json:"Labels"`
	} `json:"ContainerConfig"`
	DockerVersion string `json:"DockerVersion"`
	Author        string `json:"Author"`
	Config        struct {
		Hostname     string      `json:"Hostname"`
		Domainname   string      `json:"Domainname"`
		User         string      `json:"User"`
		AttachStdin  bool        `json:"AttachStdin"`
		AttachStdout bool        `json:"AttachStdout"`
		AttachStderr bool        `json:"AttachStderr"`
		Tty          bool        `json:"Tty"`
		OpenStdin    bool        `json:"OpenStdin"`
		StdinOnce    bool        `json:"StdinOnce"`
		Env          []string    `json:"Env"`
		Cmd          []string    `json:"Cmd"`
		ArgsEscaped  bool        `json:"ArgsEscaped"`
		Image        string      `json:"Image"`
		Volumes      interface{} `json:"Volumes"`
		WorkingDir   string      `json:"WorkingDir"`
		Entrypoint   []string    `json:"Entrypoint"`
		OnBuild      interface{} `json:"OnBuild"`
		Labels       struct {
			Maintainer string `json:"Maintainer"`
		} `json:"Labels"`
	} `json:"Config"`
	Architecture string `json:"Architecture"`
	Os           string `json:"Os"`
	Size         int    `json:"Size"`
	VirtualSize  int    `json:"VirtualSize"`
	GraphDriver  struct {
		Data struct {
			LowerDir  string `json:"LowerDir"`
			MergedDir string `json:"MergedDir"`
			UpperDir  string `json:"UpperDir"`
			WorkDir   string `json:"WorkDir"`
		} `json:"Data"`
		Name string `json:"Name"`
	} `json:"GraphDriver"`
	RootFS struct {
		Type   string   `json:"Type"`
		Layers []string `json:"Layers"`
	} `json:"RootFS"`
	Metadata struct {
		LastTagTime time.Time `json:"LastTagTime"`
	} `json:"Metadata"`
}
