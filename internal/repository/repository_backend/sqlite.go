package repository_backend

import (
	"time"

	"github.com/RG1ee/gobot/pkg/domain"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sqlite struct {
	DB_name string
	db      *gorm.DB
}

// GetIncoming implements repository.Cloth.
func (s *Sqlite) GetIncoming() []domain.Cloth {
	var clothes []domain.Cloth
	s.db.Where("status = ?", domain.ClothIncoming).Find(&clothes)
	return clothes
}

// GetOutgoing implements repository.Cloth.
func (s *Sqlite) GetOutgoing() []domain.Cloth {
	var clothes []domain.Cloth
	s.db.Where("status = ?", domain.ClothOutgoing).Find(&clothes)
	return clothes
}

// Insert implements repository.Cloth.
func (s *Sqlite) Insert(c domain.Cloth) {
	s.clearRotten()
	s.db.Create(c)
}

// Out implements repository.Cloth.
func (s *Sqlite) Out(c domain.Cloth) error {
	return s.db.Model(&c).Updates(map[string]interface{}{"Status": domain.ClothOutgoing, "OutgoingDate": time.Now()}).Error
}

// clearRotten implements repository.Cloth.
func (s *Sqlite) clearRotten() {
	var cloth domain.Cloth
	const hours_per_day = 24
	const days = 7
	duration := days * hours_per_day
	s.db.Where("incoming_date >= ?", time.Now().Add(time.Duration(-duration)*time.Hour)).Delete(&cloth)
}

func (s *Sqlite) Init() {
	db, err := gorm.Open(sqlite.Open(s.db_name), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	var cloth domain.Cloth
	db.AutoMigrate(&cloth)
	s.db = db
}
