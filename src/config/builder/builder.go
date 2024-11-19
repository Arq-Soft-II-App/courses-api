package builder

import (
	"courses-api/src/clients"
	"courses-api/src/config/db"
	"courses-api/src/config/envs"
	rabbitmq "courses-api/src/config/rabbitMQ"
	"courses-api/src/controllers"
	"courses-api/src/routes"
	"courses-api/src/services"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppBuilder struct {
	mongoClient *mongo.Client
	database    *mongo.Database
	clients     *clients.Clients
	services    *services.Services
	controllers *controllers.Controllers
	router      *gin.Engine
	envs        envs.Envs
	rabbitMQ    *rabbitmq.RabbitMQ
}

func NewAppBuilder() *AppBuilder {
	return &AppBuilder{}
}

func BuildApp() *AppBuilder {
	builder := NewAppBuilder()
	builder.envs = envs.LoadEnvs()
	return builder.
		BuildDBConnection().
		BuildRabbitMQConnection().
		BuildClients().
		BuildServices().
		BuildControllers().
		BuildRouter()
}

func (b *AppBuilder) BuildDBConnection() *AppBuilder {
	var err error
	b.mongoClient, err = db.ConnectDB()
	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	b.database = b.mongoClient.Database("courses-api")
	return b
}

func (b *AppBuilder) BuildRabbitMQConnection() *AppBuilder {
	b.rabbitMQ = rabbitmq.NewRabbitMQ()
	return b
}

func (b *AppBuilder) BuildClients() *AppBuilder {
	b.clients = clients.NewClients(b.database)
	return b
}

func (b *AppBuilder) BuildServices() *AppBuilder {
	b.services = services.NewServices(b.clients, b.rabbitMQ)
	return b
}

func (b *AppBuilder) BuildControllers() *AppBuilder {
	b.controllers = controllers.NewControllers(b.services)
	return b
}

func (b *AppBuilder) BuildRouter() *AppBuilder {
	b.router = gin.Default()
	routes.SetupRoutes(b.router, *b.controllers)
	return b
}

func (b *AppBuilder) GetRouter() *gin.Engine {
	return b.router
}

func (b *AppBuilder) GetPort() string {
	port := b.envs.Get("PORT")
	if port == "" {
		port = "4002"
	}
	return ":" + port
}
