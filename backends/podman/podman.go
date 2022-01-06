package podman

// The podman backend

import (
	"encoding/json"
	"fmt"

	"github.com/vsoch/compenv/config"
	"github.com/vsoch/compenv/lib/command"
	"github.com/vsoch/compenv/lib/errors"
	"github.com/vsoch/compenv/lib/fs"
	"github.com/vsoch/compenv/lib/logger"
	"github.com/vsoch/compenv/lib/options"
	"github.com/vsoch/compenv/libcompenv/backend"

	// Shared container backend / functions
	"github.com/vsoch/compenv/backends/container"
)

var (
	// info prints messages in colors
	info = logger.Logger{}
)

// The Podman backend is a controller to Podman containers
type Podman struct {
	container.Container
}

// Add the backend to be known to libpak
func init() {

	pm := Podman{
		Container: container.Container{
			Name:            "podman",
			Description:     "Podman is a rootless container technology",
			Options:         options.Options{},
			Executable:      fs.Which("podman"),
			ShellExecutable: config.Shell,
			GetImages:       GetImages,
			Ps:              Ps,
		},
	}

	// Add specific commands
	pm.ImagesCommand = []string{pm.Executable, "images", "--format", "json"}

	// This is backend info for Podman
	backend.Add(pm)
}

// GetImages returns Podman container images
func GetImages(c container.Container, args ...string) container.ContainerImages {

	p := Podman{Container: c}
	fmt.Println(p.ImagesCommand)
	output := p.RunCommand(p.ImagesCommand, args...)
	var images container.ContainerImages
	err := json.Unmarshal([]byte(output.Out), &images)
	errors.Check(err)

	// TODO have this return common shared containers format
	return images
}

// Ps runs podman ps to return a listing of running containers
func Ps(c container.Container, args ...string) container.Containers {

	p := Podman{Container: c}
	p.Check()
	cmd := append([]string{p.Executable, "ps", "--format", "json"}, args...)
	err, output := command.RunCommand(cmd, nil)
	errors.Check(err)

	var containers container.Containers
	err = json.Unmarshal([]byte(output.Out), &containers)
	errors.Check(err)
	return containers
}

// Info Inspects a particular podman image
// TODO eventually will want to customize output format to be generic
// TODO redo these functions to have better organization and return types
// additional args can be --type container|image
func (p Podman) Info(args ...string) error {
	p.Check()
	output := p.Inspect(args...)
	fmt.Println(output.Out)
	return nil
}

// Inspect is a shared function for Info and Get Labels
func (p Podman) Inspect(args ...string) command.Output {
	cmd := append([]string{p.Executable, "inspect", "--format", "json"}, args...)
	err, output := command.RunCommand(cmd, nil)
	errors.Check(err)
	return output
}

// GetLabels from a container image inspect
func (p Podman) GetLabels(args ...string) map[string]string {
	output := p.Inspect(args...)
	var manifests Manifests
	err := json.Unmarshal([]byte(output.Out), &manifests)
	errors.Check(err)
	labels := make(map[string]string)
	for _, manifest := range manifests {
		for key, label := range manifest.Labels {
			labels[key] = label
		}
	}
	return labels
}

// TODO write commands for rm and rmi
// args := []string{"rm"}
// args = append(args, "--force", image)
// exit code 0 and err != nil == unexpected error
// exit code 1: container does not exist
// exit code 2: container is running
// default == failed to remove
// return err

// args := []string{"rmi"}
// args = append(args, "--force", image)
// exit code 0 and err != nil == unexpected error
// exit code 1: image does not exist
// exit code 2: image has dependent children
// default == failed to remove
// return err
