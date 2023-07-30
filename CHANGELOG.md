# Change Log

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

## [v2.3.0](https://github.com/YashdalfTheGray/federator/tree/v2.3.0) (2023-07-29)

### Added

- Added `credentials_process` support to the `creds` subcommand via the `--awscli` switch (credit [@petderek](https://github.com/petderek) [#126](https://github.com/YashdalfTheGray/federator/pull/126) [#131](https://github.com/YashdalfTheGray/federator/issues/131))

## [v2.2.0](https://github.com/YashdalfTheGray/federator/tree/v2.2.0) (2020-09-25)

### Added

- `trust-policy` subcommand that prints the correct trust policy for an account or an IAM resource

## [v2.1.0](https://github.com/YashdalfTheGray/federator/tree/v2.1.0) (2020-09-08)

### Added

- `--region` flag to target a specific STS endpoint, defaults to the CLI configured region

## [v2.0.0](https://github.com/YashdalfTheGray/federator/tree/v2.0.0) (2020-09-04)

### Modified

- started using the AWS SDK version 2

## [v1.2.0](https://github.com/YashdalfTheGray/federator/tree/v1.2.0) (2020-08-31)

### Added

- `federator` has the ability to pass in an External ID to the STS call using the `--external-id` flag

## [v1.1.0](https://github.com/YashdalfTheGray/federator/tree/v1.1.0) (2020-03-07)

### Added

- The ability for `federator` to output JSON using the `--json` flag

## [v1.0.0](https://github.com/YashdalfTheGray/federator/tree/v1.0.0) (2020-02-28)

First release of the `federator`
