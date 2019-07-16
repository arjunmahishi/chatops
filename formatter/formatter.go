package formatter

//Aggregator interface
type Aggregator interface {
	Parse() error
	String() string
}

// GetFormatter based on the request
func GetFormatter(output, outputFormat string) Aggregator {
	switch outputFormat {
	case "simple-table":
		return &SimpleTableFormatter{Text: output}
	default:
		return &PlainTextFormatter{Text: output}
	}
}
