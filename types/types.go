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
	ColumnSpan int       `json:"column_span"`
	Value      string    `json:"value"`
	Style      CellStyle `json:"style"`
}

type CellStyle struct {
	FontWeight      string             `json:"font_weight"`
	BackgroundColor string             `json:"background_color"`
	Alignment       CellStyleAlignment `json:"alignment"`
}

type CellStyleAlignment struct {
	Horizontal string `json:"horizontal"`
}
