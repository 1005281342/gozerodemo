package httproxy

import (
	"github.com/fullstorydev/grpchan"
	"google.golang.org/grpc"
)

// RegisterHandler 注册Handler
func RegisterHandler(reg grpchan.ServiceRegistry, srv interface{}, desc *grpc.ServiceDesc) {
	reg.RegisterService(desc, srv)
}
