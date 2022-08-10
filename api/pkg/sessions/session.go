package session

import (
	gsessions "github.com/gorilla/sessions"
	"net/http"
)

// TODO: Read secret from env
var store = gsessions.NewCookieStore([]byte("secret"))

func Get(req *http.Request) (*gsessions.Session, error) {
	return store.Get(req, "sessions")
}

func GetNamed(req *http.Request, name string) (*gsessions.Session, error) {
	return store.Get(req, name)
}

// EXAMPLE: to use this session in controllers
//
// import (
//     "net/http"
//     "github.com/yourpackage/sessions"
// )
//
// func Index(rw http.ResponseWriter, r *http.Request) {
//     session, err := sessions.Get(r)
//     if err != nil {
//         panic(err)
//     }
//     session.Values["test"] = "test"
//     session.Save(r, rw)
// }
