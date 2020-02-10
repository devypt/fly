<p align="center">
  <a href="https://goreportcard.com/report/github.com/devypt/fly"><img src="https://goreportcard.com/badge/github.com/devypt/fly?style=flat-square" alt="GoReport"></a>
  <a href="http://godoc.org/github.com/devypt/fly"><img src="http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square" alt="GoDoc"></a>
  <a href="https://raw.githubusercontent.com/devypt/fly/master/LICENSE"><img src="https://img.shields.io/aur/license/yaourt.svg?style=flat-square" alt="License"></a>
</p>
<hr>
<h2 align="center">Golang live reload and task runner </h2>
<h3 align="center"> This is a custom go module based on <a href="https://github.com/oxequa/realize"> realize </a> </h3>
<hr>




### - ⭐️ [Top Features](#top-features)
### - 💃🏻 [Get started](#get-started)
### - 📄 [Config sample](#config-sample)
### - 📚 [Commands List](#commands-list)
### - 🛠 [Suggestions](#suggestions)

## Top Features

- High performance Live Reload.
- Manage multiple projects at the same time.
- Watch by custom extensions and paths.
- All Go commands supported.
- Switch between different Go builds.
- Custom env variables for project.
- Execute custom commands before and after a file changes or globally.
- Export logs and errors to an external file.
- Any suggestion? [Suggest an amazing feature! 🕺🏻](https://github.com/devypt/fly/issues/new)


## Quickstart
```
go install github.com/devypt/fly
```

## Commands List

### Run Command
From **project/projects** root execute:

    $ fly start


It will create a **.fly.yaml** file if doesn't already exist, add the working directory as project and run your workflow.

***start*** command supports the following custom parameters:

    --name="name"               -> Run by name on existing configuration
    --path="fly/server"         -> Custom Path (if not specified takes the working directory name)
    --generate                  -> Enable go generate
    --fmt                       -> Enable go fmt
    --test                      -> Enable go test
    --vet                       -> Enable go vet
    --install                   -> Enable go install
    --build                     -> Enable go build
    --run                       -> Enable go run
    --open                      -> Open web ui in default browser
    --no-config                 -> Ignore an existing config / skip the creation of a new one

Some examples:

    $ fly start
    $ fly start --path="mypath"
    $ fly start --name="fly" --build
    $ fly start --path="fly" --run --no-config
    $ fly start --install --test --fmt --no-config
    $ fly start --path="/Users/username/go/src/github.com/devypt/fly-examples/coin/"

If you want, you can specify additional arguments for your project:

	✅ $ fly start --path="/print/printer" --run yourParams --yourFlags // right
    ❌ $ fly start yourParams --yourFlags --path="/print/printer" --run // wrong

⚠️ The additional arguments **must go after** the params:
<br>
💡 The ***start*** command can be used with a project from its working directory without make a config file (*--no-config*).


## Color reference
💙 BLUE: Outputs of the project.<br>
💔 RED: Errors.<br>
💜 PURPLE: Times or changed files.<br>
💚 GREEN: Successfully completed action.<br>


## Config sample

*** there is no more a .fly dir, but only a .fly.yaml file ***

For more examples check: [Fly Examples](https://github.com/devypt/fly-examples)

    settings:
        resources:                  // files names
            outputs: outputs.log
            logs: logs.log
            errors: errors.log
    schema:
    - name: coin
      path: coin              // project path
      env:            // env variables available at startup
            test: test
            myvar: value
      commands:               // go commands supported
        vet:
            status: true
        fmt:
            status: true
            args:
            - -s
            - -w
        test:
            status: true
            method: gb test    // support different build tools
        generate:
            status: true
        install:
            status: true
        build:
            status: false
            method: gb build    // support differents build tool
            args:               // additional params for the command
            - -race
        run:
            status: true
      args:                     // arguments to pass at the project
      - --myarg
      watcher:
          paths:                 // watched paths
          - /
          ignore_paths:          // ignored paths
          - vendor
          extensions:                  // watched extensions
          - go
          - html
          scripts:
          - type: before
            command: echo before global
            global: true
            output: true
          - type: before
            command: echo before change
            output: true
          - type: after
            command: echo after change
            output: true
          - type: after
            command: echo after global
            global: true
            output: true
          errorOutputPattern: mypattern   //custom error pattern

## Suggestions
⭐️ Suggest a new [Feature](https://github.com/devypt/fly/issues/new)
