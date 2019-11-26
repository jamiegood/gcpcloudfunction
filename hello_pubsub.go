// Package gopher provides a set of Cloud Functions samples.
package gopher

import (
	"context"
	"log"

	"cloud.google.com/go/bigquery"
)

// PubSubMessage is the payload of a Pub/Sub event.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// Item ...
type Item struct {
	title  string `json:"string"`
	author string `json:"string"`
}

// HelloPubSub consumes a Pub/Sub message.
func HelloPubSub(ctx context.Context, m PubSubMessage) error {
	//ctx := context.Background()
	client, err := bigquery.NewClient(ctx, "jamietest")
	if err != nil {
		return err
	}
	datasetID := "jamietest"
	tableID := "hwtechradar"

	u := client.Dataset(datasetID).Table(tableID).Inserter()
	items := []*Item{
		// Item implements the ValueSaver interface.
		{title: "Phred Phlyntstone", author: "Tyson1"},
		{title: "Wylma Phlyntstone2", author: "Tyson2"},
	}
	if err := u.Put(ctx, items); err != nil {
		return err
	}
	name := string(m.Data)
	if name == "" {
		name = "World"
	}
	log.Printf("Hello, %s!", name)
	return nil
}
