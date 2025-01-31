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

func (s *Sqlite) GetOutgoingLastSevenDays() []domain.Cloth {
	var clothes []domain.Cloth
	sevenDaysAgo := time.Now().AddDate(0, 0, -7)
	s.db.Where("status = ? AND outgoing_date >= ?", domain.ClothOutgoing, sevenDaysAgo).Find(&clothes)
	return clothes
}

// GetById implements repository.Cloth.
func (s *Sqlite) GetById(id int) (domain.Cloth, error) {
	var cloth domain.Cloth
	result := s.db.First(&cloth, id)

	if result.Error != nil {
		return domain.Cloth{}, result.Error
	}

	return cloth, nil
}

// Insert implements repository.Cloth.
func (s *Sqlite) Insert(c domain.Cloth) {
	s.ClearRotten()
	s.db.Create(&c)
}

// Out implements repository.Cloth.
func (s *Sqlite) Out(c domain.Cloth) error {
	return s.db.Model(&c).Updates(map[string]interface{}{"Status": domain.ClothOutgoing, "OutgoingDate": time.Now()}).Error
}

// clearRotten implements repository.Cloth.
func (s *Sqlite) ClearRotten() {
	const hours_per_day = 24
	const days = 7
	duration := days * hours_per_day
	s.db.Where("outgoing_date <= ?", time.Now().Add(time.Duration(-duration)*time.Hour)).Delete(&domain.Cloth{})
}

func (s *Sqlite) Init() {
	db, err := gorm.Open(sqlite.Open(s.DB_name), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	var cloth domain.Cloth
	err = db.AutoMigrate(&cloth)

	if err != nil {
		panic(err)
	}

	s.db = db
}
