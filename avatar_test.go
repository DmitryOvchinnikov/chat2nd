package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	gomniauthtest "github.com/stretchr/gomniauth/test"
)

func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar

	testUser := &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return("", ErrNoAvatarURL)

	testChatUser := &chatUser{testUser}

	url, err := authAvatar.GetAvatarURL(testChatUser)
	if err != ErrNoAvatarURL {
		t.Error("AuthAvatar.GetAvatarURL should return ErrNoAvatarURL when no value present")
	}

	// set a value
	testURL := "http://url-to-gravatar/"
	testUser = &gomniauthtest.TestUser{}
	testUser.On("AvatarURL").Return(testURL, nil)

	url, err = authAvatar.GetAvatarURL(testChatUser)
	if err != nil {
		t.Error("AuthAvatar.GetAvatarURL should return no error when value present")
	}
	if url != testURL {
		t.Error("AuthAvatar.GetAvatarURL should return correct URL")
	}

}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar

	user := &chatUser{uniqueID: "abc"}

	url, err := gravatarAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("GravatarAvatar.GetAvatarURL should not return an error")
	}
	if url != "//www.gravatar.com/avatar/abc" {
		t.Errorf("GravatarAvatar.GetAvatarURL wrongly returned %s", url)
	}

}

func TestFileSystemAvatar(t *testing.T) {

	// make a test avatar file
	filename := filepath.Join("avatars", "abc.jpg")
	ioutil.WriteFile(filename, []byte{}, 0777)
	defer func() {os.Remove(filename)}()

	var fileSystemAvatar FileSystemAvatar

	user := &chatUser{uniqueID: "abc"}

	url, err := fileSystemAvatar.GetAvatarURL(user)
	if err != nil {
		t.Error("FileSystemAvatar.GetAvatarURL should not return an error")
	}
	if url != "/avatars/abs.jpg" {
		t.Errorf("FileSystemAvatar.GetAvatarURL wrongly returned %s", url)
	}

}
