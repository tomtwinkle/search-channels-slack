package slacklib

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/slack-go/slack"
	"github.com/tomtwinkle/search-channels-slack/config"
	"github.com/tomtwinkle/search-channels-slack/options/channel"
	"github.com/tomtwinkle/search-channels-slack/types"
)

type ClockType int

const (
	ClockTypeClockIn ClockType = iota + 1
	ClockTypeClockOut
	ClockTypeGoOut
	ClockTypeReturned
)

type SlackClient interface {
	PostMessage(ctx context.Context, channelName, message string) error
}

type slackClient struct {
	config *config.ConfigYaml
	client *slack.Client
}

type SlackResult struct {
	ChannelId   string
	ChannelName string
	Timestamp   string
}

func NewSlackClient(cfg *config.ConfigYaml) SlackClient {
	if cfg == nil || cfg.Slack == nil || cfg.Slack.Token == "" {
		return &slackClient{}
	}
	return &slackClient{config: cfg, client: slack.New(cfg.Slack.Token)}
}

func (s slackClient) PostMessage(ctx context.Context, channelName, message string) error {
	// TODO validation
	if channelName == "" || message == "" {
		return nil
	}
	ch, err := s.SearchChannel(ctx, channelName)
	if err != nil {
		return err
	}

	// returned channelID, timestamp, err
	if _, _, err := s.client.PostMessageContext(
		ctx,
		ch.ID,
		slack.MsgOptionText(message, false),
		slack.MsgOptionAsUser(true),
	); err != nil {
		return err
	}
	return nil
}

func (s slackClient) SearchChannel(ctx context.Context, channelName string, options ...channel.Option) (*slack.Channel, error) {
	var opt channel.ChannelOption
	// default
	opt.Types = []types.ChannelType{types.ChannelTypePrivate, types.ChannelTypePublic}
	for _, o := range options {
		o(&opt)
	}
	channelTypes := make([]string, len(opt.Types))
	for i, t := range opt.Types {
		channelTypes[i] = string(t)
	}

	var cursor string
	for {
		requestParam := &slack.GetConversationsParameters{
			Types:           channelTypes,
			Limit:           1000,
			ExcludeArchived: "true",
		}
		if cursor != "" {
			requestParam.Cursor = cursor
		}
		var channels []slack.Channel
		var err error
		channels, cursor, err = s.client.GetConversationsContext(ctx, requestParam)
		if err != nil {
			return nil, err
		}
		channelRef := &channels
		for i := 0; i < len(channels); i++ {
			ch := (*channelRef)[i]
			if ch.Name == channelName {
				return &ch, nil
			}
		}
		if cursor == "" {
			break
		}
	}
	return nil, errors.New(fmt.Sprintf("channel not found. channelName=%s", channelName))
}
