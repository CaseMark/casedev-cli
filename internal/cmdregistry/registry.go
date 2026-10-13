// Package cmdregistry is the seam between Stainless-generated CLI code
// in pkg/cmd and hand-written extension subpackages under
// pkg/cmd/extensions/.
//
// Why this exists:
//
// pkg/cmd is generated from the OpenAPI spec by Stainless and is
// rewritten on every regen. Extension subpackages that want to add
// hand-written top-level commands (for example, the webhook listen
// tunnel tool) cannot directly mutate pkg/cmd's command tree from their
// init() because doing so requires importing pkg/cmd, while pkg/cmd's
// bridge file (zz_extensions.go) needs to import them — that's a
// cycle. This package breaks the cycle: extensions Register here, and
// the bridge in pkg/cmd Drains here.
//
// Usage from an extension subpackage:
//
//	func init() {
//	    cmdregistry.Register(&cli.Command{Name: "my:tool", ...})
//	}
//
// Usage from the bridge in pkg/cmd:
//
//	func init() {
//	    cmd.Command.Commands = append(cmd.Command.Commands, cmdregistry.Drain()...)
//	}
package cmdregistry

import (
	"sync"

	"github.com/urfave/cli/v3"
)

var (
	mu      sync.Mutex
	pending []*cli.Command
)

// Register adds a hand-written top-level command to the pending set.
// Safe to call from any package's init().
func Register(c *cli.Command) {
	mu.Lock()
	defer mu.Unlock()
	pending = append(pending, c)
}

// Drain returns every command registered so far and clears the
// pending set. Intended to be called once, from the bridge file in
// pkg/cmd that splices extensions into the generated command tree.
func Drain() []*cli.Command {
	mu.Lock()
	defer mu.Unlock()
	out := pending
	pending = nil
	return out
}
