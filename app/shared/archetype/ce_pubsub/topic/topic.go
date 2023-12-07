package topic

import (
	"archetype/app/shared/archetype/ce_pubsub"
	"archetype/app/shared/archetype/container"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

var R cloudevents.Client

func init() {
	container.InjectOutboundAdapter(func() error {
		c, err := cloudevents.NewClient(ce_pubsub.Protocol, cloudevents.WithTimeNow(), cloudevents.WithUUIDs())
		if err != nil {
			log.Printf("failed to create client, %s", err.Error())
			return err
		}
		R = c
		return nil
	})
}
