package geometry

import (
	"encoding/json"
)

type GeometryType string

const (
	TPoint             GeometryType = "Point"
	TLineString        GeometryType = "LineString"
	TPolygon           GeometryType = "Polygon"
	TMultiPoint        GeometryType = "MultiPointPoint"
	TMultiLineString   GeometryType = "MultiLineString"
	TMultiPolygon      GeometryType = "MultiPolygon"
	TGeomtryCollection GeometryType = "GeometryCollection"
)

type Geometry struct {
	Type        GeometryType       `json:"type"`
	Coordinates GeoType            `json:"coordinates,omitempty"`
	Geometries  GeometryCollection `json:"geometries,omitempty"`
	BBox        []float64          `json:"bbox,omitempty"`
	CCRS        string             `json:"ccrs,omitempty"`
}

type Feature struct {
	Type     GeometryType `json:"type"`
	Geometry Geometry     `json:"geometry"`
}

func (geo *Geometry) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	var rawMessageForType string
	if len(objMap) != 0 {
		err = json.Unmarshal(*objMap["type"], &rawMessageForType)
		if err != nil {
			return err
		}

		switch rawMessageForType {
		case "Point":
			var ph Point
			err = json.Unmarshal(*objMap["coordinates"], &ph)
			if err != nil {
				return err
			}
			geo.Type = TPoint
			geo.Coordinates = ph

		case "LineString":
			var ph LineString
			err = json.Unmarshal(*objMap["coordinates"], &ph)
			if err != nil {
				return err
			}
			geo.Type = TLineString
			geo.Coordinates = ph

		case "Polygon":
			var ph Polygon
			err = json.Unmarshal(*objMap["coordinates"], &ph)
			if err != nil {
				return err
			}
			geo.Type = "Polygon"
			geo.Coordinates = ph

		case "MultiPoint":
			var ph MultiPoint
			err = json.Unmarshal(*objMap["coordinates"], &ph)
			if err != nil {
				return err
			}
			geo.Type = TMultiPoint
			geo.Coordinates = ph

		case "MultiLineString":
			var ph MultiLineString
			err = json.Unmarshal(*objMap["coordinates"], &ph)
			if err != nil {
				return err
			}
			geo.Type = TMultiLineString
			geo.Coordinates = ph

		case "MultiPolygon":
			var ph MultiPolygon
			err = json.Unmarshal(*objMap["coordinates"], &ph)
			if err != nil {
				return err
			}
			geo.Type = "MultiPolygon"
			geo.Coordinates = ph

		case "GeometryCollection":
			var ph GeometryCollection
			err = json.Unmarshal(*objMap["geometries"], &ph)
			if err != nil {
				return err
			}
			geo.Type = "GeometryCollection"
			geo.Geometries = ph

		default:
			geo.Type = ""
			geo.Coordinates = nil
		}
	}
	return nil
}

type GeoType interface {
	GetType() string
}

func (g Geometry) GetType() string {
	return string(g.Type)
}

func (g Geometry) Point() Point {
	return g.Coordinates.(Point)
}

func (g Geometry) LineString() LineString {
	return g.Coordinates.(LineString)
}

func (g Geometry) Polygon() Polygon {
	return g.Coordinates.(Polygon)
}

func (g Geometry) MultiPoint() MultiPoint {
	return g.Coordinates.(MultiPoint)
}

func (g Geometry) MultiLineString() MultiLineString {
	return g.Coordinates.(MultiLineString)
}

func (g Geometry) MultiPolygon() MultiPolygon {
	return g.Coordinates.(MultiPolygon)
}

type Point [2]float64

func (p Point) GetType() string {
	return string(TPoint)
}

type LineString [][2]float64

func (p LineString) GetType() string {
	return string(TLineString)
}

type Polygon [][][2]float64

func (p Polygon) GetType() string {
	return string(TPolygon)
}

type MultiPoint [][2]float64

func (p MultiPoint) GetType() string {
	return string(TMultiPoint)
}

type MultiLineString [][][2]float64

func (p MultiLineString) GetType() string {
	return string(TMultiLineString)
}

type MultiPolygon [][][][2]float64

func (p MultiPolygon) GetType() string {
	return string(TMultiPolygon)
}

type GeometryCollection []Geometry

func (p GeometryCollection) GetType() string {
	return "GeometryCollection"
}

type RawGeoJson map[string]interface{}

func (gj RawGeoJson) GetType() string {
	return gj["type"].(string)
}

func (g RawGeoJson) Point() Point {
	return g["coordinates"].([2]float64)
}

func (g RawGeoJson) LineString() LineString {
	return g["coordinates"].(LineString)
}

func (g RawGeoJson) Polygon() Polygon {
	return g["coordinates"].(Polygon)
}

func (g RawGeoJson) MultiPoint() MultiPoint {
	return g["coordinates"].(MultiPoint)
}

func (g RawGeoJson) MultiLineString() MultiLineString {
	return g["coordinates"].(MultiLineString)
}

func (g RawGeoJson) MultiPolygon() MultiPolygon {
	return g["coordinates"].(MultiPolygon)
}

func (g RawGeoJson) GeometryCollection() GeometryCollection {
	return g["geometries"].(GeometryCollection)
}

func (g RawGeoJson) Feature() Feature {
	geo := g["geometry"].(map[string]interface{})
	return Feature{
		Type:     "Feature",
		Geometry: RawGeoJson(geo).ToGeometry(),
	}
}

func (g RawGeoJson) ToGeometry() Geometry {
	switch g.GetType() {
	case "Point":
		_ = g.Point()
		return Geometry{GeometryType(g.GetType()), g.Point(), nil, nil, ""}
	case "LineString":
		return Geometry{GeometryType(g.GetType()), g.LineString(), nil, nil, ""}
	case "Polygon":
		return Geometry{GeometryType(g.GetType()), g.Polygon(), nil, nil, ""}
	case "MultiPoint":
		return Geometry{GeometryType(g.GetType()), g.MultiPoint(), nil, nil, ""}
	case "MultiLineString":
		return Geometry{GeometryType(g.GetType()), g.MultiLineString(), nil, nil, ""}
	case "MultiPolygon":
		return Geometry{GeometryType(g.GetType()), g.MultiPolygon(), nil, nil, ""}
	case "GeometryCollection":
		return Geometry{GeometryType(g.GetType()), nil, g.GeometryCollection(), nil, ""}
	default:
		return Geometry{}
	}
}
