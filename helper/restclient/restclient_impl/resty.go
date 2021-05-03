package restclient_impl

import (
	"github.com/go-resty/resty/v2"
	"selenium-check-awingu/helper/restclient"
)

type Resty struct {
	Url string
}

var client = resty.New()

type AuthSuccess struct {
	/* variables */
}
type AuthError struct {
	/* variables */
}

func NewResty(r Resty) restclient.RestClient {
	return &Resty{Url: r.Url}
}
