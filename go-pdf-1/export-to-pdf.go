package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"

	"github.com/signintech/gopdf"
)

type person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Gender string `json:"gender"`
}

func main() {
	persons := []person{
		{Name: "John", Age: 18, Gender: "Male"},
		{Name: "Grace", Age: 32, Gender: "Female"},
		{Name: "Kelvin", Age: 12, Gender: "Male"},
		{Name: "Milla", Age: 23, Gender: "Male"},
	}

	jsonFile, err := json.Marshal(persons)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonFile))

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{ PageSize: *gopdf.PageSizeA4 })  
	pdf.AddPage()
	// ioutil.ReadDir()
	abs, err := filepath.Abs("fonts/HLL.ttf")
	if err != nil {
		fmt.Println(err)
	}
	err = pdf.AddTTFFont("Helvetica", abs)
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("Helvetica", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	for _, v := range persons {
		rowVlue := fmt.Sprintf("v%\t v%\t v%\t", v.Name, v.Age, v.Gender)
		if err = pdf.Cell(gopdf.PageSizeA4, rowVlue); err != nil {
			fmt.Println(err)
		}
	}
	pdf.WritePdf("hello.pdf")

}
