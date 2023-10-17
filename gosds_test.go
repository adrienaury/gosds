package gosds_test

import (
	"fmt"
	"testing"

	"github.com/adrienaury/gosds"
)

func TestTest(t *testing.T) {
	t.Parallel()

	root := gosds.NewObject()

	root.SetValueForKey("name", "adrien")

	fmt.Println(root.NodeForKey("name").Parent())
}
