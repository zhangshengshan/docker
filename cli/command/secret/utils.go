package secret

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
)

func getSecretsByName(client client.APIClient, ctx context.Context, names []string) ([]swarm.Secret, error) {
	args := filters.NewArgs()
	for _, n := range names {
		args.Add("names", n)
	}

	return client.SecretList(ctx, types.SecretListOptions{
		Filters: args,
	})
}
