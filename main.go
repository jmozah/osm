package main

import (
	"context"
	"encoding/json"
	"github.com/paulmach/osm"
	"github.com/paulmach/osm/osmpbf"
	"os"
)

func main() {
	f, err := os.Open("/Users/jmozah/osm/switzerland/switzerland-latest.osm.pbf")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := osmpbf.New(context.Background(), f, 3)
	defer scanner.Close()

	nodeCount := 0
	wayCount := 0
	relationCount := 0
	otherCount := 0


	fa, err := os.Create("/Users/jmozah/osm/switzerland/processed/switzerland.json")
	if err != nil {
		return
	}
	defer fa.Close()

	for scanner.Scan() {
		switch o := scanner.Object().(type) {
		case *osm.Node:
			data, _ := json.Marshal(o)
			_, _ = fa.WriteString(string(data))
			_, _ = fa.WriteString("\n")
			nodeCount++
		case *osm.Way:
			data, _ := json.Marshal(o)
			_, _ = fa.WriteString(string(data))
			_, _ = fa.WriteString("\n")
			wayCount++
		case *osm.Relation:
			data, _ := json.Marshal(o)
			_, _ = fa.WriteString(string(data))
			_, _ = fa.WriteString("\n")
			relationCount++
		default:
			otherCount++
		}
	}
	scanErr := scanner.Err()
	if scanErr != nil {
		panic(scanErr)
	}
}
