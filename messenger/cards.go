package messenger

func reviewNotification(message, url string) interface{} {
	return struct {
		Cards interface{} `json:"cards"`
	}{
		Cards: []struct {
			Sections interface{} `json:"sections"`
		}{{
			Sections: []struct {
				Widgets interface{} `json:"widgets"`
			}{{
				Widgets: []struct {
					TextParagraph interface{} `json:"textParagraph"`
					Buttons       interface{} `json:"buttons"`
				}{
					{
						TextParagraph: struct {
							Text string `json:"text"`
						}{
							Text: message,
						},
					},
					{
						Buttons: struct {
							TextButton interface{} `json:"textButton"`
						}{
							TextButton: struct {
								Text    string      `json:"text"`
								OnClick interface{} `json:"onClick"`
							}{
								Text: "Review",
								OnClick: struct {
									OpenLink interface{} `json:"openLink"`
								}{
									OpenLink: struct {
										URL string `json:"url"`
									}{
										URL: url,
									},
								},
							},
						},
					},
				},
			}},
		}},
	}
}
