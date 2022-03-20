package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Something(t *testing.T) {

	config := GenerateConfig()
	assert.Equal(t, config.Format, "json", "expected valid struct")

}
