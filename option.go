package gurl

import "github.com/jessevdk/go-flags"

// Curl options.
type CurlOption struct {
	AnyAuth		bool	`long:"anyauth"`
	Append		bool	`short:"a" long:"append"`
	Basic		bool	`long:"basic"`
	Cacert		string	`long:"cacert"`
	CaPath		string	`long:"capath"`
}

// Gurl options.
type GurlOption struct {
	UrlEncode	bool	`short:"e" long:"encode" description:"Encode url and parameters(default: true)" default:"true"`
	File		string	`short:"f" long:"file" description:"File to read config"`
}
