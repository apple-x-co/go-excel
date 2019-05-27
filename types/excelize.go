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

	if cellStyle.FontWeight == "bold" || cellStyle.FontSize != 0 {
		instance.Font = new(ExcelizeStyleFont)
	}
	if cellStyle.FontWeight == "bold" {
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

	if cellStyle.Alignment.WrapText == true ||
		cellStyle.Alignment.ShrinkToFit == true ||
		cellStyle.Alignment.Horizontal != "" ||
		cellStyle.Alignment.Vertical != "" {
		instance.Alignment = new(ExcelizeStyleAlignment)
	}
	if cellStyle.Alignment.WrapText == true {
		instance.Alignment.WrapText = true
	}
	if cellStyle.Alignment.ShrinkToFit == true {
		instance.Alignment.ShrinkToFit = true
	}
	if cellStyle.Alignment.Horizontal != "" {
		instance.Alignment.Horizontal = cellStyle.Alignment.Horizontal
	}
	if cellStyle.Alignment.Vertical != "" {
		instance.Alignment.Vertical = cellStyle.Alignment.Vertical
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
	Horizontal  string `json:"horizontal,omitempty"`
	Vertical    string `json:"vertical,omitempty"`
	WrapText    bool   `json:"wrap_text,omitempty"`
	ShrinkToFit bool   `json:"shrink_to_fit"`
}

type ExcelizeStyleBorder struct {
	Type  string `json:"type"`
	Color string `json:"color"`
	Style int    `json:"style"`
}
