package actions

import (
	"github.com/gobuffalo/buffalo"
	"rest_helm/rest_buffalo/models"
)

var db = make(map[string]models.Download)

type DownloadResource struct{}

func (downr DownloadResource) List(c buffalo.Context) error {
	download := &models.Download {
		Name: "mariadb",
		SourceLocation: "https://kubernetes-charts.storage.googleapis.com/",
	}
	db[download.Name] = *download
	return c.Render(200, r.JSON(db))
}

/* Create Download.
func (ur DownloadResource) Create(c buffalo.Context) error {
	name, err := c.Param("name")
	if err != nil {
		return c.Render(500, r.String("name is not found"))
	}
	download := &models.Download{
		name: name,
	}
	db[download.name] = *download

	return c.Render(201, r.JSON(download))
}
*/

// Show Download by name.
func (downr DownloadResource) Show(c buffalo.Context) error {

	name := c.Param("name")
	download, ok := db[name]
	if ok {

		return c.Render(200, r.JSON(download))
	}

	return c.Render(404, r.String("download not found"))
}
