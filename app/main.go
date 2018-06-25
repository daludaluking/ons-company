package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"encoding/hex"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"encoding/json"
)

type svActionHandler func(string, http.ResponseWriter, *http.Request) error
var sv_action_name = map[int32]string{
	0:  "start",
}
var sv_action_value = map[string]int32{
	"start":        0,
}
var sv_action_handler = map[int32]svActionHandler{
	0: actionStart,
}

var onsTransactionHalder *ONSClient
var onsEventHandler *ONSEventHandler

type ONSServerInfo struct {
	TRAddr string `json:"transaction_server_ip"`
	TRPort string `json:"transaction_server_port"`
	EVAddr string `json:"event_server_ip"`
	EVPort string `json:"event_server_port"`
	ErrorCode float64 `json:"error"`
	Message string `json:"message"`
}

const (
	ERROR_NONE = iota
	ERROR_SERVER_ACTION
	ERROR_ADMIN_ACTION
	ERROR_DATASTORE_ACTION
)

type ONSErrorResponse struct {
	Action string `json:"action"`
	ErrorCode float64 `json:"error"`
	Message string `json:"message"`
	Report interface{} `json:"report"`
}

var onsServerInfo ONSServerInfo

func main() {
	err := CreateDatastoreClient()
	if err != nil {
		log.Printf("Failed to create new client : %v", err)
		os.Exit(1)
	}

	//일단 private key를 hard coding 해 놓자.
	//차후에 받는걸로...
	log.Printf("Start rest api server for ons company : http://:8080")
	registerHandlers()
	err = http.ListenAndServe(":8080", nil)

	if onsEventHandler != nil {
		onsEventHandler.Subscribe(false)
		onsEventHandler.Terminate(true)
	}

	if err != nil {
		log.Printf("Exit ons-company : %v", err)
	}
}

func registerHandlers() {

	// Use gorilla/mux for rich routing.
	// See http://www.gorillatoolkit.org/pkg/mux
	r := mux.NewRouter()
	r.Handle("/", http.RedirectHandler("/ons", http.StatusFound))
	r.Methods("POST").Path("/ons/admin").Handler(appHandler(adminRequestHandler))
	r.Methods("GET").Path("/ons/account").Handler(appHandler(accoutRequestHandler))

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
}

func accoutRequestHandler(w http.ResponseWriter, r *http.Request) *appError {
	id := r.FormValue("id")
	var level, existed string
	if id == "" {
		_ = WriteResponse("accoutRequestHandler", ONSErrorResponse{
			ErrorCode: ERROR_DATASTORE_ACTION,
			Action: "accoutRequestHandler",
			Message:  "Invalid parameter : id is needed.",
			Report: string(""),
		}, w)
		return nil
	}

	level = "Admin"
	privKey, err := GetPrivateKey(level, id)

	if err != nil {
		level = "User"
		privKey, err = GetPrivateKey(level, id)
		if err != nil {
			_ = WriteResponse("accoutRequestHandler", ONSErrorResponse{
				ErrorCode: ERROR_DATASTORE_ACTION,
				Action: "accoutRequestHandler",
				Message:  "Invalid ID : id isn't exist.",
				Report: err.Error(),
			}, w)
			return nil
		}
	}

	if privKey != "" {
		existed = "exist"
	}else{
		existed = "none"
	}

	_ = WriteResponse("accoutRequestHandler", map[string]string{
		"id": id,
		"level": level,
		"private_key": existed,
	}, w)
	return nil
}

func adminRequestHandler(w http.ResponseWriter, r *http.Request) *appError {
	action := r.FormValue("action")
	if CheckAction(action) == false {
		if CheckServerAction(action) == true {
			_ = ServerProcessAction(action, w, r)
			return nil
		}
		err := fmt.Errorf("admin: %v action doesn't exist in action list:", action)
		return appErrorf(err, "adminRequestHandler : CheckAction : %v", err)
	}

	if onsTransactionHalder == nil || onsEventHandler == nil {
		_ = WriteResponse(action, ONSErrorResponse{
			ErrorCode: ERROR_ADMIN_ACTION,
			Action: action,
			Message:  "First of all, please start ons company",
			Report: string(""),
		}, w)
		return nil
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

func CheckServerAction(action string) bool {
	_, ok := sv_action_value[action]
	return ok
}

func ServerProcessAction(action string, w http.ResponseWriter, r *http.Request) error {
	ahandler, ok := sv_action_handler[sv_action_value[action]]
	if ok == false {
		return fmt.Errorf("ServerProcessAction: %v action doesn't have handler.", action)
	}
	return ahandler(action, w, r)
}

func actionStart(action string, w http.ResponseWriter, r *http.Request) error {
	onsServerInfo.TRAddr = r.FormValue("tr_addr")
	onsServerInfo.TRPort = r.FormValue("tr_port")
	onsServerInfo.EVAddr = r.FormValue("ev_addr")
	onsServerInfo.EVPort = r.FormValue("ev_port")
	var err error
	onsEventHandler, err = NewONSEventHandler(onsServerInfo.EVAddr+":"+onsServerInfo.EVPort, "/subscriptions", nil, true)
	if err != nil {
		onsServerInfo.ErrorCode = ERROR_SERVER_ACTION
		onsServerInfo.Message = "Failed to create ons event handler"
		return WriteResponse(action, ONSErrorResponse{
			ErrorCode: ERROR_SERVER_ACTION,
			Action: action,
			Message:  "Failed to create ons event handler",
			Report: onsServerInfo,
		}, w)
	}

	//key = private key
	//

	private_key, _ := hex.DecodeString("ad661cc1acff767e4148ebf74a080a8f54c13abde64062c5cd73d65863e4dd6a")
	onsTransactionHalder = NewONSTransactionHalder(onsServerInfo.TRAddr, onsServerInfo.TRPort, private_key)

	if onsTransactionHalder == nil {
		onsServerInfo.ErrorCode = ERROR_SERVER_ACTION
		onsServerInfo.Message = "Failed to create ons event handler"
		return WriteResponse(action, ONSErrorResponse{
			ErrorCode: ERROR_SERVER_ACTION,
			Action: action,
			Message:  "Failed to create ons transaction handler",
			Report: onsServerInfo,
		}, w)
	}

	if onsEventHandler.Run() == false {
		onsServerInfo.ErrorCode = ERROR_SERVER_ACTION
		onsServerInfo.Message = "Failed to run ons event handler"
		//log.Printf("Failed to run ons event handler")
		return WriteResponse(action, ONSErrorResponse{
			ErrorCode: ERROR_SERVER_ACTION,
			Action: action,
			Message:  "Failed to run ons event handler",
			Report: onsServerInfo,
		}, w)
	}

	onsEventHandler.Subscribe(true)

	onsServerInfo.ErrorCode = 0
	onsServerInfo.Message = "OK"
	WriteResponse("actionStart", onsServerInfo, w)
	return nil
}

func WriteResponseJson(caller string, result string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, result)
	return nil
}

func WriteResponse(caller string, result interface{}, w http.ResponseWriter) error {
	response, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Fprintf(w, "%s : cannot marshal data into json", caller)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(response))
	return nil
}