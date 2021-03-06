package types

type Book struct {
	Sheets []Sheet `json:"sheets"`
}

type Sheet struct {
	Name  string `json:"name"`
	Cells []Cell `json:"cells"`
}

type Cell struct {
	Row        int       `json:"row"`
	Column     int       `json:"column"`
	ColumnSpan int       `json:"column_span,omitempty"`
	Value      string    `json:"value"`
	Style      CellStyle `json:"style,omitempty"`
}

type CellStyle struct {
	FontWeight      string             `json:"font_weight"`
	FontSize        int                `json:"font_size"`
	BackgroundColor string             `json:"background_color"`
	Alignment       CellStyleAlignment `json:"alignment"`
	Width           float64            `json:"width,omitempty"`
	Height          float64            `json:"height,omitempty"`
	Border          []CellStyleBorder  `json:"border,omitempty"`
}

type CellStyleAlignment struct {
	Horizontal  string `json:"horizontal"`
	Vertical    string `json:"vertical"`
	WrapText    bool   `json:"wrap_text"`
	ShrinkToFit bool   `json:"shrink_to_fit"`
}

type CellStyleBorder struct {
	Type  string `json:"type"`
	Color string `json:"color"`
	Style int    `json:"style"`
}
