package service

import (
	"context"
	openapi "github.com/anton1x/petstore/go"
	"gorillagrpc/internal/repo"
	"os"
)

type petAPI struct {
	openapi.PetApiService
	repo repo.PetRepo
}

func NewPetApi(repo repo.PetRepo) *petAPI {
	return &petAPI{repo: repo}
}

func (p petAPI) AddPet(ctx context.Context, pet openapi.Pet) (openapi.ImplResponse, error) {

	//TODO: Uncomment the next line to return response Response(200, Pet{}) or use other options such as http.Ok ...
	//return Response(200, Pet{}), nil

	id, err := p.repo.Add(pet)

	if err != nil {
		return openapi.Response(405, nil), err
	}

	pet.Id = int64(id)

	return openapi.Response(200, pet), nil
}

func (p petAPI) DeletePet(ctx context.Context, i int64, s string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p petAPI) FindPetsByStatus(ctx context.Context, strings []string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p petAPI) FindPetsByTags(ctx context.Context, strings []string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p petAPI) GetPetById(ctx context.Context, i int64) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p petAPI) UpdatePet(ctx context.Context, pet openapi.Pet) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p petAPI) UpdatePetWithForm(ctx context.Context, i int64, s string, s2 string) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (p petAPI) UploadFile(ctx context.Context, i int64, s string, file os.File) (openapi.ImplResponse, error) {
	//TODO implement me
	panic("implement me")
}
