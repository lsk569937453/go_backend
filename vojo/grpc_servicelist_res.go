package vojo

type GrpcRefServerInstance struct {
	Services map[string]*GrpcRefService `json:"services" `
}
type GrpcRefService struct {
	ServiceName string                    `json:"serviceName" `
	Methods     map[string]*GrpcRefMethod `json:"methods" `
}
type GrpcRefMethod struct {
	MethodName string                   `json:"methodName" `
	InputName  string                   `json:"inputName" `
	Fields     map[string]*GrpcRefField `json:"fields" `
}
type GrpcRefField struct {
	FieldName string `json:"fieldName" `
	Type      string `json:"type" `
}
