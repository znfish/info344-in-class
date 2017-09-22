# Docker Excercise Two

This is a simple challenge that involves publishing ports and setting environment variables when running a Docker container.

You will **NOT** need to modify `main.go`

You will **NOT** need to modify the `Dockerfile`

## What do I do?

### Build the go binary

Build the docker exercise go binary:

`GOOS=linux go build`

*If you're on a Linux machine you can just do `go build`.*

### Build the docker excericse image

`docker build -t <your-docker-hub-name>/docker-exercise-2 .`

### Run your newly built container

`docker run -d <OPTIONS> <your-docker-hub-name>/docker-exercise-2`

In this exercise, we will be changing what port the success server runs on. In the last exercise it would always listen on port `4000` but we can change that by setting a `PORT` environment variable on our Docker container. For reference, here is the official documentation on [docker run arguments.](https://docs.docker.com/engine/reference/commandline/run/#options)

Set the `PORT` environment variable for your docker container with the `-e` flag:

`-e PORT=<PORT>`

Remember to publish the same port when you run your containeror You won't be able to connect on that port!

Unlike the first excericse, there is a chance that your container will crash at startup (don't worry, it's not as bad as it sounds!). If run `docker ps` right after it starts and you don't see it listed, then it stopped running. Fortunately, it's common for containers to log why they stopped. You can see these logs with `docker logs <container-id>`. You can list all containers, even stopped ones with `docker ps -a`.

You may think that getting the container id of a stopped container is a lot of work, and it is! There is an easier way. You can name a container with the `--name` flag. For example, if you gave `--name my-container` to the run command you could refer to the container by `my-contaienr`. You could get the logs with `docker logs my-container`.

## Test it out

Open a browser window and go to `localhost:<port>`. If everything went well you should see a success message. If you aren't able to connect then you there is an issue with your ports!

If your container crashed, use `docker logs <container-id>` to see what the issue is and fix it!

*If you named your container, you will need to remove it before you can start another with the same name.*

## Stop and remove the container

You likely don't need or want this container to be running any more so let's stop and remove it.

Run `docker ps` to list all currently running containers and copy the `CONTAINER ID` of the docker exercise container.

Stop the conainer:

`docker stop <container-id>` or `docker stop <container-name>`

Remove the container:

`docker rm <container-id>` `docker rm <container-name>`

*When you run `docker stop` and `docker rm` you will see the container ID printed to your terminal window. You can treat this as a success message.*

Remember that you can also stop and remove a running container in one command:

`docker rm -f <container-id>` or `docker rm -f <container-name>`

This forces the removal of the container, even if it is currently running.
