// This file is hand-written and is the bridge between the
// Stainless-generated CLI tree in pkg/cmd and the hand-written
// extension subpackages under pkg/cmd/extensions/.
//
// How it works
//
//   1. Each extension subpackage uses init() to call
//      cmdregistry.Register(&cli.Command{...}). Because each subpackage
//      is blank-imported below, the Go runtime executes its init()
//      during program startup.
//
//   2. The init() in this file runs after pkg/cmd/cmd.go's init() —
//      Go runs init() functions within a package in alphabetical
//      filename order, and "zz_" sorts after every generated filename.
//      By the time we run, Command has been fully built. We splice the
//      registered extensions into Command.Commands.
//
// To add a new hand-written tool
//
//   1. Create pkg/cmd/extensions/<toolname>/ with package <toolname>.
//   2. In its cmd.go, init() should call cmdregistry.Register with a
//      *cli.Command describing the tool.
//   3. Add a blank-import below for the new subpackage.
//
// This file MUST NOT be moved into a subpackage and MUST keep its
// "zz_" filename prefix.

package cmd

import (
	"github.com/CaseMark/casedev-cli/internal/cmdregistry"

	_ "github.com/CaseMark/casedev-cli/pkg/cmd/extensions/webhooklisten"
)

func init() {
	Command.Commands = append(Command.Commands, cmdregistry.Drain()...)
}
