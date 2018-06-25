package main

import (
	"log"
	"fmt"
	"time"
	"strconv"
	"strings"
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"github.com/golang/protobuf/proto"
	"github.com/daludaluking/ons-sawtooth-sdk/ons_pb2"
	"github.com/daludaluking/ons-sawtooth-sdk/protobuf/batch_pb2"
	"github.com/daludaluking/ons-sawtooth-sdk/protobuf/transaction_pb2"
	"github.com/daludaluking/ons-sawtooth-sdk/signing"
)

var namespace = hexdigestbyString("ons")[:6]

const PARAM_GS1CODE = "gs1code"
const PARAM_ADDRESS = "address"

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

const (
	NONE = iota
	GS1_CODE
	SERVICE_TYPE
)

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

type ONSClient struct {
	Signer *signing.Signer
	Address string
	Port string
}

var iONSHandler *ONSClient = nil

func (ons *ONSClient) GetSigner() *signing.Signer {
	return ons.Signer
}

func MakeSigner(priv_key []byte) *signing.Signer {
	context := signing.CreateContext("secp256k1")
	private_key := signing.NewSecp256k1PrivateKey(priv_key)
	crypto_factory := signing.NewCryptoFactory(context)
	signer := crypto_factory.NewSigner(private_key)
	//public_key := signer.GetPublicKey()
	return signer
}

func NewONSTransactionHalder(address string, port string, privateKey []byte) *ONSClient {
	if iONSHandler == nil {
		iONSHandler = &ONSClient{
			Signer: MakeSigner(privateKey),
			Address: address,
			Port: port,
		}
	}

	return iONSHandler
}

func MakeTransactionPayload(trType int32, params map[string]interface{}) (*ons_pb2.SendONSTransactionPayload, error){
	transaction_payload := &ons_pb2.SendONSTransactionPayload {}

	switch trType {
	case REGISTER_GS1CODE:
		transaction_payload.TransactionType = ons_pb2.SendONSTransactionPayload_REGISTER_GS1CODE
		transaction_payload.RegisterGs1Code = &ons_pb2.SendONSTransactionPayload_RegisterGS1CodeTransactionData {
			Gs1Code: params[PARAM_GS1CODE].(string),
			OwnerId: params[PARAM_ADDRESS].(string),
		}
	default:
		return nil, fmt.Errorf("MakeTransactionPayload : invalid transaction type %v", trType)
	}

	return transaction_payload, nil
}

func MakeBatchList(transaction_payload *ons_pb2.SendONSTransactionPayload, signer *signing.Signer, address string) ([]byte, error) {
	gs1code_reg_payload, err := proto.Marshal(transaction_payload)
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 16)
	if err != nil {
		log.Fatal("Failed to marshal transaction payload:", err)
		return nil, err
	}

	transaction_header := &transaction_pb2.TransactionHeader{
		FamilyName:       "ons",
		FamilyVersion:    "1.0",
		Inputs:           []string{address},
		Outputs:          []string{address},
		BatcherPublicKey: signer.GetPublicKey().AsHex(),
		SignerPublicKey:  signer.GetPublicKey().AsHex(),
		Dependencies:     []string{},
		PayloadSha512:    hexdigestbyByte(gs1code_reg_payload),
		Nonce:            timestamp,
	}

	transaction_header_bytes, err := proto.Marshal(transaction_header)
	if err != nil {
		log.Fatal("Failed to marshal Transaction Header:", err)
		return nil, err
	}

	transaction := &transaction_pb2.Transaction{
		Header:          transaction_header_bytes,
		HeaderSignature: strings.ToLower(hex.EncodeToString(signer.Sign(transaction_header_bytes))),
		Payload:         gs1code_reg_payload,
	}

	batch_header := &batch_pb2.BatchHeader{
		SignerPublicKey: signer.GetPublicKey().AsHex(),
		TransactionIds:  []string{transaction.HeaderSignature},
	}

	batch_header_bytes, err := proto.Marshal(batch_header)
	if err != nil {
		log.Fatal("Failed to marshal Batch Header:", err)
		return nil, err
	}

	batch := &batch_pb2.Batch{
		Header:          batch_header_bytes,
		HeaderSignature: strings.ToLower(hex.EncodeToString(signer.Sign(batch_header_bytes))),
		Transactions:    []*transaction_pb2.Transaction{transaction},
	}

	batch_list := &batch_pb2.BatchList{
		Batches: []*batch_pb2.Batch{batch},
	}

	return proto.Marshal(batch_list)
}

func SendTransactions(ons *ONSClient, batches[]byte) ([]byte, error) {

	resp, err:= http.Post("http://"+ons.Address+":"+ons.Port+"/batches", "application/octet-stream", bytes.NewBuffer(batches))
	if err != nil {
		log.Printf("Failed to send batch list: %v", err)
		return nil, err;
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("SendTransactions: %#v", resp)
	log.Printf("SendTransactions: %#v", string(body))
	return body, err
}

func hexdigestbyString(str string) string {
	hash := sha512.New()
	hash.Write([]byte(str))
	hashBytes := hash.Sum(nil)
	return strings.ToLower(hex.EncodeToString(hashBytes))
}

func hexdigestbyByte(data []byte) string {
	hash := sha512.New()
	hash.Write(data)
	hashBytes := hash.Sum(nil)
	return strings.ToLower(hex.EncodeToString(hashBytes))
}

func MakeAddressByGS1Code(gs1_code string) string{
	return namespace + hexdigestbyString("gs1")[:8] + hexdigestbyString(gs1_code)[:56]
}

func GetONSManagerAddress() string {
	return namespace + hexdigestbyString("ons_manager")[:64]
}

func MakeAddressByServiceType(requestor string, service_type *ons_pb2.ServiceType) (string, error) {
	marshaled_service_type, err := proto.Marshal(service_type)
	if err != nil {
		return "", err
	}
	return namespace + hexdigestbyString("service-type")[:8] + hexdigestbyString(requestor)[:16] + hexdigestbyByte(marshaled_service_type)[:40], nil
}

func GetDataTypeByAddress(address string) int{
	namespace := hexdigestbyString("ons")[:6]

	target_address := namespace + hexdigestbyString("gs1")[:8]
	log.Printf("GetTableIdxByAddress : %s : %s\n", address, target_address)
	if address[:14] == target_address {
		return GS1_CODE
	}

	target_address = namespace + hexdigestbyString("service-type")[:8]
	log.Printf("GetTableIdxByAddress : %s : %s\n", address, target_address)
	if address[:14] == target_address {
		return SERVICE_TYPE
	}
	return NONE
}
