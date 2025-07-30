# jm (jump marker)

This is a lightweight CLI tool built in Go, originally created to solve a small
personal workflow problem and to deepen my understanding of the Go programming
language. While it started as a learning project, it's designed to be simple,
fast, and easy to use. Contributions and feedback are welcome!

## Usage

```shell
Jump marker (jm) lookup and management.

Usage:
  jm [command]

Available Commands:
  create      Create jump marker in current directory
  current     Show current jump marker
  help        Help about any command
  list        List jump markers
  remove      Remove jump marker from current directory
  version     Print version

Flags:
  -h, --help   help for jm
```

## Thanks

Big thanks to the developers of these CLI tools, as I every now and then had a
look into how they're doing things or what other tools and packages they're
using:

* [helm-freeze](https://github.com/Qovery/helm-freeze) - Freeze your charts in the wished versions.
* [ArgoCD](https://github.com/argoproj/argo-cd) - A declarative GitOps continuous delivery tool for Kubernetes.
* [Hugo](https://github.com/gohugoio/hugo) - A fast and flexible static site generator.
* [fzf](https://github.com/junegunn/fzf) - General-purpose command-line fuzzy finder.
