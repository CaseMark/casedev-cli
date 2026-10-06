package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

func init() {
	for _, command := range []*cli.Command{&mattersV1Delete, &mattersV1ContentPurgesCreate} {
		command.Flags = append(command.Flags, &cli.BoolFlag{
			Name:  "confirm",
			Usage: "Confirm permanent deletion of the requested matter or content",
		})
	}
}

// Validate after flags, stdin and file input have been merged, before sending a request.
func validateMatterPurgeRequest(command *cli.Command, body any) error {
	name := command.FullName()
	contentPurge := strings.HasSuffix(name, " matters:v1:content-purges create")
	if !contentPurge && !strings.HasSuffix(name, " matters:v1 delete") {
		return nil
	}
	if !command.Bool("confirm") {
		return fmt.Errorf("permanent deletion requires --confirm")
	}
	if !contentPurge {
		return nil
	}

	data, err := json.Marshal(body)
	if err != nil {
		return err
	}
	for _, field := range []string{"object_ids", "session_ids", "transcription_ids", "work_item_ids"} {
		targets := gjson.GetBytes(data, field)
		if !targets.IsArray() {
			continue
		}
		for _, id := range targets.Array() {
			if id.Type == gjson.String && strings.TrimSpace(id.String()) != "" {
				return nil
			}
		}
	}
	return fmt.Errorf("content purge requires at least one nonempty object, session, transcription, or work-item ID")
}
