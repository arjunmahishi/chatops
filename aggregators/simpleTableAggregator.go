package aggregators

import (
	"fmt"
	"strings"

	"github.com/ryanuber/columnize"
)

// SimpleTableAggregator  structure
type SimpleTableAggregator struct {
	Service string
	Text    string
	Records [][]string
}

// Parse plain text
func (sta *SimpleTableAggregator) Parse() error {
	sta.Text = strings.Trim(sta.Text, " ")
	sta.Text = strings.Trim(sta.Text, "\n")

	sta.Text = strings.Replace(sta.Text, "  ", " ", -1)

	recordArr := strings.Split(sta.Text, "\n")
	for _, record := range recordArr {
		record = strings.Trim(record, " ")
		columns := strings.Split(record, " ")
		for i, column := range columns {
			columns[i] = strings.Trim(column, " ")
		}
		sta.Records = append(sta.Records, columns)
	}

	fmt.Println(sta.Records)

	return nil
}

func (sta *SimpleTableAggregator) String() string {
	var rows []string

	for _, row := range sta.Records {
		rowText := ""
		for _, column := range row {
			rowText += fmt.Sprintf(" | %s", column)
		}

		rows = append(rows, strings.Trim(rowText, " | "))
	}

	columnizeConfig := &columnize.Config{
		Delim:  "|",
		Glue:   "     ",
		Prefix: "",
		Empty:  "-",
		NoTrim: true,
	}

	return fmt.Sprintf("%s", columnize.Format(rows, columnizeConfig))
}
