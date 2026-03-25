package types

type Stats struct {
	Lines      int    `json:"lines"`
	Errors     int    `json:"errors"`
	Keystrokes int    `json:"keystrokes"`
	FileName   string `json:"file_name"`
}

type StatsDelta struct {
	Lines      *int    `json:"lines"`
	Errors     *int    `json:"errors"`
	Keystrokes *int    `json:"keystrokes"`
	FileName   *string `json:"file_name"`
}
