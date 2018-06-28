package main

import (
	"os"
	"fmt"
	"log"
	"context"
	"strings"
	"sync"
	"crypto/sha512"
	"encoding/hex"
	"google.golang.org/api/iterator"
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
	BlockNum 	float64
	BlockId 	string `datastore:",noindex"`
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

func (db *ONSEventDatastoreDB) DBConnect(url string, db_name string, verbose bool) {
	return
}

func (db *ONSEventDatastoreDB) DBDisconnect() {
	return
}

func (db *ONSEventDatastoreDB) DBInitLatestUpdatedBlockInfo(verbose bool) (float64, error) {
	ctx := context.Background()

	if db.Client == nil {
		log.Printf("DBInitLatestUpdatedBlockInfo : Datastore client isn't exist\n")
		return 0, fmt.Errorf("DBInitLatestUpdatedBlockInfo: Datastore client isn't exist")
	}

	key := ds.NameKey(kindStrings[LATEST_BLOCK_INFO_TABLE], "LatestBlockInfo", nil)
	e := new(LatestBlockInfo)
	if err := db.Client.Get(ctx, key, e); err != nil {
		if err != ds.ErrNoSuchEntity {
			log.Printf("DBInitLatestUpdatedBlockInfo : %v\n", err)
			return 0, err
		}else{
			e.BlockNum = 0
			e.BlockId = "0000000000000000"
			e.PrevBlockId = "0000000000000000"
			_, err := db.Client.Put(ctx, key, e)
			if err != nil {
				log.Printf("DBInitLatestUpdatedBlockInfo : %v\n", err)
				return 0, err
			}
		}
		if err := db.Client.Get(ctx, key, e); err != nil {
			log.Printf("DBInitLatestUpdatedBlockInfo : %v\n", err)
			return 0, err
		}
		log.Printf("DBInitLatestUpdatedBlockInfo : latest block info was initialized. %v\n", e)
	}
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
		ctx := context.Background()
		key := ds.NameKey(kindStrings[LATEST_BLOCK_INFO_TABLE], "LatestBlockInfo", nil)
		db.LatestBlock.BlockNum = block_num
		db.LatestBlock.BlockId = block_id
		db.LatestBlock.PrevBlockId = prev_block_id
		_, err := db.Client.Put(ctx, key, db.LatestBlock)
		if err != nil {
			log.Printf("DBUpdateLatestUpdatedBlockInfo : %v\n", err)
			return err
		}
		log.Printf("DBUpdateLatestUpdatedBlockInfo : latest block info was updated : %v\n", db.LatestBlock)
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
	ctx := context.Background()
	key := ds.NameKey(kindStrings[LATEST_BLOCK_INFO_TABLE], "LatestBlockInfo", nil)
	if err := db.Client.Get(ctx, key, db.LatestBlock); err != nil {
		log.Printf("DBInitLatestUpdatedBlockInfo : %v\n", err)
		return 0, err
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

	log.Printf("DBUpdateOrInsert: key: %v, block num: %d, event: %#v\n", pk_v, block_num, v)

	var curBlockNum float64 = 0
	var err error
	var gs1Code *ONSGS1CodeEvent
	var svcType *ONSServiceTypeEvent

	key := ds.NameKey(kindStrings[table_idx], pk_v, nil)
	ctx := context.Background()
	//Transaction이 필요하다. Sync를 사용할 것인지? Transaction을 사용할 것인지? 판단필요.
	//sync보다는 transaction으로? -> sync를 사용 -> transaction이 이상하게 동작하지 않음.. ㅠㅠ
	//Thread에서 비동기적으로 호출되기 때문이다.
	if table_idx == GS1_CODE_TABLE {
		gs1Code := new(ONSGS1CodeEvent)
		if err = db.Client.Get(ctx, key, gs1Code); err == nil {
			curBlockNum = gs1Code.BlockNum
			log.Printf("DBUpdateOrInsert: GS1 Code %v, current block num: %d, updated block num: %d\n", pk_v, curBlockNum, block_num)
		}
	}else{ //SERVICE_TYPE_TABLE
		svcType := new(ONSServiceTypeEvent)
		if err = db.Client.Get(ctx, key, svcType); err == nil {
			curBlockNum = svcType.BlockNum
			log.Printf("DBUpdateOrInsert: Service type %v, current block num: %d, updated block num: %d\n", pk_v, curBlockNum, block_num)
		}
	}

	//항목이 없으면 update.
	if err == ds.ErrNoSuchEntity {
		_, err = db.Client.Put(ctx, key, v)
		if err != nil {
			log.Printf("DBUpdateOrInsert: property put error: %v\n", err)
		}else{
			log.Printf("DBUpdateOrInsert: Insert new event: %#v\n", v)
		}
		return err
	}

	//항목이 있고, current block num for item이
	//현재 update 하려는 item의 block num보다 크거나 같은 경우에는
	//update할 필요가 없다.
	if curBlockNum >= block_num {
		log.Printf("DBUpdateOrInsert: skip... current block num: %v, updated block num: %v\n", curBlockNum, block_num)
		return nil
	}

	_, err = db.Client.Put(ctx, key, v)
	log.Printf("DBUpdateOrInsert: Update event: %#v\n", v)

	if err != nil {
		log.Printf("DBUpdateOrInsert: property put error: %v, %#v, %#v\n", err, gs1Code, svcType)
	}
	return err
}

func (db *ONSEventDatastoreDB) DBDeleteAddress(address string) error {
	if db.LatestBlock == nil {
		log.Printf("DBDeleteAddress: First of all, please call DBInitLatestUpdatedBlockInfo\n")
		return fmt.Errorf("DBDeleteAddress: First of all, please call DBInitLatestUpdatedBlockInfo\n")
	}

	ctx := context.Background()
	var idx int
	if idx = db.GetTableIdxByAddress(address); idx == NONE_TABLE {
		log.Printf("DBDeleteAddress: Invalid address %v\n", address)
		return fmt.Errorf("DBDeleteAddress: Invalid address %v\n", address)
	}
	key := ds.NameKey(kindStrings[idx], address, nil)
	err := db.Client.Delete(ctx, key)
	if err != nil {
		log.Printf("DBDeleteAddress: Failed to delete %v, %v\n", address, err)
		return err
	}
	return nil
}

func CreateDatastoreClient() (*ONSEventDatastoreDB, error){
	var err error
	ctx := context.Background()
	projectID := os.Getenv("GCLOUD_DATASET_ID")
	dsClient, err = ds.NewClient(ctx, projectID)

	log.SetFlags(log.Lshortfile)

	if err != nil {
		return nil, err
	}

	onsEventDsDB = &ONSEventDatastoreDB{
		Client: dsClient,
		Lock: &sync.Mutex{},
	}
	return onsEventDsDB, err
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
	if dsClient == nil {
		return "", fmt.Errorf("GetPrivateKey: Datastore client isn't exist");
	}
	//var entity []ONSEntity
	ctx := context.Background()
	q := ds.NewQuery(kind)

	it := dsClient.Run(ctx, q)

	for ;; {
		var e ONSEntity
		key, err := it.Next(&e)
		if err == iterator.Done {
			return "", fmt.Errorf("GetPrivateKey: Failed to find private key of " + id);
		}
		if err != nil {
			return "", fmt.Errorf("GetPrivateKey: Failed to iterate entities ");
		}

		if key.Name == id {
			privKey, _ := e.V["PrivateKey"]
			return privKey.(string), nil//e.V["PrivateKey"].(string), nil
		}
	}

	return "", fmt.Errorf("GetPrivateKey: Failed to find private key of " + id);
}

func GetONSData(kind string, filter map[string][]string) ([]interface{}, error) {
	if dsClient == nil {
		return nil, fmt.Errorf("GetONSData: Datastore client isn't exist");
	}


	var onsData []interface{}
	var err error
	ctx := context.Background()

	if strings.EqualFold(kind, kindStrings[GS1_CODE_TABLE]) || strings.EqualFold(kind, "ALL")  {
		q := ds.NewQuery(kindStrings[GS1_CODE_TABLE])
		it := dsClient.Run(ctx, q)
		for ;; {
			var e ONSGS1CodeEvent
			_, err = it.Next(&e)
			//log.Printf("GetONSData: %v\n", e)
			if err != nil {
				//log.Printf("GetONSData: gs1code err %v\n", err)
				break
			}
			onsData = append(onsData, e)
		}
	}

	if strings.EqualFold(kind, kindStrings[SERVICE_TYPE_TABLE]) || strings.EqualFold(kind, "ALL")  {
		q := ds.NewQuery(kindStrings[SERVICE_TYPE_TABLE])
		it := dsClient.Run(ctx, q)
		for ;; {
			var e ONSServiceTypeEvent
			_, err = it.Next(&e)
			log.Printf("GetONSData: %v\n", e)
			if err != nil {
				//log.Printf("GetONSData: service type err %v\n", err)
				break
			}
			onsData = append(onsData, e)
		}
	}

	if len(onsData) == 0 {
		return nil, fmt.Errorf("GetONSData: %v data doesn't exist\n", kind)
	}

	if err == iterator.Done {
        err = nil
    }
	return onsData, err
}