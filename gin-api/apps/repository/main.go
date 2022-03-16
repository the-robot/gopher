package repository

import (
	"context"
	"time"

	"gingo/apps/db"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	resource *db.Resource
	repo     *mongo.Collection
}

func NewRepository(resource *db.Resource) *Repository {
	repo := resource.DB.Collection("user")
	return &Repository{resource, repo}
}

func initContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	return ctx, cancel
}
