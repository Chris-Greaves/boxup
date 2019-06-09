package service

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	pb "github.com/chris-greaves/boxup/boxup_service"
	"github.com/pkg/errors"
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
	var boxes = map[string]Box{}
	filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}

			_, filename := filepath.Split(path)
			var box = Box{Path: path, Name: filename}
			log.Printf("Adding %v to list of archives", filename)

			boxes[filename] = box
			return nil
		})
	return boxes
}

// List gets a list of all the Boxes currently stored by the server
func (s *BoxUpService) List(query *pb.SearchQuery, stream pb.BoxUpService_ListServer) error {
	s.logger.Printf("Received call to \"List\". Query string=%v", nil)
	start := time.Now()
	for _, box := range s.boxes {
		err := stream.Send(&pb.BoxInfo{Name: box.Name})
		if err != nil {
			err = errors.Wrap(err, "error sending box back to client")
			s.logger.Print(err)
			return err
		}
	}
	s.logger.Printf("Call to \"List\" took %v", time.Since(start))
	return nil
}

// Get retrieves a Box from the server
func (s *BoxUpService) Get(info *pb.BoxInfo, stream pb.BoxUpService_GetServer) error {
	s.logger.Printf("Received call to \"Get\". BoxName=%v", info.Name)
	start := time.Now()

	var writing = true
	box, ok := s.boxes[info.Name]
	if !ok {
		return errors.New("box not found")
	}

	file, err := os.Open(box.Path)
	if err != nil {
		return errors.Wrap(err, "error occurred when opening file on server")
	}

	buf := make([]byte, s.streamBitSize)
	for writing {
		n, err := file.Read(buf)

		if err != nil {
			if err == io.EOF {
				writing = false
				err = nil
				continue
			}

			return errors.Wrap(err,
				"errored while reading file")
		}

		stream.Send(&pb.BoxChunk{
			Filename: info.Name,
			Data:     buf[:n],
		})
	}

	s.logger.Printf("Call to \"Get\" took %v", time.Since(start))
	return nil
}

// Send streams a Box up to the server to be stored
func (s *BoxUpService) Send(stream pb.BoxUpService_SendServer) error {
	return nil
}
