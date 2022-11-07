package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/xuri/excelize/v2"
)

func Index(w http.ResponseWriter, r *http.Request){
		file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	styleID, err := file.NewStyle(&excelize.Style{Font: &excelize.Font{Color: "#777777"}})
	if err != nil {
		fmt.Println(err)
	}
	if err := streamWriter.SetRow("A1", []interface{}{
		excelize.Cell{StyleID: styleID, Value: "Data"}}); err != nil {
		fmt.Println(err)
	}
	for rowID := 2; rowID <= 102400; rowID++ {
		row := make([]interface{}, 50)
		for colID := 0; colID < 50; colID++ {
			row[colID] = rand.Intn(640000)
		}
		cell, _ := excelize.CoordinatesToCellName(1, rowID)
		if err := streamWriter.SetRow(cell, row); err != nil {
			fmt.Println(err)
		}
	}
	if err := streamWriter.Flush(); err != nil {
	fmt.Println(err)
	}
	//copy the relevant headers. If you want to preserve the downloaded file name, extract it with go's url parser.

	 w.Header().Set("Content-Type", "application/octet-stream")
        w.Header().Set("Content-Disposition", "attachment; filename=example.xlsx")
        w.Header().Set("Content-Transfer-Encoding", "binary")

	// if f, err := file.WriteTo(w); err != nil {
	// 	fmt.Println(err)
	// }
		file.Write(w)
}

func main() {


	http.HandleFunc("/", Index)
	err := http.ListenAndServe(":8001", nil)

	if err != nil {
		fmt.Println(err)
	}

}