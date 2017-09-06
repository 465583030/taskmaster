package controller

import (
	"github.com/timeloveboy/taskmaster/pb"

	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"time"
)

type MoeTaskMasterServer struct {
	Owner string
	Group string

	Commands chan pb.TaskResp
}

func (ms MoeTaskMasterServer) Exit() {
	close(ms.Commands)

}
func (ms MoeTaskMasterServer) CheckOwner(os *pb.OwnerState) bool {
	if ms.Owner != "" || ms.Group != "" {
		return ms.Group == os.Group && ms.Owner == os.Owner
	} else {
		if os.Owner != "" && os.Group != "" {
			ms.Owner = os.Owner
			ms.Group = os.Group
			return true
		} else {
			return false
		}
	}

}
func (ms MoeTaskMasterServer) SendCommands(cs pb.TaskMaster_ConnectServer) error {
	for {
		tr := <-ms.Commands
		err := cs.Send(&tr)
		if err != nil {
			ms.Exit()
			return err
		}
	}
}
func (ms MoeTaskMasterServer) Connect(cs pb.TaskMaster_ConnectServer) (err error) {
	go ms.SendCommands(cs)
	for {
		time.Sleep(time.Second)
		os, err := cs.Recv()
		if err != nil {
			return err
		}
		if !ms.CheckOwner(os) {
			return errors.New("you exist !")
		}
		fmt.Println("ownerstate", os)
	}
	return
}

func (this MoeTaskMasterServer) Register(ctx context.Context, req *pb.Task) (t *pb.Null, err error) {
	return
}
