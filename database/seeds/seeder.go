package seeds

import "gorm.io/gorm"

type Seeder struct {
	Name string
	Run  func(*gorm.DB) error
}

func (s *Seeder) All() []Seeder {
	return []Seeder{}
}
