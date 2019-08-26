package maker

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

const handlerStub = `package <package>

import (
	"net/http"
)

func <handler>(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}`

type HandlerMaker struct {
	Path string
}

func (m HandlerMaker) makeDirPath(handler string) string {
	rootDir := strings.Split(m.Path, "/")
	handlerDirPieces := strings.Split(handler, ".")

	return filepath.Join(
		append(rootDir, handlerDirPieces[:len(handlerDirPieces)-1]...)...,
	)
}

// Make creates a new file with the contents of the stub.
func (m HandlerMaker) Make(name string) error {
	dir := m.makeDirPath(name)

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	handlerPieces := strings.Split(name, ".")
	pkg := handlerPieces[len(handlerPieces)-2]
	file := strings.Title(strings.ToLower(handlerPieces[len(handlerPieces)-1]))

	path := dir + "/" + file + ".go"

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	return m.writeContent(f, pkg, file)
}

func (m HandlerMaker) writeContent(f io.Writer, pkg, handler string) error {

	content := strings.Replace(handlerStub, "<package>", pkg, -1)
	content = strings.Replace(content, "<handler>", handler, -1)

	if _, err := f.Write([]byte(content)); err != nil {
		return err
	}

	return nil
}
