package serializer

// Response 团队基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`   // 业务错误信息
	Error  string      `json:"error"` // 代码层面的报错，release 关闭，测试时打开
}

// 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

type DataList struct {
	Total uint        `json:"total"`
	Items interface{} `json:"items"`
}

func BuildListResponse(items interface{}, tatol uint) Response {
	return Response{
		Data: DataList{
			Total: tatol,
			Items: items,
		},
	}
}
