package web

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	cryptos "github.com/angelospillos/cryptosgateway"
	"github.com/gorilla/mux"
)

type Server struct {
	Host           string
	CryptosService cryptos.Service
}

func (s Server) Start() {

	router := mux.NewRouter()
	router.HandleFunc("/", s.Index).Methods(http.MethodGet)
	router.HandleFunc("/health", s.Health).Methods(http.MethodGet)
	router.HandleFunc("/v1/top/coins/by-market-cap", s.GetTopCoinsByMarketCap).Methods(http.MethodGet)

	err := http.ListenAndServe(s.Host, router)

	if err != nil {
		log.Fatalf("Server.Start: %v", err)
	}
}

func (s Server) Index(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		log.Printf("Server.Index: %v", err)
	}
}

func (s Server) Health(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(200)
	_, err := w.Write([]byte("Ok"))
	if err != nil {
		log.Printf("Server.Index: %v", err)
	}
}

func WriteJSON(w http.ResponseWriter, status int, value interface{}) {
	if payload, err := json.Marshal(value); err != nil {
		log.Printf("error writing response %v", err)
		Server500(w)
		return
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(status)
		_, err := w.Write(payload)
		if err != nil {
			log.Printf("error writing response %v", err)
		}
	}
}

func WriteContentString(w http.ResponseWriter, status int, contentType string, content string) {
	w.Header().Add("Content-Type", contentType)
	w.WriteHeader(status)
	_, err := w.Write([]byte(content))
	if err != nil {
		log.Printf("error writing response %v", err)
	}
}
func Server500(w http.ResponseWriter) {
	WriteContentString(w, http.StatusInternalServerError, "text/html", "<html><head/><body><h1>System error</h1></body></html>")
}

func Server404(w http.ResponseWriter) {
	WriteContentString(w, http.StatusNotFound, "text/html", "<html><head/><body><h1>Not found error</h1></body></html>")
}

func Server400(w http.ResponseWriter) {
	WriteContentString(w, http.StatusNotFound, "text/html", "<html><head/><body><h1>Bad Request</h1></body></html>")
}

func (s Server) GetTopCoinsByMarketCap(w http.ResponseWriter, req *http.Request) {

	limit, err := strconv.Atoi(req.FormValue("limit"))

	if err != nil {
		Server400(w)
		return
	}

	if limit < 10 {
		WriteJSON(w, 400, "minimum limit is 10")
		return
	}

	if limit > 100 {
		WriteJSON(w, 400, "maximum limit is 100")
		return
	}

	account, err := s.CryptosService.GetTopCoinsByMarketCap(int32(limit))

	if err != nil {
		log.Printf("Server.GetTopCoinsByMarketCap: %v", err)
		Server500(w)
		return
	}

	WriteJSON(w, 200, account)
}
