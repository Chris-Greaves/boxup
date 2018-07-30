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

package cmd

import (
	"archive/tar"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// GetBoxCmd represents the GetBox command
var GetBoxCmd = &cobra.Command{
	Use:   "GetBox [box name]",
	Short: "Get a Box from a BoxUp Server",
	Long: `This command will get the contents of a box hosted on a BoxUp Server.
	
The files and directories will be place wherever the current working directory is unless specified using the -o flag
Will assume local BoxUp Server unless specifed using --host http://domain.com or --host http://10.0.0.1
To specify a port use --port 5656`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for index, arg := range args {
			fmt.Printf("%v:%v\n", index, arg)
		}
		hostFlag := cmd.Flag("host")
		fmt.Printf("host:%v\n", hostFlag.Value.String())
		portFlag := cmd.Flag("port")
		fmt.Printf("port:%v\n", portFlag.Value.String())
		outputFlag := cmd.Flag("output")
		fmt.Printf("output:%v\n", outputFlag.Value.String())

		resp, err := http.Get("http://localhost:5950/GetBox/test-folder")
		if err != nil {
			fmt.Printf("Error occured: %v", err)
		}
		defer resp.Body.Close()

		gzipReader, err := gzip.NewReader(resp.Body)
		if err != nil {
			fmt.Printf("Error occured: %v", err)
		}
		defer gzipReader.Close()

		// err = untar(tarReader, outputFlag.Value.String())
		err = untar(gzipReader, "D:\\Projects\\boxup\\output")
		if err != nil {
			fmt.Printf("Error occured unzipping: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(GetBoxCmd)

	// Here you will define your flags and configuration settings.

	GetBoxCmd.Flags().String("host", "localhost", "Specify the host to connect")
	GetBoxCmd.Flags().Int("port", 5950, "Specify the port to use")
	GetBoxCmd.Flags().StringP("output", "o", "", "Specify the output directory (default Current Directory)")
}

func untar(reader io.Reader, target string) error {
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return errors.New(fmt.Sprintf("Error reading next: %v", err))
		}

		path := filepath.Join(target, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return errors.New(fmt.Sprintf("Error making directory: %v", err))
			}
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			return errors.New(fmt.Sprintf("Error opening file: %v", err))
		}
		defer file.Close()
		_, err = io.Copy(file, tarReader)
		if err != nil {
			return errors.New(fmt.Sprintf("Error copying file: %v", err))
		}
	}

	return nil
}
