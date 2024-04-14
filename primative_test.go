package geometry_test

import (
	"encoding/json"
	"fmt"
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
	Features []geometry.Feature
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
	os.WriteFile("./out.json", js, 0644)
}

func TestGeometryPoint(t *testing.T) {
    p := geometry.Point{2.0,2.0}
    g := geometry.Geometry{Type: "Polygon", Coordinates: p}
    t.Logf("%+v\n",g.Point())
}

func TestGeoJsonMap(t *testing.T) {
    type testMap struct {
        Type string
        Features []geometry.RawGeoJson
    }
	data, err := loadJson()
	if err != nil {
		t.Error(err)
	}
	// thing := make(map[string]interface{})
	var thing testMap
	err = json.Unmarshal(data, &thing)
	if err != nil {
		t.Error(err)
	}
    if thing.Features[0].GetType() != "Feature" {
        t.Errorf("Failed to parse RawGeoJson: %+v", thing)
    }

	js, err := json.MarshalIndent(thing, "", "  ")
	if err != nil {
		t.Error(err)
	}
	os.WriteFile("./out.json", js, 0644)

}
