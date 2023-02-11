package html

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"time"

	"github.com/leekchan/gtf"
)

type HTMLConfig struct {
	TemplatesPath string `env:"HTML_TEMPLATES_PATH" env-default:"./templates"`
	DateFormat    string `ev:"HTML_DATEFORMAT" env-default:"2 Jan 2006"`
}

type HTML struct {
	config    HTMLConfig
	Templates map[string]*template.Template
}

// NewHTML creates new HTML instance
func NewHTML(config HTMLConfig) (*HTML, error) {
	instance := &HTML{
		Templates: make(map[string]*template.Template),
		config:    config,
	}

	return instance, instance.InitTemplates()
}

// InitTemplates loads templates in memory
func (r *HTML) InitTemplates() error {

	baseFiles, err := r.GetPartials()
	if err != nil {
		return err
	}

	baseFiles = append(baseFiles, fmt.Sprintf("%s/base.tpl", r.config.TemplatesPath))

	baseTemplate := template.Must(
		template.New("base.tpl").
			Funcs(GetTmplFuncMap(r.config.DateFormat)).
			ParseFiles(baseFiles...),
	)

	for _, tmplName := range []string{"index.tpl", "post.tpl", "tag.tpl"} {
		tmplPath := fmt.Sprintf("%s/%s", r.config.TemplatesPath, tmplName)
		r.Templates[tmplName] = template.Must(template.Must(baseTemplate.Clone()).ParseFiles(tmplPath))

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *HTML) GetPartials() ([]string, error) {
	var partials []string

	walkFn := func(path string, f os.FileInfo, err error) error {
		if nil == err && !f.IsDir() {
			partials = append(partials, path)
		}
		return err
	}
	partialsPath := fmt.Sprintf("%s/partial", r.config.TemplatesPath)

	err := filepath.Walk(partialsPath, walkFn)
	if err != nil {
		return nil, err
	}

	return partials, nil
}

func GetTmplFuncMap(dateFormat string) template.FuncMap {
	funcs := gtf.GtfFuncMap

	funcs["format_date"] = func(value time.Time) string {
		return value.Format(dateFormat)
	}
	funcs["add"] = func(arg int, value int) int {
		return value + arg
	}
	funcs["noescape"] = func(value string) template.HTML {
		/* #nosec G203 -- function is supposed to be non-safe and contain JS */
		return template.HTML(value)
	}

	return funcs
}
