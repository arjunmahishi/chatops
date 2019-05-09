package aggregators

//Aggregator interface
type Aggregator interface {
	Parse() error
	String() string
}

// GetAggregator based on the request
func GetAggregator(output, outputFormat string) Aggregator {
	switch outputFormat {
	case "simple-table":
		return &SimpleTableAggregator{Text: output}
	default:
		return &PlainTextAggregator{Text: output}
	}
}
