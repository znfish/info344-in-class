# Docker Excercise One

This is a simple challenge that involves publishing ports when running a Docker container.

You will **NOT** need to modify `main.go`

You will **NOT** need to modify the `Dockerfile`

## What do I do?

### Build the go binary

You will need to build the go binary for use with the docker container. Go supports cross-compiling and you can set which OS to build for with the `GOOS` environment variable.

Build the docker exercise go binary:

`GOOS=linux go build`

*If you're on a Linux machine you can just do `go build`.*

If you get no feedback then the binary was built correctly.

### Build the docker excericse image

`docker build -t <your-docker-hub-name>/docker-exercise-1 <dir>`

Docker has version control for images. Versions of images are specified by `<docker-image-name>:<version>`. You can refer to the latest version by `<docker-image-name>:latest`. The `-t` flag automatically tags the newly built image as the latest version of that image. You will see something like:

`Successfully tagged brendankellogg/docker-exercise-1:latest`

when you build the container.

`<your-docker-hub-name>` is your username on dockerhub.

`<dir>` is the directory to build from. In most cases this is the current directory which you can specify with `.`

### Run your newly built container

`docker run -d <OPTIONS> <your-docker-hub-name>/docker-exercise-1`

The `-d` flag runs the container in detached mode. This menas that the docker container will not consume the terminal window when it's running so you can contine to run commands.

The `<OPTIONS>` are additional arguments for running your container. [There are quite a few arguments](https://docs.docker.com/engine/reference/commandline/run/#options) that you can use.

For this exercise, you will only need to publish port `4000` on the container. You can publish ports with the `--publish` or `-p` flags in the run command. Add the following flag to the run command:

`-p <host-port>:<container-port>`

Set this up so port `4000` on your host machine maps to port `4000` on your docker container.

## Test it out

Open a browser window and go to [localhost:4000](http://localhost:4000). If everything went well you should see a success message. If you aren't able to connect then you didn't publish the ports quite right.

## Stop and remove the container

You likely don't need or want this container to be running any more so let's stop and remove it.

Run `docker ps` to list all currently running containers and copy the `CONTAINER ID` of the docker exercise container.

Stop the conainer:

`docker stop <container-id>`

Remove the container:

`docker rm <container-id>`

*When you run `docker stop` and `docker rm` you will see the container ID printed to your terminal window. You can treat this as a success message.*

You can also stop and remove a running container in one command:

`docker rm -f <container-id>`

This forces the removal of the container, even if it is currently running.
