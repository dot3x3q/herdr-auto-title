package herdr_test

import (
	"strings"
	"testing"

	"github.com/kryptamine/herdr-auto-title/internal/herdr"
)

const session = "8852bfe0-8b24-4a23-a35e-7521d04da061"

func TestASessionIDIsReported(t *testing.T) {
	t.Parallel()

	ref := &herdr.AgentSessionInfo{Agent: "claude", Kind: herdr.SessionRefID, Value: session}

	if got, ok := ref.ID(); !ok || got != session {
		t.Errorf("ID() = %q, %v, want %q", got, ok, session)
	}
}

func TestAnIDThatIsNotASessionIsRefused(t *testing.T) {
	t.Parallel()

	// The id arrives over the socket and a reader makes it part of a path.
	for _, value := range []string{
		"", "../../../etc/passwd", "../annotations/" + session, "8852bfe0", strings.Repeat("a", 36),
	} {
		ref := &herdr.AgentSessionInfo{Agent: "claude", Kind: herdr.SessionRefID, Value: value}
		if got, ok := ref.ID(); ok {
			t.Errorf("value %q reported as id %q", value, got)
		}
	}
}

func TestAReferenceThatIsNotAnIDIsRefused(t *testing.T) {
	t.Parallel()

	ref := &herdr.AgentSessionInfo{Agent: "claude", Kind: "path", Value: session}

	if got, ok := ref.ID(); ok {
		t.Errorf("ID() = %q, want none for kind %q", got, ref.Kind)
	}
}

func TestAPaneWithNoSessionHasNoID(t *testing.T) {
	t.Parallel()

	var ref *herdr.AgentSessionInfo

	if got, ok := ref.ID(); ok {
		t.Errorf("ID() = %q, want none for a nil reference", got)
	}
}
