package randall

import (
	"errors"
	"net/url"
)

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func getUrlValues(providers ...QueryStringProvider) (url.Values, error) {
	if len(providers) == 0 {
		return url.Values{}, errors.New("getUrlValues() called with empty params slice")
	}

	values := make(url.Values)

	for _, provider := range providers {
		newValues, err := provider.AddQuery(values)

		if err != nil {
			return url.Values{}, err
		}

		values = newValues
	}

	return values, nil
}
