package repository

import (
	"context"
	"log"
	"os"
	"vault/internal/database/models"

	"github.com/jackc/pgx/v4/pgxpool"
)

type FileRepository struct {
	db *pgxpool.Pool
}

func (c *FileRepository) Connect() error {
	url := os.Getenv("DATABASE_URL")

	newDbConnection, err := pgxpool.Connect(context.Background(), url)

	c.db = newDbConnection

	if err != nil {
		return err
	}

	log.Println("Connected to DB")
	return nil
}

func (c *FileRepository) GetFileByName(name string) (models.File, error) {
	row := c.db.QueryRow(context.Background(), "SELECT name, size FROM files WHERE name = $1", name)

	var file models.File
	err := row.Scan(&file.Name, &file.Size)

	if err != nil {
		return models.File{}, err
	}

	return file, nil
}

func (c *FileRepository) UploadFile(file models.File) error {
	_, err := c.db.Exec(context.Background(), "INSERT INTO files (name, size) VALUES ($1, $2)", file.Name, file.Size)

	if err != nil {
		return err
	}

	return nil
}

func (c *FileRepository) Close() {
	if c.db != nil {
		c.db.Close()
	}
}
