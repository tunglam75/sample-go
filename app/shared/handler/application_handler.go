package handler

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var (
	templateDir    = os.Getenv("WORK_DIR") + "/template"
	commonDir      = templateDir + "/common/"
	templateHeader = "header.html"
	templateFooter = "footer.html"
)

// ApplicationHTTPHandler base handler struct.
type ApplicationHTTPHandler struct {
	BaseHTTPHandler
}

// DataTemplate uses all mapping template data.
type DataTemplate struct {
	Common CommonTemplate
	Data   interface{}
}

// CommonTemplate uses common mapping template data.
type CommonTemplate struct {
	ErrorCode int
	ErrorInfo string
}

// ApplicationHTTPHandlerInterface is interface.
type ApplicationHTTPHandlerInterface interface {
	ResponseHTML(http.ResponseWriter, string, interface{})
	ResponseErrorHTML(http.ResponseWriter, int, interface{})
}

// ResponseHTML responses status code 200 and html.
func (h *ApplicationHTTPHandler) ResponseHTML(w http.ResponseWriter, r *http.Request, templateFile string, data interface{}) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	commonTemplate := CommonTemplate{}

	d := &DataTemplate{
		Common: commonTemplate,
		Data:   data,
	}

	return executeTemplate(w, templateDir+"/"+templateFile+".html", d)
}

// ResponseErrorHTML calls utils.ResponseHTML.
func (h *ApplicationHTTPHandler) ResponseErrorHTML(w http.ResponseWriter, r *http.Request, code int, errorInfo string) error {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)

	commonTemplate := CommonTemplate{
		ErrorCode: code,
		ErrorInfo: errorInfo,
	}

	// set parameters.
	d := &DataTemplate{
		Common: commonTemplate,
	}
	// set common template.
	return executeTemplate(w, templateDir+"/error/error.html", d)
}

// executeTemplate executes template.
func executeTemplate(w http.ResponseWriter, templateFile string, data interface{}) error {
	funcMap := map[string]interface{}{"makeSlice": makeSlice}
	tname := filepath.Base(templateFile)
	tmpl := template.New(tname).Funcs(template.FuncMap(funcMap))
	tmpl = template.Must(tmpl.ParseFiles(
		templateFile,
	))
	err := tmpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
	return err
}

// makeSlice is template custom function.
func makeSlice(args ...interface{}) []interface{} {
	return args
}

// NewApplicationHTTPHandler returns ApplicationHTTPHandler instance.
func NewApplicationHTTPHandler(logger *logrus.Logger) *ApplicationHTTPHandler {
	return &ApplicationHTTPHandler{BaseHTTPHandler: BaseHTTPHandler{Logger: logger}}
}

// StatusServerError responses status code 500 and html.
func (h *ApplicationHTTPHandler) StatusServerError(w http.ResponseWriter, r *http.Request) error {
	// status code 500
	return h.ResponseErrorHTML(w, r, http.StatusInternalServerError, "500 Internal Server Error.")
}
