package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/projects/ai/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
