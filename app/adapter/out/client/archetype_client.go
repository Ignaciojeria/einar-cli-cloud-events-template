package client

import (
	"archetype/app/shared/archetype/ce_client"
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

var ArchetypeCeHttpClient = func(ctx context.Context, REPLACE_BY_YOUR_DOMAIN map[string]string) (err error) {
	ctxWithTarget := cloudevents.ContextWithTarget(ctx, "http://localhost:8080/")
	// Create an Event.
	event := cloudevents.NewEvent()
	//Set firestore Source for example :
	event.SetSource("eventSource")
	event.SetType("entityType")
	event.SetData(cloudevents.ApplicationJSON, REPLACE_BY_YOUR_DOMAIN)
	if result := ce_client.R.Send(ctxWithTarget, event); cloudevents.IsUndelivered(result) {
		log.Fatalf("failed to send, %v", result)
	} else {
		log.Printf("sent: %v", event)
		log.Printf("result: %v", result)
	}
	return nil
}
