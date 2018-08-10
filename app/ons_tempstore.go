package main

import (
	//"os"
	"fmt"
	"log"
	//"context"
	"strings"
	"sync"
	"crypto/sha512"
	"encoding/hex"
	//"google.golang.org/api/iterator"
	ds "cloud.google.com/go/datastore"
	//"github.com/daludaluking/ons-sawtooth-sdk/ons_pb2"
)

const (
	GS1_CODE_TABLE = iota
	SERVICE_TYPE_TABLE
	LATEST_BLOCK_INFO_TABLE
	NONE_TABLE
)

var kindStrings = []string{"GS1Code", "ServiceType", "Block"}

var dsClient *ds.Client

/*
	DBInitLatestUpdatedBlockInfo(verbose bool) (float64, error)
	DBConnect(url string, db_name string, verbose bool)
	DBDisconnect()
	DBUpdateLatestUpdatedBlockInfo(block_num float64, block_id string, prev_block_id string) error
	DBGetLatestUpdatedBlockInfo(verbose bool) (float64, error)
	DBGetLatestUpdatedBlock() (float64, string)
	DBUpdateOrInsert(table_idx int, pk_v string, block_num float64, v interface{}) error
	DBDeleteAddress(address string) error
	GetTableIdxByAddress(address string) int
*/
type LatestBlockInfo struct {
	BlockNum	float64
	BlockId         string `datastore:",noindex"`
	PrevBlockId	string `datastore:",noindex"`
}

type ONSEventDatastoreDB struct {
	onsEventDB	ONSEventDB
	Client		*ds.Client
	LatestBlock	*LatestBlockInfo
	Lock		*sync.Mutex
	SavedLatestBlock *LatestBlockInfo
}

var onsEventDsDB *ONSEventDatastoreDB
var gGS1CodeEventData map[string]*ONSGS1CodeEvent
var gSVCTypeData map[string]*ONSServiceTypeEvent


func (db *ONSEventDatastoreDB) DBConnect(url string, db_name string, verbose bool) {
	gGS1CodeEventData = make(map[string]*ONSGS1CodeEvent)
	gSVCTypeData = make(map[string]*ONSServiceTypeEvent)
	return
}

func (db *ONSEventDatastoreDB) DBDisconnect() {
	return
}


func (db *ONSEventDatastoreDB) DBInitLatestUpdatedBlockInfo(verbose bool) (float64, error) {
	e := new(LatestBlockInfo)
	e.BlockNum = 0
	e.BlockId = "0000000000000000"
	e.PrevBlockId = "0000000000000000"
	db.LatestBlock = e
	return db.LatestBlock.BlockNum, nil
}

func (db *ONSEventDatastoreDB) DBUpdateLatestUpdatedBlockInfo(block_num float64, block_id string, prev_block_id string) error {
	if db.LatestBlock == nil {
		log.Printf("DBUpdateLatestUpdatedBlockInfo : First of all, please call DBInitLatestUpdatedBlockInfo\n")
		return fmt.Errorf("DBUpdateLatestUpdatedBlockInfo : First of all, please call DBInitLatestUpdatedBlockInfo\n")
	}

	db.Lock.Lock()
	defer db.Lock.Unlock()

	if db.LatestBlock.BlockNum < block_num {
		db.LatestBlock.BlockNum = block_num
		db.LatestBlock.BlockId = block_id
		db.LatestBlock.PrevBlockId = prev_block_id
		return nil
	}
	log.Printf("DBUpdateLatestUpdatedBlockInfo : current latest block info is up-to-date : %v\n", db.LatestBlock)
	return nil
}

func (db *ONSEventDatastoreDB) DBGetLatestUpdatedBlockInfo(verbose bool) (float64, error) {
	if db.LatestBlock == nil {
		log.Printf("DBGetLatestUpdatedBlockInfo : First of all, please call DBInitLatestUpdatedBlockInfo\n")
		return 0, fmt.Errorf("DBGetLatestUpdatedBlockInfo : First of all, please call DBInitLatestUpdatedBlockInfo\n")
	}
	return db.LatestBlock.BlockNum, nil
}

func (db *ONSEventDatastoreDB) DBGetLatestUpdatedBlock() (float64, string) {
	if db.LatestBlock == nil {
		log.Printf("DBGetLatestUpdatedBlock : First of all, please call DBInitLatestUpdatedBlockInfo\n")
		return 0, ""
	}
	return db.LatestBlock.BlockNum, db.LatestBlock.BlockId
}

func (db *ONSEventDatastoreDB) DBSaveLatestUpdatedBlockInfo(block_num float64, block_id string, prev_block_id string) {
	db.Lock.Lock()
	defer db.Lock.Unlock()
	if db.SavedLatestBlock == nil {
		db.SavedLatestBlock = new(LatestBlockInfo)
	}
	db.SavedLatestBlock.BlockNum = block_num
	db.SavedLatestBlock.BlockId = block_id
	db.SavedLatestBlock.PrevBlockId = prev_block_id
	return
}

func (db *ONSEventDatastoreDB) DBUpdateSavedLatestUpdatedBlockInfo() {
	if db.LatestBlock != nil {
		db.DBUpdateLatestUpdatedBlockInfo(db.SavedLatestBlock.BlockNum, db.SavedLatestBlock.BlockId, db.SavedLatestBlock.PrevBlockId)
	}
}

func (db *ONSEventDatastoreDB) DBUpdateOrInsert(table_idx int, pk_v string, block_num float64, v interface{}) error {
	if db.LatestBlock == nil {
		log.Printf("DBUpdateOrInsert: First of all, please call DBInitLatestUpdatedBlockInfo\n")
		return fmt.Errorf("DBUpdateOrInsert: First of all, please call DBInitLatestUpdatedBlockInfo\n")
	}

	if len(kindStrings) < table_idx {
		log.Printf("DBUpdateOrInsert: Invalid table index : %d.\n", table_idx)
		return fmt.Errorf("DBUpdateOrInsert: Invalid table index")
	}

	db.Lock.Lock()
	defer db.Lock.Unlock()

        if gGS1CodeEventData == nil {
                gGS1CodeEventData = make(map[string]*ONSGS1CodeEvent)
        }

        if gSVCTypeData == nil {
                gSVCTypeData = make(map[string]*ONSServiceTypeEvent)
        }

	log.Printf("DBUpdateOrInsert: key: %v, block num: %d, event: %#v\n", pk_v, block_num, v)

	var curBlockNum float64 = 0
	var gs1Code *ONSGS1CodeEvent
	var svcType *ONSServiceTypeEvent

	key := kindStrings[table_idx]+pk_v
	var ok bool
	//Transaction이 필요하다. Sync를 사용할 것인지? Transaction을 사용할 것인지? 판단필요.
	//sync보다는 transaction으로? -> sync를 사용 -> transaction이 이상하게 동작하지 않음.. ㅠㅠ
	//Thread에서 비동기적으로 호출되기 때문이다.
	if table_idx == GS1_CODE_TABLE {
		if gs1Code, ok = gGS1CodeEventData[key]; ok == true {
			curBlockNum = gs1Code.BlockNum
		}
	}else{ //SERVICE_TYPE_TABLE
		if svcType, ok = gSVCTypeData[key]; ok == true {
			curBlockNum = svcType.BlockNum
		}
	}

	//항목이 없으면 update.
	if ok == false || curBlockNum < block_num {
		org, ok := v.(*ONSGS1CodeEvent)
		if ok == true {
			gGS1CodeEventData[key] = org
			return nil
		}
		org2, ok := v.(*ONSServiceTypeEvent)
		if ok == true {
			gSVCTypeData[key] = org2
			return nil
		}
		return fmt.Errorf("DBUpdateOrInsert: cannot casting interface\n")
	}

	log.Printf("DBUpdateOrInsert: skip... current block num: %v, updated block num: %v\n", curBlockNum, block_num)
	return nil

}

func (db *ONSEventDatastoreDB) DBDeleteAddress(address string) error {
	if db.LatestBlock == nil {
		log.Printf("DBDeleteAddress: First of all, please call DBInitLatestUpdatedBlockInfo\n")
		return fmt.Errorf("DBDeleteAddress: First of all, please call DBInitLatestUpdatedBlockInfo\n")
	}

	var idx int
	if idx = db.GetTableIdxByAddress(address); idx == NONE_TABLE {
		log.Printf("DBDeleteAddress: Invalid address %v\n", address)
		return fmt.Errorf("DBDeleteAddress: Invalid address %v\n", address)
	}
	return nil
}

func CreateDatastoreClient() (*ONSEventDatastoreDB, error){
	log.SetFlags(log.Lshortfile)
	onsEventDsDB = &ONSEventDatastoreDB{
		Lock: &sync.Mutex{},
	}
	return onsEventDsDB, nil
}

func hexdigest(str string) string {
	hash := sha512.New()
	hash.Write([]byte(str))
	hashBytes := hash.Sum(nil)
	return strings.ToLower(hex.EncodeToString(hashBytes))
}

func (db *ONSEventDatastoreDB) GetTableIdxByAddress(address string) int{
	namespace := hexdigest("ons")[:6]

	target_address := namespace + hexdigest("gs1")[:8]
	log.Printf("GetTableIdxByAddress : %s : %s\n", address, target_address)
	if address[:14] == target_address {
		return GS1_CODE_TABLE
	}

	target_address = namespace + hexdigest("service-type")[:8]
	log.Printf("GetTableIdxByAddress : %s : %s\n", address, target_address)
	if address[:14] == target_address {
		return SERVICE_TYPE_TABLE
	}

	return NONE_TABLE
}

type ONSProperty map[string]interface{}

type ONSEntity struct {
	V ONSProperty
	K *ds.Key `datastore:"__key__"`
}

func (m *ONSEntity) Load(ps []ds.Property) error {
	//fmt.Printf("%#v\n", ps)
	m.V = make(ONSProperty)
	for _, prop := range ps {
		//fmt.Printf("%#v\t%#v\t%#v\n", m.V, prop.Name, prop.Value)
		m.V[prop.Name] = prop.Value
	}
	return nil
}

func (m *ONSEntity) Save() ([]ds.Property, error) {
	var prop []ds.Property
	for key, value := range (*m).V {
		prop = append(prop, ds.Property{
			Name: key,
			Value: value,
			NoIndex: true,
		})
	}
	return prop, nil
}

func GetPrivateKey(kind string, id string) (string, error) {
	return "", fmt.Errorf("GetPrivateKey: Failed to find private key of " + id);
}

func GetONSData(kind string, filter map[string][]string) ([]interface{}, error) {

	var onsData []interface{}

	if gGS1CodeEventData == nil {
		gGS1CodeEventData = make(map[string]*ONSGS1CodeEvent)
	}

	if gSVCTypeData == nil {
		gSVCTypeData = make(map[string]*ONSServiceTypeEvent)
	}

	if strings.EqualFold(kind, kindStrings[GS1_CODE_TABLE]) || strings.EqualFold(kind, "ALL")  {
		for _, v := range gGS1CodeEventData {
			onsData = append(onsData, v)
		}
	}

	if strings.EqualFold(kind, kindStrings[SERVICE_TYPE_TABLE]) || strings.EqualFold(kind, "ALL")  {
		for _, v := range gSVCTypeData  {
			onsData = append(onsData, v)
		}
	}

	if len(onsData) == 0 {
		return nil, fmt.Errorf("GetONSData: %v data doesn't exist\n", kind)
	}

	return onsData, nil
}
