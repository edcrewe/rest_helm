package actions

import (
	"github.com/gobuffalo/buffalo"
	"rest_helm/rest_buffalo/models"

	"github.com/helm/helm"
)

var db = make(map[string]models.Download)

type DownloadResource struct{}

// @description List all available downloaded charts
// @Param name path string true "Chart name"
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Download
// @Router /download [get]
func (downr DownloadResource) List(c buffalo.Context) error {
	download := &models.Download{
		Name:           "mariadb",
		SourceLocation: "https://kubernetes-charts.storage.googleapis.com/",
	}
	db[download.Name] = *download
	return c.Render(200, r.JSON(db))
}

// @Summary Show downloaded chart
// @Description Show a download by name.
// @Param name path string true "Chart name"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Download
// @Router /download/{name} [get]
func (downr DownloadResource) Show(c buffalo.Context) error {

	name := c.Param("name")
	download, ok := db[name]
	if ok {

		return c.Render(200, r.JSON(download))
	}

	return c.Render(404, r.String("download not found"))
}

// @description Create Download.
/// ...
// @Param name path string true "Chart name"
// @Param url path int true "Chart source URL"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Download
// @Router /download [post]
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
