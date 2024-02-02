package geobesaw

import (
	"encoding/json"
	"net/http"
)

func ReturnStruct(DataStuct any) string {
	jsondata, _ := json.Marshal(DataStuct)
	return string(jsondata)
}

// ---------------------------------------------------------------------- Geojson ----------------------------------------------------------------------

func MembuatGeojsonPoint(mongoenv, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var geojsonpoint GeoJsonPoint
	err := json.NewDecoder(r.Body).Decode(&geojsonpoint)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	PostPoint(mconn, collname, geojsonpoint)
	response.Status = true
	response.Message = "Data point berhasil masuk"

	return ReturnStruct(response)
}

func MembuatGeojsonPolyline(mongoenv, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var geojsonpolyline GeoJsonLineString
	err := json.NewDecoder(r.Body).Decode(&geojsonpolyline)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	PostLinestring(mconn, collname, geojsonpolyline)
	response.Status = true
	response.Message = "Data polyline berhasil masuk"

	return ReturnStruct(response)
}

func MembuatGeojsonPolygon(mongoenv, dbname, collname string, r *http.Request) string {
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)
	var geojsonpolygon GeoJsonPolygon
	err := json.NewDecoder(r.Body).Decode(&geojsonpolygon)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	PostPolygon(mconn, collname, geojsonpolygon)
	response.Status = true
	response.Message = "Data polygon berhasil masuk"

	return ReturnStruct(response)
}

func AmbilDataGeojson(mongoenv, dbname, collname string, r *http.Request) string {
	mconn := SetConnection(mongoenv, dbname)
	datagedung := GetAllBangunan(mconn, collname)
	return ReturnStruct(datagedung)
}

func PostGeoIntersects(mongoenv, dbname, collname string, r *http.Request) string {
	var coordinate Point
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)

	err := json.NewDecoder(r.Body).Decode(&coordinate)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	response.Status = true
	response.Message = GeoIntersects(mconn, collname, coordinate)
	return ReturnStruct(response)
}

func PostGeoWithin(mongoenv, dbname, collname string, r *http.Request) string {
	var coordinate Polygon
	var response Pesan
	response.Status = false
	mconn := SetConnection(mongoenv, dbname)

	err := json.NewDecoder(r.Body).Decode(&coordinate)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	response.Status = true
	response.Message = GeoWithin(mconn, collname, coordinate)

	return ReturnStruct(response)
}

func PostNear(mongoenv, dbname, collname string, r *http.Request) string {
	var coordinate Point
	var response Pesan
	response.Status = false
	mconn := SetConnection2dsphere(mongoenv, dbname, collname)

	err := json.NewDecoder(r.Body).Decode(&coordinate)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	response.Status = true
	response.Message = Near(mconn, collname, coordinate)

	return ReturnStruct(response)
}

func PostNearSphere(mongoenv, dbname, collname string, r *http.Request) string {
	var coordinate Point
	var response Pesan
	response.Status = false
	mconn := SetConnection2dsphere(mongoenv, dbname, collname)

	err := json.NewDecoder(r.Body).Decode(&coordinate)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	response.Status = true
	response.Message = NearSphere(mconn, collname, coordinate)

	return ReturnStruct(response)
}

func PostBox(mongoenv, dbname, collname string, r *http.Request) string {
	var coordinate Polyline
	var response Pesan
	response.Status = false
	mconn := SetConnection2dsphere(mongoenv, dbname, collname)

	err := json.NewDecoder(r.Body).Decode(&coordinate)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	response.Status = true
	response.Message = Box(mconn, collname, coordinate)

	return ReturnStruct(response)
}

func PostCenter(mongoenv, dbname, collname string, r *http.Request) string {
	var coordinate Point
	var response Pesan
	response.Status = false
	mconn := SetConnection2dsphere(mongoenv, dbname, collname)

	err := json.NewDecoder(r.Body).Decode(&coordinate)

	if err != nil {
		response.Message = "Error parsing application/json: " + err.Error()
		return ReturnStruct(response)
	}

	response.Status = true
	response.Message = Center(mconn, collname, coordinate)

	return ReturnStruct(response)
}
