package response

// ubah namanya jadi GetList
type GetList struct {
	RecordTotal         int64       `json:"record_total"`
	RecordTotalFiltered int64       `json:"record_total_filtered"`
	Data                interface{} `json:"data"`
}

// ubah namanya jadi GetlistResponse
// nama file jadi GetList
func NewGetList(data interface{}, recordtotal int64, recordfiltered int64) Grid {
	return Grid{RecordTotal: recordtotal, RecordTotalFiltered: recordfiltered, Data: data}
}
