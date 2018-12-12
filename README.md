# GOLangWorkspace

Exploring the GO language.

Setup
=====

.bash_profile

sets up

export GOPATH="/Users/JohnMoshakis/Documents/develop/GOLANGWorkspace"

Downloading Packages
====================

example from the console

go get github.com/aws/aws-lambda-go/lambda


Building for AWS
================

From command prompt

From the src folder

$ env GOOS=linux GOARCH=amd64 go build -o ~/Downloads/main GatewayLambdaHelloWorld
$ zip -j ~/Downloads/main.zip ~/Downloads/main
