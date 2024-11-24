package service

import (
	"io"
	"mime/multipart"
	"os"
	"vault/internal/database/models"
	"vault/internal/database/repository"
)

type FileService struct {
	repo *repository.FileRepository
}

func NewFileService(repo *repository.FileRepository) *FileService {
	return &FileService{
		repo: repo,
	}
}

func (s *FileService) Upload(fileHeader *multipart.FileHeader) error {
	src, err := fileHeader.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	fileName := fileHeader.Filename

	dst, err := os.Create("./" + fileName)
	if err != nil {
		return err
	}

	defer dst.Close()

	var bytesCopied int64
	if bytesCopied, err = io.Copy(dst, src); err != nil {
		return err
	}

	var file *models.File = &models.File{
		Name: fileName,
		Size: int(bytesCopied),
	}

	s.repo.UploadFile(*file)

	println("Copied file:", fileHeader.Filename, ", bytes:", bytesCopied)

	return nil
}
