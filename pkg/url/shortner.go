package url

import (
	"fmt"
	"microUrl/pkg/hash"
	"microUrl/pkg/utils"
)

func Shrink(url string, length uint, usePadding bool) (string, error) {
	if utils.CheckEmpty(url) {
		return "", fmt.Errorf("url is empty")
	}

	if len(url) < int(length) {
		return url, nil
	}

	u := newUrl(url, length, usePadding)
	return u.short()
}

type url struct {
	original string
	len      uint
	usePad   bool
	sha      *hash.SHA
}

func newUrl(link string, length uint, usePadding bool) *url {
	u := &url{
		original: link,
		len:      length,
		usePad:   usePadding,
	}

	if usePadding {
		u.sha = hash.New(hash.WithPadding())
	} else {
		u.sha = hash.New()
	}

	return u
}

func (u *url) short() (string, error) {
	h, err := u.sha.Hash(u.original)
	if err != nil {
		return "", err
	}

	h = h[:u.len]
	if u.usePad {
		return utils.Base64EncodeWithPadding(h, '='), nil
	}

	return utils.Base64Encode(h), nil
}
