package pnetserver

import (
	"bufio"
	"github.com/cpssd/paranoid/pfsm/returncodes"
	pb "github.com/cpssd/paranoid/proto/paranoidnetwork"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"os/exec"
	"strconv"
)

func (s *ParanoidServer) Unlink(ctx context.Context, req *pb.UnlinkRequest) (*pb.EmptyMessage, error) {
	command := exec.Command("pfsm", "-n", "unlink", ParanoidDir, req.Path)
	stdout, err := command.StdoutPipe()
	if err != nil {
		log.Println("ERROR: Could not capture stdout of subprocess.")
	}
	read := bufio.NewReader(stdout)
	err = command.Run()
	if err != nil {
		log.Printf("ERROR: Could not unlink file %s: %v.\n", req.Path, err)
		returnError := grpc.Errorf(codes.Internal, "could not unlink file %s: %v", req.Path, err)
		return &pb.EmptyMessage{}, returnError
	}

	pfsmoutput, _, err := read.ReadLine()
	codeBytes := pfsmoutput[:2]
	code, err := strconv.Atoi(string(codeBytes))
	if err != nil {
		log.Println("ERROR: Could not interpret PFSM error code.")
		returnError := grpc.Errorf(codes.Internal, "could not unlink file %s: %v", req.Path, err)
		return &pb.EmptyMessage{}, returnError
	}

	switch code {
	case returncodes.EACCES:
		log.Printf("INFO: Do not have permission to edit %s.\n", req.Path)
		returnError := grpc.Errorf(codes.PermissionDenied,
			"do not have permission to edit %s",
			req.Path)
		return &pb.EmptyMessage{}, returnError
	case returncodes.ENOENT:
		log.Printf("INFO: File %s does not exist.\n", req.Path)
		returnError := grpc.Errorf(codes.NotFound,
			"file %s does not exist",
			req.Path)
		return &pb.EmptyMessage{}, returnError
	}

	return &pb.EmptyMessage{}, nil
}
