package login

import (
    "net/http"
    "sample/app/shared/handler"
    "github.com/gorilla/sessions"
)

// HTTPHandler struct.
type HTTPHandler struct {
    handler.ApplicationHTTPHandler
}

type templateData struct {
	LoginError string
	UserName string
}

var (
    key   = []byte("secret-key")
    store = sessions.NewCookieStore(key)
)
var loginError = ""

func (h *HTTPHandler) LoginDemo(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	userName, _ := session.Values["userName"].(string)
    err := h.ResponseHTML(w, r, "login/login", templateData{
		LoginError: loginError,
		UserName: userName,
    })
    if err != nil {
        _ = h.StatusServerError(w, r)
    }
}
func CheckUser(userName, pass string) bool {
    uName, pwd := "lamnt", "123123"
    if uName == userName && pwd == pass {
        return true
    }
    return false
}

func (h *HTTPHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
    loginError = ""
    userName := r.FormValue("user-name")
    password := r.FormValue("password")
    if !(len(userName) <= 0)  && !(len(password) <= 0) {
        userIsValid := CheckUser(userName, password)
        if userIsValid {
            session, _ := store.Get(r, "user")
            session.Values["userName"] = userName
            session.Values["password"] = password
            session.Save(r, w)
        } else {
            loginError = "Incorrect name and password"
        }
    } else {
        loginError = "Please, input name and password!"
    }
    http.Redirect(w, r, "/login", 302)

}

func (h *HTTPHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "user")
    session.Options.MaxAge = -1
    session.Save(r, w)
    http.Redirect(w, r, "/login", 302)
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewLoginHTTPHandler(ah *handler.ApplicationHTTPHandler) *HTTPHandler {
    // item set.
    return &HTTPHandler{ApplicationHTTPHandler: *ah}
}
