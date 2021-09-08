package balanced_http_client

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticRoundRobinBalancer_Select(t *testing.T) {
	url1, _ := url.Parse("http://1.example.com/upstream/")
	url2, _ := url.Parse("http://2.example.com/upstream/")
	balancer := &StaticRoundRobinBalancer{
		backends: []*url.URL{
			url1,
			url2,
		},
	}

	// 1st
	{
		got, err := balancer.Select(nil)
		assert.Equal(t, url1, got)
		assert.Nil(t, err)
	}
	// 2nd
	{
		got, err := balancer.Select(nil)
		assert.Equal(t, url2, got)
		assert.Nil(t, err)
	}
	// 3rd
	{
		got, err := balancer.Select(nil)
		assert.Equal(t, url1, got)
		assert.Nil(t, err)
	}
	// replace
	{
		balancer.SetBackends([]*url.URL{url2, url1})
		got, err := balancer.Select(nil)
		assert.Equal(t, url2, got)
		assert.Nil(t, err)
	}

}
