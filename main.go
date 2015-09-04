package main

import (
	"encoding/xml"
	"flag"
	"log"
	"os"
)

func main() {
	// Flags handling
	metadataName := flag.String("m", "Merged GPX", "String used in <metadata><name>")
	trackName := flag.String("t", "Merged track", "String used in <trk><name>")
	outfname := flag.String("o", "", "Output file")

	flag.Parse()

	if len(*outfname) == 0 {
		log.Println("No output filename specified")
		os.Exit(1)
	}

	// Read input GPXs
	inp_gpxs := make([]GPX, flag.NArg())
	ntrkpts := 0
	for i, fname := range flag.Args() {
		fi, err := os.Open(fname)
		if err != nil {
			log.Fatal(err)
		}
		defer fi.Close()

		rdr := xml.NewDecoder(fi)
		err = rdr.Decode(&inp_gpxs[i])
		if err != nil {
			log.Fatal(err)
		}
		ntrkpts += len(inp_gpxs[i].Track.TrackSegment)
	}

	// Struct which is used for merging all the trackpoints
	out_gpx := GPX{}
	out_gpx.Metadata.Name = *metadataName
	out_gpx.Track.Name = *trackName
	out_gpx.Track.TrackSegment = make([]TrackPoint, ntrkpts)

	// Performs the merging
	i := 0
	for _, gpx := range inp_gpxs {
		copy(out_gpx.Track.TrackSegment[i:i+len(gpx.Track.TrackSegment)], gpx.Track.TrackSegment)
		i += len(gpx.Track.TrackSegment)
	}

	// Writes output file
	fo, err := os.Create(*outfname)
	if err != nil {
		log.Fatal(err)
	}
	defer fo.Close()

	fo.Write([]byte(xml.Header))

	enc := xml.NewEncoder(fo)
	enc.Indent("", "  ")
	enc.Encode(out_gpx)
}
