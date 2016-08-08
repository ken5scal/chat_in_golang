package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
)

type Avatar interface {
	GetAvatarURL(ChatUser) (string, error)
}

type TryAvatars []Avatar

func (a TryAvatars) GetAvatarURL(u ChatUser) (string, error) {
	for _, avatar := range a {
		if url, err := avatar.GetAvatarURL(u); err == nil {
			return url, nil
		}
	}
	return "", ErrNoAvatarURL
}

var ErrNoAvatarURL = errors.New("chat: Failed fetching avatart URL")

type AuthAvatar struct {}
var UseAuthAvatar AuthAvatar

func (AuthAvatar) GetAvatarURL(u ChatUser) (string, error) {
	url := u.AvatarURL()
	if len(url) == 0 {
		return "", ErrNoAvatarURL
	}
	return u.AvatarURL(), nil
}

type GravatarAvatar struct {}
var UseGravatar GravatarAvatar
func (_ GravatarAvatar) GetAvatarURL(u ChatUser) (string, error) {
	//if userid, ok := c.userData["userid"]; ok {
	//	if useridStr, ok := userid.(string); ok {
	//		return "//www.gravatar.com/avatar/" + useridStr, nil
	//	}
	//}
	//return "", ErrNoAvatarURL
	return "//www.gravatar.com/avatar/" + u.UniqueID(), nil
}


type FileSystemAvatar struct{}

var UseFileSystemAvatar FileSystemAvatar

func (FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
	files, err := ioutil.ReadDir("avatars")
	if err != nil {
		return "", ErrNoAvatarURL
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if match, _ := filepath.Match(u.UniqueID() + "*", file.Name()); match {
			return "/avatars/" + file.Name(), nil
		}
		//fname := file.Name()
		//if u.UniqueID() == strings.TrimSuffix(fname, filepath.Ext(fname)) {
		//	return "/avatars/" + fname, nil
		//}
	}
	return "", ErrNoAvatarURL
}

//func (_ FileSystemAvatar) GetAvatarURL(u ChatUser) (string, error) {
//
//
//
//	if userid, ok := c.userData["userid"]; ok {
//		if useridStr, ok := userid.(string); ok {
//			if files, err := ioutil.ReadDir("avatars"); err == nil {
//				for _, file := range files {
//					if file.IsDir() {
//						continue
//					}
//					if match, _ := filepath.Match(useridStr + "*", file.Name());
//					match{
//						return "/avatars/" + file.Name(), nil
//					}
//				}
//			}
//		}
//	}
//	return "", ErrNoAvatarURL
//}
