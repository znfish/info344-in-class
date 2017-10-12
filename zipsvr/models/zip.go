package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Zip struct {
	Code  string
	City  string
	State string
}

// * --> pointer to a type
// slice of pointers point to the Zip struct
type ZipSlice []*Zip

// map: string --> slice of pointers point to the Zip struct
type ZipIndex map[string]ZipSlice

// can have multiple return types
func LoadZips(fileName string) (ZipSlice, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}

	reader := csv.NewReader(f)
	_, err = reader.Read() // _ --> ignore something
	// no := because there is no new value here needs to be created
	if err != nil {
		return nil, fmt.Errorf("error reading header row: %v", err)
	}

	zips := make(ZipSlice, 0, 43000)

	// return error when it hits the end of file
	for {
		fields, err := reader.Read()
		if err == io.EOF {
			// signaling I'm done with the parsing string
			return zips, nil
		}
		if err != nil {
			return nil, fmt.Errorf("error reading record %v", err)
		}
		z := &Zip{ // use & because zips is a slice of pointers point to the Zip struct?
			Code:  fields[0],
			City:  fields[3],
			State: fields[6],
		}
		zips = append(zips, z)
	}
}
