package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_readConfig(t *testing.T) {
	testConfig := &config{ConfigPath: "testdata/fixture.yaml"}
	actual, err := testConfig.readConfig()
	assert.NoError(t, err)
	assert.Equal(t, &ConfigYaml{
		Slack: &ConfigSlack{
			Token: "test_slack_token",
		},
	}, actual)
}
