package ce_pubsub

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/config"
	"context"

	cepubsub "github.com/cloudevents/sdk-go/protocol/pubsub/v2"
)

var (
	Protocol *cepubsub.Protocol
)

func init() {
	container.InjectInstallation(func() error {
		projectId := config.GOOGLE_PROJECT_ID.Get()
		t, err := cepubsub.New(context.Background(),
			cepubsub.WithProjectID(projectId))
		Protocol = t
		return err
	})
}
