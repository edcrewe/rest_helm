package grifts

import (
  "github.com/gobuffalo/buffalo"
  "rest_helm/rest_buffalo/actions"
)

func init() {
  buffalo.Grifts(actions.App())
}
