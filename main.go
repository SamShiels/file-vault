package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("", "public")
	e.POST("/upload", upload)

	e.Logger.Fatal(e.Start(":8082"))
}

func upload(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	fileName := file.Filename

	dst, err := os.Create("./" + fileName)
	if err != nil {
		return err
	}

	defer dst.Close()

	var bytesCopied int64
	if bytesCopied, err = io.Copy(dst, src); err != nil {
		return err
	}

	println("Copied file:", file.Filename, ", bytes:", bytesCopied)

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>File %s uploaded successfully with fields name=%s and email=%s.</p>", file.Filename, name, email))
}
