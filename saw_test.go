package geobesaw

import (
	"fmt"
	"testing"
)

var dbname = "gis"
var collname = "gis"

func TestGeoIntersects(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Point{
		Coordinates: []float64{
			106.82524765255425, -6.4764977535096335,
		},
	}
	datagedung := GeoIntersects(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestGeoWithin(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", dbname)
	coordinates := Polygon{
		Coordinates: [][][]float64{
			{
				{106.82419629896253, -6.475156621396806},
				{106.82371151727801, -6.4775568189086385},
				{106.82481252764393, -6.476063227359916},
				{106.82499276799916, -6.475289390836494},
				{106.82587780809916, -6.476708948350165},
			},
		},
	}
	datagedung := GeoWithin(mconn, collname, coordinates)
	fmt.Println(datagedung)
}

func TestNear(t *testing.T) {
	mconn := SetConnection2dsphere("MONGOSTRING", "gis", "gis")
	coordinates := Point{
		Coordinates: []float64{
			106.82422765859621, -6.477255989541547,
		},
	}
	datagedung := Near(mconn, "gis", coordinates)
	fmt.Println(datagedung)
}

func TestNearSphere(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gis")
	coordinates := Point{
		Coordinates: []float64{
			106.82398588225232, -6.47670045038231,
		},
	}
	datagedung := NearSphere(mconn, "gis", coordinates)
	fmt.Println(datagedung)
}

func TestBox(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gis")
	coordinates := Polyline{
		Coordinates: [][]float64{
			{106.82491558268919, -6.475313159288788},
			{106.82443167179292, -6.477590851097517},
		},
	}
	datagedung := Box(mconn, collname, coordinates)
	fmt.Println(datagedung)
}