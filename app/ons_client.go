package main

import (
	"log"
	"fmt"
	"time"
	"strconv"
	"strings"
	"bytes"
	"reflect"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
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
const PARAM_FLAGS = "flags"
const PARAM_SVC = "service"
const PARAM_REGEXP = "regexp"
const PARAM_INDEX = "index"

const (
	NONE = iota
	GS1_CODE
	SERVICE_TYPE
)

const (
	REGISTER_GS1CODE       = 	ons_pb2.SendONSTransactionPayload_REGISTER_GS1CODE
	DEREGISTER_GS1CODE     = 	ons_pb2.SendONSTransactionPayload_DEREGISTER_GS1CODE
	ADD_RECORD             = 	ons_pb2.SendONSTransactionPayload_ADD_RECORD
	REMOVE_RECORD          = 	ons_pb2.SendONSTransactionPayload_REMOVE_RECORD
	REGISTER_SERVICETYPE   = 	ons_pb2.SendONSTransactionPayload_REGISTER_SERVICETYPE
	DEREGISTER_SERVICETYPE = 	ons_pb2.SendONSTransactionPayload_DEREGISTER_SERVICETYPE
	CHANGE_GS1CODE_STATE   = 	ons_pb2.SendONSTransactionPayload_CHANGE_GS1CODE_STATE
	CHANGE_RECORD_STATE    = 	ons_pb2.SendONSTransactionPayload_CHANGE_RECORD_STATE
	ADD_MANAGER            = 	ons_pb2.SendONSTransactionPayload_ADD_MANAGER
	REMOVE_MANAGER         = 	ons_pb2.SendONSTransactionPayload_REMOVE_MANAGER
	ADD_SUMANAGER          = 	ons_pb2.SendONSTransactionPayload_ADD_SUMANAGER
	REMOVE_SUMANAGER       = 	ons_pb2.SendONSTransactionPayload_REMOVE_SUMANAGER
	OP_MANAGER             = 	ons_pb2.SendONSTransactionPayload_OP_MANAGER
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
	log.SetFlags(log.Lshortfile)
	return iONSHandler
}

/*
type SendONSTransactionPayload struct {
	TransactionType       SendONSTransactionPayload_ONSTransactionType                    `protobuf:"varint,1,opt,name=transaction_type,json=transactionType,enum=SendONSTransactionPayload_ONSTransactionType" json:"transaction_type,omitempty"`
	RegisterGs1Code       *SendONSTransactionPayload_RegisterGS1CodeTransactionData       `protobuf:"bytes,2,opt,name=register_gs1_code,json=registerGs1Code" json:"register_gs1_code,omitempty"`
	DeregisterGs1Code     *SendONSTransactionPayload_DeregisterGS1CodeTransactionData     `protobuf:"bytes,3,opt,name=deregister_gs1_code,json=deregisterGs1Code" json:"deregister_gs1_code,omitempty"`
	AddRecord             *SendONSTransactionPayload_AddRecordTransactionData             `protobuf:"bytes,4,opt,name=add_record,json=addRecord" json:"add_record,omitempty"`
	RemoveRecord          *SendONSTransactionPayload_RemoveRecordTransactionData          `protobuf:"bytes,5,opt,name=remove_record,json=removeRecord" json:"remove_record,omitempty"`
	RegisterServiceType   *SendONSTransactionPayload_RegisterServiceTypeTransactionData   `protobuf:"bytes,6,opt,name=register_service_type,json=registerServiceType" json:"register_service_type,omitempty"`
	DeregisterServiceType *SendONSTransactionPayload_DeregisterServiceTypeTransactionData `protobuf:"bytes,7,opt,name=deregister_service_type,json=deregisterServiceType" json:"deregister_service_type,omitempty"`
	ChangeGs1CodeState    *SendONSTransactionPayload_ChangeGS1CodeStateTransactionData    `protobuf:"bytes,8,opt,name=change_gs1_code_state,json=changeGs1CodeState" json:"change_gs1_code_state,omitempty"`
	ChangeRecordState     *SendONSTransactionPayload_ChangeRecordStateTransactionData     `protobuf:"bytes,9,opt,name=change_record_state,json=changeRecordState" json:"change_record_state,omitempty"`
	AddManager            *SendONSTransactionPayload_AddManagerTransactionData            `protobuf:"bytes,10,opt,name=add_manager,json=addManager" json:"add_manager,omitempty"`
	RemoveManager         *SendONSTransactionPayload_RemoveManagerTransactionData         `protobuf:"bytes,11,opt,name=remove_manager,json=removeManager" json:"remove_manager,omitempty"`
	AddSumanager          *SendONSTransactionPayload_AddSUManagerTransactionData          `protobuf:"bytes,12,opt,name=add_sumanager,json=addSumanager" json:"add_sumanager,omitempty"`
	RemoveSumanager       *SendONSTransactionPayload_RemoveSUManagerTransactionData       `protobuf:"bytes,13,opt,name=remove_sumanager,json=removeSumanager" json:"remove_sumanager,omitempty"`
	OpManager             *SendONSTransactionPayload_OPManagerTransactionData             `protobuf:"bytes,14,opt,name=op_manager,json=opManager" json:"op_manager,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                                                        `json:"-"`
	XXX_unrecognized      []byte                                                          `json:"-"`
	XXX_sizecache         int32                                                           `json:"-"`
}
*/
func MakeTransactionPayload(trType int32, params map[string]interface{}) (*ons_pb2.SendONSTransactionPayload, error){
	transaction_payload := &ons_pb2.SendONSTransactionPayload {}
	transaction_payload.TransactionType = ons_pb2.SendONSTransactionPayload_ONSTransactionType(trType)
	switch transaction_payload.TransactionType {
	case REGISTER_GS1CODE:
		transaction_payload.RegisterGs1Code = &ons_pb2.SendONSTransactionPayload_RegisterGS1CodeTransactionData {
			Gs1Code: params[PARAM_GS1CODE].(string),
			OwnerId: params[PARAM_ADDRESS].(string),
		}
	case DEREGISTER_GS1CODE:
		transaction_payload.DeregisterGs1Code = &ons_pb2.SendONSTransactionPayload_DeregisterGS1CodeTransactionData {
			Gs1Code: params[PARAM_GS1CODE].(string),
		}
	case ADD_RECORD:
		transaction_payload.AddRecord = &ons_pb2.SendONSTransactionPayload_AddRecordTransactionData {
			Gs1Code: params[PARAM_GS1CODE].(string),
			Record: &ons_pb2.SendONSTransactionPayload_RecordTranactionData {
				Flags: params[PARAM_FLAGS].(int32),
				Service: params[PARAM_SVC].(string),
				Regexp: params[PARAM_REGEXP].(string),
			},
		}
	case REMOVE_RECORD:
		transaction_payload.RemoveRecord = &ons_pb2.SendONSTransactionPayload_RemoveRecordTransactionData {
			Gs1Code: params[PARAM_GS1CODE].(string),
			Index: params[PARAM_GS1CODE].(uint32),
		}
	case REGISTER_SERVICETYPE:

	case DEREGISTER_SERVICETYPE:
	default:
		return nil, fmt.Errorf("MakeTransactionPayload : invalid transaction type %v", trType)
	}

	return transaction_payload, nil
}

func retirveGS1Code(payload *ons_pb2.SendONSTransactionPayload) (string, error) {
	vals := reflect.ValueOf(payload).Elem()
	for i := 0; i < vals.NumField(); i++ {
		valueField := vals.Field(i)
		//typeField := vals.Type().Field(i)
		if valueField .Type().Kind().String() == "ptr" && valueField.IsNil() == false {
			gs1codeField := valueField.Elem().FieldByName("Gs1Code")
				if gs1codeField.IsValid() {
					return gs1codeField.String(), nil
				}
		}
	}
	return "", fmt.Errorf("Field to retrive gs1 code from payload.")
}

func MakeTransactionPayloadFromJSON(d *json.Decoder, requestor string) (*ons_pb2.SendONSTransactionPayload, string, error) {
	if d == nil {
		return nil, "", fmt.Errorf("MakeTransactionPayloadFromJSON: decoder is nil.")
	}
	var transaction_payload ons_pb2.SendONSTransactionPayload
	var address string
	var err error
	err = d.Decode(&transaction_payload)

	if err != nil {
		return nil, "", err
	}

	if transaction_payload.TransactionType >= ADD_MANAGER && transaction_payload.TransactionType <= OP_MANAGER {
		address = GetONSManagerAddress()
	}else if transaction_payload.TransactionType == REGISTER_SERVICETYPE{
		address, err = MakeAddressByServiceType(requestor, transaction_payload.RegisterServiceType.ServiceType)
		if err != nil {
			return nil, "", err
		}
	}else if transaction_payload.TransactionType == DEREGISTER_SERVICETYPE {
		address = transaction_payload.DeregisterServiceType.Address
	}else{
		address, err = retirveGS1Code(&transaction_payload)
		log.Printf("retirveGS1Code : gs1code : %v", address)
		address = MakeAddressByGS1Code(address)
		log.Printf("retirveGS1Code : address : %v", address)
		if err != nil {
			return nil, "", err
		}
	}
	return &transaction_payload, address, nil
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
