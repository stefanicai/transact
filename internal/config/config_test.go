package config

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoadFromFile_Success(t *testing.T) {
	cfg, err := LoadFromFile("./test_config.yaml")
	assert.Nil(t, err)
	require.NotNil(t, cfg)
	require.NotNil(t, cfg.Forex)
	assert.Equal(t, "http://test", cfg.Forex.TreasuryURL)
}

func TestLoadFromFile_ShouldNotCreateConfigIfMissingFile(t *testing.T) {
	_, err := LoadFromFile("/incorrect_file_path.yaml")
	assert.NotNil(t, err)
}
