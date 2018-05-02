package bios

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiscoveryDir(t *testing.T) {
	ts := newFileServer()
	defer ts.Close()

	net := NewNetwork("/tmp/disco", ts.URL+"/bp1.yaml", &IPFS{})
	assert.NoError(t, net.traverseGraph())
	assert.NoError(t, net.verifyGraph())
	assert.NoError(t, net.calculateWeights())
	// assert.Equal(t, 0.5, net.discoveredPeers[ts.URL+"/bp1.yaml"].TotalWeight)
	// assert.Equal(t, 1.0, net.discoveredPeers[ts.URL+"/bp2.yaml"].TotalWeight)
	// assert.Equal(t, 1.0, net.discoveredPeers[ts.URL+"/bp3.yaml"].TotalWeight)
	// assert.Equal(t, ts.URL+"/bp2.yaml", net.orderedPeers[0].DiscoveryRef)
	// assert.Equal(t, ts.URL+"/bp3.yaml", net.orderedPeers[1].DiscoveryRef)
	// assert.Equal(t, ts.URL+"/bp1.yaml", net.orderedPeers[2].DiscoveryRef)
}

func newFileServer() *httptest.Server {
	curDir, _ := os.Getwd()
	ts := httptest.NewServer(http.FileServer(http.Dir(filepath.Join(curDir, "test-data"))))
	return ts
}
