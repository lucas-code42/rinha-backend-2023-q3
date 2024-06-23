package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePerson(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "CreatePerson - NOT IMPLEMENTED!!")
}

func GetPersonById(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "GetPersonById - NOT IMPLEMENTED!!")
}

func SearchPerson(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "SearchPerson - NOT IMPLEMENTED!!")
}

func CountPeople(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "CountPeople - NOT IMPLEMENTED!!")
}
