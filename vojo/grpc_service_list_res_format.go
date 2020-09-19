package vojo

import "sort"

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
type grpcServiceFormatSlice []*GrpcRefServiceFormat

func (s grpcServiceFormatSlice) Len() int           { return len(s) }
func (s grpcServiceFormatSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s grpcServiceFormatSlice) Less(i, j int) bool { return s[i].ServiceName < s[j].ServiceName }

func ExchangeGrpcServiceLisFormat(res *GrpcRefServerInstance) *GrpcRefServerInstanceFormat {
	grpcServiceList := make(grpcServiceFormatSlice, 0)

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
	sort.Stable(grpcServiceList)
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
