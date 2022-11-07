package main

import (
	"fmt"
)

func main() {

	r := NewRequestPdf("")

	//html template path
	templatePath := "templates/sample.html"

	//path for download pdf
	outputPath := "./storage/pdf.pdf"
	

	//html template data
	templateData := struct {
	    Data string
	}{
	   Data: "data",
	}

	if err := r.ParseTemplate(templatePath, templateData); err == nil {
		ok, _ := r.GeneratePDF(outputPath)
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}
}