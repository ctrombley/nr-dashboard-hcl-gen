package main

type JSONDashboard struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Permissions string `json:"permissions"`

	Pages []JSONPage `json:"pages"`
}

type JSONPage struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	Widgets []JSONWidget `json:"widgets"`
}

type JSONWidget struct {
	Title            string               `json:"title"`
	Visualization    JSONVisualization    `json:"visualization"`
	Layout           JSONLayout           `json:"layout"`
	RawConfiguration JSONRawConfiguration `json:"rawConfiguration"`
}

type JSONVisualization struct {
	ID string `json:"id"`
}

type JSONLayout struct {
	Column int `json:"column"`
	Row    int `json:"row"`
	Height int `json:"height"`
	Width  int `json:"width"`
}

type JSONRawConfiguration struct {
	DataFormatters    []string        `json:"dataFormatters"`
	Text              string          `json:"text"`
	Facet             JSONFacet       `json:"facet"`
	NRQLQueries       []JSONNRQLQuery `json:"nrqlQueries"`
	Legend            JSONLegend      `json:"legend"`
	YAxisLeft         JSONYAxisLeft   `json:"yAxisLeft"`
	LinkedEntityGUIDs []string        `json:"linkedEntityGuids"`
}

type JSONFacet struct {
	ShowOtherSeries bool `json:"showOtherSeries"`
}

type JSONNRQLQuery struct {
	AccountID int    `json:"accountId"`
	Query     string `json:"query"`
}

type JSONLegend struct {
	Enabled bool `json:"enabled"`
}

type JSONYAxisLeft struct {
	Zero bool `json:"zero"`
}
