package db

import (
	"context"
	"log"
	"sync"
	"time"

	"courses-api/src/config/envs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var once sync.Once
var dbInstance *mongo.Client

func ConnectDB() (*mongo.Client, error) {
	var err error
	env := envs.LoadEnvs(".env")
	MONGO_URI := env.Get("MONGO_URI")

	once.Do(func() {
		clientOptions := options.Client().ApplyURI(MONGO_URI)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		dbInstance, err = mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatalf("Error al conectar con MongoDB: %v", err)
		}

		err = dbInstance.Ping(ctx, nil)
		if err != nil {
			log.Fatalf("No se pudo conectar a MongoDB: %v", err)
		}

		log.Println("Conexión a MongoDB establecida")
	})

	return dbInstance, err
}
