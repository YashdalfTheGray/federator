我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

![Build](https://github.com/YashdalfTheGray/federator/workflows/Build/badge.svg)

# federator

A utility to federate into an AWS account using AWS Security Token Service and then get a link to go directly into the console.

## Usage

### Installation

The easiest way to install this tool is to grab the built binary from the releases, put it somewhere that is in your path, run a quick `chmod +x <path_to_federator>` and you'll be good to go.

The other way to get this tool is to pull down the repository and run `make`. This will put a file called `federator` in the `bin` folder in the root of the project directory.

## Running

The federator tool has two modes, `link` and `creds`. These take the form of subcommands. The `link` subcommand gives you a link to sign into the AWS Console and the `creds` subcommand just gives you temporary credentials with the right commands to import it into a shell instance. There is also a JSON output mode so that it can be machine readable or piped into something like [`jq`](https://stedolan.github.io/jq/).

```shell
federator <subcommand> <options>
```

The arguments that each subcommand can take are listed below with the subcommands that they are compatible with.

| Parameter       | Subcommand      | Defaults                                    | Description                                                        |
| --------------- | --------------- | ------------------------------------------- | ------------------------------------------------------------------ |
| `--role-arn`    | `link`, `creds` | --                                          | The ARN of the role to assume                                      |
| `--external-id` | `link`, `creds` | ""                                          | The external ID, if necessary, to be provided                      |
| `--json`        | `link`, `creds` | false                                       | Whether to print out the results in JSON or plain text             |
| `--issuer`      | `link`          | https://aws.amazon.com                      | The link where the user will be taken when the session has expired |
| `--destination` | `link`          | https://console.aws.amazon.com/console/home | The link that the user will be redirected to after login           |

## Development

### Installation

Once the code is cloned, you're going to need to run a preliminary `make build` so that all the dependencies come down and a build runs. Once that is done, an executable in `./bin` will be added.

### Running

The one requirement for running this tool is having installed and fully configured AWS CLI. Being able to run `aws sts get-caller-identity` successfully wil tell you that it is installed and configured properly.

The executable in `bin` has the right permissions, so running `./bin/federator` will let you run it in development.

### Contributing

Contributions to this project are welcome and encouraged. This project uses the standard golang tooling, `go-fmt`, and go modules. There is a `make` target for tests and one for generating coverage. Make sure to run these commands before creating a pull request.
