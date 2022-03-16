package repository

import (
	"gingo/apps/db"

	"go.mongodb.org/mongo-driver/mongo"
)

type userEntity struct {
	resource *db.Resource
	repo     *mongo.Collection
}

type IUser interface {
	GetOneById(id string) (int, error)
}

func (r *Repository) NewUserEntity() IUser {
	repo := r.resource.DB.Collection("user")
	return &userEntity{r.resource, repo}
}

func (e *userEntity) GetOneById(id string) (int, error) {
	return 0, nil
}
