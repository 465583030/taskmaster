package controller

import (
	"github.com/timeloveboy/taskmaster/pb"

	"golang.org/x/net/context"
)

type MoeTaskMasterServer struct {
}

func (this MoeTaskMasterServer) Stop(*pb.Owner, pb.TaskMaster_StopServer) (err error) {
	return
}
func (this MoeTaskMasterServer) Get(*pb.Owner, pb.TaskMaster_GetServer) (err error) {
	return
}
func (this MoeTaskMasterServer) Report(ctx context.Context, req *pb.TaskState) (gc *pb.RespTask, err error) {
	return
}
