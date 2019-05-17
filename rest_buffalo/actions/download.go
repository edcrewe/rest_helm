package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/edcrewe/rest_helm/rest_buffalo/models"
)

var db = make(map[string]models.Download)

type DownloadResource struct{}

func (ur DownloadResource) List(c buffalo.Context) error {
	return c.Render(200, r.JSON(db))
}
