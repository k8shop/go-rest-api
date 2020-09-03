package informer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/k8shop/go-rest-api/pkg/models"
	kafka "github.com/segmentio/kafka-go"
)

//Informer for events
type Informer struct {
	kafkas map[string]*kafka.Writer
}

//NewInformer informed informer
func NewInformer(brokers []string) *Informer {
	config := kafka.WriterConfig{}
	config.Brokers = brokers

	i := &Informer{kafkas: map[string]*kafka.Writer{}}
	config.Topic = "products"
	i.kafkas["products"] = kafka.NewWriter(config)
	return i
}

//InformProducts informer info
func (i *Informer) InformProducts(p *models.Product) error {
	value, err := json.Marshal(p)
	if err != nil {
		return err
	}
	log.Printf("informing topic: products of value: %+v", string(value))
	msg := kafka.Message{Value: value}
	return i.kafkas["products"].WriteMessages(context.Background(), msg)
}

//Close all open informers
func (i *Informer) Close() {
	for _, kafka := range i.kafkas {
		kafka.Close()
	}
}
