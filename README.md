# `myply` Server
> This is a repository for `myply` server

<img width="1728" alt="myply" src="https://user-images.githubusercontent.com/37536298/177248803-64893e61-827b-4a54-81f0-fd8ee664fe9f.png">


- `Go` 1.18 or later
  - To use generics
- [`wire`](https://github.com/google/wire) for Dependency injection
  - [wire mocking test example](https://github.com/google/wire/tree/main/internal/wire/testdata/ExampleWithMocks/foo)
- [`Fiber`](https://github.com/gofiber?type=source) web framework
- `Hexagonal` arch
- [`Youtube data api`](https://developers.google.com/youtube/v3/getting-started)
- [`Ginkgo`](https://github.com/onsi/ginkgo) for BDD
- `GCP`

## refs
- unit test guide: [fiber/app_test.go](https://github.com/gofiber/fiber/blob/master/app_test.go)
- [fiber recipes](https://github.com/gofiber/recipes)
- related to youtube api
  - [Unofficial youtube music library](https://ytmusicapi.readthedocs.io/en/latest/#)
  - [Electron wrapper around YouTube Music](https://github.com/th-ch/youtube-music)
    - It also supports music streaming
