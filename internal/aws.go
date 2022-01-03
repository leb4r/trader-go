package internal

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SnsMessage struct {
	Type  string            `json:"type"`
	Title string            `json:"title"`
	Model map[string]string `json:"model"`
}

func CreateAwsSession(region string) (*session.Session, error) {
	if sess, err := session.NewSession(&aws.Config{Region: aws.String(region)}); err != nil {
		return sess, nil
	} else {
		return nil, nil
	}
}

func CreateTopicArn(topicName string) string {
	return "arn:aws:sns:account-id:topicName"
}

func PublishToSnsTopic(message SnsMessage, topic string) error {
	snsMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	region := "SNS_REGION"
	topicArn := CreateTopicArn("testing")

	sess, err := CreateAwsSession(region)
	if err != nil {
		return err
	}

	svc := sns.New(sess)
	params := &sns.PublishInput{
		Message:  aws.String(string(snsMessage)),
		TopicArn: aws.String(topicArn),
	}

	_, err = svc.Publish(params)
	if err != nil {
		return err
	}

	return nil
}
