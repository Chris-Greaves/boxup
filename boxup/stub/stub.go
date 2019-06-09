// Copyright © 2018 Christopher Greaves <cjgreaves97@hotmail.co.uk>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stub

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"

	pb "github.com/chris-greaves/boxup/boxup_service"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

// ServiceClient is a stub object to connect to the BoxUp server
type ServiceClient struct {
	client pb.BoxUpServiceClient
	conn   *grpc.ClientConn
}

// New attempts to connect to the server at the url provided, and returns a new instance of ServiceClient connected to that server if successful
func New(url string) (*ServiceClient, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to connect to server at '%v'", url)
	}

	client := pb.NewBoxUpServiceClient(conn)
	return &ServiceClient{client: client, conn: conn}, nil
}

// Close closes the connection to the server
func (c *ServiceClient) Close() {
	c.conn.Close()
}

// List prints all the Boxes currently stored on the server
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

// Get downloads a Box from the server
func (c *ServiceClient) Get(name string) error {
	stream, err := c.client.Get(context.Background(), &pb.BoxInfo{Name: name})
	defer stream.CloseSend()
	wd, err := os.Getwd()
	if err != nil {
		return errors.Wrap(err, "error getting current directory")
	}

	file, err := os.Create(path.Join(wd, name))
	if err != nil {
		return errors.Wrapf(err, "error occurred creating file on client. BoxName=%v", name)
	}
	defer file.Close()

	for {
		chunk, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}

			return errors.Wrapf(err, "failed unexpectedely while retrieving chunks from stream. BoxName=%v", name)
		}

		_, err = file.Write(chunk.Data)
		if err != nil {
			return errors.Wrapf(err, "failed unexpectedely while writing chunks to file. BoxName=%v", name)
		}
	}

	return nil
}
