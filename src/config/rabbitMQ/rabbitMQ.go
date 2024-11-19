package rabbitmq

import (
	"courses-api/src/config/envs"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	QueueName  string
	amqpURL    string
	mu         sync.RWMutex
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

		instance = &RabbitMQ{
			QueueName: queueName,
			amqpURL:   amqpURL,
		}

		go instance.connectWithRetry()
	})

	return instance
}

func (r *RabbitMQ) connectWithRetry() {
	for {
		conn, err := amqp.Dial(r.amqpURL)
		if err != nil {
			log.Printf("Error al conectar con RabbitMQ: %v. Reintentando en 5 segundos...", err)
			time.Sleep(5 * time.Second)
			continue
		}

		ch, err := conn.Channel()
		if err != nil {
			log.Printf("Error al abrir un canal en RabbitMQ: %v. Reintentando en 5 segundos...", err)
			conn.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		_, err = ch.QueueDeclare(
			r.QueueName,
			true,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Printf("Error al declarar la cola en RabbitMQ: %v. Reintentando en 5 segundos...", err)
			ch.Close()
			conn.Close()
			time.Sleep(5 * time.Second)
			continue
		}

		// Conexión y canal exitosos
		r.mu.Lock()
		r.connection = conn
		r.channel = ch
		r.mu.Unlock()

		log.Println("Conexión a RabbitMQ establecida y cola declarada.")

		// Manejar la reconexión si la conexión se pierde
		closeChan := make(chan *amqp.Error)
		r.connection.NotifyClose(closeChan)

		err = <-closeChan
		if err != nil {
			log.Printf("Conexión a RabbitMQ cerrada: %v. Reintentando conexión...", err)
		}

		// Limpiar y volver a intentar la conexión
		r.mu.Lock()
		r.channel = nil
		r.connection = nil
		r.mu.Unlock()

		// Esperar antes de reintentar
		time.Sleep(5 * time.Second)
	}
}

func (r *RabbitMQ) PublishMessage(message string) error {
	r.mu.RLock()
	ch := r.channel
	r.mu.RUnlock()

	if ch == nil {
		log.Println("Canal de RabbitMQ no está listo. No se puede publicar el mensaje.")
		return fmt.Errorf("canal de RabbitMQ no está establecido")
	}

	err := ch.Publish(
		"",
		r.QueueName,
		false,
		false,
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
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.channel != nil {
		r.channel.Close()
	}
	if r.connection != nil {
		r.connection.Close()
	}
	log.Println("Conexión a RabbitMQ cerrada")
}
