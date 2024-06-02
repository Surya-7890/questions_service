package database

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/Surya-7890/questions_service/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository() (*Repository, error) {
	db_uri := os.Getenv("POSTGRES_URI")
	db, err := gorm.Open(postgres.Open(db_uri))
	db.Logger = logger.Default.LogMode(logger.Error)
	if err != nil {
		return nil, err
	}
	return &Repository{
		DB: db,
	}, nil
}

func (r *Repository) CreateTables() {
	migrator := r.DB.Migrator()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		exists := migrator.HasTable(&models.Company{})
		if !exists {
			if err := migrator.CreateTable(&models.Company{}); err != nil {
				fmt.Println("error", err)
			}
		}
		exists = migrator.HasTable(&models.Question{})
		if !exists {
			if err := migrator.CreateTable(&models.Question{}); err != nil {
				fmt.Println("error", err)
			}
		}
		wg.Done()
	}()
	wg.Wait()
}

func (r *Repository) CreateNewCompany(name string) error {
	var company models.Company

	err := r.DB.Where("name = ?", name).First(&company).Error

	if err == nil {
		return fmt.Errorf("company with the name already exists")
	}

	company.Name = name
	err = r.DB.Create(&company).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (r *Repository) CreateNewQuestion(company_name, question string, year time.Time) {

}
