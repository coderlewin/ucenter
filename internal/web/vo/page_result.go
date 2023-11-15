package vo

type PageResult struct {
	Records any   `json:"records"`
	Total   int64 `json:"total"`
}
