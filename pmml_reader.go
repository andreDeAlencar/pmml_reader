package pmmlReader

import (
	"log"
	"os"
	"io/ioutil"
	"encoding/xml"
)

type Model struct {
	PMML 
}

type PMML struct {
	Version string `xml:"version,attr"`
	XMLNs string `xml:"xmlns,attr"`
	Header Header
	DataDictionary DataDictionary
}

type Header struct {
	Copyright string `xml:"copyright,attr"`
	Application Application 
}

type Application struct {
	Name string `xml:"name,attr"`
	Version string `xml:"version,attr"`
}

type DataDictionary struct {
	NumberOfFields int `xml:"numberOfFields,attr"`
	DataFields []DataField `xml:"DataField"`
}

type DataField struct {
	Name string `xml:"name,attr"`
	Optype string `xml:"optype,attr"`
	DataType string `xml:"dataType,attr"`
	Interval Interval
	Values []Value `xml:"Value"`
}

type Interval struct {
	Closure string `xml:"closure,attr"`
	LeftMargin float64 `xml:"leftMargin,attr"`
	RightMargin float64 `xml:"rightMargin,attr"`
}

type Value struct {
	Value string `xml:"value,attr"`
}

func Read(path string) (Model){

	file, err := os.Open(path) 
	if err != nil {
		log.Fatal(err)
	}
	
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	
	var p Model
	
	err = xml.Unmarshal(data, &p)
    if err != nil {
		log.Fatal(err)
	}
	return p
}