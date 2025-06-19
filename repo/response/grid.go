package response

type Grid struct {
	RecordTotal         int64       `json:"record_total"`
	RecordTotalFiltered int64       `json:"record_total_filtered"`
	Data                interface{} `json:"data"`
}

func GridResponse(data interface{}, recordtotal int64, recordfiltered int64) Grid {
	return Grid{RecordTotal: recordtotal, RecordTotalFiltered: recordfiltered, Data: data}
}
