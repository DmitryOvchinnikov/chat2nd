package main

import (
	"errors"
	"io/ioutil"
	"path"
)

// ErrNoAvatarURL is the error that is returned when the
// Avatar instance is unable to provide an avatar URL.
var ErrNoAvatarURL = errors.New("chat: unable to get an avatar URL")

// Avatar represents types capable of representing
// user profile pictures.
type Avatar interface {
	// GetAvatarURL gets the avatar URL for the specified client,
	// or returns an error if something goes wrong.
	// ErrNoAvatarURL is returned if the object is unable to get
	// a URL for the specified client.
	GetAvatarURL(ChatUser) (string, error)
}

type AuthAvatar struct {}

var UseAuthAvatar AuthAvatar

func (AuthAvatar) GetAvatarURL(c *client) (string, error) {
	if url, ok := c.userData["avatar_url"]; ok {
		if urlStr, ok := url.(string); ok {
			return urlStr, nil
		}
	}

	return "", ErrNoAvatarURL
}

type GravatarAvatar struct {}

var UseGravatar GravatarAvatar

func (GravatarAvatar) GetAvatarURL(c *client) (string, error) {
	if userID, ok := c.userData["userid"]; ok {
		if userIDStr, ok := userID.(string); ok {
			return "//www.gravatar.com/avatar/" + userIDStr, nil
		}
	}

	return "", ErrNoAvatarURL
}

type FileSystemAvatar struct {}

var UseFileSystemAvatar FileSystemAvatar

func (FileSystemAvatar) GetAvatarURL(c *client) (string, error) {
	if userID, ok := c.userData["userid"]; ok {
		if userIDStr, ok := userID.(string); ok {
			files, err := ioutil.ReadDir("avatars")
			if err != nil {
				return "", ErrNoAvatarURL
			}

			for _, file := range files {
				if file.IsDir() {
					continue
				}
				if match, _ := path.Match(userIDStr + "*", file.Name()); match {
					return "/avatars/" + file.Name(), nil
				}
			}
		}
	}

	return "", ErrNoAvatarURL
}
