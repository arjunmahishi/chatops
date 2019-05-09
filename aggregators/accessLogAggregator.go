package aggregators

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ryanuber/columnize"
)

// AccessLogAggregator  structure
type AccessLogAggregator struct {
	Service string
	Text    string
	Logs    []*LogRecord
}

// LogRecord structure
type LogRecord struct {
	Count  uint64
	Method string
	Path   string
	Status int
}

// String form of the Aggregator access log
func (al *AccessLogAggregator) String() string {

	var recordArray []string
	for _, record := range al.Logs {
		recordArray = append(recordArray, record.String())
	}

	return fmt.Sprintf("```%s```", columnize.SimpleFormat(recordArray))
}

// Parse the text and put into a struct
func (al *AccessLogAggregator) Parse() error {
	al.Text = strings.Replace(al.Text, "\"", "", -1)
	al.Text = strings.Replace(al.Text, "//", "/", -1)

	al.Text = strings.Replace(al.Text, "/v1/Accounts/", "", -1)
	al.Text = strings.Replace(al.Text, "HTTP/1.1", "", -1)

	rows := strings.Split(strings.Trim(al.Text, "\n"), "\n")
	rowArray := make([][]string, len(rows))
	for i, ele := range rows {
		rowArray[i] = strings.Split(strings.Trim(ele, " "), " ")
	}

	for _, row := range rowArray {

		count, err := strconv.ParseUint(row[0], 10, 64)
		if err != nil {
			return err
		}
		status, err := strconv.ParseInt(row[4], 10, 32)
		if err != nil {
			return err
		}

		al.Logs = append(al.Logs, &LogRecord{
			Count:  count,
			Method: row[1],
			Path:   row[2],
			Status: int(status),
		})
	}

	return nil
}

func (lr LogRecord) String() string {
	return fmt.Sprintf("%d | %s | %s", lr.Count, lr.Method, lr.Path)
}
