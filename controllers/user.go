package controllers

import (
	"encoding/json"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type User struct {
	Name string `json:name`
	Age  int    `json:age`
}

type Data struct {
	Uptime  int    `json:uptime`
	Version string `json:version`
}

type Meta struct {
	Count int `json:count`
}

type Response struct {
	Meta Meta `json:meta`
	Data Data `json:data`
}

// returns all users
func GetUsers(c echo.Context) error {
	users := &User{
		Name: "Jimmy",
		Age:  35,
	}
	log.Printf("%v", users)
	return c.JSON(http.StatusOK, users)
}

// returns
func GetHellos(c echo.Context) error {

	resp, err := http.Get("http://cgb-backend.sweatworks.tk/version")
	if err != nil {
		return echo.NewHTTPError(404, "wrong url")
	}

	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)

	health := &Response{}
	if err := dec.Decode(health); err != nil {
		log.Fatalf("error: can't decode - %s", err)
	}
	return c.JSON(http.StatusOK, health)

}
