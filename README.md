# go-client




## Initial Setup

This section is intended to help developers and contributors get a working copy of
`go-client` on their end

<details>
<summary>
    1. Clone this repository
</summary><br>

```sh
git clone https://github.com/ripple.mq/go-client
cd go-client
```
</details>

<details>
<summary>
    2. Install `golangci-lint`
</summary><br>

Install `golangci-lint` from the [official website][golangci-install] for your OS
</details>
<br>


## Local Development

This section will guide you to setup a fully-functional local copy of `go-client`
on your end and teach you how to use it! Make sure you have installed
[golangci-lint][golangci-install] before following this section!

**Note:** This section relies on the usage of [Makefile][makefile-official]. If you
can't (or don't) use Makefile, you can follow along by running the internal commands
from [`go-client's` Makefile][makefile-file] (most of which are
OS-independent)!

### Installing dependencies

To install all dependencies associated with `go-client`, run the
command

```sh
make install
```



### Using Code Formatters

Code formatters format your code to match pre-decided conventions. To run automated code
formatters, use the Makefile command

```sh
make codestyle
```

### Using Code Linters

Linters are tools that analyze source code for possible errors. This includes typos,
code formatting, syntax errors, calls to deprecated functions, potential security
vulnerabilities, and more!

To run pre-configured linters, use the command

```sh
make lint
```

### Running Tests

Tests in `go-client` are classified as *fast* and *slow* - depending
on how quick they are to execute.

To selectively run tests from either test group, use the Makefile command

```sh
make fast-test

OR

make slow-test
```

Alternatively, to run the complete test-suite -- i.e. *fast* and *slow* tests at one
go, use the command

```sh
make test
```

### Running the Test-Suite

The *`test-suite`* is simply a wrapper to run linters, stylecheckers and **all** tests
at once!

To run the test-suite, use the command

```sh
make test-suite
```

In simpler terms, running the test-suite is a combination of running [linters](#using-code-linters)
and [all tests](#running-tests) one after the other!
<br>


## Additional Resources

### Makefile help

<details>
<summary>
    Tap for a list of Makefile commands
</summary><br>

|     Command    	|                               Description                               	| Prerequisites 	|
|:--------------:	|:-----------------------------------------------------------------------:	|:-------------:	|
|     `help`     	| Generate help dialog listing all Makefile commands with description     	|       NA      	|
|    `install`   	| Fetch project dependencies                                              	|       NA      	|
|   `codestyle`  	| Run code-formatters                                                     	| golangci-lint 	|
|     `lint`     	| Check codestyle and run linters                                         	| golangci-lint 	|
|     `test`     	| Run **all** tests                                                       	|       NA      	|
|  `fast-tests`  	| Selectively run *fast* tests                                            	|       NA      	|
|  `slow-tests`  	| Selectively run *slow* tests                                            	|       NA      	|
|  `test-suite`  	| Check codestyle, run linters and **all** tests                        	| golangci-lint 	|
|      `run`     	| Run *go-client*                           	|       NA      	|
|  `docker-gen`  	| Create production docker image for *go-client* 	|     docker    	|
|  `docker-debug`  	| Create debug-friendly docker image for *go-client* 	|     docker    	|
| `clean-docker` 	| Remove docker image generated by `docker-gen`                           	|     docker    	|

<br>
</details>

Optionally, to see a list of all Makefile commands, and a short description of what they
do, you can simply run

```sh
make
```

Which is equivalent to;

```sh
make help
```

Both of which will list out all Makefile commands available, and a short description
of what they do!

### Generating Binaries

To generate binaries for multiple OS/architectures, simply run

```sh
bash build-script.sh
```

The command will generate binaries for Linux, Windows and Mac targetting multiple
architectures at once! The binaries, once generated will be stored in the `bin`
directory inside the project directory.

The binaries generated will be named in the format

```text
go-client_<version>_<target-os>_<architecture>.<extension>
```

The `<extension>` is optional. By default, `version` is an empty string. A custom
version can be passed as an argument while running the script. As an example;

```sh
bash build-script.sh v1.2.1
```

An example of the files generated by the previous command will be;

```sh
go-client_v1.2.1_windows_x86_64.exe
```



### Using Docker

To run `go-client` in a docker container, read the instructions in
[docker section](./docker).

### Running `go-client`

To run go-client, use the command

```sh
make run
```

Additionally, you can pass any additional command-line arguments (if needed) as the
argument "`q`". For example;

```sh
make run q="--help"

OR

make run q="--version"
```
<br>



<br>


</div>

[makefile-file]: ./Makefile
[project-license]: ./LICENSE
[github-actions]: ../../actions
[github-releases]: ../../releases
[precommit-config]: ./.pre-commit-config.yaml
[gomod-file]: ../main/go.mod
[github-actions-tests]: ../../actions/workflows/tests.yml
[dependabot-pulls]: ../../pulls?utf8=%E2%9C%93&q=is%3Apr%20author%3Aapp%2Fdependabot

[semver-link]: https://semver.org
[pre-commit]: https://pre-commit.com
[github-repo]: https://github.com/new
[gitlab-repo]: https://gitlab.com/new
[dependabot-link]: https://dependabot.com
[githooks]: https://git-scm.com/docs/githooks
[python-install]: https://www.python.org/downloads
[release-drafter-config]: ./.github/release-drafter.yml
[makefile-official]: https://www.gnu.org/software/make
[pip-install]: https://pip.pypa.io/en/stable/installation
[codecov-docs]: https://docs.codecov.com/docs#basic-usage
[go-template-link]: https://github.com/notsatan/go-template
[golangci-install]: https://golangci-lint.run/usage/install
[cookiecutter-link]: https://github.com/cookiecutter/cookiecutter
[release-drafter]: https://github.com/marketplace/actions/release-drafter
[creating-secrets]: https://docs.github.com/en/actions/security-guides/encrypted-secrets#creating-encrypted-secrets-for-a-repository