<div align="center">

[![License](https://img.shields.io/badge/License-MIT-blue)](#license)
[![Go Report Card](https://goreportcard.com/badge/github.com/daytonaio/daytona-provider-hetzner)](https://goreportcard.com/report/github.com/daytonaio/daytona-provider-hetzner)
[![Issues - daytona](https://img.shields.io/github/issues/daytonaio/daytona-hetzner-provider)](https://github.com/daytonaio/daytona-provider-hetzner/issues)
![GitHub Release](https://img.shields.io/github/v/release/daytonaio/daytona-hetzner-provider)

</div>

<h1 align="center">Daytona Hetzner Provider</h1>
<div align="center">
This repository is the home of the <a href="https://github.com/daytonaio/daytona">Daytona</a> Hetzner Provider.
</div>
</br>

<p align="center">
  <a href="https://github.com/daytonaio/daytona-provider-hetzner/issues/new?assignees=&labels=bug&projects=&template=bug_report.md&title=%F0%9F%90%9B+Bug+Report%3A+">Report Bug</a>
    ·
  <a href="https://github.com/daytonaio/daytona-provider-hetzner/issues/new?assignees=&labels=enhancement&projects=&template=feature_request.md&title=%F0%9F%9A%80+Feature%3A+">Request Feature</a>
    ·
  <a href="https://go.daytona.io/slack">Join Our Slack</a>
    ·
  <a href="https://x.com/Daytonaio">X</a>
</p>

> [!TIP]
> Write a description of your Provider here.

## Target Options

| Property        | Type     | Optional | DefaultValue            | InputMasked | DisabledPredicate |
| --------------- | -------- | -------- | ----------------------- | ----------- | ----------------- |
| Required String | String   | false    | default-required-string | false       |                   |
| Optional String | String   | true     |                         | true        |                   |
| Optional Int    | Int      | true     |                         | false       |                   |
| FilePath        | FilePath | true     | ~/.ssh                  | false       | ^default-target$  |

### Default Targets

#### Local

| Property        | Value                   |
| --------------- | ----------------------- |
| Required String | default-required-string |

## Code of Conduct

This project has adapted the Code of Conduct from the [Contributor Covenant](https://www.contributor-covenant.org/). For more information see the [Code of Conduct](CODE_OF_CONDUCT.md) or contact [codeofconduct@daytona.io.](mailto:codeofconduct@daytona.io) with any additional questions or comments.

## Contributing

The Daytona Docker Provider is Open Source under the [MIT License](LICENSE). If you would like to contribute to the software, you must:

1. Read the Developer Certificate of Origin Version 1.1 (https://developercertificate.org/)
2. Sign all commits to the Daytona Docker Provider project.

This ensures that users, distributors, and other contributors can rely on all the software related to Daytona being contributed under the terms of the [License](LICENSE). No contributions will be accepted without following this process.

Afterwards, navigate to the [contributing guide](CONTRIBUTING.md) to get started.

## Questions

For more information on how to use and develop Daytona, talk to us on
[Slack](https://go.daytona.io/slack).
