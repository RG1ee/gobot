package domain

import "time"

type ClothStatus int

const (
	ClothIncoming ClothStatus = iota
	ClothOutgoing
)

type Cloth struct {
	Name         string
	PhotoId      string
	IncomingDate time.Time
	OutgoingDate time.Time
	Status       ClothStatus
}