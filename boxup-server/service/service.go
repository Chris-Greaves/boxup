package service

import (
	"log"
	"os"
	"path/filepath"

	pb "github.com/chris-greaves/boxup/boxup_service"
)

// Box is an Archive stored on the server
type Box struct {
	Name string
	Path string
}

// BoxUpService is the server object for the service
type BoxUpService struct {
	logger        *log.Logger
	storagePath   string
	streamBitSize int
	boxes         map[string]Box
}

// New creates a new instance of a BoxUpService
//
// It takes a storagePath which is an absolute path to the area Boxes should be stored.
// And a logger to specify where the logs should go.
func New(storagePath string, logger *log.Logger) *BoxUpService {
	boxes := getExistingBoxes(storagePath)
	return &BoxUpService{
		logger:        logger,
		boxes:         boxes,
		storagePath:   storagePath,
		streamBitSize: 1024}
}

func getExistingBoxes(path string) map[string]Box {
	var archs = map[string]Box{}
	filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			_, filename := filepath.Split(path)
			var arch = Box{Path: path, Name: filename}
			log.Printf("Adding %v to list of archives", filename)

			archs[filename] = arch
			return nil
		})
	return archs
}

func (s *BoxUpService) List(query *pb.SearchQuery, stream pb.BoxUpService_ListServer) error {
	return nil
}

func (s *BoxUpService) Get(box *pb.BoxInfo, stream pb.BoxUpService_GetServer) error {
	return nil
}

func (s *BoxUpService) Send(stream pb.BoxUpService_SendServer) error {
	return nil
}
