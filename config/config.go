package config

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
)

type (
	Spec struct {
		BucketPath       string `required:"true" split_words:"true"`
		EchoDebug        bool   `default:"true"  split_words:"true"`
		ListenPort       int    `default:"8080"  split_words:"true"`
		LogJSON          bool   `default:"false" split_words:"true"`
		S3KeyId          string `required:"true" split_words:"true"`
		S3SecretKey      string `required:"true" split_words:"true"`
		TelegramBotToken string `required:"true" split_words:"true"`
	}
)

var (
	cfg Spec
)

func GetConfig() *Spec {
	if cfg == (Spec{}) {
		err := envconfig.Process("", &cfg)
		if err != nil {
			panic(fmt.Sprintf("error parsing config: %s", err))
		}
	}

	return &cfg
}
