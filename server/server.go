package server

import (
	"log"
	"net/http"

	da "github.com/krishnakantha1/expenseTrackerBackend/dataaccess"
	h "github.com/krishnakantha1/expenseTrackerBackend/handlers"
	"github.com/krishnakantha1/expenseTrackerBackend/utils"
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
	s.smux.HandleFunc("/api/ping/{id}", s.bindDA(h.Ping, http.MethodGet))

	//auth
	s.smux.HandleFunc("/api/auth/v1/login", s.bindDA(h.Login, http.MethodPost))
	s.smux.HandleFunc("/api/auth/v1/login-jwt", s.bindDA(h.LoginWithJWT, http.MethodPost))

	//expense
	s.smux.HandleFunc("api/expense/v1/ingest", s.bindDA(h.ExpenseIngestion, http.MethodPost))
}

/*
returns a http.HandlerFunc which calls the function provided in the argument.
This is used to provide DataAccessInterface to the handler func

Params:

	f : function of the signature handlerWithDA
	allowedMethod (string): the current http method thats allowed
*/
func (s *Server) bindDA(f handlerWithDA, allowedMethod string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, FETCH")
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if r.Method == http.MethodOptions {
			return
		}

		if r.Method != allowedMethod {
			utils.RequestNotAllowedResponse(w)
			return
		}

		f(s.dataAccess, w, r)
	}
}
