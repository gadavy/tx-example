// Code generated by protoc-gen-tron. DO NOT EDIT.
// source: tx-example/api/teachers/v1/teachers.proto

package teachers

import (
	"context"
	"embed"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/loghole/tron/transport"
	"google.golang.org/grpc"
)

//go:embed teachers.swagger.json
var swagger embed.FS

// TeachersAPIServiceDesc is description for the TeachersAPIServer.
type TeachersAPIServiceDesc struct {
	svc TeachersAPIServer
}

func NewTeachersAPIServiceDesc(s TeachersAPIServer) transport.ServiceDesc {
	return &TeachersAPIServiceDesc{svc: s}
}

func (d *TeachersAPIServiceDesc) RegisterGRPC(s *grpc.Server) {
	RegisterTeachersAPIServer(s, d.svc)
}

func (d *TeachersAPIServiceDesc) RegisterHTTP(mux *runtime.ServeMux) {
	RegisterTeachersAPIHandlerServer(context.Background(), mux, d.svc)
}

func (d *TeachersAPIServiceDesc) SwaggerDef() []byte {
	b, _ := swagger.ReadFile("teachers.swagger.json")

	return b
}