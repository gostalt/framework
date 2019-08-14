package maker

import (
	"io"
	"os"
	"strings"
)

// entityStub is the basic implementation of an entity.
//
// TODO: The go generate code here isn't great as it relies on shell.
// However, the generate code is called relatively, and invoking the
// gostalt binary with ../../gostalt doesn't pull in env files. Hmm...
const entityStub = `package entity

//go:gen sh -c "cd ../../ && go run main.go migrate magic $GOFILE"

type <Entity> struct {
	// Fields here
}

// Methods here
`

// EntityMaker creates a stub of an entity.
type EntityMaker struct {
	Path string
}

// Make creates a new file with the contents of the stub.
func (m EntityMaker) Make(name string) error {
	entity := strings.Title(strings.ToLower(name))

	filepath := m.Path + entity + ".go"

	f, err := os.Create(filepath)
	if err != nil {
		return err
	}

	return m.writeContent(f, entity)
}

func (m EntityMaker) writeContent(f io.Writer, entity string) error {
	content := strings.Replace(entityStub, "<Entity>", entity, -1)

	// Annoyingly, running `go generate ./...` will trigger the
	// command in the entityStub, so we must replace it here.
	content = strings.Replace(content, "go:gen", "go:generate", -1)
	if _, err := f.Write([]byte(content)); err != nil {
		return err
	}

	return nil
}
