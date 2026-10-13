// Package baseurl validates base URLs supplied via flags or environment
// variables. It lives outside pkg/cmd so that hand-written extension
// subpackages can import it without creating an import cycle through
// the Stainless-generated pkg/cmd package.
package baseurl

import (
	"fmt"
	"strings"
)

// Validate returns an error if value is non-empty but does not start
// with http:// or https://. The source argument names the flag or
// environment variable for inclusion in the error message.
func Validate(value, source string) error {
	if value != "" && !strings.HasPrefix(value, "http://") && !strings.HasPrefix(value, "https://") {
		return fmt.Errorf("%s %q is missing a scheme (expected http:// or https://)", source, value)
	}
	return nil
}
