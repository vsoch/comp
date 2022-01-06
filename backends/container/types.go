package container

import (
	"time"
)

// Container images available via podman images, we parse docker in too
type ContainerImages []ContainerImage
type ContainerImage struct {
	ID          string      `json:"Id"`
	ParentID    string      `json:"ParentId"`
	RepoTags    interface{} `json:"RepoTags"`
	RepoDigests []string    `json:"RepoDigests"`
	Size        int         `json:"Size"`
	StringSize  string
	SharedSize  int               `json:"SharedSize"`
	VirtualSize int               `json:"VirtualSize"`
	Labels      map[string]string `json:"Labels"`
	Containers  int               `json:"Containers"`
	Names       []string          `json:"Names"`
	Digest      string            `json:"Digest"`
	History     []string          `json:"History"`
	Created     int               `json:"Created"`
	CreatedAt   time.Time         `json:"CreatedAt"`
}

// Containers is output json (list) of container objects
type Containers []ContainerPs

// Container represnts an output row of a ps
type ContainerPs struct {
	AutoRemove bool              `json:"AutoRemove"`
	Command    interface{}       `json:"Command"`
	CreatedAt  string            `json:"CreatedAt"`
	Exited     bool              `json:"Exited"`
	ExitedAt   int64             `json:"ExitedAt"`
	ExitCode   int               `json:"ExitCode"`
	ID         string            `json:"Id"`
	Image      string            `json:"Image"`
	ImageID    string            `json:"ImageID"`
	IsInfra    bool              `json:"IsInfra"`
	Labels     map[string]string `json:"Labels"`
	Mounts     []interface{}     `json:"Mounts"`
	Names      []string          `json:"Names"`
	Namespaces Namespaces        `json:"Namespaces"`
	Networks   interface{}       `json:"Networks"`
	Pid        int               `json:"Pid"`
	Pod        string            `json:"Pod"`
	PodName    string            `json:"PodName"`
	Ports      interface{}       `json:"Ports"`
	Size       string            `json:"Size"`
	StartedAt  int               `json:"StartedAt"`
	State      string            `json:"State"`
	Status     string            `json:"Status"`
	Created    int               `json:"Created"`
}

type Namespaces struct{}
