package converter

import (
	"html/template"
	"net/http"
)

var tpls *template.Template

func init() {
	tpls = template.Must(template.ParseGlob("templates/*.html"))
	http.HandleFunc("/", handleindex)
	http.HandleFunc("/convert", handleconverter)
}

func handleindex(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}

	err := tpls.ExecuteTemplate(res, "index.gohtml", nil)
	if err != nil {
		http.Error(res, err.Error(), 500)
	}

}
