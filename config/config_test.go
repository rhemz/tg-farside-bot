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
	os.Setenv("BUCKET_PATH", "bucket.path/")
	defer os.Unsetenv("BUCKET_PATH")
	os.Setenv("S3_KEY_ID", "s3key")
	defer os.Unsetenv("S3_KEY_ID")
	os.Setenv("S3_SECRET_KEY", "s3secret")
	defer os.Unsetenv("S3_SECRET_KEY")
	os.Setenv("TELEGRAM_BOT_TOKEN", "tg:key")

	tests := []struct {
		name string
		want *Spec
	}{
		{name: "loaded-defaults", want: &Spec{
			BucketPath:       "bucket.path/",
			EchoDebug:        true,
			ListenPort:       8080,
			LogJSON:          false,
			S3KeyId:          "s3key",
			S3SecretKey:      "s3secret",
			TelegramBotToken: "tg:key",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Exactly(t, GetConfig(), tt.want)
		})
	}
}
