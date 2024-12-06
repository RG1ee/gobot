package repository

import "github.com/RG1ee/gobot/pkg/domain"

type Cloth interface {
	GetIncoming() []domain.Cloth
	GetOutgoing() []domain.Cloth
	GetById(id int) (domain.Cloth, error)
	Insert(domain.Cloth)
	Out(domain.Cloth) error
	Init()
	ClearRotten()
}
