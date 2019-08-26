package maker

import (
	"bytes"
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

func TestWriteContents(t *testing.T) {
	var b bytes.Buffer
	m := HandlerMaker{Path: "/"}

	if err := m.writeContent(&b, "example", "ExampleHandler"); err != nil {
		t.Fatalf("Got error from writeContent: %s", err)
	}

	got := b.String()
	want := handlerTestStub
	if got != want {
		t.Errorf("Got  %s\nWant %s\n", got, want)
	}
}

const handlerTestStub = `package example

import (
	"net/http"
)

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}`
