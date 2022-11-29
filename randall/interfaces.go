package randall

import "net/url"

type QueryStringProvider interface {
	AddQuery(v url.Values) (url.Values, error)
}
