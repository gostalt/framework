package service

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gostalt/framework/route"
	"github.com/gostalt/router"
	"github.com/sarulabs/di/v2"
)

func NewViewServiceProvider(path string, cached bool) *viewServiceProvider {
	return &viewServiceProvider{
		path:   path,
		cached: cached,
	}
}

type viewServiceProvider struct {
	path   string
	cached bool
}

func (provider viewServiceProvider) Register(builder *di.Builder) error {
	if err := builder.Add(di.Def{
		Name: "views",
		Build: func(c di.Container) (interface{}, error) {
			return provider.load(provider.path), nil
		},

		Unshared: !provider.cached,
	}); err != nil {
		return fmt.Errorf("unable to register view service: %w", err)
	}

	return nil

}

func (provider viewServiceProvider) Boot(container di.Container) error {
	resp, err := container.SafeGet("router")
	if err != nil {
		return fmt.Errorf("unable to boot view service: cannot retrieve router from container: %w", err)
	}

	rtr, ok := resp.(*router.Router)
	if !ok {
		return fmt.Errorf("unable to boot view service: router is not of type *router.Router, got %T", rtr)
	}

	rtr.AddHandlerTransformer(provider.viewHandlerTransformer)

	return nil
}

// load walks through the directory provided and loads all the
// `.html` files.
func (provider viewServiceProvider) load(path string) *template.Template {
	path = filepath.Clean(path)

	tmpls, err := provider.findAndParseTemplates(path, provider.viewFunctions())
	if err != nil {
		log.Fatalln("unable to load templates:", err)
	}

	return tmpls
}

func (provider viewServiceProvider) viewFunctions() template.FuncMap {
	return template.FuncMap{
		"asset": func(path string) string {
			return "/assets/" + path
		},
	}
}

func (provider viewServiceProvider) findAndParseTemplates(
	path string,
	funcMap template.FuncMap,
) (*template.Template, error) {
	// TODO: Tidy this up and make it clearer - comments, and split into funcs, for eg.
	pfx := len(path) + 1
	root := template.New("")

	err := filepath.Walk(
		path,
		func(path string, info os.FileInfo, e1 error) error {
			if !info.IsDir() && strings.HasSuffix(path, ".html") {
				if e1 != nil {
					return e1
				}

				b, e2 := ioutil.ReadFile(path)
				if e2 != nil {
					return e2
				}

				// Strip the `.html` string from the end of the
				// template so we can execute it using `name`
				// rather than `name.html`.
				name := path[pfx : len(path)-5]

				name = strings.Join(
					strings.Split(name, "/"),
					".",
				)

				t := root.New(name).Funcs(funcMap)
				t, e2 = t.Parse(string(b))
				if e2 != nil {
					return e2
				}
			}

			return nil
		},
	)

	return root, err
}

func (provider viewServiceProvider) viewHandlerTransformer(val route.View) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := r.Form
		views := di.Get(r, "views").(*template.Template)

		if err := views.ExecuteTemplate(w, string(val), params); err != nil {
			// Something went wrong either finding or executing the template.
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Oops, something went wrong"))
			log.Printf("unable to execute template `%s`: %s", val, err.Error())
			return
		}
	})
}
