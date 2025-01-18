package component_middlewares

import (
	tele "gopkg.in/telebot.v3"
)

type UserState int

type State struct {
	UserState         UserState
	CurrentMessages   []tele.Editable
	PreviousMessages  *[]tele.Editable
	DeletedIdMessages []int
}

const (
	NullState UserState = iota
	StateWaitPhoto
	StateSaveChanges
)
