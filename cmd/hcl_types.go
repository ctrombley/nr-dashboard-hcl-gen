package main

type HCLDashboard struct {
	Resources struct {
		ResourceName  string `hcl:"resourceName,label"`
		ResourceLabel string `hcl:"resoruceLabel,label"`
		Name          string `hcl:"name"`
		Description   string `hcl:"description,optional"`

		Pages []HCLPage `hcl:"page,block"`
	} `hcl:"resource,block"`
}

type HCLPage struct {
	Name        string `hcl:"name"`
	Description string `hcl:"description,optional"`

	BillboardWidgets []HCLWidget `hcl:"widget_billboard,block"`
	BarWidgets       []HCLWidget `hcl:"widget_bar,block"`
	MarkdownWidgets  []HCLWidget `hcl:"widget_markdown,block"`
}

type HCLWidget struct {
	Title             string         `hcl:"title"`
	Column            string         `hcl:"column,optional"`
	Row               string         `hcl:"row,optional"`
	Height            string         `hcl:"height,optional"`
	Width             string         `hcl:"width,optional"`
	NRQLQueries       []HCLNRQLQuery `hcl:"nrql_query,block"`
	LinkedEntityGUIDs []string       `hcl:"linked_entity_guids,optional"`
	Text              string         `hcl:"text,optional"`
}

type HCLVisualization struct {
	ID string `hcl:"id"`
}

type HCLLayout struct {
	Column string `hcl:"column"`
	Row    string `hcl:"row"`
	Height string `hcl:"height"`
	Width  string `hcl:"width"`
}

type HCLRawConfiguration struct {
	DataFormatters    []string       `hcl:"dataFormatters"`
	Facet             HCLFacet       `hcl:"facet"`
	NRQLQueries       []HCLNRQLQuery `hcl:"nrqlQueries"`
	Legend            HCLLegend      `hcl:"legend"`
	YAxisLeft         HCLYAxisLeft   `hcl:"yAxisLeft"`
	LinkedEntityGUIDs []string       `hcl:"linkedEntityGuids"`
}

type HCLFacet struct {
	ShowOtherSeries bool `hcl:"showOtherSeries"`
}

type HCLNRQLQuery struct {
	AccountID int    `hcl:"account_id,optional"`
	Query     string `hcl:"query"`
}

type HCLLegend struct {
	Enabled bool `hcl:"enabled"`
}

type HCLYAxisLeft struct {
	Zero bool `hcl:"zero"`
}
