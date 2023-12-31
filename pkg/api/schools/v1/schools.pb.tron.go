// Code generated by protoc-gen-tron. DO NOT EDIT.
// source: tx-example/api/schools/v1/schools.proto

package schools

import (
	"context"
	"embed"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/loghole/tron/transport"
	"google.golang.org/grpc"
)

//go:embed schools.swagger.json
var swagger embed.FS

// SchoolsAPIServiceDesc is description for the SchoolsAPIServer.
type SchoolsAPIServiceDesc struct {
	svc SchoolsAPIServer
}

func NewSchoolsAPIServiceDesc(s SchoolsAPIServer) transport.ServiceDesc {
	return &SchoolsAPIServiceDesc{svc: s}
}

func (d *SchoolsAPIServiceDesc) RegisterGRPC(s *grpc.Server) {
	RegisterSchoolsAPIServer(s, d.svc)
}

func (d *SchoolsAPIServiceDesc) RegisterHTTP(mux *runtime.ServeMux) {
	RegisterSchoolsAPIHandlerServer(context.Background(), mux, d.svc)
}

func (d *SchoolsAPIServiceDesc) SwaggerDef() []byte {
	b, _ := swagger.ReadFile("schools.swagger.json")

	return b
}
