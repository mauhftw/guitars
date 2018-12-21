package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	"go-guitarists/helpers"
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

	// This has to be managed using a configuration file
	version := os.Getenv("GUITARISTS_RELEASE_VERSION")
	if version == "" {
		version = "latest"
	}

	v := &Version{
		Status:  "ok",
		Version: version,
	}

	return c.JSON(http.StatusOK, v)
}
