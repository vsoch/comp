package container

// This is an abstract backend that is shared by Podman and Docker

import (
	"fmt"
	"log"
	"strings"

	"github.com/vsoch/comp/lib/command"
	"github.com/vsoch/comp/lib/errors"
	"github.com/vsoch/comp/lib/logger"
	"github.com/vsoch/comp/lib/options"
	"github.com/vsoch/comp/libcomp/env"
	"github.com/vsoch/comp/libcomp/uri"
)

var (
	// info prints messages in colors
	info = logger.Logger{}
)

// The Container Backend provides basic shared functions for containers
type Container struct {
	Name        string
	Description string

	// Options specific to podman
	Options options.Options

	// Full path to podman executable
	Executable string

	// Default container shell
	ShellExecutable string

	// Container technology specific commands
	ImagesCommand  []string
	InspectCommand []string
	ExistsCommand  []string

	// Getimages and Ps functions MUST be defined
	GetImages func(c Container, args ...string) ContainerImages
	Ps        func(c Container, args ...string) Containers
}

// Redundant functions to retrieve metadata
func (c Container) GetName() string {
	return c.Name
}

// GetDescriptions returns the container description
func (c Container) GetDescription() string {
	return c.Description
}

// Get Options returns container options
func (c Container) GetOptions() options.Options {
	return c.Options
}

// submit a Docker container (runs detached)
func (c Container) Submit(args ...string) error {

	// Run the container, not detached, and don't remove
	return c.RunContainer(args[0], true, false)
}

// Get ensures that we have an image pulled (or does nothing if it exists)
func (c Container) Get(image string) {
	// pull the image if it doesn't exist
	if !c.Exists(image) {
		info.Blue("Pulling image " + image)
		c.Pull(image)
	}
}

// Images shows all images available
func (c Container) Images(args ...string) {
	c.ImagesTable(c.GetImages(c, args...))
}

// GetURI returns the full unique resource identifier for a container image
func (c Container) GetURI(image string) *uri.URI {

	// If we use podman container exists, it doesn't seem to work
	images := c.GetImages(c)
	for _, containerImage := range images {
		for _, name := range containerImage.Names {
			containeruri := uri.New(name)
			containeruri.Digest = containerImage.Digest
			if containeruri.Matches(name) {
				return containeruri
			}
		}
	}
	return nil
}

// Exists determines if a container exists, returns true or false
func (c Container) Exists(image string) bool {

	// If we use podman container exists, it doesn't seem to work
	images := c.GetImages(c)
	for _, containerImage := range images {
		for _, name := range containerImage.Names {
			containeruri := uri.New(name)
			if containeruri.Matches(name) {
				return true
			}
		}
	}
	return false
}

// run a Docker container (does not remove)
func (c Container) Run(args ...string) error {

	// Run the container, not detached, and remove
	return c.RunContainer(args[0], false, true)
}

// Shell runs a container image, either detached or removable (not a shell)
func (c Container) Shell(image string, remove bool) error {

	// check that podman is installed and pull the image if necessary
	c.Check()
	c.Get(image)

	// Prepare the command and show to the user
	cmd := []string{c.Executable, "run", "-it", "--entrypoint", c.ShellExecutable, image}
	info.Cyan(strings.Join(cmd, " "))
	err := command.RunInteractive(cmd, nil)
	errors.Check(err)
	return err
}

// Pull an image by name
func (c Container) Pull(image string) error {
	err, output := command.RunCommand([]string{c.Executable, "pull", image}, nil)
	if err != nil || output.ExitCode != 0 {
		log.Fatalf("running %s pull %s", c.GetName(), image)
	}
	return err
}

// Check to see if container technology is installed
func (c Container) Check() {
	if c.Executable == "" {
		log.Fatalf("Cannot use podman backend without podman installed.")
	}
}

// runContainer runs a container image, either detached or removable (not a shell)
func (c Container) RunContainer(image string, detached bool, remove bool) error {
	c.Check()
	cmd := []string{c.Executable, "run"}

	// Run in detached mode?
	if detached {
		cmd = append(cmd, "-d")
	}

	// Remove after done?
	if remove {
		cmd = append(cmd, "--rm")
	}

	// Finish up the command
	cmd = append(cmd, image)
	info.Cyan(strings.Join(cmd, " "))
	output := c.RunCommand(cmd)
	fmt.Println(output.Out)
	return nil
}

// RunCommand is a general function to run a command for a container
func (c Container) RunCommand(cmd []string, args ...string) command.Output {
	c.Check()
	cmd = append(cmd, args...)
	err, output := command.RunCommand(cmd, nil)
	errors.Check(err)
	return output
}

// Env returns the output
func (c Container) Env(image string) *env.Environment {
	c.Check()
	cmd := []string{c.Executable, "run", "-it", "--entrypoint", "env", image}
	info.Cyan(strings.Join(cmd, " "))
	err, output := command.RunCommand(cmd, nil)
	errors.Check(err)
	environ := env.Parse(output.Out)
	return environ
}

// PsTable is a shared function for printing images from a ps command
func (c Container) PsTable(images Containers) {

	// Print into table
	table := info.Table()
	table.SetHeader([]string{"IMAGE ID", "STATUS"})
	for _, image := range images {
		table.AddRow([]string{image.Image, image.Status})
	}
	table.Print()
}

// List lists *running* podman images (podman ps --format json)
// Additional args can include -a --filter <pattern>
func (c Container) List(args ...string) {
	c.PsTable(c.Ps(c, args...))
}

// ImagesTable is a shared function for printing an images table from ContainerImages
func (c Container) ImagesTable(images ContainerImages) {

	// Print images into table
	table := info.Table()
	table.SetHeader([]string{"REPOSITORY", "TAG", "IMAGE ID", "SIZE"})

	// Keep track of those we have seen
	seen := make(map[string]bool)
	for _, image := range images {
		for _, name := range image.Names {
			// If it's a sha reference, skip, we want to show tags only
			if strings.Contains(name, "@") {
				continue
			}
			// Split into name and tag
			containeruri := uri.New(name)

			// Do we have an int or string size?
			size := image.StringSize
			if size == "" {
				size = fmt.Sprintf("%d", image.Size)
			}
			// Only print those we haven't seen
			if _, ok := seen[containeruri.WithTag()]; !ok {
				table.AddRow([]string{containeruri.Image, containeruri.Tag, image.ID, size})
				seen[containeruri.WithTag()] = true
			}
		}
	}
	table.Print()
}
