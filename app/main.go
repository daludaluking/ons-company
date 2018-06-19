package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"encoding/hex"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var onsTransactionHalder *ONSClient

func main() {
	/*
	   r.Handle("/", http.RedirectHandler("/ons", http.StatusFound))
	   http.HandleFunc("/", handle)
	   http.HandleFunc("/_ah/health", healthCheckHandler)
	   http.HandleFunc("/admin", adminHandler)
	   log.Print("Listening on port 8080")

	   conn, _, err := websocket.DefaultDialer.Dial("ws://www.example.com/socketserver", nil)
	   if err != nil {
	       log.Printf("Websocket dial error: %v", err)
	   }
	       conn.WriteMessage(websocket.TextMessage, []byte("Connect"))
	*/
	//일단 private key를 hard coding 해 놓자.
	//차후에 받는걸로...
	private_key, _ := hex.DecodeString("ad661cc1acff767e4148ebf74a080a8f54c13abde64062c5cd73d65863e4dd6a")
	onsTransactionHalder = NewONSTransactionHalder(private_key)
	registerHandlers()
	log.Printf("Start ons-company : http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func registerHandlers() {
	// Use gorilla/mux for rich routing.
	// See http://www.gorillatoolkit.org/pkg/mux
	r := mux.NewRouter()
	r.Handle("/", http.RedirectHandler("/ons", http.StatusFound))
	r.Methods("POST").Path("/ons/admin").Handler(appHandler(adminRequestHandler))

	// Respond to App Engine and Compute Engine health checks.
	// Indicate the server is healthy.
	r.Methods("GET").Path("/_ah/health").HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {w.Write([]byte("ok"))})


	http.Handle("/", handlers.CombinedLoggingHandler(os.Stderr, r))

}
func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello world!")
}

func adminRequestHandler(w http.ResponseWriter, r *http.Request) *appError {
	action := r.FormValue("action")
	if CheckAction(action) == false {
		err := fmt.Errorf("admin: %v action doesn't exist in action list:", action)
		return appErrorf(err, "adminRequestHandler : CheckAction : %v", err)
	}

	err := AdminProcessAction(onsTransactionHalder, action, w, r)

	if err != nil {
		return appErrorf(err, "adminRequestHandler : AdminProcessAction : %v", err)
	}
	return nil
}

// http://blog.golang.org/error-handling-and-go
type appHandler func(http.ResponseWriter, *http.Request) *appError

type appError struct {
	Error   error
	Message string
	Code    int
}

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		log.Printf("Handler error: status code: %d, message: %s, underlying err: %#v",
			e.Code, e.Message, e.Error)
		http.Error(w, e.Message, e.Code)
	}
}

func appErrorf(err error, format string, v ...interface{}) *appError {
	return &appError{
		Error:   err,
		Message: fmt.Sprintf(format, v...),
		Code:    500,
	}
}
