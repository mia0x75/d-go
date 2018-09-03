package g

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"
	"github.com/mia0x75/dashboard-go/utils"
)

var (
	templates map[string]*template.Template
)

type Template struct{}

func init() {
	path, err := utils.GetCurrentPath()
	if err != nil {
		fmt.Printf("cannot startup project, error: %s\r\n", err.Error())
		os.Exit(1)
	}
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templatesDir := path + "/templates/"
	pages := []string{}
	err = filepath.Walk(templatesDir+"views/", func(page string, f os.FileInfo, err error) error {
		if !f.IsDir() && strings.HasSuffix(page, ".html") {
			if err != nil {
				return err
			}
			if _, err := ioutil.ReadFile(page); err != nil {
				return err
			}
			pages = append(pages, page)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	layouts, err := filepath.Glob(templatesDir + "layouts/*.html")
	if err != nil {
		log.Fatal(err)
	}

	partials, err := filepath.Glob(templatesDir + "partials/*.html")
	if err != nil {
		log.Fatal(err)
	}
	includes := append(layouts, partials...)
	// Generate our templates map from our layouts/ and partials/ directories
	for _, page := range pages {
		files := append(includes, page)
		key := page[len(templatesDir+"views/"):len(page)]
		templates[key] = template.Must(parseFiles(path, nil, files...))
	}

}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}

	err := tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (t *Template) Templates() map[string]*template.Template {
	return templates
}

func parseFiles(path string, t *template.Template, filenames ...string) (*template.Template, error) {
	if len(filenames) == 0 {
		// Not really a problem, but be consistent.
		return nil, fmt.Errorf("html/template: no files named in call to ParseFiles")
	}
	for _, filename := range filenames {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		s := string(b)
		name := filepath.Base(filename)
		// TODO: Remove hardcode
		if strings.Index(filename, "/templates/views/docs/") > 0 {
			name = filename[len(path)+1+len("templates/views/"):] //
		}
		fmt.Printf("filename: %-90s -> name: %s\n", filename, name)
		// First template becomes return value if not already defined,
		// and we use that one for subsequent New calls to associate
		// all the templates together. Also, if this file has the same name
		// as t, this file becomes the contents of t, so
		//  t, err := New(name).Funcs(xxx).ParseFiles(name)
		// works. Otherwise we create a new template associated with t.
		var tmpl *template.Template
		if t == nil {
			t = template.New(name)
		}
		if name == t.Name() {
			tmpl = t
		} else {
			tmpl = t.New(name)
		}
		_, err = tmpl.Parse(s)
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
