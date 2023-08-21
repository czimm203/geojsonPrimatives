package primative_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/czimm203/geojsonPrimatives"
)

func loadJson() ([]byte, error) {
	data, err := os.ReadFile("./test.json")
	if err != nil {
		return nil, err
	}
	return data, nil
}

type testJson struct {
	Type     string
	Features []feature
}

type feature struct {
	Type     string
	Geometry primative.Geometry
}

func TestGeoJson(t *testing.T) {
	data, err := loadJson()
	if err != nil {
		t.Error(err)
	}

	// thing := make(map[string]interface{})
	var thing testJson

	err = json.Unmarshal(data, &thing)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", thing)

	js, err := json.MarshalIndent(thing, "", "  ")
	if err != nil {
		t.Error(err)
	}

	ioutil.WriteFile("./out.json", js, 0644)
}
