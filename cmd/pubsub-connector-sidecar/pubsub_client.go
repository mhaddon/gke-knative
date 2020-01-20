package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"sync"
)

type pubsubClient struct {
	context           context.Context
	contextWithCancel context.Context
	client            *pubsub.Client
	subscription      *pubsub.Subscription
	egressTopic       *pubsub.Topic
	lock              sync.Mutex
	cancel            context.CancelFunc
}

var (
	pubsubClientInstance *pubsubClient
	pubsubClientOnce     sync.Once
)

func getPubsubClient() *pubsubClient {
	pubsubClientOnce.Do(func() {
		config := getConfig()

		psClient, err := pubsub.NewClient(context.Background(), config.ProjectId); if err != nil {
			panic("unable to create pubsub client: " + err.Error())
		}

		cctx, cancel := context.WithCancel(context.Background())

		pubsubClientInstance = &pubsubClient{
			context:           context.Background(),
			contextWithCancel: cctx,
			cancel:            cancel,
			client:            psClient,
			subscription:      psClient.Subscription(config.SubscriptionId),
			egressTopic:       psClient.TopicInProject(config.EgressTopicName, config.ProjectId),
		}
	})

	return pubsubClientInstance
}
