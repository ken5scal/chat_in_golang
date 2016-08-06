package main

import "testing"

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)

	// No URL has been set
	url, err := authAvatar.GetAvatarURL(client);
	if err != ErrNoAvatarURL {
		t.Error("If No Value exists, AuthAvata.GetAvatarURL should return ErrorNoAvatarURL")
	}

	// URL has been set
	testUrl := "http://url-to-avatar/"
	client.userData = map[string]interface{}{"avatar_url": testUrl}
	url, err = authAvatar.GetAvatarURL(client); if err != nil {
		t.Error("If Value Exists, GetAvatarURL should not return Error")
	} else {
		if url != testUrl {
			t.Error("GetAvatarURL should return ")
		}
	}

}
