package response

type BaseResponse struct {
	Message string       `json:"message"`
	Meta    *interface{} `json:"meta,omitempty"`
	Data    *interface{} `json:"data,omitempty"`
}

func NewBaseResponse(msg *string, data interface{}) BaseResponse {
	resp := BaseResponse{
		Message: "OK",
	}

	if data != nil {
		resp.Data = &data
	}

	if msg != nil {
		resp.Message = *msg
	}

	return resp
}

func NewBaseResponseWithMeta(msg *string, data interface{}, meta interface{}) BaseResponse {
	resp := BaseResponse{
		Message: "OK",
	}

	if data != nil {
		resp.Data = &data
	}

	if msg != nil {
		resp.Message = *msg
	}

	if meta != nil {
		resp.Meta = &meta
	}

	return resp
}
