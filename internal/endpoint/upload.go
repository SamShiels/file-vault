package endpoint

import (
	"fmt"
	"net/http"
	"vault/internal/service"

	"github.com/labstack/echo/v4"
)

type Api struct {
	fileService *service.FileService
}

func NewApi(fileService *service.FileService) *Api {
	return &Api{
		fileService: fileService,
	}
}

func (a *Api) Upload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Unable to retrieve file",
		})
	}

	err = a.fileService.Upload(file)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("Unable to upload file: %s", err),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "File uploaded successfully",
	})
}
