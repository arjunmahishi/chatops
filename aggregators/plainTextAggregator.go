package aggregators

import "fmt"

// PlainTextAggregator  structure
type PlainTextAggregator struct {
	Service string
	Text    string
}

// Parse plain text
func (pta PlainTextAggregator) Parse() error {
	return nil
}

func (pta PlainTextAggregator) String() string {
	return fmt.Sprintf("```%s```", pta.Text)
}
