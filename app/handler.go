package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type adminHandler func(*ONSClient, http.ResponseWriter, *http.Request) error

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
	ahandler, ok := action_handler[action_value[action]]
	if ok == false {
		return fmt.Errorf("AdminProcessAction: %v action doesn't have handler.", action)
	}
	return ahandler(h, w, r)
}

//request is post type.
func actionRegisterHandler(h *ONSClient, w http.ResponseWriter, r *http.Request) error {
	gs1Code := r.FormValue("gs1code")
	params := map[string]interface{} {
		PARAM_GS1CODE: gs1Code,
		PARAM_ADDRESS: h.GetSigner().GetPublicKey().AsHex(),
	}

	log.Printf("actionRegisterHandler : params with %#v\n", params)

	payload, err := MakeTransactionPayload(REGISTER_GS1CODE, params)
	if err != nil {
		log.Printf("actionRegisterHandler :Failed to call MakeTransactionPayload")
		writeResponse("actionRegisterHandler", err, w)
		return err
	}

	address := MakeAddressByGS1Code(gs1Code)
	batch_list_bytes, err := MakeBatchList(payload, h.GetSigner(), address)
	if err != nil {
		log.Printf("actionRegisterHandler :Failed to call MakeBatchList")
		writeResponse("actionRegisterHandler", err, w)
		return err
	}

	result, err := SendTransactions(h, batch_list_bytes)
	if err != nil {
		log.Printf("actionRegisterHandler :Failed to call SendTransactions")
		writeResponse("actionRegisterHandler", err, w)
		return err
	}

	err = writeResponseJson("actionRegisterHandler", string(result), w)
	return err
}

func writeResponseJson(caller string, result string, w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, result)
	return nil
}

func writeResponse(caller string, result interface{}, w http.ResponseWriter) error {
	response, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Fprintf(w, "%s : cannot marshal data into json", caller)
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(response))
	return nil
}