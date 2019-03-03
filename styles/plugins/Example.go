/* WordDensity.go

 */
package main

import (
	"strings"

	"github.com/errata-ai/vale/core"
)

// Scope determines where this rule applies.
//
// See https://errata-ai.github.io/vale/formats/.
var Scope = "summary"

// Level specifies the importance of this rule.
//
// See https://errata-ai.github.io/vale/styles/#extension-points.
var Level = "warning"

// Rule is the entry point to your custom rule.
//
//
func Rule(text string, file *core.File) []core.Alert {
	alerts := []core.Alert{}
	pos := strings.Index(text, "The")
	if pos >= 0 {
		alerts = append(alerts,
			core.Alert{
				Check: "plugins.Example",
				Severity: Level,
				Span: []int{pos, pos + 3},
				Message: "Don't use 'The'!"})
	}
	return alerts
}
