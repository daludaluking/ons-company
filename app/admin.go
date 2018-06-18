package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

const action_register = "reg"
const action_deregister = "dereg"
const action_add = "add_rec"
const action_remove = "rm_rec"
const action_register_svc = "reg_svc"
const action_deregister_svc = "dereg_svc"
const action_change_gstate = "chx_gstate"
const action_change_rstate = "chx_rstate"
const action_add_mngr = "add_mngr"
const action_remove_mngr = "rm_mngr"
const action_add_sumngr = "add_sumngr"
const action_remove_sumngr = "rm_sumngr"
const action_op_sumngr = "op_mngr"

type adminHandler func(*ONSClient, http.ResponseWriter, *http.Request) error

const (
	REGISTER_GS1CODE       = 0
	DEREGISTER_GS1CODE     = 1
	ADD_RECORD             = 2
	REMOVE_RECORD          = 3
	REGISTER_SERVICETYPE   = 4
	DEREGISTER_SERVICETYPE = 5
	CHANGE_GS1CODE_STATE   = 6
	CHANGE_RECORD_STATE    = 7
	ADD_MANAGER            = 8
	REMOVE_MANAGER         = 9
	ADD_SUMANAGER          = 10
	REMOVE_SUMANAGER       = 11
	OP_MANAGER             = 12
)

var action_name = map[int32]string{
	0:  "reg",
	1:  "dereg",
	2:  "add_rec",
	3:  "rm_rec",
	4:  "reg_svc",
	5:  "dereg_svc",
	6:  "chx_gstate",
	7:  "chx_rstate",
	8:  "add_mngr",
	9:  "rm_mngr",
	10: "add_sumngr",
	11: "rm_sumngr",
	12: "op_mngr",
}

var action_value = map[string]int32{
	"reg":        0,
	"dereg":      1,
	"add_rec":    2,
	"rm_rec":     3,
	"reg_svc":    4,
	"dereg_svc":  5,
	"chx_gstate": 6,
	"chx_rstate": 7,
	"add_mngr":   8,
	"rm_mngr":    9,
	"add_sumngr": 10,
	"rm_sumngr":  11,
	"op_mngr":    12,
}

var action_handler = map[int32]adminHandler{
	0: actionRegisterHandler,
}

func CheckAction(action string) bool {
	_, ok := action_value[action]

	return ok
}

func AdminProcessAction(h *ONSClient, action string, w http.ResponseWriter, r *http.Request) error {
	_, ok := action_handler[action_value[action]]
	if ok == false {
		return fmt.Errorf("AdminProcessAction: %v action doesn't have handler.", action)
	}
	gs1Code := r.FormValue("gs1code")
	log.Printf("AdminProcessAction : called with %v\n", gs1Code)
	result := make(map[string]string)
	result["action"] = "register gs1 code"
	result["gs1 code"] = gs1Code
	result["public key"] = h.GetSigner().GetPublicKey().AsHex()
	err := writeResponseJson("AdminProcessAction", result, w)
	return err
}

//request is post type.
func actionRegisterHandler(h *ONSClient, w http.ResponseWriter, r *http.Request) error {
	_ = w
	_ = r
	_ = h
	return nil
}

func writeResponseJson(caller string, result interface{}, w http.ResponseWriter) error {
	response, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Fprintf(w, "%s : cannot marshal data into json", caller)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(response))
	return nil
}