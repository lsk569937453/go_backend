package vojo

type GrpcRefServerInstanceFormat struct {
	Services []*GrpcRefServiceFormat `json:"services" `
}
type GrpcRefServiceFormat struct {
	ServiceName string `json:"serviceName" `

	Methods []*GrpcRefMethodFormat `json:"methods" `
}
type GrpcRefMethodFormat struct {
	MethodName string `json:"methodName" `
	InputName  string `json:"inputName" `

	Fields []*GrpcRefField `json:"fields" `
}

func ExchangeGrpcServiceLisFormat(res *GrpcRefServerInstance) *GrpcRefServerInstanceFormat {
	grpcServiceList := make([]*GrpcRefServiceFormat, 0)

	for _, serviceItem := range res.Services {

		grpcMethodList := make([]*GrpcRefMethodFormat, 0)
		for _, methodItem := range serviceItem.Methods {

			fieldList := make([]*GrpcRefField, 0)
			for _, fieldItem := range methodItem.Fields {
				fieldItem.Type = fieldNameFormat(fieldItem.Type)
				fieldList = append(fieldList, fieldItem)
			}
			methodFormat := &GrpcRefMethodFormat{
				Fields:     fieldList,
				MethodName: methodItem.MethodName,
				InputName:  methodItem.InputName,
			}
			grpcMethodList = append(grpcMethodList, methodFormat)
		}
		serviceInstance := &GrpcRefServiceFormat{
			Methods:     grpcMethodList,
			ServiceName: serviceItem.ServiceName,
		}
		grpcServiceList = append(grpcServiceList, serviceInstance)

	}
	realresult := &GrpcRefServerInstanceFormat{
		Services: grpcServiceList,
	}

	return realresult
}

//fix the field name from TYPE_STRING to string
func fieldNameFormat(src string) string {
	if src == "TYPE_STRING" {
		return "string"
	} else {
		return src
	}

}
