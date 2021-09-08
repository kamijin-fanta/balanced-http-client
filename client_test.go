package balanced_http_client

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
	}))
	defer ts.Close()

	client := NewClient()
	res, err := client.Get(ts.URL)
	require.Nil(t, err)
	content, err := ioutil.ReadAll(res.Body)
	require.Nil(t, err)
	require.Equal(t, []byte("Hello"), content)
}
