// Code generated by goctl. DO NOT EDIT.
package types

type Rsp struct {
	Result bool        `json:"result, omitempty"`
	Code   int32       `json:"code, omitempty"`
	Msg    string      `json:"msg, omitempty"`
	Data   interface{} `json:"data, omitempty"`
}

type SayReq struct {
	Name string `json:"name"`
}
