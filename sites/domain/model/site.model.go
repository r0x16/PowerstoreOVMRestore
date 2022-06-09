package model

import "github.com/r0x16/PowerstoreOVMRestore/shared/domain"

type Site struct {
	domain.Model
	name        string
	description string
}
