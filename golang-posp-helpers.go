package golangposphelpers

import (
	"database/sql"
	"fmt"
)

func TestGolangHelper() {

	fmt.Println("asdf")
}

func GetDBResponse(rows *sql.Rows, ret []map[string]interface{}) []map[string]interface{} {
	cols, _ := rows.Columns()

	for rows.Next() {
		colVals := make([]interface{}, len(cols))
		for i := range colVals {
			colVals[i] = new(interface{})
		}
		err := rows.Scan(colVals...)
		if err != nil {
			CheckErr(err)
		}
		colNames, err := rows.Columns()
		if err != nil {
			CheckErr(err)
		}
		these := make(map[string]interface{})
		for idx, name := range colNames {
			these[name] = *colVals[idx].(*interface{})
		}
		ret = append(ret, these)
	}
	if err := rows.Err(); err != nil {
		CheckErr(err)
	}

	return ret

}

// Function for handling errors
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

// Function for handling messages
func PrintMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

func Deferring() {
	if err := recover(); err != nil {
		fmt.Println("An error occurred:", err)
	}
}
