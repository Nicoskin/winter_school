package unionfind

import (
	"fmt"
	"testing"
)

// go test -v graph.go fs_test.go fs.go
func TestMain2(t *testing.T) {
	ds := NewDisjointSet(10)
	fmt.Println(ds.Find(1))
}
