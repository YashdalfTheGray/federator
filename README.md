![Build](https://github.com/YashdalfTheGray/federator/workflows/Build/badge.svg)

# federator

A utility to federate into an AWS account using AWS Security Token Service and then get a link to go directly into the console.

## Usage

### Installation

The best way to "install" this tool is to pull down the repository and run `make`. This will put a file called `federator` in the `bin` folder in the root of the project directory.

## Running

The federator tool has two modes, `link` and `creds`. These take the form of subcommands. The `link` subcommand gives you a link to sign into the AWS Console and the `creds` subcommand just gives you temporary credentials with the right commands to import it into a shell instance.

```shell
federator <subcommand> <options>
```

The arguments that each subcommand can take are listed below with the subcommands that they are compatible with.

| Parameter       | Subcommand      | Defaults                                    | Description                                                        |
| --------------- | --------------- | ------------------------------------------- | ------------------------------------------------------------------ |
| `--role-arn`    | `link`, `creds` | --                                          | The ARN of the role to assume                                      |
| `--issuer`      | `link`          | https://aws.amazon.com                      | The link where the user will be taken when the session has expired |
| `--destination` | `link`          | https://console.aws.amazon.com/console/home | The link that the user will be redirected to after login           |

## Development

### Installation

Once the code is cloned, you're going to need to run a preliminary `make build` so that all the dependencies come down and a build runs. Once that is done, an executable in `./bin` will be added.

### Running

The one requirement for running this tool is having installed and fully configured AWS CLI. Being able to run `aws sts get-caller-identity` successfully wil tell you that it is installed and configured properly.

The executable in `bin` has the right permissions, so running `./bin/federator` will let you run it in development.
