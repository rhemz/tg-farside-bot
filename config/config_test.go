package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetConfig(t *testing.T) {
	// env vars aren't present
	assert.Panics(t, func() { GetConfig() }, "missing required env vars did not panic")

	// process env vars and test defaults
	os.Setenv("S3_BUCKET_NAME", "bucket.name")
	defer os.Unsetenv("S3_BUCKET_NAME")
	os.Setenv("S3_BUCKET_PATH", "bucket.path/")
	defer os.Unsetenv("S3_BUCKET_PATH")
	os.Setenv("S3_ENDPOINT", "https://test")
	defer os.Unsetenv("S3_ENDPOINT")
	os.Setenv("S3_KEY_ID", "s3key")
	defer os.Unsetenv("S3_KEY_ID")
	os.Setenv("S3_SECRET_KEY", "s3secret")
	defer os.Unsetenv("S3_SECRET_KEY")
	os.Setenv("TELEGRAM_BOT_TOKEN", "tg:key")
	defer os.Unsetenv("TELEGRAM_BOT_TOKEN")

	tests := []struct {
		name string
		want *Spec
	}{
		{name: "loaded-defaults", want: &Spec{
			EchoDebug:        true,
			ListenPort:       8080,
			LogJSON:          false,
			S3BucketName:     "bucket.name",
			S3BucketPath:     "bucket.path/",
			S3Endpoint:       "https://test",
			S3KeyId:          "s3key",
			S3SecretKey:      "s3secret",
			TelegramBotToken: "tg:key",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := GetConfig()
			assert.Exactly(t, x, tt.want)
		})
	}
}
