package maker

import (
	"fmt"
	"testing"
)

func TestMakeDirPath(t *testing.T) {
	m := HandlerMaker{Path: "a/long/path/"}

	path := m.makeDirPath("nested.to.handler")
	expected := "a/long/path/nested/to"

	if path != expected {
		fmt.Printf("Got  %s\nWant %s\n", path, expected)
		t.FailNow()
	}
}
