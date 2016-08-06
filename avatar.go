package main

import "errors"

type Avatar interface {
	GetAvatarURL(c *client) (string, error)
}
var ErrNoAvatarURL = errors.New("chat: Failed fetching avatart URL")