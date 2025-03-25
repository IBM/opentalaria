# OpenTalaria

## Overview
OpenTalaria is a lightweight message broker that implements the Kafka protocol and persists messages in different storage systems like PostgreSQL, SQLite and more.

## Supported APIs
The project is still a work in progress, please see the [list](docs/apis.md) of currently supported APIs which is updated when new APIs are implemented.

## Running the project
The project is still in early stages of development, so there are no scheduled releases. To build from source you will need the Go compiler v1.21 or later and Make. To build run the following command `make build`. This will generate the binary `bin/opentalaria`. Note that debug information is stripped from the resulting binary. If you wish to debug the project, you can run it with `make run`. 

To run the binary, call it with an optional flag `-c` to specify a configuration file `bin/opentalaria -c config.yaml`. If no configuration file is provided, the broker will look for a file named `config.yaml` in the current directory. If the file does not exist the broker falls back on environment variables. Please see the [documentation](docs/configuration.md) for more info.

## Contributing

Contributions to this project are released to the public under the project's [opensource license](LICENSE.md).
By contributing to this project you agree to the [Developer Certificate of Origin](https://developercertificate.org/) (DCO).
The DCO was created by the Linux Kernel community and is a simple statement that you, as a contributor, wrote or otherwise have the legal right to contribute those changes.

Contributors must _sign-off_ that they adhere to these requirements by adding a `Signed-off-by` line to all commit messages with an email address that matches the commit author:

```
feat: this is my commit message

Signed-off-by: Random J Developer <random@developer.example.org>
```

Git even has a `-s` command line option to append this automatically to your
commit message:

```
$ git commit -s -m 'This is my commit message'
```

*Note* if you want to automatically sign-off all commits, follow the steps in [this](https://stackoverflow.com/a/46536244) StackOverflow answer.
