package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

// StaticDiscovery is a static service discovery.
// It always returns the registered servers in static yaml file.
type StaticDiscovery struct {
	basePath    string
	servicePath string
	pairs       []*KVPair
	chans       []chan []*KVPair
	mu          sync.Mutex

	stopCh chan struct{}
}

// NewStaticDiscovery returns a new StaticDiscovery.
func NewStaticDiscovery(basePath, servicePath, configFile string) ServiceDiscovery {
	discovery := &StaticDiscovery{
		basePath:    basePath,
		servicePath: servicePath,
	}

	// format:  [{Key: "", Value: ""}, {Key: "", Value: ""}]
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal(err)
	}

	discovery.pairs = make([]*KVPair, 0, 1)
	if err := json.Unmarshal(data, &discovery.pairs); err != nil {
		log.Fatal(err)
	}

	return discovery
}

// Clone clones this ServiceDiscovery with new servicePath.
func (d *StaticDiscovery) Clone(servicePath string) ServiceDiscovery {
	return d
}

// GetServices returns the servers
func (d StaticDiscovery) GetServices() []*KVPair {
	return d.pairs
}

// WatchService returns a nil chan.
func (d *StaticDiscovery) WatchService() chan []*KVPair {
	return nil
}

func (d *StaticDiscovery) RemoveWatcher(ch chan []*KVPair) {
}

func (d *StaticDiscovery) Close() {
}
