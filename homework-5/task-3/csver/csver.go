package csver

import (
	"encoding/csv"
	"log"
	"os"
	"reflect"
	"strconv"
)

// Converts struct instance attributes to strings
func SerializeStruct(row interface{}) (values []string) {
	v := reflect.ValueOf(row)
	numberOfFields := v.NumField()

	for i := 0; i < numberOfFields; i++ {
		val := v.Field(i)
		var fieldValue string
		switch val.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fieldValue = strconv.FormatInt(val.Int(), 10)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			fieldValue = strconv.FormatUint(val.Uint(), 10)
		default:
			fieldValue = val.String()
		}

		values = append(values, fieldValue)
	}
	return values
}

// Prepare headers for CSV
func GetHeaders(str interface{}) (fields []string) {
	t := reflect.TypeOf(str)
	for i := 0; i < t.NumField(); i++ {
		fieldName, _ := t.Field(i).Tag.Lookup("csv")

		fields = append(fields, fieldName)
	}
	return fields
}

// Write the struct to the file in CSV format
func WriteToFile(rows []interface{}, fileName string) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	csvHeaders := GetHeaders(rows[0])
	csvWriter.Write(csvHeaders)

	for _, applicant := range rows {
		csvWriter.Write(SerializeStruct(applicant))
	}
}
