package g

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
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
		name := page[len(templatesDir+"views/"):len(page)]
		// TODO:
		if name == "index.html" {
			for i, v := range files {
				if strings.HasSuffix(v, "/layouts/single.html") {
					DeleteSlice(files, i)
				}
			}
		}
		templates[name] = template.Must(parseFiles(path, nil, files...))
	}

}

// DeleteSlice
func DeleteSlice(slice interface{}, index int) (interface{}, error) {
	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	if slice == nil || length == 0 || (length-1) < index {
		return nil, errors.New("error")
	}
	if length-1 == index {
		return sliceValue.Slice(0, index).Interface(), nil
	} else if (length - 1) >= index {
		return reflect.AppendSlice(sliceValue.Slice(0, index), sliceValue.Slice(index+1, length)).Interface(), nil
	}
	return nil, errors.New("error")
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
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
