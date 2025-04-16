package torrent

import (
	"net/url"
)

type Torrent struct {
	Announce url.URL `json:"announce"`
	Info     Info    `json:"info"`
}

type Info struct {
	Name    string  `json:"name"`
	PLength uint32  `json:"piece length"`
	Pieces  []uint8 `json:"pieces"`
	Key     Key     `json:"key"`
}

type Key interface {
	Type() string
}

type SingleFileKey struct {
	Length uint32 `json:"length"`
}

func (s SingleFileKey) Type() string {
	return "SingleFile"
}

type MultiFileKey struct {
	File []File `json:"files"`
}

func (s MultiFileKey) Type() string {
	return "MultiFile"
}

type File struct {
	Length uint32   `json:"length"`
	Path   []string `json:"path"`
}
