package utils

import (
	"github.com/jezman/yapdd/pdd"
	"github.com/levigross/grequests"
)

var (
	// Headers for requests
	Headers = map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
		"PddToken":     pdd.Token,
	}

	// RequestOptions include headers and body params
	RequestOptions = &grequests.RequestOptions{
		Headers: Headers,
		Params:  map[string]string{},
	}
)
