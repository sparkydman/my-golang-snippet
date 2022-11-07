package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

func generteTemplate(w http.ResponseWriter, r *http.Request){
	// path, err := filepath.Abs("./pdf-from-template/pdf.gohtml")
	// if err != nil {
	// 	fmt.Println("Couldn't find absolute", err)
	// }
	// t, err := template.ParseFiles(path)
	// if err != nil {
	// 	fmt.Println("Failed to parse template: ", err)
	// }
	// file, err := os.CreateTemp("temp","template_pdf_*.pdf")
	// if err != nil {
	// 	fmt.Println("Couldn't create temporary file: ", err)
	// }
	// defer func (){
	// 	if err := file.Close(); err != nil {
	// 		fmt.Println("Couldn't close temporary file: ", err)
	// 	}
	// 	if err := os.Remove(file.Name()); err != nil {
	// 		fmt.Println("Couldn't mot remove file", err)
	// 	}
	// }()

	// Create new PDF generator
  pdfg, err := wkhtmltopdf.NewPDFGenerator()
  if err != nil {
    log.Fatal(err)
  }

  // Set global options
  pdfg.Dpi.Set(300)
//   pdfg.Orientation.Set(OrientationLandscape)
  pdfg.Grayscale.Set(true)

  // Create a new input page from an URL
  page := wkhtmltopdf.NewPageReader(strings.NewReader(getTagHTML()))

  workingDir, err := os.Getwd()
  if err != nil {
	log.Fatal(err)
  }

  page.Allow.Set(workingDir)
  // Set options for this page
  page.FooterRight.Set("[page]")
  page.FooterFontSize.Set(10)
  page.Zoom.Set(0.95)

  // Add to document
  pdfg.AddPage(page)

  // Create PDF document in internal buffer
  err = pdfg.Create()
  if err != nil {
    log.Fatal(err)
  }

  // Write buffer contents to file on disk
//   err = pdfg.WriteFile(file.Name())
//   if err != nil {
//     log.Fatal(err)
//   }

	// if err := t.Execute(file, nil); err != nil {
	// 	fmt.Println("Failed to execute template: ", err)
	// }
		w.Header().Set("Content-Type", "application/octet-stream")
        w.Header().Set("Content-Disposition", "attachment; filename=sample.pdf")
        w.Header().Set("Content-Transfer-Encoding", "binary")
		w.Write(pdfg.Bytes())
		// http.ServeFile(w, r, file.Name())
}

func getTagHTML() string {
	file, err := os.Open("pdf.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func main() {
	http.HandleFunc("/", generteTemplate)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println(err)
	}

}