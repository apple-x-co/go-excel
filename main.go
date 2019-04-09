package main

import (
	"apple-x-co/go-excel/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
	flag "github.com/spf13/pflag"
)

func main() {
	var (
		inputPath  = flag.StringP("file", "f", "book.json", "file path of input json.")
		outputPath = flag.StringP("out", "o", "book.xlsx", "file path of output excel.")
		help       = flag.BoolP("help", "h", false, "show help message")
		sample     = flag.BoolP("sample", "s", false, "show sample json")
	)
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return
	}
	if *sample {
		book := types.Book{}
		book.Sheets = append([]types.Sheet{}, types.Sheet{Name: "Sheet1", Cells: []types.Cell{}})
		book.Sheets[0].Cells = append([]types.Cell{}, types.Cell{Row: 1, Column: 1, Value: "A1"})
		j, _ := json.MarshalIndent(book, "", " ")
		fmt.Println(string(j))
		return
	}

	f, err := os.Open(*inputPath)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	var decoded types.Book
	bytes := []byte(string(b))
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		fmt.Println("error:", err)
		return
	}

	xlsx := excelize.NewFile()

	for _, sheet := range decoded.Sheets {
		index := xlsx.NewSheet(sheet.Name)

		for _, cell := range sheet.Cells {
			cellName, err := convertCellName(cell.Column, cell.Row)
			if err != nil {
				fmt.Println(err)
				continue
			}

			xlsx.SetCellValue(sheet.Name, cellName, string(cell.Value))

			cellFormat := ""
			if cell.Style.FontWeight == "bold" {
				cellFormat += `"font":{"bold":true}`
			}
			if cell.Style.BackgroundColor != "" {
				if cellFormat != "" {
					cellFormat += ","
				}
				cellFormat += `"fill":{"type":"pattern","color":["` + cell.Style.BackgroundColor + `"],"pattern":1}`
			}
			if cell.Style.Alignment.Horizontal == "center" {
				if cellFormat != "" {
					cellFormat += ","
				}
				cellFormat += `"alignment":{"horizontal":"center"}`
			}

			if cellFormat != "" {
				style, err := xlsx.NewStyle(`{` + cellFormat + `}`)
				if err != nil {
					fmt.Println(err)
					continue
				}
				xlsx.SetCellStyle(sheet.Name, cellName, cellName, style)
			}

			if cell.ColumnSpan != 0 {
				mergeCellName, err := convertCellName(cell.Column+cell.ColumnSpan, cell.Row)
				if err != nil {
					fmt.Println(err)
					continue
				}

				xlsx.MergeCell(sheet.Name, cellName, mergeCellName)
			}

			//fmt.Printf("%v\n", cell.Style.BackgroundColor)
		}

		xlsx.SetActiveSheet(index)
	}

	if err := xlsx.SaveAs(*outputPath); err != nil {
		fmt.Println(err)
		return
	}
}

func convertCellName(col int, row int) (string, error) {
	if col < 1 {
		return "", fmt.Errorf("incorrect column number %d", col)
	}
	if row < 1 {
		return "", fmt.Errorf("incorrect row number %d", row)
	}

	var axis string
	for col > 0 {
		axis = string((col-1)%26+65) + axis
		col = (col - 1) / 26
	}

	return fmt.Sprintf("%s%d", axis, row), nil
}
