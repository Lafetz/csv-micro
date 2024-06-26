package parser

import (
	"reflect"
	"testing"
)

func TestParseCSV(t *testing.T) {
	csvData := `Rank,Name,Platform,Year,Genre,Publisher,NA_Sales,EU_Sales,JP_Sales,Other_Sales,Global_Sales
1,Wii Sports,Wii,2006,Sports,Nintendo,41.49,29.02,3.77,8.46,82.74
2,Super Mario Bros.,NES,1985,Platform,Nintendo,29.08,3.58,6.81,0.77,40.24
3,Mario Kart Wii,Wii,2008,Racing,Nintendo,15.85,12.88,3.79,3.31,35.82
4,Wii Sports Resort,Wii,2009,Sports,Nintendo,15.75,11.01,3.28,2.96,33
5,Pokemon Red/Pokemon Blue,GB,1996,Role-Playing,Nintendo,11.27,8.89,10.22,1,31.37`

	expected := &CSVData{
		TableName:  "Rank",
		FieldNames: []string{"Rank", "Name", "Platform", "Year", "Genre", "Publisher", "NA_Sales", "EU_Sales", "JP_Sales", "Other_Sales", "Global_Sales"},
		Data: [][]string{
			{"1", "Wii Sports", "Wii", "2006", "Sports", "Nintendo", "41.49", "29.02", "3.77", "8.46", "82.74"},
			{"2", "Super Mario Bros.", "NES", "1985", "Platform", "Nintendo", "29.08", "3.58", "6.81", "0.77", "40.24"},
			{"3", "Mario Kart Wii", "Wii", "2008", "Racing", "Nintendo", "15.85", "12.88", "3.79", "3.31", "35.82"},
			{"4", "Wii Sports Resort", "Wii", "2009", "Sports", "Nintendo", "15.75", "11.01", "3.28", "2.96", "33"},
			{"5", "Pokemon Red/Pokemon Blue", "GB", "1996", "Role-Playing", "Nintendo", "11.27", "8.89", "10.22", "1", "31.37"},
		},
	}

	result, err := parseCSV(csvData)
	if err != nil {
		t.Fatalf("parseCSV error: %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("parseCSV result mismatch:\nGot: %#v\nExpected: %#v", result, expected)
	}
}
