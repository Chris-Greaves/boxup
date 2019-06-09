package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path"

	"github.com/Chris-Greaves/boxup/boxup-server/service"
	pb "github.com/chris-greaves/boxup/boxup_service"
	homedir "github.com/mitchellh/go-homedir"
	"google.golang.org/grpc"
)

var (
	logger = log.New(os.Stderr, "BoxUp: ", log.Lshortfile|log.Ldate|log.Ltime|log.Lmicroseconds|log.LUTC)
)

func main() {
	lis := getListener(3535)

	path, err := getStoragePath()
	if err != nil {
		logger.Fatalf("Unable to get Storage Path: %v", err)
	}
	logger.Printf("Storage path is %v", path)

	s := service.New(path, logger)
	gs := grpc.NewServer()
	pb.RegisterBoxUpServiceServer(gs, s)

	logger.Print("Starting server on port 3535")
	gs.Serve(lis)
}

func getListener(port int) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
	return lis
}

func getStoragePath() (string, error) {
	hdir, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	dir := path.Join(hdir, ".boxup")

	_, statErr := os.Stat(dir)
	if statErr != nil {
		if os.IsNotExist(statErr) {
			os.MkdirAll(dir, os.ModeDir)
		} else {
			return "", statErr
		}
	}

	return dir, nil
}
