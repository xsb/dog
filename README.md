# Dog

Dog is a command line application that executes tasks. It works in a similar way as GNU Make or ruby's Rake but it is a more generic task runner, not a build tool. It can be used as a layer on top of your Makefile or your shell scripts. Dog's default script syntax is `sh` but it also supports BASH, Python or Ruby so you can write your tasks in any language.

Roadmap for v0.1:

- [x] List and run tasks
- [x] Support multiple languages
- [x] Allow multiple Dogfiles per directory
- [x] Show status code after running a task
- [ ] Pass environment variables to tasks
- [ ] Pre-hooks and post-hooks

## What is a Dogfile?

Dogfile is a specification that uses YAML to describe the tasks related to a project. We think that the Spec will be finished (no further breaking changes) by the v1.0 version of Dog.

- Read the [Dogfile Spec](https://github.com/xsb/dog/blob/master/DOGFILE_SPEC.md)

This is Dog's own Dogfile.yml:

```yml
- task: build
  description: Build dog binary
  run: |
    [ -d bin ] || mkdir bin
    go get -u ./...
    go build -o bin/dog

- task: clean
  description: Clean compiled binaries
  run: rm -rf bin

- task: run-test-dogfiles
  description: Run all Tasks in testdata Dogfiles
  run: ./scripts/test-dogfiles.sh
```

## Contributing

At this moment we are focused on implementing the basics that will allow us to publish v0.1. This project is organized using GitHub [Issues](https://github.com/xsb/dog/issues) and [Pull Requests](https://github.com/xsb/dog/pulls).

If you want to help, take a look at:

- Dogfile Spec [discussion](https://github.com/xsb/dog/issues?q=is%3Aissue+is%3Aopen+label%3A%22dogfile+spec%22)
- Open [bugs](https://github.com/xsb/dog/issues?q=is%3Aissue+is%3Aopen+label%3Abug)
- Lacking features for [Milestone v0.1](https://github.com/xsb/dog/milestones/v0.1)
- Lacking features for [Milestone v0.2](https://github.com/xsb/dog/milestones/v0.2)
