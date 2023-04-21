package ghandler

import (
	sysSvc "github.com/jettjia/go-ddd-hertz/application/service/sys"
	grpcGoodsProto "github.com/jettjia/go-ddd-hertz/interfaces/grpc/proto/goods"
)

type GrpcGoodsServer struct {
	grpcGoodsProto.UnimplementedGoodsServer
	SysMenuSrv *sysSvc.SysMenuService
}
