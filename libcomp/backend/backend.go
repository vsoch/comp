package backend

// A backend is a cluster controller module that can be registered with Libpak
import (
	"fmt"
	"github.com/vsoch/comp/libcomp/env"
	"github.com/vsoch/comp/libcomp/uri"
	"log"
)

// Lookup of Backends. If we need more variables alongside, we can turn this into a
// struct that holds BackendInfo
var (
	Backends map[string]BackendInfo
)

// BackendInfo provides information about a backend
type BackendInfo interface {

	// Name of the backend is populated by a config
	GetName() string

	// Full Name of unique resource identifier
	GetURI(name string) *uri.URI

	// Description of this fs - defaults to Name
	GetDescription() string
	Env(image string) *env.Environment
	Shell(image string, remove bool) error
	Run(args ...string) error

	// Images available
	Images(args ...string)

	// List running containers
	List(args ...string)
	Info(args ...string) error
}

func List() map[string]BackendInfo {
	return Backends
}

func Add(backend BackendInfo) {
	if Backends == nil {
		Backends = make(map[string]BackendInfo)
	}
	Backends[backend.GetName()] = backend
}

// Get a backend by name
func Get(name string) (BackendInfo, error) {
	for backendName, entry := range Backends {
		if backendName == name {
			return entry, nil
		}
	}
	return nil, fmt.Errorf("did not find backend named %q", name)
}

// GetOrFail ensures we can find the entry
func GetOrFail(name string) BackendInfo {
	backend, err := Get(name)
	if err != nil {
		log.Fatalf("Failed to get backend: %v", err)
	}
	return backend
}
