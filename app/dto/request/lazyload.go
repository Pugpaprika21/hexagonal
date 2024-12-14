package request

type LazyLoad struct {
	Limit        int32  `json:"limit"`
	Offset       int32  `json:"offset"`
	SearchText   string `json:"search_text,omitempty"`
	SearchNumber int32  `json:"search_number,omitempty"`
}
