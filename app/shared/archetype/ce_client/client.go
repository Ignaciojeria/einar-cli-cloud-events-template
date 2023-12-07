package ce_client

import (
	"archetype/app/shared/archetype/container"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

var R cloudevents.Client
var _ = container.InjectInboundAdapter(func() error {
	c, err := cloudevents.NewClientHTTP()
	R = c
	return err
})
