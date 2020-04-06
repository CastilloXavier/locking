package views

import (
	"html/template"
	"path/filepath"
)

var LayoutDir = "views/layouts/"
var LayoutExtenxion = "gohtml"

func NewView(layout string, files ...string) *View {
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

func filesLayouts() []string {
	files, err := filepath.Glob(LayoutDir + "*" + LayoutExtenxion)
	if err != nil {
		panic(err)
	}
	return files
}
