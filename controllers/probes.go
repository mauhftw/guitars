package controllers

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/mauhftw/go-guitarists/config"
	"github.com/mauhftw/go-guitarists/helpers"
)

type Health struct {
	Status string `json:status`
	Uptime int    `json:uptime`
}

type Version struct {
	Status  string `json:status`
	Version string `json:version`
}

// returns system health and uptime
func GetHealth(c echo.Context) error {
	u := int(time.Since(helpers.StartTime).Seconds())
	uptime := &Health{
		Status: "ok",
		Uptime: u,
	}

	return c.JSON(http.StatusOK, uptime)
}

// returns the release software version
func GetVersion(c echo.Context) error {

	// This is being managed by config.go file. I don't like using empty interfaces
	version := config.GuitaristReleaseVersion

	v := &Version{
		Status:  "ok",
		Version: version,
	}

	return c.JSON(http.StatusOK, v)
}
