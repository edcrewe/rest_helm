package actions

import (
	"github.com/gobuffalo/buffalo"
	"rest_helm/rest_buffalo/models"
)

var db = make(map[string]models.Download)

type DownloadResource struct{}

// List all available downloaded charts
// @Router /api/v1/download [get]
func (downr DownloadResource) List(c buffalo.Context) error {
	download := &models.Download{
		Name:           "mariadb",
		SourceLocation: "https://kubernetes-charts.storage.googleapis.com/",
	}
	db[download.Name] = *download
	return c.Render(200, r.JSON(db))
}

// Show a download by name.
// @Param name path string true "Chart name"
// @Router /api/v1/download/{name} [get]
func (downr DownloadResource) Show(c buffalo.Context) error {

	name := c.Param("name")
	download, ok := db[name]
	if ok {

		return c.Render(200, r.JSON(download))
	}

	return c.Render(404, r.String("download not found"))
}

// Create Download.
/// ...
// @Param name path string true "Chart name"
// @Param url path int true "Chart source URL"
// ...
// @Router /api/v1 [post]
func (ur DownloadResource) Create(c buffalo.Context) error {
	name := c.Param("name")
	url := c.Param("url")

	download := &models.Download{
		Name:           name,
		SourceLocation: url,
	}
	db[name] = *download

	return c.Render(201, r.JSON(download))
}
