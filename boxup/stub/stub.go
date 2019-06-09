package stub

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/chris-greaves/boxup/boxup_service"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	client pb.BoxUpServiceClient
	conn   *grpc.ClientConn
}

func New(url string) *ServiceClient {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}

	client := pb.NewBoxUpServiceClient(conn)
	return &ServiceClient{client: client, conn: conn}
}

func (c *ServiceClient) Close() {
	c.conn.Close()
}

func (c *ServiceClient) List() error {
	stream, err := c.client.List(context.Background(), &pb.SearchQuery{})
	if err != nil {
		return err
	}
	for {
		box, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Printf("Box: %v\n", box)
	}
	return nil
}
