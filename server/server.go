package server

import (
	"log"
	"net/http"

	da "github.com/krishnakantha1/expenseTrackerBackend/dataaccess"
	h "github.com/krishnakantha1/expenseTrackerBackend/handlers"
)

type handlerWithDA func(da.DataAccess, http.ResponseWriter, *http.Request)

type Server struct {
	smux       *http.ServeMux
	dataAccess da.DataAccess
}

/*
Initiates and starts a http server

dataAccess : Implementation of dataAccessInterface
port : port on which the server should listen to
*/
func Init(dataAccess da.DataAccess, port string) {
	smux := http.NewServeMux()

	server := &Server{
		smux:       smux,
		dataAccess: dataAccess,
	}

	server.BindHandlers()

	log.Println("Listening on port", port)
	if err := http.ListenAndServe(port, server.smux); err != nil {
		log.Fatal(err)
	}
}

/*
Binds the urls to a handler
*/
func (s *Server) BindHandlers() {
	s.smux.HandleFunc("GET /api/ping/{id}", s.bindDA(h.Ping))

	//auth
	s.smux.HandleFunc("POST /api/auth/v1/login", s.bindDA(h.Login))
	s.smux.HandleFunc("POST /api/auth/v1/login-jwt", s.bindDA(h.LoginWithJWT))

	//expense
	s.smux.HandleFunc("POST api/expense/v1/ingest", s.bindDA(h.ExpenseIngestion))
}

/*
returns a http.HandlerFunc which calls the function provided in the argument.
This is used to provide DataAccessInterface to the handler func

f : function of the signature handlerWithDA
*/
func (s *Server) bindDA(f handlerWithDA) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if r.Method == http.MethodOptions {
			return
		}

		f(s.dataAccess, w, r)
	}
}
