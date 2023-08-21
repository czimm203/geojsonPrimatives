package primative 

import (
	"encoding/json"
)

type GeometryType string

const (
	TPoint           GeometryType = "Point"
	TLineString      GeometryType = "LineString"
	TPolygon         GeometryType = "Polygon"
	TMultiPoint      GeometryType = "MultiPointPoint"
	TMultiLineString GeometryType = "MultiLineString"
	TMultiPolygon    GeometryType = "MultiPolygon"
)

type Geometry struct {
	Type        GeometryType  `json:"type"`
	Coordinates GeoType       `json:"coordinates"`
    BBox        []float64    `json:"bbox,omitempty"`
    CCRS        string        `json:"ccrs,omitempty"`
}

func (geo *Geometry) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}
	var rawMessageForType string
	// var rawMessageForBbox string
	if len(objMap) != 0 {
		err = json.Unmarshal(*objMap["type"], &rawMessageForType)
		if err != nil {
			return err
		}

        // if *objMap["bbox"] != nil {
        //     err = json.Unmarshal(*objMap["bbox"], &rawMessageForBbox)
        //     if err != nil {
        //         return err
        //     }
        // } else {
        //     geo.BBox = nil
        // }

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
                        
		default:
			geo.Type = ""
			geo.Coordinates = nil
		}
	}
	return nil
}

type GeoType interface {
    GetType() GeometryType
}

type Point [2]float64

func (p Point) GetType() GeometryType {
	return TPoint
}

type LineString [][2]float64

func (p LineString) GetType() GeometryType {
	return TLineString
}

type Polygon [][][2]float64 

func (p Polygon) GetType() GeometryType {
	return TPolygon
}

type MultiPoint [][2]float64

func (p MultiPoint) GetType() GeometryType {
	return TMultiPoint
}

type MultiLineString [][][2]float64

func (p MultiLineString) GetType() GeometryType {
	return TMultiLineString
}

type MultiPolygon [][][][2]float64

func (p MultiPolygon) GetType() GeometryType {
	return TMultiPolygon
}
