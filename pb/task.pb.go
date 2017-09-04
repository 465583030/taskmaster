// Code generated by protoc-gen-go. DO NOT EDIT.
// source: task.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	task.proto

It has these top-level messages:
	Owner
	TaskState
	Task
	RespTask
	GetTaskResp
	StopTaskResp
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Owner struct {
	Owner string `protobuf:"bytes,1,opt,name=Owner" json:"Owner,omitempty"`
	Group string `protobuf:"bytes,2,opt,name=Group" json:"Group,omitempty"`
}

func (m *Owner) Reset()                    { *m = Owner{} }
func (m *Owner) String() string            { return proto.CompactTextString(m) }
func (*Owner) ProtoMessage()               {}
func (*Owner) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Owner) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Owner) GetGroup() string {
	if m != nil {
		return m.Group
	}
	return ""
}

type TaskState struct {
	Owner *Owner `protobuf:"bytes,1,opt,name=Owner" json:"Owner,omitempty"`
	Task  *Task  `protobuf:"bytes,2,opt,name=Task" json:"Task,omitempty"`
}

func (m *TaskState) Reset()                    { *m = TaskState{} }
func (m *TaskState) String() string            { return proto.CompactTextString(m) }
func (*TaskState) ProtoMessage()               {}
func (*TaskState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TaskState) GetOwner() *Owner {
	if m != nil {
		return m.Owner
	}
	return nil
}

func (m *TaskState) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

type Task struct {
	TaskId int64 `protobuf:"zigzag64,1,opt,name=TaskId" json:"TaskId,omitempty"`
}

func (m *Task) Reset()                    { *m = Task{} }
func (m *Task) String() string            { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()               {}
func (*Task) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Task) GetTaskId() int64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

type RespTask struct {
	Task          *Task `protobuf:"bytes,1,opt,name=Task" json:"Task,omitempty"`
	EnableSeconds int64 `protobuf:"zigzag64,2,opt,name=EnableSeconds" json:"EnableSeconds,omitempty"`
}

func (m *RespTask) Reset()                    { *m = RespTask{} }
func (m *RespTask) String() string            { return proto.CompactTextString(m) }
func (*RespTask) ProtoMessage()               {}
func (*RespTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RespTask) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func (m *RespTask) GetEnableSeconds() int64 {
	if m != nil {
		return m.EnableSeconds
	}
	return 0
}

type GetTaskResp struct {
	RespTask *RespTask `protobuf:"bytes,1,opt,name=RespTask" json:"RespTask,omitempty"`
}

func (m *GetTaskResp) Reset()                    { *m = GetTaskResp{} }
func (m *GetTaskResp) String() string            { return proto.CompactTextString(m) }
func (*GetTaskResp) ProtoMessage()               {}
func (*GetTaskResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetTaskResp) GetRespTask() *RespTask {
	if m != nil {
		return m.RespTask
	}
	return nil
}

type StopTaskResp struct {
	Tasks *Task `protobuf:"bytes,1,opt,name=Tasks" json:"Tasks,omitempty"`
}

func (m *StopTaskResp) Reset()                    { *m = StopTaskResp{} }
func (m *StopTaskResp) String() string            { return proto.CompactTextString(m) }
func (*StopTaskResp) ProtoMessage()               {}
func (*StopTaskResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *StopTaskResp) GetTasks() *Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func init() {
	proto.RegisterType((*Owner)(nil), "pb.Owner")
	proto.RegisterType((*TaskState)(nil), "pb.TaskState")
	proto.RegisterType((*Task)(nil), "pb.Task")
	proto.RegisterType((*RespTask)(nil), "pb.RespTask")
	proto.RegisterType((*GetTaskResp)(nil), "pb.GetTaskResp")
	proto.RegisterType((*StopTaskResp)(nil), "pb.StopTaskResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TaskMaster service

type TaskMasterClient interface {
	// Stop 服务端向下推送,终止任务
	Stop(ctx context.Context, in *Owner, opts ...grpc.CallOption) (TaskMaster_StopClient, error)
	// Get 服务端向下推送,接收任务
	Get(ctx context.Context, in *Owner, opts ...grpc.CallOption) (TaskMaster_GetClient, error)
	// Report 客户端请求,续租任务
	Report(ctx context.Context, in *TaskState, opts ...grpc.CallOption) (*RespTask, error)
}

type taskMasterClient struct {
	cc *grpc.ClientConn
}

func NewTaskMasterClient(cc *grpc.ClientConn) TaskMasterClient {
	return &taskMasterClient{cc}
}

func (c *taskMasterClient) Stop(ctx context.Context, in *Owner, opts ...grpc.CallOption) (TaskMaster_StopClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TaskMaster_serviceDesc.Streams[0], c.cc, "/pb.TaskMaster/Stop", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskMasterStopClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TaskMaster_StopClient interface {
	Recv() (*GetTaskResp, error)
	grpc.ClientStream
}

type taskMasterStopClient struct {
	grpc.ClientStream
}

func (x *taskMasterStopClient) Recv() (*GetTaskResp, error) {
	m := new(GetTaskResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskMasterClient) Get(ctx context.Context, in *Owner, opts ...grpc.CallOption) (TaskMaster_GetClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_TaskMaster_serviceDesc.Streams[1], c.cc, "/pb.TaskMaster/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &taskMasterGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TaskMaster_GetClient interface {
	Recv() (*StopTaskResp, error)
	grpc.ClientStream
}

type taskMasterGetClient struct {
	grpc.ClientStream
}

func (x *taskMasterGetClient) Recv() (*StopTaskResp, error) {
	m := new(StopTaskResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *taskMasterClient) Report(ctx context.Context, in *TaskState, opts ...grpc.CallOption) (*RespTask, error) {
	out := new(RespTask)
	err := grpc.Invoke(ctx, "/pb.TaskMaster/Report", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TaskMaster service

type TaskMasterServer interface {
	// Stop 服务端向下推送,终止任务
	Stop(*Owner, TaskMaster_StopServer) error
	// Get 服务端向下推送,接收任务
	Get(*Owner, TaskMaster_GetServer) error
	// Report 客户端请求,续租任务
	Report(context.Context, *TaskState) (*RespTask, error)
}

func RegisterTaskMasterServer(s *grpc.Server, srv TaskMasterServer) {
	s.RegisterService(&_TaskMaster_serviceDesc, srv)
}

func _TaskMaster_Stop_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Owner)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskMasterServer).Stop(m, &taskMasterStopServer{stream})
}

type TaskMaster_StopServer interface {
	Send(*GetTaskResp) error
	grpc.ServerStream
}

type taskMasterStopServer struct {
	grpc.ServerStream
}

func (x *taskMasterStopServer) Send(m *GetTaskResp) error {
	return x.ServerStream.SendMsg(m)
}

func _TaskMaster_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Owner)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TaskMasterServer).Get(m, &taskMasterGetServer{stream})
}

type TaskMaster_GetServer interface {
	Send(*StopTaskResp) error
	grpc.ServerStream
}

type taskMasterGetServer struct {
	grpc.ServerStream
}

func (x *taskMasterGetServer) Send(m *StopTaskResp) error {
	return x.ServerStream.SendMsg(m)
}

func _TaskMaster_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TaskState)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskMasterServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TaskMaster/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskMasterServer).Report(ctx, req.(*TaskState))
	}
	return interceptor(ctx, in, info, handler)
}

var _TaskMaster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TaskMaster",
	HandlerType: (*TaskMasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Report",
			Handler:    _TaskMaster_Report_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stop",
			Handler:       _TaskMaster_Stop_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Get",
			Handler:       _TaskMaster_Get_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "task.proto",
}

func init() { proto.RegisterFile("task.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xd1, 0x4a, 0xf3, 0x40,
	0x10, 0x85, 0x9b, 0xfe, 0x6d, 0x68, 0x26, 0x2d, 0xbf, 0x0c, 0x22, 0xa5, 0x48, 0x95, 0x45, 0x6a,
	0xaf, 0x82, 0xa4, 0x17, 0x3e, 0x81, 0x04, 0x05, 0x15, 0x36, 0xbe, 0x40, 0x62, 0xe7, 0xaa, 0x92,
	0x5d, 0xb2, 0x2b, 0x3e, 0x81, 0xef, 0x2d, 0x33, 0x49, 0x9a, 0x14, 0xbd, 0xca, 0x99, 0x33, 0x73,
	0xbe, 0x19, 0xb2, 0x00, 0xbe, 0x70, 0x87, 0xc4, 0xd6, 0xc6, 0x1b, 0x1c, 0xdb, 0x52, 0xed, 0x60,
	0xfa, 0xfa, 0x55, 0x51, 0x8d, 0xe7, 0xad, 0x58, 0x06, 0xd7, 0xc1, 0x36, 0xd2, 0xbd, 0x9b, 0xd5,
	0xe6, 0xd3, 0x2e, 0xc7, 0x8d, 0x2b, 0x85, 0x7a, 0x82, 0xe8, 0xad, 0x70, 0x87, 0xdc, 0x17, 0x9e,
	0xf0, 0x6a, 0x18, 0x8c, 0xd3, 0x28, 0xb1, 0x65, 0x22, 0x46, 0xc7, 0xb8, 0x84, 0x09, 0x4f, 0x0b,
	0x22, 0x4e, 0x67, 0xdc, 0xe7, 0x5a, 0x8b, 0xab, 0xd6, 0x4d, 0x17, 0x2f, 0x20, 0xe4, 0xef, 0xe3,
	0x5e, 0x38, 0xa8, 0xdb, 0x4a, 0xbd, 0xc0, 0x4c, 0x93, 0xb3, 0x32, 0xd3, 0x91, 0x82, 0xbf, 0x48,
	0x78, 0x03, 0x8b, 0x87, 0xaa, 0x28, 0x3f, 0x28, 0xa7, 0x77, 0x53, 0xed, 0x9d, 0x2c, 0x44, 0x7d,
	0x6a, 0xaa, 0x7b, 0x88, 0x33, 0xf2, 0x12, 0x23, 0x67, 0x71, 0xdb, 0xe3, 0x5b, 0xec, 0x9c, 0xb1,
	0x9d, 0xa7, 0x8f, 0x5d, 0x95, 0xc0, 0x3c, 0xf7, 0xc6, 0x1e, 0x93, 0x6b, 0x98, 0xb2, 0x76, 0xbf,
	0xae, 0x69, 0xec, 0xf4, 0x3b, 0x00, 0x60, 0xf5, 0x5c, 0x38, 0x4f, 0x35, 0x6e, 0x60, 0xc2, 0x71,
	0xec, 0xff, 0xcf, 0xea, 0x3f, 0xcb, 0xc1, 0x31, 0x6a, 0x74, 0x17, 0xe0, 0x06, 0xfe, 0x65, 0xe4,
	0x87, 0x63, 0x67, 0x2c, 0x87, 0xab, 0x65, 0xee, 0x16, 0x42, 0x4d, 0xd6, 0xd4, 0x1e, 0x17, 0xdd,
	0x66, 0x79, 0x8f, 0xd5, 0xc9, 0xfd, 0x6a, 0x54, 0x86, 0xf2, 0xd8, 0xbb, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xe7, 0xdd, 0xfd, 0x51, 0xfa, 0x01, 0x00, 0x00,
}
