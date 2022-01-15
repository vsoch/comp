package diff

import (
	"fmt"
	"github.com/vsoch/comp/lib/logger"
	"sort"
)

// Get sorted keys for a dict to print
func sortedKeys(env map[string]string) []string {

	// Ensure we print sorted
	keys := make([]string, 0, len(env))
	for key := range env {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// Get sorted keys for a dict to print
func sortedChanges(env map[string]Change) []string {

	// Ensure we print sorted
	keys := make([]string, 0, len(env))
	for key := range env {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// A change represnts a change between envars
type Change struct {
	Name     string `json:"name"`
	Original string `json:"original"`
	New      string `json:"new"`
}

// A diff holds changes, additions, deletions
type Diff struct {
	Added     map[string]string
	Removed   map[string]string
	Unchanged map[string]string
	Changed   map[string]Change
}

func (d *Diff) PrintRemoved() {
	keys := sortedKeys(d.Removed)
	log := logger.Logger{}
	for _, key := range keys {
		name := fmt.Sprintf("%-25s", key)
		log.Red(" - " + name + " " + d.Removed[key])
	}
}

func (d *Diff) PrintAdded() {
	keys := sortedKeys(d.Added)
	log := logger.Logger{}
	for _, key := range keys {
		name := fmt.Sprintf("%-25s", key)
		log.Green(" + " + name + " " + d.Added[key])
	}
}

func (d *Diff) PrintChanged() {
	keys := sortedChanges(d.Changed)
	log := logger.Logger{}
	for _, key := range keys {
		name := fmt.Sprintf("%-25s", key)
		// TODO this should be the change in colors
		log.Yellow(" + " + name + " " + d.Changed[key].Original)
		// log.Yellow(" + " + name + " " + d.GetDiff(key))
	}
}

// Create a new Diff
func NewDiff() *Diff {
	added := map[string]string{}
	removed := map[string]string{}
	unchanged := map[string]string{}
	changed := map[string]Change{}
	return &Diff{Added: added, Removed: removed, Unchanged: unchanged, Changed: changed}
}