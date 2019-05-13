package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	//"github.com/olivere/elastic"
	elastic "gopkg.in/olivere/elastic.v5"
)

// Monitoring is a structure used for serializing/deserializing data in Elasticsearch.
type Monitoring struct {
	Project    string    `json:"project"`
	Body       string    `json:"message"`
	Counter    int       `json:"counter"`
	Created    time.Time `json:"created,omitempty"`
	ProbePoint string    `json:"probepoint,omitempty"`
}

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings" : {
		"http" : {
			"properties" : {
				"project":{
					"type":"keyword"
				},
				"body":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"created":{
					"type":"date"
				},
				"probepoint":{
					"type":"geo_point"
				}
			}
		},
		"dns" : {
			"properties" : {
				"project":{
					"type":"keyword"
				},
				"body":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"created":{
					"type":"date"
				},
				"probepoint":{
					"type":"geo_point"
				}
			}
		}
	}
}`

func connectES() (context.Context, *elastic.Client) {
	// Starting with elastic.v5, you must pass a context to execute each service
	ctx := context.Background()

	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
		panic(err)
	}
	return ctx, client
}

func createIndex(ctx context.Context, client *elastic.Client, _index string) {
	// Starting with elastic.v5, you must pass a context to execute each service
	fmt.Println("** createIndex()")

	exists, err := client.IndexExists(_index).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		fmt.Println("Creating index..")
		// Create a new index.
		createIndex, err := client.CreateIndex(_index).BodyString(mapping).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}
}

func registerDocumentJSON(ctx context.Context, client *elastic.Client, _index, _type, _id string, _body Monitoring) {
	// Starting with elastic.v5, you must pass a context to execute each service
	fmt.Println("** registDocumentJSON()")

	put1, err := client.Index().
		Index(_index).
		Type(_type).
		Id(_id).
		BodyJson(_body).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed http %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}

func registerDocumentString(ctx context.Context, client *elastic.Client, _index, _type, _id, _body string) {
	// Starting with elastic.v5, you must pass a context to execute each service
	fmt.Println("** registDocumentString()")

	put1, err := client.Index().
		Index(_index).
		Type(_type).
		Id(_id).
		BodyJson(_body).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed http %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}

func showDocument(ctx context.Context, client *elastic.Client, _index, _type, _id string) {
	// Starting with elastic.v5, you must pass a context to execute each service
	fmt.Println("** showDocument()")

	// Get http with specified ID
	get1, err := client.Get().
		Index(_index).
		Type(_type).
		Id(_id).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}
}

func flush(ctx context.Context, client *elastic.Client, _index string) {
	// Starting with elastic.v5, you must pass a context to execute each service
	fmt.Println("** flush()")

	// Flush to make sure the documents got written.
	client.Flush().Index(_index).Do(ctx)
}

func searchWithTermQuery(ctx context.Context, client *elastic.Client, _index string, _term_query *elastic.TermQuery, _sort_field string, _sort_asc bool) {
	// Starting with elastic.v5, you must pass a context to execute each service
	fmt.Println("** searchWithTermQuery()")

	searchResult, err := client.Search().
		Index(_index).                // search in index "twitter"
		Query(_term_query).           // specify the query
		Sort(_sort_field, _sort_asc). // sort by "user" field, ascending
		From(0).Size(10).             // take documents 0-9
		Pretty(true).                 // pretty print request and response JSON
		Do(ctx)                       // execute
	if err != nil {
		// Handle error
		panic(err)
	}

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// Each is a convenience function that iterates over hits in a search result.
	// It makes sure you don't need to check for nil values in the response.
	// However, it ignores errors in serialization. If you want full control
	// over iterating the hits, see below.
	var ttyp Monitoring
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if t, ok := item.(Monitoring); ok {
			fmt.Printf("Monitoring by %s: %s\n", t.Project, t.Body)
		}
	}
	// TotalHits is another convenience function that works even when something goes wrong.
	fmt.Printf("Found a total of %d http\n", searchResult.TotalHits())

	// Here's how you iterate through results with full control over each step.
	if searchResult.Hits.TotalHits > 0 {
		fmt.Printf("Found a total of %d http\n", searchResult.Hits.TotalHits)

		// Iterate through results
		for _, hit := range searchResult.Hits.Hits {
			// hit.Index contains the name of the index

			// Deserialize hit.Source into a Monitoring (could also be just a map[string]interface{}).
			var t Monitoring
			err := json.Unmarshal(*hit.Source, &t)
			if err != nil {
				// Deserialization failed
			}

			// Work with http
			fmt.Printf("Monitoring by %s: %s\n", t.Project, t.Body)
		}
	} else {
		// No hits
		fmt.Print("Found no http\n")
	}
}

func updateDocument(ctx context.Context, client *elastic.Client, _index, _type, _id string) {
	// Starting with elastic.v5, you must pass a context to execute each service
	fmt.Println("** updateDocument()")

	update, err := client.Update().Index(_index).Type(_type).Id(_id).
		Script(elastic.NewScriptInline("ctx._source.counter += params.num").Lang("painless").Param("num", 1)).
		Upsert(map[string]interface{}{"counter": 0}).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("New version of http %q is now %d\n", update.Id, update.Version)
}

func deleteIndex(ctx context.Context, client *elastic.Client, _index string) {
	// Starting with elastic.v5, you must pass a context to execute each service
	fmt.Println("** deleteIndex()")

	deleteIndex, err := client.DeleteIndex(_index).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !deleteIndex.Acknowledged {
		// Not acknowledged
	}
}

func main() {
	ctx, client := connectES()

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Use the IndexExists service to check if a specified index exists.
	createIndex(ctx, client, "anitya")

	// Index a http (using JSON serialization)
	http1 := Monitoring{Project: "demo", Body: "Take Five", Counter: 0}
	registerDocumentJSON(ctx, client, "anitya", "http", "1", http1)
	// Index a second http (by string)
	http2 := `{"project" : "demo", "body" : "It's a Raggy Waltz"}`
	registerDocumentString(ctx, client, "anitya", "http", "2", http2)
	http3 := Monitoring{Project: "demo", Body: "Hi Five", Counter: 0}
	registerDocumentJSON(ctx, client, "anitya", "http", "3", http3)

	dns1 := Monitoring{Project: "demo", Body: "www.yahoo.co.jp IN A 182.22.25.124", Counter: 0}
	registerDocumentJSON(ctx, client, "anitya", "dns", "1", dns1)
	dns2 := Monitoring{Project: "demo", Body: "www.yahoo.co.jp IN A 182.22.25.124", Counter: 0}
	registerDocumentJSON(ctx, client, "anitya", "dns", "2", dns2)
	dns3 := Monitoring{Project: "demo", Body: "www.yahoo.co.jp IN A 182.22.25.124", Counter: 0}
	registerDocumentJSON(ctx, client, "anitya", "dns", "3", dns3)

	//showDocument(ctx, client, "anitya", "http", "1")
	//flush(ctx, client, "anitya")

	// Search with a term query
	//termQuery := elastic.NewTermQuery("project", "demo")
	//searchWithTermQuery(ctx, client, "anitya", termQuery, "created", true)

	// Update a http by the update API of Elasticsearch.
	// We just increment the number of counter.
	//updateDocument(ctx, client, "anitya", "http", "1")
	/*
		// Delete an index.
		//deleteIndex(ctx, client, "anitya")
	*/
}
