package rabbitmq

import (
	"courses-api/src/config/envs"
	"log"
	"sync"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	QueueName  string
}

var instance *RabbitMQ
var once sync.Once

func NewRabbitMQ() *RabbitMQ {
	once.Do(func() {
		env := envs.LoadEnvs()
		amqpURL := env.Get("RABBITMQ_URL")
		queueName := env.Get("RABBITMQ_QUEUE_NAME")
		if queueName == "" {
			queueName = "course_updates"
		}

		conn, err := amqp.Dial(amqpURL)
		if err != nil {
			log.Fatalf("Error al conectar con RabbitMQ: %v", err)
		}

		ch, err := conn.Channel()
		if err != nil {
			log.Fatalf("Error al abrir un canal en RabbitMQ: %v", err)
		}

		_, err = ch.QueueDeclare(
			queueName, // name
			true,      // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments
		)
		if err != nil {
			log.Fatalf("Error al declarar la cola en RabbitMQ: %v", err)
		}

		instance = &RabbitMQ{
			connection: conn,
			channel:    ch,
			QueueName:  queueName,
		}

		log.Println("Conexión a RabbitMQ establecida y cola declarada")
	})

	return instance
}

func (r *RabbitMQ) PublishMessage(message string) error {
	err := r.channel.Publish(
		"",          // exchange
		r.QueueName, // routing key (queue name)
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Printf("Error al publicar mensaje en RabbitMQ: %v", err)
	}
	return err
}

func (r *RabbitMQ) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.connection != nil {
		r.connection.Close()
	}
	log.Println("Conexión a RabbitMQ cerrada")
}
