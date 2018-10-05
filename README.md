# Family Tree Backend

[![Build Status](https://img.shields.io/travis/fredliang44/family-tree/master.svg?style=flat-square)](https://travis-ci.org/fredliang44/family-tree)
[![Codecov](https://img.shields.io/codecov/c/github/fredliang44/family-tree.svg?style=flat-square)](https://codecov.io/gh/fredliang44/family-tree)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/fredliang44/family-tree)
[![Go Report Card](https://goreportcard.com/badge/github.com/fredliang44/family-tree?style=flat-square)](https://goreportcard.com/report/github.com/fredliang44/family-tree)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Ffredliang44%2Ffamily-tree.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Ffredliang44%2Ffamily-tree?ref=badge_shield)
## 1. Docs

Request|Document
---|---
Restful| <https://fmt.fredliang.cn>
GraphQL| [schema.graphql](schema.graphql)

## 2. URL

Service|URL
---|---
Frontend | <https://fmt.hustunique.com>
Backend | <https://fmt.fredliang.cn>
GraphQL | <https://fmt.fredliang.cn/graphql>

## 3. Develope & Deploy

1. Please make sure your `go version` >= 1.11
2. Edit `config.example.yml` and save it as `config.yml` for develope usage.
3. Then edit `config.example.yml` and save it as `config.deploy.yml` for deploy usage.
4. When deploying, mv `config.deploy.yml` to the same folder with your binary file you've built.

## 4. Project Structure
```shell
├── Dockerfile           // docker support
├── Gopkg.lock
├── Gopkg.toml
├── Makefile
├── README.md
├── config.example.yml   // example config
├── db
│   ├── mongo.go
│   └── redis.go
├── family-tree.go       // main func and router
├── graphql
│   ├── base.go
│   ├── mutation.go
│   ├── query.go
│   ├── resolvers.go
│   ├── schema.go
│   ├── types
│   │   ├── company.go
│   │   ├── group.go
│   │   ├── project.go
│   │   └── user.go
│   ├── types.go
│   └── utils.go
├── handler              // handler json requests
│   ├── base.go
│   ├── base_test.go
│   ├── db.go
│   ├── db_test.go
│   ├── register.go
│   ├── register_test.go
│   ├── reset_password.go
│   └── reset_password_test.go
├── middleware
│   ├── auth.go
│   └── cors.go
├── schema.graphql       // graphql doc
└── utils
    ├── config.go        // load config
    ├── recovery.go
    └── sms.go           // handle message request with Dayu of Aliyun and Tencent Cloud SMS Service
```

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Ffredliang44%2Ffamily-tree.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Ffredliang44%2Ffamily-tree?ref=badge_large)