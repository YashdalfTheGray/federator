# federator

A utility to federate into an AWS account using AWS Security Token Service and then get a link to go directly into the console.

## Using the tool

Instructions to come later on running the tool.

## Development

### Installation

Once the code is cloned, you're going to need to run a preliminary `make build` so that all the dependencies come down and a build runs. Once that is done, an executable in `./bin` will be added.

### Running

The one requirement for running this tool is having installed and fully configured AWS CLI. Being able to run `aws sts get-caller-identity` successfully wil tell you that it is installed and configured properly.

The executable in `bin` has the right permissions, so running `./bin/federator` will let you run it in development.
