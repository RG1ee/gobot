package repository

import (
	"time"

	"github.com/RG1ee/gobot/pkg/domain"
)

type Mock struct {
}

func (Mock) GetIncoming() []domain.Cloth {
	r := make([]domain.Cloth, 1)
	r[0] = domain.Cloth{Name: "test", PhotoId: "123", IncomingDate: time.Now(), OutgoingDate: time.Now(), Status: domain.ClothIncoming}
	return r
}

func (Mock) GetOutgoing() []domain.Cloth {
	r := make([]domain.Cloth, 1)
	r[0] = domain.Cloth{Name: "test", PhotoId: "123", IncomingDate: time.Now(), OutgoingDate: time.Now(), Status: domain.ClothOutgoing}
	return r
}

func (Mock) Insert(domain.Cloth) {

}

func (Mock) Out(name string) error {
	return nil
}
