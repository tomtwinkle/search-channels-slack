package slacklib

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/tomtwinkle/attendance-client/config"
	"os"
	"testing"
)

func TestSlackClient_Action(t *testing.T) {
	if err := godotenv.Load("../.env"); err != nil {
		t.Log(".env not found")
		t.SkipNow()
	}
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		t.Error(".env SLACK_TOKEN not set")
		t.FailNow()
	}
	channelName := os.Getenv("SLACK_CHANNEL_NAME")
	if channelName == "" {
		t.Error(".env SLACK_CHANNEL_NAME not set")
		t.FailNow()
	}

	t.Run("success", func(t *testing.T) {
		ctx := context.Background()
		c := NewSlackClient(&config.ConfigSlack{
			Token: token,
			Channels: []config.ConfigSlackChannel{
				{
					Name:    channelName,
					ClockIn: &config.ConfigSlackAction{Message: "test!!"},
				},
			},
		})
		results, err := c.Action(ctx, ClockTypeClockIn)
		if !assert.NoError(t, err) {
			t.Errorf("%+v\n", err)
			t.FailNow()
		}
		t.Log(results)
	})
}
