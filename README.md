![Build](https://github.com/YashdalfTheGray/federator/workflows/Build/badge.svg)

# federator

A utility to federate into an AWS account using AWS Security Token Service and then get a link to go directly into the console.

## Usage

### Installation

The easiest way to install this tool is to grab the built binary from the releases, put it somewhere that is in your path, run a quick `chmod +x <path_to_federator>` and you'll be good to go.

The other way to get this tool is to pull down the repository and run `make`. This will put a file called `federator` in the `bin` folder in the root of the project directory. Since we use the AWS SDK v2 for Go, we are required to use Go 1.13 or higher. Additionally, since this package uses Go modules, it also requires Go 1.13.

## Running

The federator tool has two modes, `link` and `creds`. These take the form of subcommands. The `link` subcommand gives you a link to sign into the AWS Console and the `creds` subcommand just gives you temporary credentials with the right commands to import it into a shell instance. There is also a JSON output mode so that it can be machine readable or piped into something like [`jq`](https://stedolan.github.io/jq/).

```shell
federator <subcommand> <options>
```

### Command line flags

The arguments that each subcommand can take are listed below with the subcommands that they are compatible with.

#### Flags for the `link` subcommand

| Parameter       | Defaults                                    | Description                                                                                     |
| --------------- | ------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `--role-arn`    | --                                          | The ARN of the role to assume                                                                   |
| `--external-id` | "" (empty string)                           | The external ID, if necessary, to be provided, it will be added to the trust policy if provided |
| `--region`      | from the CLI config                         | The region to make the STS call against                                                         |
| `--json`        | `false`                                     | Whether to print out the results in JSON or plain text                                          |
| `--issuer`      | https://aws.amazon.com                      | The link where the user will be taken when the session has expired                              |
| `--destination` | https://console.aws.amazon.com/console/home | The link that the user will be redirected to after login                                        |

#### Flags for the `creds` subcommand

| Parameter       | Defaults            | Description                                                                                     |
| --------------- | ------------------- | ----------------------------------------------------------------------------------------------- |
| `--role-arn`    | --                  | The ARN of the role to assume                                                                   |
| `--external-id` | "" (empty string)   | The external ID, if necessary, to be provided, it will be added to the trust policy if provided |
| `--region`      | from the CLI config | The region to make the STS call against                                                         |
| `--json`        | `false`             | Whether to print out the results in JSON or plain text                                          |

#### Flags for the `trust-policy` subcommand

| Parameter       | Defaults          | Description                                                                                                          |
| --------------- | ----------------- | -------------------------------------------------------------------------------------------------------------------- |
| `--arn`         | --                | The IAM resource ARN to trust in the policy, either this or the account ID must be specified for the command to work |
| `--account-id`  | --                | The AWS account ID to trust in the policy, either this or the resource ARN must be specified for the command to work |
| `--external-id` | "" (empty string) | The external ID, if necessary, to be provided, it will be added to the trust policy if provided                      |
| `--json`        | `false`           | Whether to print out the results in JSON or plain text                                                               |

### Examples

These examples assume that you have `federator` in your path already.

```sh
# get creds for a role without an external ID requirement
federator creds --role-arn arn:aws:iam::000000000000:role/someRole

# get a console link and provide an external ID
federator link --role-arn arn:aws:iam::000000000000:role/someRole --external-id "some external id"

# Use a regional STS endpoint different from the one configured with the CLI
federator creds --role-arn arn:aws:iam::000000000000:role/someRole --region us-east-1

# Change the output to json
federator link --role-arn arn:aws:iam::000000000000:role/someRole --json

# Output a trust policy for an account ID
federator trust-policy --account-id 000000000000

# Output a trust policy for an IAM user with an external ID provided
federator trust-policy --arn arn:aws:iam::000000000000:user/myUser --external-id "some external id"
```

## Development

### Installation

Once the code is cloned, you're going to need to run a preliminary `make build` so that all the dependencies come down and a build runs. Once that is done, an executable in `./bin` will be added.

### Running

The one requirement for running this tool is having installed and fully configured AWS CLI. Being able to run `aws sts get-caller-identity` successfully wil tell you that it is installed and configured properly.

The executable in `bin` has the right permissions, so running `./bin/federator` will let you run it in development.

### Contributing

Contributions to this project are welcome and encouraged. This project uses the standard golang tooling, `go-fmt`, and go modules. There is a `make` target for tests and one for generating coverage. Make sure to run these commands before creating a pull request.
