package go_excel

import (
	"apple-x-co/go-excel/types"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/pkg/errors"
	flag "github.com/spf13/pflag"
	"io/ioutil"
	"os"
)

type go_excel struct {
	version  string
	revision string
}

func NewGoExcel(version string, revision string) *go_excel {
	instance := new(go_excel)
	instance.version = version
	instance.revision = revision
	return instance
}

func (go_excel *go_excel) Execute() error {
	var (
		inputPath   = flag.StringP("in", "i", "book.json", "file path of input json.")
		outputPath  = flag.StringP("out", "o", "book.xlsx", "file path of output excel.")
		showHelp    = flag.BoolP("help", "h", false, "show help message")
		showSample  = flag.BoolP("sample", "s", false, "show sample json")
		showVersion = flag.BoolP("version", "v", false, "show version")
	)
	flag.Parse()

	if *showHelp {
		flag.PrintDefaults()
		return nil
	}
	if *showVersion {
		fmt.Println("version:", go_excel.version+"."+go_excel.revision)
		return nil
	}
	if *showSample {
		book := types.Book{}
		book.Sheets = append([]types.Sheet{}, types.Sheet{Name: "Sheet1", Cells: []types.Cell{}})
		book.Sheets[0].Cells = append([]types.Cell{}, types.Cell{Row: 1, Column: 1, Value: "A1"})
		j, _ := json.MarshalIndent(book, "", " ")
		fmt.Println(string(j))
		return nil
	}

	f, err := os.Open(*inputPath)
	if err != nil {
		return errors.Wrap(err, "failed to open input file")
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return errors.Wrap(err, "failed to read input file")
	}

	var decoded types.Book
	bytes := []byte(string(b))
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return errors.Wrap(err, "failed to unmarshal input file")
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
			colName, err := convertColName(cell.Column)
			if err != nil {
				fmt.Println(err)
				continue
			}

			xlsx.SetCellValue(sheet.Name, cellName, string(cell.Value))

			style := types.NewExcelizeStyleByCellStyle(&cell.Style)
			if style.HasStyles() {
				styleJson, _ := json.Marshal(style)
				//fmt.Printf("%v\n", string(styleJson))
				style, err := xlsx.NewStyle(string(styleJson))
				if err != nil {
					fmt.Println(err)
					continue
				}
				xlsx.SetCellStyle(sheet.Name, cellName, cellName, style)
			}

			if cell.Style.Width != 0 {
				xlsx.SetColWidth(sheet.Name, colName, colName, cell.Style.Width)
			}
			if cell.Style.Height != 0 {
				xlsx.SetRowHeight(sheet.Name, cell.Row, cell.Style.Height)
			}

			if columnSpan := cell.ColumnSpan; columnSpan != 0 {
				mergeCellName, err := convertCellName(cell.Column+columnSpan, cell.Row)
				if err != nil {
					fmt.Println(err)
					continue
				}

				xlsx.MergeCell(sheet.Name, cellName, mergeCellName)
			}
		}

		xlsx.SetActiveSheet(index)
	}

	if err := xlsx.SaveAs(*outputPath); err != nil {
		return errors.Wrap(err, "failed to save output file")
	}

	return nil
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

func convertColName(col int) (string, error) {
	if col < 1 {
		return "", fmt.Errorf("incorrect column number %d", col)
	}

	var axis string
	for col > 0 {
		axis = string((col-1)%26+65) + axis
		col = (col - 1) / 26
	}

	return fmt.Sprintf("%s", axis), nil
}
