package vojo

type SearchRes struct {
	Time    string      `json:"resCode"`
	Message interface{} `json:"message"`
}
type TableRow struct {
	KeyWord string `json:"keyWord"`
	UUID    string `json:"uuid"`
}
