// Various structures used for unmarshaling GPX files
package main

import (
	"encoding/xml"
	"io"
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

func (g *GPX) Read(r io.Reader) error {
	rdr := xml.NewDecoder(r)
	return rdr.Decode(g)
}

func (g *GPX) Write(w io.Writer) error {
	w.Write([]byte(xml.Header))
	enc := xml.NewEncoder(w)
	enc.Indent("", "  ")
	return enc.Encode(*g)
}
