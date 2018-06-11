package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {

}

var rootCmd = &cobra.Command{
	Use:   "boxup",
	Short: "BoxUp is a directory and file hosting provider",
	Long: `This is the CLi tool for BoxUp. 
Use this tool to connect to and command a BoxUp server.append
The BoxUp server can be found at https://bitbucket.org/ChristopherGreaves/boxup/src/master/boxup-server/.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
