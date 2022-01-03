package internal

import (
	"context"

	"github.com/leb4r/trader-go/internal/models"
)

type PubSubClient interface {
	// Create makes a new optic and returns its ARN
	Create(ctx context.Context, topic string) (string, error)
	// ListTopics lists all of the Topics
	ListTopics(ctx context.Context) ([]*models.Topic, error)
	// Subscribe adds a subscription to a topic
	Subscribe(ctx context.Context, endpoint, protocol, topicARN string) (string, error)
	// ListTopicSubscriptions returns a list of all subscriptions of a given topicARN
	ListTopicSubscriptions(ctx context.Context, topicARN string) ([]*models.TopicSubscription, error)
	// Publish writes a message to the topic
	Publish(ctx context.Context, message, topicARN string) (string, error)
	// Unsubscribe removes a subscription from a topic
	Unsubscribe(ctx context.Context, subscriptionARN string) error
}
