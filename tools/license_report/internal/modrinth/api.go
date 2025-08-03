package modrinth

import (
	"time"

	"github.com/ghifari160/zei"
)

const apiURL = "https://api.modrinth.com/"

type Config struct {
	UserAgent string
	Timeout   time.Duration
}

type API struct {
	client *zei.Client
}

func (API) v2() string {
	return apiURL + "v2"
}

func New(conf Config) *API {
	client := zei.New(&zei.Config{
		UserAgent: conf.UserAgent,
		Timeout:   conf.Timeout,
	})

	return &API{
		client: client,
	}
}
