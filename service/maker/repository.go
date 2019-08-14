package maker

import (
	"io"
	"os"
	"strings"
)

// repositoryStub is the basic implementation of a repository.
const repositoryStub = `package repository

import (
	"gostalt/app/entity"

	"github.com/jmoiron/sqlx"
)

type <Repository> struct {
	*sqlx.DB
}

func (r <Repository>) Fetch(id int) (entity.<Repository>, error) {
	<repository> := entity.<Repository>{}
	err := r.Get(&<repository>, "select * from <repository>s where id = $1 limit 1", id)
	if err != nil {
		r.Logger.Warning([]byte(err.Error()))
		return <repository>, err
	}

	return <repository>, nil
}

func (r <Repository>) FetchAll() []entity.<Repository> {
	<repository>s := []entity.<Repository>{}
	<repository>s, err := r.Select(&<repository>s, "select * from <repository>s")
	if err != nil {
		return []entity.<Repository>{}
	}

	return <repository>s
}
`

// RepositoryMaker creates a stub of a repository for interacting
// with a database for a particular entity.
type RepositoryMaker struct {
	Path string
}

// Make creates a new file with the contents of the stub.
func (m RepositoryMaker) Make(name string) error {
	repository := strings.Title(strings.ToLower(name))

	filepath := m.Path + repository + ".go"

	f, err := os.Create(filepath)
	if err != nil {
		return err
	}

	return m.writeContent(f, repository)
}

func (m RepositoryMaker) writeContent(f io.Writer, repository string) error {

	content := strings.Replace(repositoryStub, "<repository>", strings.ToLower(repository), -1)
	content = strings.Replace(content, "<Repository>", repository, -1)

	if _, err := f.Write([]byte(content)); err != nil {
		return err
	}

	return nil
}
