package common

type ResultInfo struct {
	ResultCode        string `json:"resultCode"`
	ResultDescription string `json:"resultDescription"`
}

func (r ResultInfo) IsSuccess() bool {
	return r.ResultCode == "00" || r.ResultCode == "0"
}

type BaseResponse[T any] struct {
	Result       ResultInfo `json:"result"`
	ResponseData T          `json:"responseData"`
}
