// Package antigravity reads what an Antigravity CLI session is about from the
// annotation the agent keeps beside it. Herdr says which session a pane holds;
// the annotation is where the agent stores the title it generated for it.
package antigravity

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

// Agent is the label Herdr gives Antigravity CLI.
const Agent = "agy"

// titlePattern matches the title field of an annotation, which is protobuf
// text format holding one quoted string.
var titlePattern = regexp.MustCompile(`(?m)^\s*title\s*:\s*("(?:[^"\\\n]|\\.)*")`)

type Reader struct {
	root string
}

// NewReader builds a reader over the user's home directory, where Antigravity
// CLI keeps its state without letting the user move it. An empty home reads
// nothing.
func NewReader(home string) *Reader {
	if home == "" {
		return &Reader{}
	}

	return &Reader{root: filepath.Join(home, ".gemini", "antigravity-cli")}
}

// Title reports what the agent named the session, or "" when nothing can be
// read. The id becomes part of a path, so it must be one
// herdr.AgentSessionInfo.ID accepted.
func (r *Reader) Title(sessionID string) string {
	if r.root == "" {
		return ""
	}

	//nolint:gosec // the root is the user's home and herdr refuses an id that is not a UUID
	content, err := os.ReadFile(filepath.Join(r.root, "annotations", sessionID+".pbtxt"))
	if err != nil {
		return ""
	}

	match := titlePattern.FindSubmatch(content)
	if match == nil {
		return ""
	}

	title, err := strconv.Unquote(string(match[1]))
	if err != nil {
		return ""
	}

	return title
}
