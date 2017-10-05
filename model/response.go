package model

type ServiceResponse struct {
	Speech      string `json:"speech"`
	DisplayText string `json:"displayText"`
	Data struct {
		Google struct {
			ExpectUserResponse bool `json:"expect_user_response"`
			IsSsml             bool `json:"is_ssml"`
			PermissionsRequest struct {
				OptContext  string   `json:"opt_context"`
				Permissions []string `json:"permissions"`
			} `json:"permissions_request"`
			ExpectedInput []ExpectedInput `json:"expectedInputs, omitempty"`
		} `json:"google"`
	} `json:"data"`
	ContextOut []struct {
		Name     string `json:"name"`
		Lifespan int    `json:"lifespan"`
		Parameters struct {
			City string `json:"city"`
		} `json:"parameters"`
	} `json:"contextOut"`

	//TODO
	//add contextOut
}

type ExpectedInput struct {
	InputPrompt                      `json:"inputPrompt"`
	PossibleIntents []PossibleIntent `json:"possibleIntents"`
}

type InputPrompt struct {
	InitialPrompts []InitialPrompt `json:"initialPrompts"`
	NoInputPrompts []interface{}   `json:"noInputPrompts"`
}

type InitialPrompt struct {
	TextToSpeech string `json:"textToSpeech"`
}

type PossibleIntent struct {
	Intent string  `json:"intent"`
	InputValueData `json:"inputValueData"`
}

type InputValueData struct {
	Type        string   `json:"@type"`
	OptContext  string   `json:"optContext"`
	Permissions []string `json:"permissions"`
}
