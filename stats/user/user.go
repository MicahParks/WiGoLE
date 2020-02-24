package user

import (
	"net/url"
)

type Parameters struct {
}

type Response struct {
	Success       bool
	ImageBadgeUrl url.URL
}
