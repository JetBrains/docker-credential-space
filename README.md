# docker-credential-space

## Introduction

`docker-credential-space` is a Docker credential helper for [JetBrains Space Container Registry](https://www.jetbrains.com/help/space/container-registry.html). It allows for **v18.03+ Docker clients** to easily make authenticated requests to JetBrains Space repositories (*.registry.jetbrains.space, etc.).

The helper implements the [Docker Credential Store](https://docs.docker.com/engine/reference/commandline/login/#/credentials-store) API, but enables more advanced authentication schemes for JetBrains Space users.

## JetBrains Space Credentials

The helper searches for credentials in the `JB_SPACE_CLIENT_ID` and `JB_SPACE_CLIENT_SECRET` environment variables.

**Examples:**

To verify that credentials are being returned for a given registry, e.g. for `https://org.registry.jetbrains.space`:

```shell
echo "https://org.registry.jetbrains.space" | docker-credential-space get
```

### Building from Source

The program in this repository is written with the Go programming language and built with `make`. These instructions assume that [**Go 1.11+**](https://golang.org/) and `make` are installed on a \*nix system.

You can download the source code, compile the binary, and put it in your `$GOPATH` with `go get`.

```shell
go get -u github.com/JetBrains/docker-credential-space
```

If `$GOPATH/bin` is in your system `$PATH`, this will also automatically install the compiled binary. You can confirm using `which docker-credential-space` and continue to the [section on Configuration and Usage](#configuration-and-usage).

Alternatively, you can use `make` to build the program. The executable will be output to the `bin` directory inside the repository.

```shell
cd $GOPATH/src/github.com/JetBrains/docker-credential-space
make
```

Then, you can put that binary in your `$PATH` to make it visible to `docker`. For example, if `/usr/bin` is present in your system path:

```shell
sudo mv ./bin/docker-credential-space /usr/bin/docker-credential-space
```

## Configuration and Usage

Add a `credHelpers` entry in the Docker config file (usually `~/.docker/config.json` on OSX and Linux, `%USERPROFILE%\.docker\config.json` on Windows) for each JetBrains Space registry that you care about. The key should be the domain of the registry (**without** the "https://") and the value should be the suffix of the credential helper binary (everything after "docker-credential-").

	e.g. for `docker-credential-space`:

  <pre>
    {
      "credHelpers": {
            "coolregistry.com": ... ,
            <b>"org.registry.jetbrains.space": "space",
            ...</b>
      }
    }
  </pre>

## License

Apache 2.0. See [LICENSE](LICENSE) for more information.
