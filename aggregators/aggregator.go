package aggregators

//Aggregator interface
type Aggregator interface {
	Parse() error
	String() string
}

// GetAggregator based on the request
func GetAggregator(output, outputFormat string) Aggregator {
	switch outputFormat {
	case "access-log":
		return &AccessLogAggregator{Text: output}
	default:
		return &PlainTextAggregator{Text: output}
	}
}
