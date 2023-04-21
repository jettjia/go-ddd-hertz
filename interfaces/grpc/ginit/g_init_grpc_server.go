package ginit

import (
	"github.com/jettjia/go-ddd-hertz/boot"
	"google.golang.org/grpc"

	"github.com/jettjia/go-ddd-hertz/interfaces/grpc/ghandler"
	grpcGoodsProto "github.com/jettjia/go-ddd-hertz/interfaces/grpc/proto/goods"
)

// RegisterGrpcSrv 初始化grpc的服务
func RegisterGrpcSrv(app *boot.App, server *grpc.Server) {
	grpcGoodsProto.RegisterGoodsServer(server, &ghandler.GrpcGoodsServer{
		SysMenuSrv: app.SysMenuSvc,
	})
}
