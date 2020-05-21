package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets950a25363eee220f7d8ce234bcc0b349e4ea9072 = "<!doctype html>\r\n<body>\r\n  <p>Hello, {{.Foo}}</p>\r\n</body>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"html"}, "/html": []string{"index.html"}}, map[string]*assets.File{
	"/html/index.html": &assets.File{
		Path:     "/html/index.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1590029610, 1590029610562858200),
		Data:     []byte(_Assets950a25363eee220f7d8ce234bcc0b349e4ea9072),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1589885180, 1589885180576114000),
		Data:     nil,
	}, "/html": &assets.File{
		Path:     "/html",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1589885187, 1589885187104049800),
		Data:     nil,
	}}, "")
