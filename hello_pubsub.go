// Package gopher provides a set of Cloud Functions samples.
package gopher

import (
	"context"
	"fmt"
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

//Item2 ...
type Item2 struct {
	author string
	title  string
	count  int
}

// HelloPubSub consumes a Pub/Sub message.
func HelloPubSub(ctx context.Context, m PubSubMessage) error {
	//ctx := context.Background()
	client, err := bigquery.NewClient(ctx, "jamietest")
	if err != nil {
		return err
	}
	datasetID := "hwtechradar"
	tableID := "mytable"

	// running a select
	// 	q := client.Query(`
	//     select * from hwtechradar.mytable
	// `)
	// 	it, err := q.Read(ctx)

	// 	for {
	// 		var values []bigquery.Value
	// 		err := it.Next(&values)
	// 		if err == iterator.Done {
	// 			break
	// 		}
	// 		if err != nil {
	// 			return err
	// 			// TODO: Handle error.
	// 		}
	// 		fmt.Println(values)
	// 		log.Printf("Hello, %s!", values)
	// 	}

	//

	// Item implements the ValueSaver interface.
	//u := client.Dataset(datasetID).Table(tableID).Inserter()
	u := client.Dataset(datasetID).Table(tableID).Uploader()

	items2 := []*Item2{
		{title: "n1", author: "n11", count: 7},
		{title: "n2", author: "n22", count: 2},
	}
	// if err := u.Put(ctx, items2); err != nil {
	// 	log.Printf("Hello, %s!", err)
	// 	return err
	// }

	//err := u.Put(ctx, items2)
	if err := u.Put(ctx, items2); err != nil {
		if multiError, ok := err.(bigquery.PutMultiError); ok {
			for _, err1 := range multiError {
				for _, err2 := range err1.Errors {
					fmt.Println(err2)
					log.Printf("Err2: %s", err2)
				}
			}
		} else {
			fmt.Println(err)
			log.Printf("Err: %s", err)

		}
	}

	// if err != nil {
	// 	// TODO: Handle error.
	// }
	// u := client.Dataset(datasetID).Table(tableID).Inserter()
	// items := []*Item{
	// 	// Item implements the ValueSaver interface.
	// 	{title: "Phred Phlyntstone", author: "Tyson1"},
	// }
	// if err := u.Put(ctx, items); err != nil {
	// 	return err
	// }
	// name := string(m.Data)
	// if name == "" {
	// 	name = "World"
	// }
	// log.Printf("Hello, %s!", name)
	return nil
}
