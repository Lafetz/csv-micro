package parser

import (
	"encoding/csv"
	"errors"
	"strings"
)

var (
	ErrInvalidFile = errors.New("invalid csv file")
)

type parserService struct {
	downloader downloaderApi
}

func (srv *parserService) parse(data string) error {
	parsedData, err := parseCSV(data)
	if err != nil {
		return err
	}
	if err := srv.downloader.sendData(parsedData); err != nil {
		return err
	}
	return nil
}

func parseCSV(data string) (*CSVData, error) {
	reader := csv.NewReader(strings.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, ErrInvalidFile
	}
	if len(records) == 0 {
		return nil, ErrInvalidFile
	}
	tableName := records[0][0]
	fieldNames := records[0]
	dataRows := records[1:]

	csvDataStruct := &CSVData{
		TableName:  tableName,
		FieldNames: fieldNames,
		Data:       dataRows,
	}
	return csvDataStruct, nil
}
