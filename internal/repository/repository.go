package repository

import "github.com/RG1ee/gobot/pkg/domain"

type Cloth interface {
	GetIncoming() []domain.Cloth
	GetOutgoing() []domain.Cloth
	Insert(domain.Cloth)
	Out(domain.Cloth) error
	Init()
	clearRotten()
}
