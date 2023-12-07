package pubsub

import (
	"archetype/app/shared/archetype/ce_pubsub/topic"
	"context"
	"errors"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	cecontext "github.com/cloudevents/sdk-go/v2/context"
)

var ArchetypePublisher = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) error {
	event := cloudevents.NewEvent()
	//example source
	event.SetType("eventType")
	event.SetSource("eventSource")
	_ = event.SetData("application/json", REPLACE_BY_YOUR_DOMAIN)
	ctxWithTopic := cecontext.WithTopic(ctx, "INSERT YOUR TOPIC NAME HERE")
	result := topic.R.Send(ctxWithTopic, event)
	if cloudevents.IsUndelivered(result) {
		log.Printf("failed to send: %v", result.Error())
		return errors.New(result.Error())
	}
	return nil
}
