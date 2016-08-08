package main

import (
	"testing"
	"path/filepath"
	"io/ioutil"
	"os"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	testUser := &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return("", ErrNoAvatarURL)
	testChatUser := &chatUser{User: testUser}
	url, err := authAvatar.GetAvatarURL(testChatUser)

	if err != ErrNoAvatarURL {
		t.Error("If No Value exists, AuthAvata.GetAvatarURL should return ErrorNoAvatarURL")
	}

	// URL has been set
	testUrl := "http://url-to-avatar/"
	testUser = &gomniauthtest.TestUser{}
	testChatUser.User = testUser
	testUser.On("AvatarURL").Return(testURL, nil)
	url, err = authAvatar.GetAvatarURL(testChatUser)
	if err != nil {
		t.Error("If Value Exists, GetAvatarURL should not return Error")
	} else {
		if url != testUrl {
			t.Error("GetAvatarURL should return ")
		}
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gAvatar GravatarAvatar
	user := &chatUser{uniqueID: "abc"}
	url, err := gAvatar.GetAvatarURL(user)

	if err != nil {
		t.Error("GravatarAvatar.GetAvatarURL should not return Error")
	}
	if url != "//www.gravatar.com/avatar/" + "abc" {
		t.Errorf("Wrong url: %s", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {
	// Create avatar file for test
	filename := filepath.Join("avatars", "abc.jpg")
	ioutil.WriteFile(filename, []byte{}, 0777)
	defer func() {os.Remove(filename)}()

	var fileSystemAvatar FileSystemAvatar
	user := &chatUser{uniqueID: "abc"}
	url, err := fileSystemAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("FileSystemAdaptar.GetAvatarURL Should not return error")
	}
	if url != "/avatars/abc.jpg" {
		t.Errorf("Wrong URL returned : %s", url)
	}
}