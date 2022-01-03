package aws

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/leb4r/trader-go/internal"
	"github.com/leb4r/trader-go/internal/models"
)

var _ internal.PubSubClient = SNS{}

type SNS struct {
	timeout time.Duration
	client  *sns.SNS
}

func NewSnsSession(session *session.Session, timeout time.Duration) SNS {
	return SNS{
		timeout: timeout,
		client:  sns.New(session),
	}
}

func (s SNS) Create(ctx context.Context, topic string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	res, err := s.client.CreateTopicWithContext(ctx, &sns.CreateTopicInput{
		Name: aws.String(topic),
	})
	if err != nil {
		return "", err
	}

	return *res.TopicArn, nil
}

func (s SNS) ListTopics(ctx context.Context) ([]*models.Topic, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	res, err := s.client.ListTopicsWithContext(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("list topics: %w", err)
	}

	topics := make([]*models.Topic, len(res.Topics))

	for i, topic := range res.Topics {
		topics[i] = &models.Topic{
			ARN: *topic.TopicArn,
		}
	}

	return topics, nil
}

func (s SNS) Subscribe(ctx context.Context, endpoint, protocol, topicARN string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	res, err := s.client.SubscribeWithContext(ctx, &sns.SubscribeInput{
		Endpoint:              aws.String(endpoint),
		Protocol:              aws.String(protocol),
		ReturnSubscriptionArn: aws.Bool(true),
		TopicArn:              aws.String(topicARN),
	})
	if err != nil {
		return "", err
	}

	return *res.SubscriptionArn, nil
}

func (s SNS) ListTopicSubscriptions(ctx context.Context, topicARN string) ([]*models.TopicSubscription, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	res, err := s.client.ListSubscriptionsByTopicWithContext(ctx, &sns.ListSubscriptionsByTopicInput{
		NextToken: nil,
		TopicArn:  aws.String(topicARN),
	})
	if err != nil {
		return nil, fmt.Errorf("list topic subscriptions: %w", err)
	}

	subs := make([]*models.TopicSubscription, len(res.Subscriptions))

	for i, sub := range res.Subscriptions {
		subs[i] = &models.TopicSubscription{
			ARN:      *sub.SubscriptionArn,
			TopicARN: *sub.TopicArn,
			Endpoint: *sub.Endpoint,
			Protocol: *sub.Protocol,
		}
	}

	return subs, nil
}

func (s SNS) Publish(ctx context.Context, message, topicARN string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	res, err := s.client.PublishWithContext(ctx, &sns.PublishInput{
		Message:  &message,
		TopicArn: aws.String(topicARN),
	})
	if err != nil {
		return "", fmt.Errorf("publish: %w", err)
	}

	return *res.MessageId, nil
}

func (s SNS) Unsubscribe(ctx context.Context, subscriptionARN string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	if _, err := s.client.UnsubscribeWithContext(ctx, &sns.UnsubscribeInput{
		SubscriptionArn: aws.String(subscriptionARN),
	}); err != nil {
		return fmt.Errorf("unsubscribe: %w", err)
	}

	return nil
}
