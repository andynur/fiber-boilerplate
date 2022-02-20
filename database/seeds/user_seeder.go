package seeds

import (
	"github.com/andynur/fiber-boilerplate/app/models"
	"gorm.io/gorm"
)

type UserSeeder struct{}

func (s *UserSeeder) CreateUser(db *gorm.DB, firstName string, lastName string) error {
	return db.Create(&models.User{FirstName: firstName, LastName: lastName}).Error
}

func (s *UserSeeder) Run() []Seeder {
	return []Seeder{}
}
