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

To get the client, simply type `go get -u github.com/chris-greaves/boxup/boxup`. This will pull down the code from github, along with all its dependencies.

Once this has finished, go into the newly created "boxup" folder, and run the command `go build`. You should see a executable file called boxup (The file created will vary depending on the OS)

You can move this file wherever, but the best place would be a directory referenced by PATH. this means you can run the commands without having to reference the path to the file.

### Installing Server

To get the client, simply type `go get -u github.com/chris-greaves/boxup/boxup-server`. This will pull down the code from github, along with all its dependencies.

Once this has finished, go into the newly created "boxup-server" folder, and run the command `go build` You should see a executable file called "boxup-server" (The file created will vary depending on the OS)

Once this file exists, you can run it and it will start hosting the boxup API for any boxup CLIs to connect to.