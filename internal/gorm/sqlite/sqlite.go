package sqlite

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/0B1t322/items/internal/models"
	"github.com/glebarez/sqlite"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"io"
	"os"
)

// ConnectToDatabase connect to sqlite database
func ConnectToDatabase(sqliteDbFilePath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(sqliteDbFilePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&models.Item{}); err != nil {
		return nil, err
	}

	return db, nil
}

// ApplyMigrations apply new migrations to database or launch init schema
func ApplyMigrations(db *gorm.DB) error {
	m := gormigrate.New(
		db, gormigrate.DefaultOptions, []*gormigrate.Migration{},
	)

	m.InitSchema(
		func(db *gorm.DB) error {
			if err := db.AutoMigrate(&models.Item{}); err != nil {
				return err
			}

			if err := InitItemsFromCsv(db, "db/ueba.csv"); err != nil {
				return err
			}

			return nil
		},
	)

	return m.Migrate()
}

func InitItemsFromCsv(db *gorm.DB, csvFilePath string) error {
	file, err := os.Open(csvFilePath)
	if err != nil {
		return fmt.Errorf("failed to open csv file: %w", err)
	}
	defer file.Close()

	r := csv.NewReader(file)

	if _, err := r.Read(); err != nil {
		return fmt.Errorf("failed to read header: %w", err)
	}

	err = db.Transaction(
		func(tx *gorm.DB) error {
			for {
				record, err := r.Read()
				if errors.Is(err, io.EOF) {
					break
				}

				// MARK: в csv файле пое is_admin бывает пустое
				if record[48] == "" {
					record[48] = "0"
				}

				if err := tx.
					Model(&models.Item{}).
					Exec(
						`INSERT INTO items VALUES (?)`, record[1:],
					).Error; err != nil {
					return err
				}
			}

			return nil
		},
	)
	if err != nil {
		return fmt.Errorf("failed to init data from csv: %w", err)
	}

	return nil
}
