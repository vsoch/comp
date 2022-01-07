package docker

// The podman backend

import (
	"fmt"
	"strings"

	"github.com/vsoch/comp/lib/command"
	"github.com/vsoch/comp/lib/errors"
	"github.com/vsoch/comp/lib/fs"
	"github.com/vsoch/comp/lib/logger"
	"github.com/vsoch/comp/lib/options"
	"github.com/vsoch/comp/lib/str"
	"github.com/vsoch/comp/libcomp/backend"

	"github.com/docker/docker/client"

	// Shared container backend / functions
	"github.com/vsoch/comp/backends/container"
)

var (
	// info prints messages in colors
	info = logger.Logger{}
)

// The Docker backend is a controller to Docker containers
type Docker struct {
	container.Container

	// Client is the Docker client
	Client *client.Client
}

// Add the backend to be known to libpak
func init() {

	// Create a Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv)
	errors.Check(err)

	docker := Docker{
		Container: container.Container{
			Name:        "docker",
			Description: "Docker is a popular, widely used container technology",
			Options:     options.Options{},
			Executable:  fs.Which("docker"),

			// TODO shell wanted should come from config, not hard coded! cc @alecbcs
			ShellExecutable: "/bin/sh",
			GetImages:       GetImages,
			Ps:              Ps,
		},
		Client: cli,
	}

	// Add specific commands
	docker.ImagesCommand = []string{docker.Executable, "images"}

	// This is backend info for Docker
	backend.Add(docker)
}

// Getimages is the Docker version of GetImages placed as a func field on the container struct
func GetImages(c container.Container, args ...string) container.ContainerImages {

	d := Docker{Container: c}
	output := d.RunCommand(d.ImagesCommand, args...)
	images := container.ContainerImages{}

	// Parsed table lines
	for _, line := range d.parseTableRows(output.Out) {
		parts := strings.Split(line, " ")

		// Should have length 7
		// [test latest 5000d9ff2e9a 2 days ago 1.24MB]
		if len(parts) != 7 {
			continue
		}

		// Combine container and tag into name
		name := fmt.Sprintf("%s:%s", strings.Trim(parts[0], " "), strings.Trim(parts[1], " "))
		imageId := strings.Trim(parts[2], " ")
		size := strings.Trim(parts[6], " ")

		image := container.ContainerImage{Names: []string{name}, ID: imageId, StringSize: size}
		images = append(images, image)
	}
	return images
}

// Ps runs docker ps to return a listing of running containers
func Ps(c container.Container, args ...string) container.Containers {

	d := Docker{Container: c}
	d.Check()
	cmd := append([]string{d.Executable, "ps"}, args...)
	err, output := command.RunCommand(cmd, nil)
	errors.Check(err)

	// Prepare list of containers to return
	images := container.Containers{}

	// Parse table
	for _, line := range d.parseTableRows(output.Out) {
		parts := strings.Split(line, " ")

		status := strings.Join(parts[3:5], " ")
		// [c14bb0ce158c redis:latest "docker-entrypoint.s…" 9 days ago Up 47 hours 6379/tcp sregistry_redis_1]
		image := container.ContainerPs{Image: parts[1], Status: status}
		images = append(images, image)
	}
	return images
}

// Parse table will parse some number of entries in list of lines, removing extra space
func (d Docker) parseTableRows(output string) []string {
	rows := []string{}
	for _, line := range strings.Split(output, "\n")[1:] {
		line = str.CleanSpace(line)
		rows = append(rows, line)
	}
	return rows
}

// TODO write me vsoch
func (d Docker) Info(args ...string) error {
	d.Check()
	// TODO NOT written yet
	return nil
}

// TODO write me vsoch
// Inspect is a shared function for Info and Get Labels
func (d Docker) Inspect(args ...string) command.Output {
	cmd := append([]string{d.Executable, "inspect"}, args...)
	err, output := command.RunCommand(cmd, nil)
	errors.Check(err)
	return output
}

// TODO write me vsoch
// GetLabels from a container image inspect
func (d Docker) GetLabels(args ...string) map[string]string {
	//	output := d.Inspect(args...)
	labels := make(map[string]string)
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
