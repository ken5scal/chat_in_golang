package main

import (
	"errors"
	"fmt"
	"crypto/md5"
	"io"
	"strings"
)

type Avatar interface {
	GetAvatarURL(c *client) (string, error)
}
var ErrNoAvatarURL = errors.New("chat: Failed fetching avatart URL")

type AuthAvatar struct {}
//var UseAuthAvatar AuthAvatar

// receiver (_ AuthAvatar) means AuthAvatar will not be referenced within the method
func (_ AuthAvatar) GetAvatarURL(c *client) (string, error) {
	fmt.Println(c.userData["avatar_url"])
	if url, ok := c.userData["avatar_url"]; ok {
		fmt.Println(c.userData["avatar_url"])
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}
	return "", ErrNoAvatarURL
}

type GravatarAvatar struct {}
var UseGravatar GravatarAvatar
func (_ GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userid, ok := c.userData["userid"]; ok {
		if useridStr, ok := userid.(string); ok {
			m := md5.New()
			io.WriteString(m, strings.ToLower(useridStr))
			return fmt.Sprintf("//www.gravatar.com/avatar/%x", m.Sum(nil)), nil
		}
	}
	return "", ErrNoAvatarURL
}