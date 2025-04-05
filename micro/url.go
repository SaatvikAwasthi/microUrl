package micro

import "microUrl/pkg/url"

type URL interface {
	Shrink(url string) (string, error)
}

type microURL struct {
	length   uint
	isPadded bool
}

func New(length uint, isPadded bool) URL {
	return &microURL{
		length:   length,
		isPadded: isPadded,
	}
}

func Default() URL {
	return &microURL{
		length:   5,
		isPadded: false,
	}
}

func (m *microURL) Shrink(link string) (string, error) {
	return url.Shrink(link, m.length, m.isPadded)
}
