package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/systemfiles/stay-up/api"
)

type User struct {
	FirstName string
	LastName string
	Email string
}

func (u *User) New(firstName, lastName, email string) (error) {
	switch {
	case firstName == "":
		return errors.New("cannot create a new user with no first name")
	case lastName == "":
		return errors.New("cannot create a new user with no last name")
	case email == "":
		return errors.New("cannot create a new user with no email address")
	}

	u.FirstName = firstName
	u.LastName = lastName
	u.Email = email

	return nil
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users", func (c echo.Context) error {
		users := [...]User{
			{"Bob", "Smith", "bob@smith.ca"},
			{"Arthur", "James", "arthur@smith.ca"},
			{"Tod", "Matthews", "todm@gmail.com"},
			{"Bailey", "Jim", "jim@lindsey.co.uk"},
		}

		// Make the response in HTML
		resp := "<html><body><h1>Users:</h1><ul>"
		for _,user := range users {
			resp += "<li>" + user.FirstName + " " + user.LastName + " -> <strong>" + user.Email + "</strong></li>"
		}
		resp += "</ul></body></html>"

		return c.HTML(http.StatusOK, resp)
	})
	e.POST("/users", func (c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}

		user := new(User)
		
		if err := user.New(u.FirstName, u.LastName, u.Email); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to create user. Reason: " + err.Error())
		}

		return c.JSON(http.StatusCreated, user)
	})

	e.GET("status", api.GetStatus)
	e.Logger.Fatal(e.Start(":1323"))
}