// Package webhooklisten implements the `casedev webhook:listen` tool —
// an ephemeral ngrok-backed webhook tunnel that creates a temporary
// webhook endpoint pointing at a local HTTP server, and tears the
// endpoint down on Ctrl-C.
//
// This is a hand-written CLI extension. It registers itself via
// internal/cmdregistry; the bridge file pkg/cmd/zz_extensions.go
// blank-imports this package and splices the registered command into
// the Stainless-generated root command tree.
package webhooklisten

import (
	"fmt"
	"strings"

	"github.com/CaseMark/casedev-cli/internal/baseurl"
	"github.com/CaseMark/casedev-cli/internal/cmdregistry"
	"github.com/urfave/cli/v3"
)

func init() {
	cmdregistry.Register(&cli.Command{
		Name:     "webhook:listen",
		Category: "TOOLS",
		Usage:    "Stream platform webhooks to your local machine via ngrok tunnel",
		Suggest:  true,
		Flags:    runCmd.Flags,
		Action:   handleRun,
		Commands: []*cli.Command{
			&runCmd,
			&initCmd,
			&showCmd,
		},
	})
}

var runCmd = cli.Command{
	Name:            "run",
	Usage:           "Start an ngrok-backed webhook listener and optionally forward deliveries locally.",
	Suggest:         true,
	HideHelpCommand: true,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "events",
			Usage: "Comma-separated event-type glob patterns",
		},
		&cli.StringFlag{
			Name:      "forward-to",
			Usage:     "Forward received webhooks to this local URL",
			Validator: func(value string) error { return validateOptionalURL(value, "--forward-to") },
		},
		&cli.IntFlag{
			Name:  "tunnel-port",
			Usage: "Local port for the ngrok tunnel (default: auto-assign)",
		},
		&cli.BoolFlag{
			Name:  "show-secret",
			Usage: "Print the webhook signing secret to stderr after endpoint registration. Off by default to avoid leaks via shell history or scrollback; opt in only when you need the secret for signature verification.",
		},
		&cli.StringFlag{
			Name:  "print",
			Usage: "Output format for received webhooks (pretty, raw)",
			Value: "pretty",
			Validator: func(value string) error {
				switch strings.ToLower(strings.TrimSpace(value)) {
				case "pretty", "raw":
					return nil
				default:
					return fmt.Errorf("--print must be one of: pretty, raw")
				}
			},
		},
	},
	Action: handleRun,
}

var initCmd = cli.Command{
	Name:            "init",
	Usage:           "Create or update saved webhook listen configuration.",
	Suggest:         true,
	HideHelpCommand: true,
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "global", Aliases: []string{"g"}, Usage: "Use ~/.config/casedev/webhooks.json instead of .casedev/webhooks.json"},
		&cli.StringSliceFlag{Name: "add", Usage: "Add an event pattern to saved config"},
		&cli.StringSliceFlag{Name: "remove", Usage: "Remove an event pattern from saved config"},
		&cli.StringFlag{
			Name:      "endpoint",
			Usage:     "Set the default forward-to URL",
			Validator: func(value string) error { return validateOptionalURL(value, "--endpoint") },
		},
	},
	Action: handleInit,
}

var showCmd = cli.Command{
	Name:            "show",
	Usage:           "Print resolved webhook listen configuration.",
	Suggest:         true,
	HideHelpCommand: true,
	Flags: []cli.Flag{
		&cli.BoolFlag{Name: "global", Aliases: []string{"g"}, Usage: "Show only ~/.config/casedev/webhooks.json"},
	},
	Action: handleShow,
}

func validateOptionalURL(value, flagName string) error {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}
	return baseurl.Validate(value, flagName)
}
