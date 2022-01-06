package uri

import (
	"github.com/vsoch/compenv/lib/str"
	"os"
	"path/filepath"
	"strings"
)

// Parse a string into a container unique resource identifier
type URI struct {
	Raw   string
	Image string

	// A URI can be parsed from a path
	Path string

	// A tag is separated with :
	Tag string

	// A digest is determed with @
	// Technically a digest can be a sha256 or tag, but this works
	Digest string
}

// New parses a URI from a standard string
func New(name string) *URI {
	uri := URI{Raw: name}
	uri.Parse()
	return &uri
}

// NewFromPath parses a URI from a filesystem path
func NewFromPath(path string) *URI {

	uri := URI{}
	uri.Path = path
	parts := strings.Split(path, string(os.PathSeparator))

	// Parse backwards - assume last could be digest
	parts = str.ReverseStringList(parts)

	// Do we have a digest?
	if strings.HasPrefix(parts[0], "sha256") {
		uri.Digest = parts[0]
		parts = parts[1:]
	}
	// Only add the tag if it isn't <none> (docker	uri.Tag := parts[0]
	uri.Tag = parts[0]
	parts = parts[1:]
	uri.Image = strings.Join(str.ReverseStringList(parts), "/")
	return &uri
}

// URI returns the most appropriate resource identifier
func (c *URI) URI() string {
	if c.Tag != "" && !strings.Contains(c.Tag, "<none>") {
		return c.WithTag()
	}
	if c.Digest != "" {
		return c.WithDigest()
	}
	return c.Image
}

// Storage Path returns a full storage relative path
func (c *URI) StoragePath() string {

	path := c.Image
	if c.Tag != "" {
		path = filepath.Join(path, c.Tag)
	}
	if c.Digest != "" {
		path = filepath.Join(path, c.Digest)
	}
	return path
}

// Parse the RAW URI
func (c *URI) Parse() {
	if c.Raw == "" {
		return
	}

	// <registry?>/<image>@<hash>
	if strings.Contains(c.Raw, "@") {
		parts := strings.SplitN(c.Raw, "@", 2)
		c.Image = parts[0]
		c.Digest = parts[1]

		// We have a name and tag
		// <registry?>/<image>:<tag>
	} else if strings.Contains(c.Raw, ":") {
		parts := strings.SplitN(c.Raw, ":", 2)
		c.Image = parts[0]
		c.Tag = parts[1]

	} else {
		c.Image = c.Raw
	}
	// TODO more complicated cases
}

// The tagged image
func (c *URI) WithTag() string {
	return c.Image + ":" + c.Tag
}

// The digest
func (c *URI) WithDigest() string {
	return c.Image + "@" + c.Digest
}

// Determine if a contender name matches the known name
func (c *URI) Matches(name string) bool {

	// Case 1: the name matches exactly
	if name == c.Image || name == c.WithTag() || name == c.WithDigest() {
		return true
	}

	// For each of the image, with docker, and with localhost, see if matches
	for _, reg := range []string{"localhost", "docker.io"} {
		if name == reg+"/"+c.Image || name == reg+"/"+c.WithTag() || name == reg+"/"+c.WithDigest() {
			return true
		}
	}
	return false
}
