package grifts

import (
  "github.com/gobuffalo/buffalo"
	"rest_helm/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
