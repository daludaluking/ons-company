package main

import (
	"os"
	"fmt"
	"context"
	"google.golang.org/api/iterator"
	ds "cloud.google.com/go/datastore"
)

var dsClient *ds.Client

func CreateDatastoreClient() error{
	var err error
	ctx := context.Background()
	projectID := os.Getenv("GCLOUD_DATASET_ID")
	dsClient, err = ds.NewClient(ctx, projectID)
	return err
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
/*
	keys, err := dsClient.GetAll(ctx, q, &entity)
	if err != nil {
		fmt.Printf("%v, %#v\n", kind, err)
		return "", err
	}
	for i, key := range keys {
		if key.Name == id {
			return entity[i].V["PrivateKey"].(string), nil
		}
	}
	//fmt.Printf("%v\n", properties)
*/
	return "", fmt.Errorf("GetPrivateKey: Failed to find private key of " + id);
}