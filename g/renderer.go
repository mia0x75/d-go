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
	"sync"

	"github.com/labstack/echo"
	"github.com/mia0x75/dashboard-go/utils"
)

var (
	templates map[string]*template.Template
	renderer  *Renderer
	lock      = new(sync.Mutex)
)

type Renderer struct{}

func NewRenderer() *Renderer {
	lock.Lock()
	defer lock.Unlock()
	if renderer != nil {
		return renderer
	}
	renderer = &Renderer{}
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
		name := page[len(templatesDir+"views/"):len(page)]
		templates[name] = template.Must(parse(path, name, files...))
	}
	return renderer
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}

	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		log.Fatal(err)
		return nil
	} else {
		return nil
	}
}

func (t *Renderer) GetTemplate(name string) (*template.Template, error) {
	if tmpl, ok := templates[name]; ok {
		return tmpl, nil
	}
	return nil, fmt.Errorf("specified template: %s does not exist.", name)
}

func parse(path string, name string, files ...string) (*template.Template, error) {
	if len(files) == 0 {
		// Not really a problem, but be consistent.
		return nil, fmt.Errorf("html/template: no files named in call to parse")
	}
	tmpl := template.New(name)
	for _, file := range files {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		if _, err := tmpl.Parse(string(b)); err != nil {
			return nil, err
		}
	}
	return tmpl, nil
}
