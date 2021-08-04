package api

import (
	"GoNews/pkg/storage"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Default message handler
type Message struct {
	Api     string
	Version float32
	Message string
}

// API
type API struct {
	db     storage.Interface
	router *mux.Router
}

// API constructor
func New(db storage.Interface) *API {
	api := API{
		db: db,
	}
	api.router = mux.NewRouter()
	api.endpoints()
	return &api
}

func (api *API) endpoints() {
	api.router.HandleFunc("/version", api.DefaultHandler).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/news/{n}", api.NewsHandler).Methods(http.MethodGet, http.MethodOptions)
	api.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./webapp"))))
}

func (api *API) Router() *mux.Router {
	return api.router
}

func (api *API) DefaultHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Api:     "gonews",
		Version: 1.0,
		Message: "bitch",
	}
	bytes, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func (api *API) NewsHandler(w http.ResponseWriter, r *http.Request) {
	var n int = 10
	s := strings.Split(r.URL.Path, "/")
	if len(s) > 0 {
		n, _ = strconv.Atoi(s[len(s)-1])
	}
	posts, err := api.db.News(n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}
