package main

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	"encoding/json"
	"net/http"
	"github.com/daludaluking/ons-sawtooth-sdk/ons_pb2"
)

type adminHandler func(*ONSClient, http.ResponseWriter, *http.Request) error

/*
	parameter list
		1. action
		2. gs1code
		3. address
*/
var action_name = map[int32]string{
	0:  "reg",
	1:  "dereg",
	2:  "addrec",
	3:  "rmrec",
	4:  "regsvc",
	5:  "deregsvc",
	6:  "chxgstate",
	7:  "chxrstate",
	8:  "addmngr",
	9:  "rmmngr",
	10: "addsumngr",
	11: "rmsumngr",
	12: "opmngr",
}

var action_value = map[string]int32{
	"reg":        0,
	"dereg":      1,
	"addrec":    2,
	"rmrec":     3,
	"regsvc":    4,
	"deregsvc":  5,
	"chxgstate": 6,
	"chxrstate": 7,
	"addmngr":   8,
	"rmmngr":    9,
	"addsumngr": 10,
	"rmsumngr":  11,
	"opmngr":    12,
}

var action_handler = map[int32]adminHandler{
	0: actionHandlerForGS1Code,
	1: actionHandlerForGS1Code,
	2: actionHandlerForRecord,
	3: actionHandlerForRecord,
}

func CheckAction(action string) bool {
	_, ok := action_value[strings.ToLower(action)]
	return ok
}

func TxProcessAction(h *ONSClient, action string, w http.ResponseWriter, r *http.Request) error {
	contentType := r.Header.Get("Content-type")
	if contentType == "application/json" {
		//json형태로 받아서 처리하자... 너무 쉬워진다...
		var requestor string
		decoder := json.NewDecoder(r.Body)
		if h == nil {
			//just for test...
			requestor = "test"
		}else{
			requestor = h.GetSigner().GetPublicKey().AsHex()
		}
		payload, address, err:= MakeTransactionPayloadFromJSON(decoder, requestor)
		if err != nil {
			return err
		}
		log.Printf("%v, %v\n", payload, address)
		return sendTX(h, "actionHandlerForGS1Code", address, payload, w)
	}
	ahandler, ok := action_handler[action_value[strings.ToLower(action)]]
	if ok == false {
		return fmt.Errorf("TxProcessAction: %v action doesn't have handler.", action)
	}
	return ahandler(h, w, r)
}

func sendTX(h *ONSClient, action string, address string, payload *ons_pb2.SendONSTransactionPayload, w http.ResponseWriter) error {
	batch_list_bytes, err := MakeBatchList(payload, h.GetSigner(), address)
	if err != nil {
		log.Printf("sendTX (action : %v) : Failed to call MakeBatchList", action)
		WriteResponse("sendTX", err, w)
		return err
	}

	result, err := SendTransactions(h, batch_list_bytes)
	if err != nil {
		log.Printf("sendTX (action : %v) : Failed to call SendTransactions", action)
		WriteResponse("sendTX", err, w)
		return err
	}
	err = WriteResponseJson(action, string(result), w)
	return err
}

//request is post type.
func actionHandlerForGS1Code(h *ONSClient, w http.ResponseWriter, r *http.Request) error {
	var trType ons_pb2.SendONSTransactionPayload_ONSTransactionType
	action := r.FormValue("action")
	if strings.EqualFold(action, "reg") {
		trType = REGISTER_GS1CODE
	}else if strings.EqualFold(action, "dereg") {
		trType = DEREGISTER_GS1CODE
	}else{
		err := fmt.Errorf("actionHandlerForGS1Code : Invalid action : The action must be one of 'reg' or 'dreg'.")
		WriteResponse("actionHandlerForGS1Code", err, w)
		return err
	}

	gs1Code := r.FormValue("gs1code")
	params := map[string]interface{} {
		PARAM_GS1CODE: gs1Code,
		PARAM_ADDRESS: h.GetSigner().GetPublicKey().AsHex(), //if trType == DEREGISTER_GS1CODE, it will be ignored.
	}

	log.Printf("actionHandlerForGS1Code : trType = %v, params with %#v\n", trType, params)

	payload, err := MakeTransactionPayload(int32(trType), params)
	if err != nil {
		log.Printf("actionHandlerForGS1Code :Failed to call MakeTransactionPayload")
		WriteResponse("actionHandlerForGS1Code", err, w)
		return err
	}
	address := MakeAddressByGS1Code(gs1Code)
	return sendTX(h, "actionHandlerForGS1Code", address, payload, w)
}

func actionHandlerForRecord(h *ONSClient, w http.ResponseWriter, r *http.Request) error {
	var trType ons_pb2.SendONSTransactionPayload_ONSTransactionType
	var index uint64 = 0
	var flags uint64 = 0
	var err error
	action := r.FormValue("action")
	if strings.EqualFold(action, "addRec") {
		trType = ADD_RECORD
		sflags := r.FormValue("flags")
		if sflags == "" {
			err = fmt.Errorf("actionHandlerForRecord : Flags value doesn't exist.")
			WriteResponse("actionHandlerForRecord", err, w)
			return err
		}
		if flags, err = strconv.ParseUint(sflags, 10, 32); err != nil {
			err = fmt.Errorf("actionHandlerForRecord : Flags value is not a numeric type.")
			WriteResponse("actionHandlerForRecord", err, w)
			return err
		}
	}else if strings.EqualFold(action, "rmRec")  {
		trType = REMOVE_RECORD
		sidx := r.FormValue("index")
		if sidx == "" {
			err = fmt.Errorf("actionHandlerForRecord : Index value doesn't exist.")
			WriteResponse("actionHandlerForRecord", err, w)
			return err
		}
		if index, err = strconv.ParseUint(sidx, 10, 32); err != nil {
			err = fmt.Errorf("actionHandlerForRecord : Index value is not a numeric type.")
			WriteResponse("actionHandlerForRecord", err, w)
			return err
		}
	}else{
		err = fmt.Errorf("actionHandlerForRecord : Invalid action : The action must be one of 'addRec' or 'rmRec'.")
		WriteResponse("actionHandlerForRecord", err, w)
		return err
	}

	gs1Code := r.FormValue("gs1code")
	params := map[string]interface{} {
		PARAM_GS1CODE: gs1Code,
		PARAM_FLAGS: uint32(flags),
		PARAM_SVC: r.FormValue("service"),
		PARAM_REGEXP: r.FormValue("regexp"),
		PARAM_INDEX: uint32(index),
	}

	log.Printf("actionHandlerForRecord : trType = %v, params with %#v\n", trType, params)

	payload, err := MakeTransactionPayload(int32(trType), params)
	if err != nil {
		log.Printf("actionHandlerForGS1Code :Failed to call MakeTransactionPayload")
		WriteResponse("actionHandlerForGS1Code", err, w)
		return err
	}
	address := MakeAddressByGS1Code(gs1Code)
	return sendTX(h, "actionHandlerForRecord", address, payload, w)
}