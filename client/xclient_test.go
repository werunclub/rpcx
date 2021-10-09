package client

import (
	"testing"
)

func TestXClient_filterByStateAndGroup(t *testing.T) {
	servers := map[string]string{"a": "", "b": "state=inactive&ops=10", "c": "ops=20", "d": "group=test&ops=20"}
	filterByStateAndGroup("test", servers)
	if _, ok := servers["b"]; ok {
		t.Error("has not remove inactive node")
	}
	if _, ok := servers["a"]; ok {
		t.Error("has not remove inactive node")
	}
	if _, ok := servers["c"]; ok {
		t.Error("has not remove inactive node")
	}
	if _, ok := servers["d"]; !ok {
		t.Error("node must be removed")
	}
}
