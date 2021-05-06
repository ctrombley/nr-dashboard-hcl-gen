package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

var (
	widgetTypes = map[string]string{
		"viz.area":      "widget_area",
		"viz.bar":       "widget_bar",
		"viz.billboard": "widget_billboard",
		"viz.bullet":    "widget_bullet",
		"viz.funnel":    "widget_funnel",
		"viz.heatmap":   "widget_heatmap",
		"viz.histogram": "widget_histogram",
		"viz.json":      "widget_json",
		"viz.line":      "widget_line",
		"viz.markdown":  "widget_markdown",
		"viz.pie":       "widget_pie",
		"viz.table":     "widget_table",
	}
)

func generateDashboardHCL(resourceLabel string, input []byte) {
	var d JSONDashboard
	if err := json.Unmarshal(input, &d); err != nil {
		log.Fatal(err)
	}

	resourceName := "newrelic_one_dashboard"

	h := strings.Builder{}
	i := ""

	h.WriteString(fmt.Sprintf("%sresource \"%s\" \"%s\" {\n", i, resourceName, resourceLabel))
	i += "  "

	h.WriteString(fmt.Sprintf("%sname = \"%s\"\n", i, d.Name))

	if d.Description != "" {
		h.WriteString(fmt.Sprintf("%sdescription = \"%s\"\n", i, d.Description))
	}

	if d.Permissions != "" {
		h.WriteString(fmt.Sprintf("%spermissions = \"%s\"\n", i, strings.ToLower(d.Permissions)))
	}

	for _, p := range d.Pages {
		h.WriteString(fmt.Sprintf("\n%spage {\n", i))
		i += "  "

		h.WriteString(fmt.Sprintf("%sname = \"%s\"\n", i, p.Name))
		if p.Description != "" {
			h.WriteString(fmt.Sprintf("%sdescription = \"%s\"\n", i, p.Description))
		}

		for _, w := range p.Widgets {
			if widgetTypes[w.Visualization.ID] == "" {
				log.Fatalf("unrecognized widget type \"%s\"\n", w.Visualization.ID)
			}

			h.WriteString(fmt.Sprintf("\n%s%s {\n", i, widgetTypes[w.Visualization.ID]))
			i += "  "

			h.WriteString(fmt.Sprintf("%stitle = \"%s\"\n", i, strings.ReplaceAll(w.Title, "\"", "\\\"")))

			if w.Layout.Row != 0 {
				h.WriteString(fmt.Sprintf("%srow = %d\n", i, w.Layout.Row))
			}

			if w.Layout.Column != 0 {
				h.WriteString(fmt.Sprintf("%scolumn = %d\n", i, w.Layout.Column))
			}

			if w.Layout.Height != 0 {
				h.WriteString(fmt.Sprintf("%sheight = %d\n", i, w.Layout.Height))
			}

			if w.Layout.Width != 0 {
				h.WriteString(fmt.Sprintf("%swidth = %d\n", i, w.Layout.Width))
			}

			if len(w.RawConfiguration.LinkedEntityGUIDs) > 0 {
				h.WriteString(fmt.Sprintf("%slinked_entity_guids = [\"%s\"]", i, strings.Join(w.RawConfiguration.LinkedEntityGUIDs, "\",\"")))
			}

			if w.RawConfiguration.Text != "" {
				h.WriteString(fmt.Sprintf("%stext = <<EOT\n%s\nEOT\n", i, w.RawConfiguration.Text))
			}

			for _, q := range w.RawConfiguration.NRQLQueries {
				h.WriteString(fmt.Sprintf("\n%snrql_query {\n", i))
				i += "  "

				if q.AccountID != 0 {
					h.WriteString(fmt.Sprintf("%saccount_id = %d\n", i, q.AccountID))
				}

				h.WriteString(fmt.Sprintf("%squery = <<EOT\n%s\nEOT\n", i, q.Query))

				i = i[0 : len(i)-2]
				h.WriteString(fmt.Sprintf("%s}\n", i))
			}

			i = i[0 : len(i)-2]
			h.WriteString(fmt.Sprintf("%s}\n", i))
		}

		i = i[0 : len(i)-2]
		h.WriteString(fmt.Sprintf("%s}\n", i))
	}

	i = i[0 : len(i)-2]
	h.WriteString(fmt.Sprintf("}\n"))

	fmt.Println(h.String())
}
