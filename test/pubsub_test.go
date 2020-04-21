package gcplocal_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"cloud.google.com/go/pubsub"
)

const (
	topicName = "topic-test"
	subID     = "subscid"
	msg       = "pubsub: ok"
)

func init() {
	os.Setenv("PUBSUB_EMULATOR_HOST", PubSubEmulatorHost)
}
func TestPubsub(t *testing.T) {
	// emulator非起動時とメッセージ受信の終了のため
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		t.Fatal(err)
	}

	topic, err := client.CreateTopic(ctx, topicName)
	if err != nil {
		// emulator非起動時に発生
		if err == context.DeadlineExceeded {
			t.Fatal("pubsub emulator is not available")
		}
		t.Fatal(err)
	}
	// topic,subsc削除用のcontext(実行されるときにctxは終了済みのため)
	delCtx := context.Background()
	defer topic.Delete(delCtx)

	subsc, err := client.CreateSubscription(ctx, subID, pubsub.SubscriptionConfig{Topic: topic})
	defer subsc.Delete(delCtx)
	if err != nil {
		t.Fatal(err)
	}

	// Pub
	if err := publish(ctx, topic); err != nil {
		t.Fatal(err)
	}

	// Sub
	if err := subscribe(ctx, subsc); err != nil {
		t.Fatal(err)
	}
}

func publish(ctx context.Context, topic *pubsub.Topic) error {
	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(fmt.Sprintf(msg)),
	})
	_, err := res.Get(ctx)
	return err
}

func subscribe(ctx context.Context, sub *pubsub.Subscription) error {
	var receivedMsg string
	err := sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		receivedMsg = string(m.Data)
		m.Ack()
	})
	if err == nil && receivedMsg != msg {
		err = errors.New("unexpected message '" + receivedMsg + "'")
	}
	return err
}
