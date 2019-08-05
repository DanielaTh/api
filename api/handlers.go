package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DanielaTh/api/models"
	"github.com/DanielaTh/api/repositories"

	"github.com/gorilla/mux"
)

// Handler : configuration
type Handler struct {
	SubjectDB repositories.SubjectStore
}

// Foo : test
func Foo() string {
	return "hola mundo"
}

// NewHandler : Init new Handler
func NewHandler(conn string) (Handler, error) {
	db, err := repositories.NewDB(conn)
	if err != nil {
		var h Handler
		return h, err
	}
	return Handler{SubjectDB: db}, nil
}

// GetSubject : get a subject
func (h Handler) GetSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	sid := mux.Vars(r)["id"]
	id64, err := strconv.ParseInt(sid, 10, 64)

	if err != nil {
		http.Error(w, getJSONErrInfo(err.Error()), http.StatusInternalServerError)
		return
	}

	p, err := h.SubjectDB.GetSubject(int(id64))

	if err != nil {
		http.Error(w, getJSONErrInfo(err.Error()), http.StatusInternalServerError)
		return
	}

	if p.ID == 0 {
		http.Error(w, getJSONErrInfo("Not Found"), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(p)
	if err != nil {
		http.Error(w, getJSONErrInfo(err.Error()), http.StatusInternalServerError)
		return
	}
	w.Write(js)
}

// AddSubject : add a subject
func (h Handler) AddSubject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(r.Body)
	var p models.Subject
	err := decoder.Decode(&p)
	if err != nil {
		http.Error(w, getJSONErrInfo(err.Error()), http.StatusInternalServerError)
		return
	}

	p, err = h.SubjectDB.AddSubject(p)
	if err != nil {
		http.Error(w, getJSONErrInfo(err.Error()), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(p)
	if err != nil {
		http.Error(w, getJSONErrInfo(err.Error()), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func getJSONErrInfo(msg string) string {
	errinfo := struct {
		Info string
	}{
		Info: msg,
	}
	j, _ := json.Marshal(errinfo)
	return string(j)
}
