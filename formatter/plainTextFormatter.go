package formatter

import "fmt"

// PlainTextFormatter  structure
type PlainTextFormatter struct {
	Service string
	Text    string
}

// Parse plain text
func (pta PlainTextFormatter) Parse() error {
	return nil
}

func (pta PlainTextFormatter) String() string {
	return fmt.Sprintf("```%s```", pta.Text)
}
