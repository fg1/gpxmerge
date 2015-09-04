// Various structures used for unmarshaling GPX files
package main

import (
	"encoding/xml"
)

type TrackPoint struct {
	Lat float64 `xml:"lat,attr"`
	Lon float64 `xml:"lon,attr"`
	Ele float64 `xml:"ele"`
}

type Track struct {
	Name         string       `xml:"name"`
	TrackSegment []TrackPoint `xml:"trkseg>trkpt"`
}

type Metadata struct {
	Name string `xml:"name"`
	// Desc string `xml:"desc"`
}

type GPX struct {
	XMLName  xml.Name `xml:"gpx"`
	Metadata Metadata `xml:"metadata"`
	Track    Track    `xml:"trk"`
}
