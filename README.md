# path-template

![help wanted](https://img.shields.io/badge/-help%20wanted-success)
[![godoc reference](https://godoc.org/github.com/quenbyako/path-template?status.svg)](https://godoc.org/github.com/quenbyako/path-template)
[![Go Report Card](https://goreportcard.com/badge/github.com/quenbyako/path-template)](https://goreportcard.com/report/github.com/quenbyako/path-template)
[![license MIT](https://img.shields.io/badge/license-MIT-green)](https://github.com/quenbyako/path-template/blob/master/README.md)
![version v1.0.0](https://img.shields.io/badge/version-v1.0.0-success)
![unstable](https://img.shields.io/badge/stability-stable-success)
[![chat telegram](https://img.shields.io/badge/chat-telegram-0088cc?logo=telegram&labelColor=ffffff)](https://bit.ly/2xlsVsQ)
<!--
code quality
golangci
contributors
go version
gitlab pipelines
-->

english [русский](https://github.com/quenbyako/path-template/blob/master/doc/ru_RU/README.md)

Simple and amazing. Templates for any folder!

<p align="center">
<img src="./docs/assets/logo.gif"/>
</p>

## Oh my god, why???

Because why not? Yes, this package is really small, but why not to have yet another universal package?

Path templates using every where: in [openapi](https://github.com/OAI/OpenAPI-Specification), [mux router](https://github.com/gorilla/mux), and many other projects! So, why not to create unified, production-ready and just damn simple implementation?

## Getting started

Am i required to say how to install? Okay, okay...

```sh
go get -u -v github.com/quenbyako/path-template
```

God bless a guy, who invent so simple package fetching.

## How to use

![demo](./docs/assets/demo.gif)

Code examples are [here](https://github.com/quenbyako/path-template/blob/master/examples)

### Make your first path checker

So, the main architecture of all package has two amazing functions: `FillTemplate` and `MatchPath`. First one is, well, fills your template with your variables, and second one, guess what? You goddamn right, matching input path with template, and, if matched, extracting variables. Note that `MatchPath` **does not** have wildcards, and if you try to match `/a/b/c` into template `/{long_path}/c`, `MatchPath` will throw error that path doesn't matched with template.

Here is a good example:

```go
package main

import "github.com/quenbyako/path-template"

func main() {
    res, err := tpl.FillTemplate(
        "/home/vasya/{project_path}/assets",
            map[string]string{
                "project_path":"go/src/gihub.com/xelaj/telegram_client",
        },
    )
    check(err)

    println(res)
    // will print: "/home/vasya/go/src/gihub.com/xelaj/telegram_client/assets"
}
```

## RFC6570-compliant?

Is it? Well, no, it's _partly_ compliant. But main components (variables, regex, and shiny mustache brackets `{}`) of this specification is implemented already. And, as always, PRs are really welcome! I want to improve `path-template` as 6570 implementation.

Wanna full compliancy? Use package like [this one](https://github.com/jtacoma/uritemplates), maybe it'll help you. But note, that this one was committed last time 5 (!!!) years ago.

## Contributing

Please read [contributing guide](https://github.com/quenbyako/path-template/blob/master/.github//CONTRIBUTING.md) if you want to help. And the help is very necessary!

## TODO

- [ ] idk, add some tests?

## Authors

* **Richard Cooper** <[rcooper.xelaj@protonmail.com](mailto:rcooper.xelaj@protonmail.com)>

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/quenbyako/path-template/blob/master/LICENSE.md) file for details