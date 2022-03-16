package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getURI() string {
	return fmt.Sprint(
		"mongodb://",
		viper.GetString("database.config.username"), ":", viper.GetString("database.config.password"),
		"@", viper.GetString("database.config.host"), ":", viper.GetString("database.config.port"),
	)
}

func getDBName() string {
	return viper.GetString("database.config.database_name")
}

type Resource struct {
	DB *mongo.Database
}

// Close use this method to close database connection
func (r *Resource) Close() {
	logrus.Warning("Closing all db connections")
}

func Init() (*Resource, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(getURI()))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer cancel()

	return &Resource{DB: client.Database(getDBName())}, nil
}
