package main

import (
  "net/http"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  // Routes
  e.GET("/", hello)
  e.GET("/users", userIndexHandler)

  // Start server
  e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}

type User struct {
    Name  string `json:"name" form:"name"`
    Email string `json:"email" form:"email"`
}

type Users []User

func userIndexHandler(c echo.Context) error {
  var users Users

  users = append(users, User {
    Name:  "Taro",
    Email: "taro@example.com",
  })
  users = append(users, User {
    Name:  "Jiro",
    Email: "jiro@example.com",
  })

  return c.JSON(http.StatusOK, users)
}
