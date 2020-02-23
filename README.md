![Build](https://github.com/YashdalfTheGray/federator/workflows/Build/badge.svg)

# federator

A utility to federate into an AWS account using AWS Security Token Service and then get a link to go directly into the console.

## Using the tool

The federator tool has two modes, `link` and `creds`. These take the form of subcommands. The `link` subcommand gives you a link to sign into the AWS Console and the `creds` subcommand just gives you temporary credentials with the right commands to import it into a shell instance. The only argument required for both subcommands is `--role-arn` which is the ARN of the role to assume.

## Development

### Installation

Once the code is cloned, you're going to need to run a preliminary `make build` so that all the dependencies come down and a build runs. Once that is done, an executable in `./bin` will be added.

### Running

The one requirement for running this tool is having installed and fully configured AWS CLI. Being able to run `aws sts get-caller-identity` successfully wil tell you that it is installed and configured properly.

The executable in `bin` has the right permissions, so running `./bin/federator` will let you run it in development.
