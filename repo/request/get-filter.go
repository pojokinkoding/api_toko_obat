package request

type GetFilter struct {
	PageOffset int    `form:"page_offset"`
	PageLimit  int    `form:"page_limit"`
	Dir        string `form:"dir"`
	Field      string `form:"field"`
}
