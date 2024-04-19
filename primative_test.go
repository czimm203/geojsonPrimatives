package geometry_test

import (
	"encoding/json"
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
	var thing testJson
	err = json.Unmarshal(data, &thing)
	if err != nil {
		t.Error(err)
	}

	js, err := json.MarshalIndent(thing, "", "  ")
	if err != nil {
		t.Error(err)
	}
	os.WriteFile("./out.json", js, 0644)
}

func TestGeometryPoint(t *testing.T) {
	p := geometry.Point{2.0, 2.0}
	_ = geometry.Geometry{Type: "Polygon", Coordinates: p}
}

func TestGeometryCollection(t *testing.T) {
	type GCFeat struct {
		Type     string `json:"type"`
		Geometry struct {
			Type       string                      `json:"type"`
			Geometries geometry.GeometryCollection `json:"geometries"`
		} `json:"geometry"`
	}
	raw, err := os.ReadFile("./test/SCC053.json")
	if err != nil {
		t.Error(err)
	}
	var feat geometry.Feature
	err = json.Unmarshal(raw, &feat)
	if err != nil {
		t.Error(err)
	}
    for _, f := range feat.Geometry.Geometries {
        t.Logf("%+v", f.MultiPolygon())
    }

    raw2, err := json.MarshalIndent(feat,"","    ")
    os.WriteFile("./out2.json", raw2, 0644)
}
