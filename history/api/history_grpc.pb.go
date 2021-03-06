// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ActionsAndHistoryClient is the client API for ActionsAndHistory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActionsAndHistoryClient interface {
	//Getters
	GetActions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetActionsResponse, error)
	GetHistory(ctx context.Context, in *HistoryRequest, opts ...grpc.CallOption) (*HistoryResponse, error)
	//Создание записей
	CreateCard(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*StatusWithProperties, error)
	InstallOnWorkspace(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error)
	RemoveFromWorkspace(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error)
	NeedRepair(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error)
	SendToRepair(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error)
	ReceiveFromRepair(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error)
	Archive(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error)
	DeArchive(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error)
	ChangeName(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error)
	ChangeInventory(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error)
	AddToGroup(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error)
	RemoveFromGroup(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error)
	SendToWarehouse(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error)
	GiveToEmployee(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error)
	TakeFromEmployee(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error)
	//Проверки на существование и доступность ввода
	IsCreated(ctx context.Context, in *PropertySerial, opts ...grpc.CallOption) (*StatusWithProperty, error)
	IsOnWorkspace(ctx context.Context, in *PropertySerial, opts ...grpc.CallOption) (*Status, error)
	IsNeedsRepair(ctx context.Context, in *PropertySerial, opts ...grpc.CallOption) (*Status, error)
	IsUnderRepair(ctx context.Context, in *PropertySerial, opts ...grpc.CallOption) (*Status, error)
	IsInStock(ctx context.Context, in *PropertySerial, opts ...grpc.CallOption) (*Status, error)
	IsInArchive(ctx context.Context, in *PropertySerial, opts ...grpc.CallOption) (*Status, error)
}

type actionsAndHistoryClient struct {
	cc grpc.ClientConnInterface
}

func NewActionsAndHistoryClient(cc grpc.ClientConnInterface) ActionsAndHistoryClient {
	return &actionsAndHistoryClient{cc}
}

func (c *actionsAndHistoryClient) GetActions(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetActionsResponse, error) {
	out := new(GetActionsResponse)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/GetActions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) GetHistory(ctx context.Context, in *HistoryRequest, opts ...grpc.CallOption) (*HistoryResponse, error) {
	out := new(HistoryResponse)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/GetHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) CreateCard(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*StatusWithProperties, error) {
	out := new(StatusWithProperties)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/CreateCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) InstallOnWorkspace(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/InstallOnWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) RemoveFromWorkspace(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/RemoveFromWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) NeedRepair(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/NeedRepair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) SendToRepair(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/SendToRepair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) ReceiveFromRepair(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/ReceiveFromRepair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) Archive(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/Archive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) DeArchive(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/DeArchive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) ChangeName(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/ChangeName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) ChangeInventory(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/ChangeInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) AddToGroup(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/AddToGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) RemoveFromGroup(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/RemoveFromGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) SendToWarehouse(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/SendToWarehouse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) GiveToEmployee(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/GiveToEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) TakeFromEmployee(ctx context.Context, in *ActionRequest, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/TakeFromEmployee", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) IsCreated(ctx context.Context, in *PropertySerial, opts ...grpc.CallOption) (*StatusWithProperty, error) {
	out := new(StatusWithProperty)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/IsCreated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) IsOnWorkspace(ctx context.Context, in *PropertySerial, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/IsOnWorkspace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) IsNeedsRepair(ctx context.Context, in *PropertySerial, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/IsNeedsRepair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) IsUnderRepair(ctx context.Context, in *PropertySerial, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/IsUnderRepair", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) IsInStock(ctx context.Context, in *PropertySerial, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/IsInStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *actionsAndHistoryClient) IsInArchive(ctx context.Context, in *PropertySerial, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/ActionsAndHistory/IsInArchive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionsAndHistoryServer is the server API for ActionsAndHistory service.
// All implementations must embed UnimplementedActionsAndHistoryServer
// for forward compatibility
type ActionsAndHistoryServer interface {
	//Getters
	GetActions(context.Context, *Empty) (*GetActionsResponse, error)
	GetHistory(context.Context, *HistoryRequest) (*HistoryResponse, error)
	//Создание записей
	CreateCard(context.Context, *ActionRequest) (*StatusWithProperties, error)
	InstallOnWorkspace(context.Context, *ActionRequest) (*Status, error)
	RemoveFromWorkspace(context.Context, *ActionRequest) (*Status, error)
	NeedRepair(context.Context, *ActionRequest) (*Status, error)
	SendToRepair(context.Context, *ActionRequest) (*Status, error)
	ReceiveFromRepair(context.Context, *ActionRequest) (*Status, error)
	Archive(context.Context, *ActionRequest) (*Status, error)
	DeArchive(context.Context, *ActionRequest) (*Status, error)
	ChangeName(context.Context, *ActionRequest) (*Status, error)
	ChangeInventory(context.Context, *ActionRequest) (*Status, error)
	AddToGroup(context.Context, *ActionRequest) (*Status, error)
	RemoveFromGroup(context.Context, *ActionRequest) (*Status, error)
	SendToWarehouse(context.Context, *ActionRequest) (*Status, error)
	GiveToEmployee(context.Context, *ActionRequest) (*Status, error)
	TakeFromEmployee(context.Context, *ActionRequest) (*Status, error)
	//Проверки на существование и доступность ввода
	IsCreated(context.Context, *PropertySerial) (*StatusWithProperty, error)
	IsOnWorkspace(context.Context, *PropertySerial) (*Status, error)
	IsNeedsRepair(context.Context, *PropertySerial) (*Status, error)
	IsUnderRepair(context.Context, *PropertySerial) (*Status, error)
	IsInStock(context.Context, *PropertySerial) (*Status, error)
	IsInArchive(context.Context, *PropertySerial) (*Status, error)
	mustEmbedUnimplementedActionsAndHistoryServer()
}

// UnimplementedActionsAndHistoryServer must be embedded to have forward compatible implementations.
type UnimplementedActionsAndHistoryServer struct {
}

func (UnimplementedActionsAndHistoryServer) GetActions(context.Context, *Empty) (*GetActionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActions not implemented")
}
func (UnimplementedActionsAndHistoryServer) GetHistory(context.Context, *HistoryRequest) (*HistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistory not implemented")
}
func (UnimplementedActionsAndHistoryServer) CreateCard(context.Context, *ActionRequest) (*StatusWithProperties, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCard not implemented")
}
func (UnimplementedActionsAndHistoryServer) InstallOnWorkspace(context.Context, *ActionRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InstallOnWorkspace not implemented")
}
func (UnimplementedActionsAndHistoryServer) RemoveFromWorkspace(context.Context, *ActionRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromWorkspace not implemented")
}
func (UnimplementedActionsAndHistoryServer) NeedRepair(context.Context, *ActionRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NeedRepair not implemented")
}
func (UnimplementedActionsAndHistoryServer) SendToRepair(context.Context, *ActionRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToRepair not implemented")
}
func (UnimplementedActionsAndHistoryServer) ReceiveFromRepair(context.Context, *ActionRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveFromRepair not implemented")
}
func (UnimplementedActionsAndHistoryServer) Archive(context.Context, *ActionRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Archive not implemented")
}
func (UnimplementedActionsAndHistoryServer) DeArchive(context.Context, *ActionRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeArchive not implemented")
}
func (UnimplementedActionsAndHistoryServer) ChangeName(context.Context, *ActionRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeName not implemented")
}
func (UnimplementedActionsAndHistoryServer) ChangeInventory(context.Context, *ActionRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeInventory not implemented")
}
func (UnimplementedActionsAndHistoryServer) AddToGroup(context.Context, *ActionRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToGroup not implemented")
}
func (UnimplementedActionsAndHistoryServer) RemoveFromGroup(context.Context, *ActionRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromGroup not implemented")
}
func (UnimplementedActionsAndHistoryServer) SendToWarehouse(context.Context, *ActionRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToWarehouse not implemented")
}
func (UnimplementedActionsAndHistoryServer) GiveToEmployee(context.Context, *ActionRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GiveToEmployee not implemented")
}
func (UnimplementedActionsAndHistoryServer) TakeFromEmployee(context.Context, *ActionRequest) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TakeFromEmployee not implemented")
}
func (UnimplementedActionsAndHistoryServer) IsCreated(context.Context, *PropertySerial) (*StatusWithProperty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsCreated not implemented")
}
func (UnimplementedActionsAndHistoryServer) IsOnWorkspace(context.Context, *PropertySerial) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsOnWorkspace not implemented")
}
func (UnimplementedActionsAndHistoryServer) IsNeedsRepair(context.Context, *PropertySerial) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsNeedsRepair not implemented")
}
func (UnimplementedActionsAndHistoryServer) IsUnderRepair(context.Context, *PropertySerial) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsUnderRepair not implemented")
}
func (UnimplementedActionsAndHistoryServer) IsInStock(context.Context, *PropertySerial) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsInStock not implemented")
}
func (UnimplementedActionsAndHistoryServer) IsInArchive(context.Context, *PropertySerial) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsInArchive not implemented")
}
func (UnimplementedActionsAndHistoryServer) mustEmbedUnimplementedActionsAndHistoryServer() {}

// UnsafeActionsAndHistoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActionsAndHistoryServer will
// result in compilation errors.
type UnsafeActionsAndHistoryServer interface {
	mustEmbedUnimplementedActionsAndHistoryServer()
}

func RegisterActionsAndHistoryServer(s grpc.ServiceRegistrar, srv ActionsAndHistoryServer) {
	s.RegisterService(&ActionsAndHistory_ServiceDesc, srv)
}

func _ActionsAndHistory_GetActions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).GetActions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/GetActions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).GetActions(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_GetHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).GetHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/GetHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).GetHistory(ctx, req.(*HistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_CreateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).CreateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/CreateCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).CreateCard(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_InstallOnWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).InstallOnWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/InstallOnWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).InstallOnWorkspace(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_RemoveFromWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).RemoveFromWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/RemoveFromWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).RemoveFromWorkspace(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_NeedRepair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).NeedRepair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/NeedRepair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).NeedRepair(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_SendToRepair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).SendToRepair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/SendToRepair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).SendToRepair(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_ReceiveFromRepair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).ReceiveFromRepair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/ReceiveFromRepair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).ReceiveFromRepair(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_Archive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).Archive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/Archive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).Archive(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_DeArchive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).DeArchive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/DeArchive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).DeArchive(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_ChangeName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).ChangeName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/ChangeName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).ChangeName(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_ChangeInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).ChangeInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/ChangeInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).ChangeInventory(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_AddToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).AddToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/AddToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).AddToGroup(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_RemoveFromGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).RemoveFromGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/RemoveFromGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).RemoveFromGroup(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_SendToWarehouse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).SendToWarehouse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/SendToWarehouse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).SendToWarehouse(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_GiveToEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).GiveToEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/GiveToEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).GiveToEmployee(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_TakeFromEmployee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).TakeFromEmployee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/TakeFromEmployee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).TakeFromEmployee(ctx, req.(*ActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_IsCreated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PropertySerial)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).IsCreated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/IsCreated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).IsCreated(ctx, req.(*PropertySerial))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_IsOnWorkspace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PropertySerial)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).IsOnWorkspace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/IsOnWorkspace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).IsOnWorkspace(ctx, req.(*PropertySerial))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_IsNeedsRepair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PropertySerial)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).IsNeedsRepair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/IsNeedsRepair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).IsNeedsRepair(ctx, req.(*PropertySerial))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_IsUnderRepair_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PropertySerial)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).IsUnderRepair(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/IsUnderRepair",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).IsUnderRepair(ctx, req.(*PropertySerial))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_IsInStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PropertySerial)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).IsInStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/IsInStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).IsInStock(ctx, req.(*PropertySerial))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActionsAndHistory_IsInArchive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PropertySerial)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsAndHistoryServer).IsInArchive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ActionsAndHistory/IsInArchive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsAndHistoryServer).IsInArchive(ctx, req.(*PropertySerial))
	}
	return interceptor(ctx, in, info, handler)
}

// ActionsAndHistory_ServiceDesc is the grpc.ServiceDesc for ActionsAndHistory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActionsAndHistory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ActionsAndHistory",
	HandlerType: (*ActionsAndHistoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetActions",
			Handler:    _ActionsAndHistory_GetActions_Handler,
		},
		{
			MethodName: "GetHistory",
			Handler:    _ActionsAndHistory_GetHistory_Handler,
		},
		{
			MethodName: "CreateCard",
			Handler:    _ActionsAndHistory_CreateCard_Handler,
		},
		{
			MethodName: "InstallOnWorkspace",
			Handler:    _ActionsAndHistory_InstallOnWorkspace_Handler,
		},
		{
			MethodName: "RemoveFromWorkspace",
			Handler:    _ActionsAndHistory_RemoveFromWorkspace_Handler,
		},
		{
			MethodName: "NeedRepair",
			Handler:    _ActionsAndHistory_NeedRepair_Handler,
		},
		{
			MethodName: "SendToRepair",
			Handler:    _ActionsAndHistory_SendToRepair_Handler,
		},
		{
			MethodName: "ReceiveFromRepair",
			Handler:    _ActionsAndHistory_ReceiveFromRepair_Handler,
		},
		{
			MethodName: "Archive",
			Handler:    _ActionsAndHistory_Archive_Handler,
		},
		{
			MethodName: "DeArchive",
			Handler:    _ActionsAndHistory_DeArchive_Handler,
		},
		{
			MethodName: "ChangeName",
			Handler:    _ActionsAndHistory_ChangeName_Handler,
		},
		{
			MethodName: "ChangeInventory",
			Handler:    _ActionsAndHistory_ChangeInventory_Handler,
		},
		{
			MethodName: "AddToGroup",
			Handler:    _ActionsAndHistory_AddToGroup_Handler,
		},
		{
			MethodName: "RemoveFromGroup",
			Handler:    _ActionsAndHistory_RemoveFromGroup_Handler,
		},
		{
			MethodName: "SendToWarehouse",
			Handler:    _ActionsAndHistory_SendToWarehouse_Handler,
		},
		{
			MethodName: "GiveToEmployee",
			Handler:    _ActionsAndHistory_GiveToEmployee_Handler,
		},
		{
			MethodName: "TakeFromEmployee",
			Handler:    _ActionsAndHistory_TakeFromEmployee_Handler,
		},
		{
			MethodName: "IsCreated",
			Handler:    _ActionsAndHistory_IsCreated_Handler,
		},
		{
			MethodName: "IsOnWorkspace",
			Handler:    _ActionsAndHistory_IsOnWorkspace_Handler,
		},
		{
			MethodName: "IsNeedsRepair",
			Handler:    _ActionsAndHistory_IsNeedsRepair_Handler,
		},
		{
			MethodName: "IsUnderRepair",
			Handler:    _ActionsAndHistory_IsUnderRepair_Handler,
		},
		{
			MethodName: "IsInStock",
			Handler:    _ActionsAndHistory_IsInStock_Handler,
		},
		{
			MethodName: "IsInArchive",
			Handler:    _ActionsAndHistory_IsInArchive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/history.proto",
}
