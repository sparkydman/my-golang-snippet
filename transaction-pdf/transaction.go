package main

import (
	"fmt"
	"os"

	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func main() {
m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(20, 10, 20)

	buildHeading(m)
	// buildFooter(m)
	// buildFruitList(m)
	// buildSignature(m)

	err := m.OutputFileAndClose("templates/transactions.pdf")
	if err != nil {
		fmt.Println("⚠️  Could not save PDF:", err)
		os.Exit(1)
	}

	fmt.Println("PDF saved successfully")
}

func buildHeading(m pdf.Maroto){
	m.RegisterHeader(func() {
		// m.Row(50, func() {
		// 	m.Col(12, func(){
		// 		if err := m.FileImage("images/appzone-logo.png", props.Rect{
		// 			Center: true,
		// 			Percent: 70,
		// 		});err	 != nil {
		// 			fmt.Println(err)
		// 		}
		// 	})
		// })
		m.Row(8, func() {
			m.Col(12, func(){
				m.Text("12b, Admiralty Way, Lekki Phase 1, Lagos.", props.Text{
					Top:   5,
					Align: consts.Left,
					Style: consts.Italic,
					Size: 8,
					Family: consts.Courier,
					Color: getGreyColor(),
				})
			})
		})
	})
}

func getGreyColor() color.Color {
	return color.Color{
		Red:   206,
		Green: 206,
		Blue:  206,
	}
}