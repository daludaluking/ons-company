package main

import (
	"github.com/daludaluking/ons-company/ons_pb2"
	"github.com/daludaluking/ons-company/sawtooth_sdk/protobuf/batch_pb2"
	"github.com/daludaluking/ons-company/sawtooth_sdk/protobuf/transaction_pb2"
	"github.com/daludaluking/ons-company/sawtooth_sdk/signing"

	"crypto/sha512"
	"encoding/hex"
//	"encoding/json"

	"strings"
//	"net/http"
//	"bytes"
	"log"
//	"fmt"
	"time"
	"strconv"
	"github.com/golang/protobuf/proto"
)

type ONSClient struct {
	Signer *signing.Signer
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

func NewONSTransactionHalder(privateKey []byte) *ONSClient {
	if iONSHandler == nil {
		iONSHandler = &ONSClient{
			Signer: MakeSigner(privateKey),
		}
	}

	return iONSHandler
}

/*
func ProcessRegisterGS1Code(h *ONSTransactionHalder, gs1Code string) error {

}
*/
func MakeRegisterGS1CodePayload(h *ONSClient, gs1Code string) (*ons_pb2.SendONSTransactionPayload, error) {
	register_gs1_code_payload := &ons_pb2.SendONSTransactionPayload{
		TransactionType: ons_pb2.SendONSTransactionPayload_REGISTER_GS1CODE,
		RegisterGs1Code: &ons_pb2.SendONSTransactionPayload_RegisterGS1CodeTransactionData{
			Gs1Code: gs1Code,
			OwnerId: h.Signer.GetPublicKey().AsHex(),
		},
	}
	return register_gs1_code_payload, nil
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

func hexdigestbyByte(data []byte) string {
	hash := sha512.New()
	hash.Write(data)
	hashBytes := hash.Sum(nil)
	return strings.ToLower(hex.EncodeToString(hashBytes))
}