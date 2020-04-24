package main

import (
  "net/http"
  "os"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/joho/godotenv"
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
  gorm.Model
  Account string
  Password string
}

type Users []User

func userIndexHandler(c echo.Context) error {
  var users Users

  err := godotenv.Load()
  if err != nil {
    panic("環境変数の読込に失敗しました")
  }

  dbHost := os.Getenv("MYSQL_HOST")
  dbName := os.Getenv("MYSQL_DBNAME")
  dbUser := os.Getenv("MYSQL_USER")
  dbPass := os.Getenv("MYSQL_PASSWORD")

  db, err := gorm.Open("mysql", dbUser + ":" + dbPass + "@tcp(" + dbHost + ":3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local")
  if err != nil {
    panic("データベースへの接続に失敗しました")
  }
  defer db.Close()

  db.Find(&users)

  return c.JSON(http.StatusOK, users)
}
