package link

import (
	"context"
	"strings"
)

type ProviderId int

const (
	Yandex ProviderId = iota + 1
	Apple
	Spotify
	Youtube
	Vk
)

var ProviderMap = map[ProviderId][]string{
	Yandex:  []string{"yandex.ru"},
	Apple:   []string{"apple.co"},
	Spotify: []string{"spotify.com"},
	Youtube: []string{"youtube.com"},
	Vk:      []string{"vk.com"},
}

type Link struct {
	Id         int    `db:"id"`
	Url        string `db:"url"`
	ProviderId ProviderId
}

type RsLink struct {
	Id         int        `json:"id"`
	Url        string     `json:"url"`
	ProviderId ProviderId `json:"provId"`
}

type Repository interface {
	LinksFor(ctx context.Context, id int) ([]Link, error)
}

func detectProvider(url string) ProviderId {
	for key, value := range ProviderMap {
		for _, val := range value {
			if strings.Contains(url, val) {
				return key
			}
		}
	}
	return 0
}
