package go_config

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetCurrentProfiles(t *testing.T) {
	parseProfiles("test,test1")
	assert.Equal(t, "test", GetCurrentProfiles()[0])
	assert.Equal(t, "test1", GetCurrentProfiles()[1])
}

func TestIsProfileActive(t *testing.T) {
	parseProfiles("test,test1")
	assert.True(t, IsProfileActive("test"))
	assert.True(t, IsProfileActive("test1"))
	assert.False(t, IsProfileActive("test2"))
}

func TestLoad(t *testing.T) {
	type Config struct {
		Value1 *string
		Value2 *string
	}
	val1 := "value1"
	val2 := "value2"
	config := Config{
		Value1: &val1,
		Value2: &val2,
	}
	d, err := json.Marshal(&config)
	if !assert.NoError(t, err) {
		return
	}

	err = ioutil.WriteFile("config.json", d, os.ModePerm)
	if !assert.NoError(t, err) {
		return
	}
	defer func() {
		os.Remove("config.json")
	}()

	var loaded Config
	err = Load(&loaded)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "value1", *loaded.Value1)
	assert.Equal(t, "value2", *loaded.Value2)
}

func TestLoadWithProfile(t *testing.T) {
	type Config struct {
		Value1 *string `json:"value_1,omitempty"`
		Value2 *string `json:"value_2,omitempty"`
	}
	val1 := "value1"
	val2 := "value2"
	config := Config{
		Value1: &val1,
		Value2: &val2,
	}

	// write config.json
	d, err := json.Marshal(&config)
	if !assert.NoError(t, err) {
		return
	}

	err = ioutil.WriteFile("config.json", d, os.ModePerm)
	if !assert.NoError(t, err) {
		return
	}
	defer func() {
		os.Remove("config.json")
	}()

	// write config-test.json
	val2 = "fromTestJson"
	config.Value1 = nil
	config.Value2 = &val2
	d, err = json.Marshal(&config)
	if !assert.NoError(t, err) {
		return
	}

	err = ioutil.WriteFile("config-TestLoadWithProfile.json", d, os.ModePerm)
	if !assert.NoError(t, err) {
		return
	}
	defer func() {
		os.Remove("config-TestLoadWithProfile.json")
	}()

	var loaded Config
	parseProfiles("TestLoadWithProfile")
	err = Load(&loaded)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "value1", *loaded.Value1)
	assert.Equal(t, "fromTestJson", *loaded.Value2)
}
