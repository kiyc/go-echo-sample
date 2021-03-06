package main

import (
  "net/http"
  "os"
  "time"
  "html/template"
	"io"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/joho/godotenv"
  "github.com/dgrijalva/jwt-go"
)

func main() {
  // Echo instance
  e := echo.New()

  // Middleware
  e.Use(middleware.Logger())
  e.Use(middleware.Recover())

  renderer := &Template{
    templates: template.Must(template.ParseGlob("views/*.html")),
  }
  e.Renderer = renderer

  // Routes
  e.Static("/js", "public/js")
  e.GET("/", func(c echo.Context) error {
    return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
  })
  e.POST("/register", userRegisterHandler)
  e.POST("/login", userLoginHandler)

  e.GET("/dashboard", func(c echo.Context) error {
    return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
  })
  e.GET("/users", func(c echo.Context) error {
    return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
  })

  api := e.Group("/api")
  api.Use(middleware.JWT([]byte("secret")))
  api.GET("/users", userIndexHandler)
  api.POST("/users", userRegisterHandler)
  api.GET("/users/:id", userShowHandler)
  api.PATCH("/users/:id", userUpdateHandler)
  api.DELETE("/users/:id", userDeleteHandler)

  // Start server
  e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
  return c.String(http.StatusOK, "Hello, World!")
}

type Template struct {
  templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
  return t.templates.ExecuteTemplate(w, name, data)
}

type User struct {
  gorm.Model
  Account string `form:"account"`
  Password string `form:"password"`
}

type Users []User

type Errors map[string]string

func dbConnect() *gorm.DB {
  err := godotenv.Load()
  if err != nil {
    panic(err.Error())
  }

  dbHost := os.Getenv("MYSQL_HOST")
  dbName := os.Getenv("MYSQL_DBNAME")
  dbUser := os.Getenv("MYSQL_USER")
  dbPass := os.Getenv("MYSQL_PASSWORD")
  dbPort := os.Getenv("MYSQL_PORT")

  db, err := gorm.Open("mysql", dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=true&loc=Asia%2FTokyo")

  if err != nil {
    panic(err.Error())
  }

  return db
}

func userIndexHandler(c echo.Context) error {
  var users Users

  err := godotenv.Load()
  if err != nil {
    panic("環境変数の読込に失敗しました")
  }

  db := dbConnect()
  defer db.Close()

  db.Find(&users)

  return c.JSON(http.StatusOK, users)
}

func userRegisterHandler(c echo.Context) (err error) {
  user := new(User)

  if err = c.Bind(user); err != nil {
    panic(err.Error())
  }

  //var errors Errors
  errors := make(Errors)

  if len(user.Account) == 0 {
    errors["Account"] = "Account is required."
  }
  if len(user.Password) == 0 {
    errors["Password"] = "Password is required."
  }
  if len(errors) > 0 {
    return c.JSON(422, errors)
  }

  db := dbConnect()
  defer db.Close()

  db.Create(&user)

  return c.JSON(http.StatusCreated, user)
}

func userLoginHandler(c echo.Context) error {
  var user User

  account := c.FormValue("account")
  password := c.FormValue("password")

  db := dbConnect()
  defer db.Close()

  result := db.Where("account = ?", account).Where("password = ?", password).First(&user)

  if result.Error != nil {
    return echo.ErrUnauthorized
  }

  token := jwt.New(jwt.SigningMethodHS256)

  claims := token.Claims.(jwt.MapClaims)
  claims["account"] = user.Account
  claims["admin"] = true
  claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

  t, err := token.SignedString([]byte("secret"))
  if err != nil {
    return err
  }

  return c.JSON(http.StatusOK, map[string]string{
    "token": t,
  })
}

func userShowHandler(c echo.Context) error {
  var user User
  id := c.Param("id")

  db := dbConnect()
  defer db.Close()

  result := db.First(&user, id)

  if result.Error != nil {
    return echo.ErrNotFound
  }

  return c.JSON(http.StatusOK, user)
}

func userUpdateHandler(c echo.Context) error {
  var user User
  id := c.Param("id")

  db := dbConnect()
  defer db.Close()

  result := db.First(&user, id)

  if result.Error != nil {
    return echo.ErrNotFound
  }

  account := c.FormValue("account")
  password := c.FormValue("password")

  if len(account) > 0 {
    user.Account = account
  }
  if len(password) > 0 {
    user.Password = password
  }

  db.Save(&user)

  return c.JSON(http.StatusOK, user)
}

func userDeleteHandler(c echo.Context) error {
  var user User
  id := c.Param("id")

  db := dbConnect()
  defer db.Close()

  result := db.First(&user, id)

  if result.Error != nil {
    return echo.ErrNotFound
  }

  db.Delete(&user)

  return c.JSON(http.StatusOK, "Deleted.")
}
