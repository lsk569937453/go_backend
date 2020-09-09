package vojo

type GrpcInvokeReq struct {
	Url         string `form:"url" json:"url" `
	ServiceName string `form:"serviceName" json:"serviceName" `
	MethodName  string `form:"methodName" json:"methodName" `
	ReqJson     string `form:"reqJson" json:"reqJson" `
}
