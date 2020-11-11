package client

import (
	"encoding/json"
	"io/ioutil"
	"sync"

	"github.com/werunclub/rpcx/log"
)

// StaticDiscovery is a static service discovery.
// It always returns the registered servers in static yaml file.
type StaticDiscovery struct {
	basePath    string
	servicePath string
	configPath  string

	pairs []*KVPair
	chans []chan []*KVPair
	mu    sync.Mutex

	stopCh chan struct{}
}

// NewStaticDiscovery returns a new StaticDiscovery.
func NewStaticDiscovery(basePath, servicePath, configPath string) ServiceDiscovery {
	discovery := &StaticDiscovery{
		basePath:    basePath,
		servicePath: servicePath,
		configPath:  configPath,
	}

	// format:  {"serviceName": ["addr1", "addr1"]}
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Panicf("read config file error: %v", err)
	}

	servicesData := make(map[string][]string)
	if err := json.Unmarshal(data, &servicesData); err != nil {
		log.Panicf("parse config file error: %v", err)
	}

	// find service
	for k, pairs := range servicesData {
		if k == servicePath {
			for _, v := range pairs {
				discovery.pairs = append(discovery.pairs, &KVPair{
					Key:   v,
					Value: "tps=",
				})

				log.Infof("found server for %s: %s", servicePath, v)
			}
			break
		}
	}

	return discovery
}

// NewStaticDiscoveryTemplate returns a new StaticDiscovery template.
func NewStaticDiscoveryTemplate(basePath, configPath string) ServiceDiscovery {
	return &StaticDiscovery{
		basePath:   basePath,
		configPath: configPath,
	}
}

// Clone clones this ServiceDiscovery with new servicePath.
func (d *StaticDiscovery) Clone(servicePath string) ServiceDiscovery {
	return NewStaticDiscovery(d.basePath, servicePath, d.configPath)
}

// GetServices returns the servers
func (d StaticDiscovery) GetServices() []*KVPair {
	return d.pairs
}

// WatchService returns a nil chan.
func (d *StaticDiscovery) WatchService() chan []*KVPair {
	ch := make(chan []*KVPair, 10)
	d.chans = append(d.chans, ch)
	return ch
}

func (d *StaticDiscovery) RemoveWatcher(ch chan []*KVPair) {
	d.mu.Lock()
	defer d.mu.Unlock()

	var chans []chan []*KVPair
	for _, c := range d.chans {
		if c == ch {
			continue
		}

		chans = append(chans, c)
	}

	d.chans = chans
}

func (d *StaticDiscovery) Close() {
	close(d.stopCh)
}

func (d *StaticDiscovery) load() {
	close(d.stopCh)
}
