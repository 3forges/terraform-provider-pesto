# The CI/CD

In this folder, you will find all design and concept informations about the CI/CD system at work for this project.

## Global Design

### About Golang-CI

#### Run it

```bash
golangci-lint run
```

Will give in git bash for windows, or in powershell, an stdout starting by:

```bash
$ golangci-lint run
level=warning msg="[lintersdb] The name \"vet\" is deprecated. The linter has been renamed to: govet."
level=warning msg="The linter 'exportloopref' is deprecated (since v1.60.2) due to: Since Go1.22 (loopvar) this linter is no longer relevant. Replaced by copyloopvar."

```

And here is the example stdout of the first time I ran the `golangci-lint run` command:

![first golang-ci run](./images/golang-ci-run.ex1.PNG)

To solve th issues mentioned above, I changed the content of the `./.golang-ci.yml`, from:

```Yaml
# Visit https://golangci-lint.run/ for usage documentation
# and information on other useful linters
issues:
  max-per-linter: 0
  max-same-issues: 0

linters:
  disable-all: true
  enable:
    - durationcheck
    - errcheck
    - exportloopref
    - forcetypeassert
    - godot
    - gofmt
    - gosimple
    - ineffassign
    - makezero
    - misspell
    - nilerr
    - predeclared
    - staticcheck
    - tenv
    - unconvert
    - unparam
    - unused
    - vet
```

to:

```Yaml
# Visit https://golangci-lint.run/ for usage documentation
# and information on other useful linters
issues:
  max-per-linter: 0
  max-same-issues: 0

linters:
  disable-all: true
  enable:
    - durationcheck
    - errcheck
    # - exportloopref
    - copyloopvar
    - forcetypeassert
    - godot
    - gofmt
    - gosimple
    - ineffassign
    - makezero
    - misspell
    - nilerr
    - predeclared
    - staticcheck
    - tenv
    - unconvert
    - unparam
    - unused
    # - vet
    - govet
```

And I end up with a last new issue advertised by `golangci-lint`, about the go version I have installed on my computer:

![go version issue](./images/golang-ci-run.ex1.go.version.issue.PNG)

At this point, I have two options, to solve this last issue:

* Either I upgrade my golang version (to `1.22`), and I can use `copyloopvar` linter.
* Or, I downgrade the installed version of `golangci-lint` to `v1.60.1` (the most latest released version before `v1.60.2`), and and I use the `exportloopref` linter instead of the `copyloopvar` linter.

Choosing between the two options is very important: I have designed the code of my terrraform provider with go version `1.21`, and I want, eventually, that my provider supports being built with both go version `1.21` **_and_** go version `1.22`.

This is why I will, on my computer, I chose:

* to edit the `./.golang-ci.yml` to use the `exportloopref` linter instead of the `copyloopvar` linter,
* and downgrade version of Golang CI Lint to version `v1.60.1`, and  by running:

```bash
export GOLANGCI_LINT_VERSION='v1.60.1'
./tools/utils/installation/golangci-lint/windows/golangci-lint.sh

golangci-lint run

# Et voilà, no more warning
```

Et voilà, no more warning, only the 3 lint error messages are left, and I am not changing my  installed version of golang `go1.21`:

![Et voilà, no more versions warning](./images/golang-ci-run.ex1.all.issues.solved.for.go.1.21.PNG)

I then edited the 3 sources files to solve the clearly statedlint errors, and git commit and pushed.

Finally, when I will upgrade version of golang to `go1.22`, I will then:

* Upgrade `golangci-lint` to version `v1.62.0`
* edit the `./.golang-ci.yml` to use the `copyloopvar` linter instead of the `exportloopref` linter.

### About GoReleaser

## References

* <https://golangci-lint.run>
* <https://golangci-lint.run/welcome/install/>
* <https://golangci-lint.run/welcome/quick-start/>
