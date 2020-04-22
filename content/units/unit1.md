---
title: "Unit 1"
date: 2020-04-17T15:16:56+02:00
weight: 10
---

## What's happening in this unit?
The first unit mainly exists to get you familar with the format and style of this guide.

## Action items

* open example.go to familiarize yourself with the basic structure of code
* build the example application
* build the example docker image
* run the example docker image

## Instructions

### Task 1
Open the example app in this unit and familiarize yourself with the code.

This is not attempting to be a go tutorial but might be able to help you along the way.

Therefore the applications will contain a lot more comments than usual to make sure that nothing is left unclear

### Task 2
Make sure your go dev environment is set up and then run `go build example.go` you can of course use the `-o` flag
to name your compiled binary what you like `go build -o rira12621-superawesome-go-app example.go`

### Task 3
Assuming that you have set up docker locally you can just from within this directory run the following command `docker build -t mygreatgoapp:latest .`

Let's quickly go over what the command is doing in case it has been a while since you last used docker from the CLI.

* `docker` is exectuing the docker binary
* `build` is sending the context to the docker daemon and instructing it to build
* `-t` is telling docker to tag the newly built image `mygreatgoapp:latest` where `mygreatgoapp` is the name of your image and `latest` tags this build as "latest"
* `.` is the instruction in which directory to look for the Dockerfile which in this case is the current directory

### Task 4
Now that we have manually built the app as well as built the docker image, all that is left is to run it.

Do so by executing `docker run --rm mygreatgoapp:latest`

It will let you know if everything worked fine.

In case you are wondering, the `--rm` flag is only to have docker clean up the container after it's done.

