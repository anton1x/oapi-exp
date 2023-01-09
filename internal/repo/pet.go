package repo

import openapi "github.com/anton1x/petstore/go"

type PetRepo interface {
	Add(pet openapi.Pet) (int, error)
}

type petRepoInmem struct {
	Pets []openapi.Pet
}

func (p *petRepoInmem) Add(pet openapi.Pet) (int, error) {
	p.Pets = append(p.Pets, pet)
	return len(p.Pets), nil
}

func NewPetRepoInmem() *petRepoInmem {
	return &petRepoInmem{}
}
