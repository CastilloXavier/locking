package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var LayoutDir = "views/layouts/"
var TemplateDir = "views/"
var LayoutExtenxion = ".gohtml"

func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files, filesLayouts()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

//Render is used to render the view with the predefined layout.
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}

func filesLayouts() []string {
	files, err := filepath.Glob(LayoutDir + "*" + LayoutExtenxion)
	if err != nil {
		panic(err)
	}
	return files
}

// addTemplatePath takes in a slice og strings
// representig fle paths fot templates,an d it prepends
// the TemplateDir directory to each string in the slice
//
// Eg the input {"home"} would result in the output
// {"views/home"} if TemplateDir == "views/"
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

// addTempleateExt takes in a slice of strings
// representig file paths for templates and it appends
// the TemplateExt extension to each string in the slice
//
// Eg the input {"home"} would result in the output
// {"home.gohtml"} if TempleateExt == ".gohtml"
func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + LayoutExtenxion
	}
}
