# `myply` Server
> This is a repository for `myply` server

<img width="1163" alt="Screen Shot 2022-07-31 at 7 59 11 PM" src="https://user-images.githubusercontent.com/37536298/182022976-061e3bbb-0b5d-4c00-8e18-ce683c22ff0e.png">


- spec
  - `Go` 1.18 or later, to use generics
  - [`wire`](https://github.com/google/wire) for Dependency injection
    - [wire mocking test example](https://github.com/google/wire/tree/main/internal/wire/testdata/ExampleWithMocks/foo)
  - [`Fiber`](https://github.com/gofiber?type=source) web framework
  - `Hexagonal` arch
  - [`Youtube data api`](https://developers.google.com/youtube/v3/getting-started)
  - [`Ginkgo`](https://github.com/onsi/ginkgo) for BDD
  - `GCP`


## local
- pre-commit

```console
$ cp pre-commit.example .git/hooks/pre-commit
$ chmod +x .git/hooks/pre-commit
```

- setup

```console
make setup
```

- run
```console
$ make local
GO111MODULE=on go run ./application/cmd/main.go

 ┌───────────────────────────────────────────────────┐ 
 │                   Fiber v2.34.1                   │ 
 │               http://127.0.0.1:3000               │ 
 │       (bound on host 0.0.0.0 and port 3000)       │ 
 │                                                   │ 
 │ Handlers ............. 2  Processes ........... 1 │ 
 │ Prefork ....... Disabled  PID ............. 11610 │ 
 └───────────────────────────────────────────────────┘ 

# 127.0.0.1:8080/swagger/index.html
```

- docker
```
$ make docker.fiber
```

## prod
- https://myply-server-rwwy3wj4sa-du.a.run.app
- https://myply-server-rwwy3wj4sa-du.a.run.app/swagger


## directory structure

```js
├── application // Interface layer and Application Services in hexagonal architecture
│   ├── cmd // command line interface
│   ├── controller // http controller
│   ├── middleware // http middleware, it can wrap errors or set request uuid or jwt authorization
│   └── routes // http router
├── domain // domain layer
│   ├── entity
│   ├── repository // same as Port in hexagonal architecture
│   ├── service // domain service layer
│   └── vo // value object
├── go.mod
├── infrastructure // infrastructure layer in hexagonal
│   ├── clients // external APIs (E.g. google client, firebase client ..)
│   ├── configs // configuration of fiber app 
│   ├── logger
│   └── persistence // Impl of /domain/repository interface, database persistence layer(E.g. mysql, postgreSQL, mongo, redis ..)
└── pre-commit.example // go fmt before commit
```

- `/application`: https://github.com/Sairyss/domain-driven-hexagon#application-layer
  - ** Note that `cmd` does not represent CQRS's Command. (It's different, `cmd` just represents command line interface)
- `/doamin`: https://github.com/Sairyss/domain-driven-hexagon#domain-layer
- `/infrastructure`: https://github.com/Sairyss/domain-driven-hexagon#infrastructure-layer

## uml
![schema](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/Nexters/myply/main/docs/schema.puml)

## refs
- unit test guide: [fiber/app_test.go](https://github.com/gofiber/fiber/blob/master/app_test.go)
- [fiber recipes](https://github.com/gofiber/recipes)
- related to youtube api
  - [Unofficial youtube music library](https://ytmusicapi.readthedocs.io/en/latest/#)
  - [Electron wrapper around YouTube Music](https://github.com/th-ch/youtube-music)
    - It also supports music streaming
