package events

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
)

type EventStream struct {
	// TODO: add number of components using stream and conncetion status
	Stream *gochannel.GoChannel
}

var (
	singlePubSubInstance      *gochannel.GoChannel
	singleEventStreamInstance *EventStream
	once                      sync.Once
)

func NewEventStream() *EventStream {
	if singleEventStreamInstance == nil {
		once.Do(func() {
			singleEventStreamInstance = &EventStream{
				Stream: NewChannel(),
			}
		})
		fmt.Println("✅ Event stream created")
	} else {
		fmt.Println("Event stream already created")
	}
	return singleEventStreamInstance
}
func NewChannel() *gochannel.GoChannel {
	singlePubSubInstance = gochannel.NewGoChannel(
		gochannel.Config{},
		watermill.NewStdLogger(false, false))
	return singlePubSubInstance
}

func (ev *EventStream) PublishMessage(payload []byte, topic string) {
	
		msg := message.NewMessage(watermill.NewUUID(), payload)
		err := ev.Stream.Publish(topic, msg)
		if err != nil {
			log.Printf("unable to publish message with topic: %v, :%v", topic, err.Error())
		}
		time.Sleep(time.Second)
		fmt.Println("✅ publishing service started...")
	

		
		// fmt.Println(payload)
	

}

func (ev *EventStream) SubscribeMessage(ctx context.Context, topic string) <-chan *message.Message {
	messages, err := ev.Stream.Subscribe(ctx, topic)
	if err != nil {
		log.Printf("no subscriber to topic : %v, %v", topic, err.Error())
	}
	fmt.Println("✅ subscription service started")
	return messages
}

func (ev *EventStream) Process(messages <-chan *message.Message) {
	for msg := range messages {
		// fmt.Printf("received message: %s, payload: %s\n", msg.UUID, string(msg.Payload))
		msg.Ack()
	}
}
