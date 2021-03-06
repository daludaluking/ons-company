package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"net/url"
	"encoding/hex"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	//"github.com/daludaluking/ons-sawtooth-sdk/ons_pb2"
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
	ERROR_TX_ACTION
	ERROR_DATASTORE_ACTION
)

type ONSErrorResponse struct {
	Action string `json:"action"`
	ErrorCode float64 `json:"error"`
	Message string `json:"message"`
	Report interface{} `json:"report"`
}

var onsServerInfo ONSServerInfo
//ONSEventDB is interface.
var onsDB *ONSEventDatastoreDB

func main() {
	var err error
	log.SetFlags(log.Lshortfile)

	onsDB, err = CreateDatastoreClient()
	if err != nil {
		log.Printf("Failed to create new client : %v", err)
		os.Exit(1)
	}

	log.Printf("CreateDatastoreClient : %#v", onsDB)

	_, err = onsDB.DBInitLatestUpdatedBlockInfo(false)
	if err != nil {
		log.Printf("Failed to DBInitLatestUpdatedBlockInfo : %v", err)
		os.Exit(1)
	}

	blockNum, blockId :=  onsDB.DBGetLatestUpdatedBlock()
	log.Printf("DBInitLatestUpdatedBlockInfo : Block num: %v, Block Id: %v, DB Info: %#v", blockNum, blockId, onsDB)

	//return;
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
	//r.Methods("POST").Path("/ons/tx").Handler(appHandler(txRequestHandler))
	r.Methods("POST").Path("/ons/tx").Handler(appHandler(txRequestHandler))
	r.Methods("POST").Path("/ons/admin").Handler(appHandler(adminRequestHandler))
	r.Methods("GET").Path("/ons/account").Handler(appHandler(accoutRequestHandler))
	r.Methods("GET").Path("/ons/data").Handler(appHandler(dataRequestHandler))

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

func dataRequestHandler(w http.ResponseWriter, r *http.Request) *appError {
	kind := r.FormValue("kind")

	filter, err :=  url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		//make empty map.
		filter = make(map[string][]string)
	}

	log.Printf("dataRequestHandler: %v, filter(%v) = %#v\n", kind, r.URL.RawQuery, filter)
	d, err := GetONSData(kind, filter)
	if err != nil {
		_ = WriteResponse("dataRequestHandler", ONSErrorResponse{
			ErrorCode: ERROR_DATASTORE_ACTION,
			Action: "dataRequestHandler",
			Message:  "Failed to get ons data",
			Report: err.Error(),
		}, w)
		return nil
	}
	log.Printf("dataRequestHandler: %v\n", d)
	_ = WriteResponseMulti("dataRequestHandler", d, w)
	return nil
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
/*
func txRequestHandler(w http.ResponseWriter, r *http.Request) *appError {
	log.Printf("%#v\n", r)
	contentType := r.Header.Get("Content-type")
	if contentType != "application/json" {
		_ = WriteResponse("txRequestHandler", ONSErrorResponse{
			ErrorCode: ERROR_TX_ACTION,
			Action: "txRequestHandler",
			Message:  "transaction data only supports the application/json content type.",
			Report: string(""),
		}, w)
		return nil
	}

	decoder := json.NewDecoder(r.Body)
	log.Println(decoder)
	var t ons_pb2.SendONSTransactionPayload
	err := decoder.Decode(&t)

	if err != nil {
		panic(err)
	}

	log.Printf("%v\n", t.RegisterGs1Code.Gs1Code)

	_ = WriteResponse("txRequestHandler", ONSErrorResponse{
		ErrorCode: ERROR_TX_ACTION,
		Action: "txRequestHandler",
		Message:  "Not implemented",
		Report: string(""),
	}, w)
	return nil
}
*/
func txRequestHandler(w http.ResponseWriter, r *http.Request) *appError {
	action := r.FormValue("action")
	err := TxProcessAction(onsTransactionHalder, action, w, r)
	if err != nil {
		return appErrorf(err, "txRequestHandler : AdminProcessAction : %v", err)
	}
	return nil
}

func adminRequestHandler(w http.ResponseWriter, r *http.Request) *appError {
	action := r.FormValue("action")
	if CheckAdminAction(action) == true {
		_ = ServerProcessAction(action, w, r)
		return nil
	}
	err := fmt.Errorf("admin: %v action doesn't exist in action list:", action)
	return appErrorf(err, "adminRequestHandler : CheckAction : %v", err)
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

func CheckAdminAction(action string) bool {
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
	onsEventHandler, err = NewONSEventHandler(onsServerInfo.EVAddr+":"+onsServerInfo.EVPort, "/subscriptions", onsDB, true)
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

	//for test....
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
	return WriteResponseMulti(caller, []interface{}{result}, w)
}

func WriteResponseMulti(caller string, results []interface{}, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	response, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Fprintf(w, "%s : cannot marshal data into json", caller)
		return err
	}
	fmt.Fprintf(w, string(response))
	return nil
}