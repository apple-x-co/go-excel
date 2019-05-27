package types

// For excelize
type ExcelizeStyle struct {
	Font      *ExcelizeStyleFont      `json:"font,omitempty"`
	Fill      *ExcelizeStyleFill      `json:"fill,omitempty"`
	Alignment *ExcelizeStyleAlignment `json:"alignment,omitempty"`
	Border    []*ExcelizeStyleBorder  `json:"border,omitempty"`
}

func NewExcelizeStyleByCellStyle(cellStyle *CellStyle) *ExcelizeStyle {
	instance := new(ExcelizeStyle)

	if cellStyle.IsBold() || cellStyle.FontSize != 0 {
		instance.Font = new(ExcelizeStyleFont)
	}
	if cellStyle.IsBold() {
		instance.Font.Bold = true
	}
	if cellStyle.FontSize != 0 {
		instance.Font.Size = cellStyle.FontSize
	}

	if cellStyle.BackgroundColor != "" {
		instance.Fill = new(ExcelizeStyleFill)
		instance.Fill.Type = "pattern"
		instance.Fill.Color = []string{cellStyle.BackgroundColor}
		instance.Fill.Pattern = 1
	}

	if cellStyle.IsAlignmentHorizontalCenter() {
		instance.Alignment = new(ExcelizeStyleAlignment)
		instance.Alignment.Horizontal = "center"
	}

	if len(cellStyle.Border) != 0 {
		for _, cellStyleBorder := range cellStyle.Border {
			border := new(ExcelizeStyleBorder)
			border.Type = cellStyleBorder.Type
			border.Color = cellStyleBorder.Color
			border.Style = cellStyleBorder.Style
			instance.Border = append(instance.Border, border)
		}
	}

	return instance
}

func (E *ExcelizeStyle) HasStyles() bool {
	if E.Font != nil || E.Fill != nil || E.Alignment != nil || E.Border != nil {
		return true
	}
	return false
}

type ExcelizeStyleFont struct {
	Bold bool `json:"bold,omitempty"`
	Size int  `json:"size,omitempty"`
}

type ExcelizeStyleFill struct {
	Type    string   `json:"type,omitempty"`
	Color   []string `json:"color,omitempty"`
	Pattern int      `json:"pattern,omitempty"`
}

type ExcelizeStyleAlignment struct {
	Horizontal string `json:"horizontal,omitempty"`
}

type ExcelizeStyleBorder struct {
	Type  string `json:"type"`
	Color string `json:"color"`
	Style int    `json:"style"`
}