# BoxUp

A cross platform directory hosting service, that allows you to box up and retrieve the contents of a directory.

For the full explanation on why this project exists see [here](https://christophergreaves.co.uk/projects/boxup)

## Installation

Installation is easy, but will require some pre-requisite downloads.  
I hope in the future to have a release in github with all the files pre-compiled

### Pre-requisite

You will need to download the Go binaries for your OS. visit https://golang.org/dl/ to download the correct binaries.

Once you have this installed and have tested its working ([see here](https://golang.org/doc/install#testing)), you can begin to install BoxUp.

### Installing Client

To get the client executable, simply enter `go install github.com/chris-greaves/boxup/boxup` into a terminal. This will pull down the code and build it.

Providing that your environment is setup correctly, you should see the boxup executable in the bin folder of your first directory on GOPATH.

### Installing Server

To get the client executable, simply enter `go install github.com/chris-greaves/boxup/boxup-server` into a terminal. This will pull down the code and build it.

Providing that your environment is setup correctly, you should see the boxup-server executable in the bin folder of your first directory on GOPATH.

Once this file exists, you can run it and it will start hosting the boxup API for any boxup CLIs to connect to.