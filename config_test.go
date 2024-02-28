package main

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// This will examine the config global, and ensure the values match config.yaml
func TestLoadConfig(t *testing.T) {

	assert.IsType(t, globalConfig{}, config, "type check")

	if reflect.DeepEqual(config, new(globalConfig)) {
		t.Error("config is empty.")
	}

	assert.False(t, config.Collector.Debug, "Configuration - Collector.Debug")
	assert.False(t, config.Collector.SampleData, "Configuration - Collector.SampleData")
	assert.NotEmpty(t, config.Collector.SampleDataPath, "Configuration - Collector.SampleDataPath")

	assert.NotEmpty(t, config.OneView.Domain, "Configuration - OneView.Domain")
	assert.NotZero(t, len(config.OneView.Hosts), "Configuration - OneView.Hosts")
	assert.NotEmpty(t, config.OneView.Pass, "Configuration - OneView.Pass")
	assert.NotEmpty(t, config.OneView.User, "Configuration - OneView.User")

}
