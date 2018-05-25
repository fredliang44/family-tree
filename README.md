# Family Tree Backend

[![Build Status](https://img.shields.io/travis/fredliang44/family-tree/master.svg?style=flat-square)](https://travis-ci.org/fredliang44/family-tree)
[![Codecov](https://img.shields.io/codecov/c/github/fredliang44/family-tree.svg?style=flat-square)](https://codecov.io/gh/fredliang44/family-tree)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/fredliang44/family-tree)
[![Go Report Card](https://goreportcard.com/badge/github.com/fredliang44/family-tree?style=flat-square)](https://goreportcard.com/report/github.com/fredliang44/family-tree)
## 1. Docs

Request|Document
---|---
Restful| <https://fmt.fredliang.cn>
GraphQL| <https://github.com/fredliang44/family-tree/blob/master/schema.graphql>

## 2. URL

Service|URL
---|---
Frontend | <https://fmt.hustunique.com>
Backend | <https://fmt.fredliang.cn>

## 3. Develope & Deploy

1. Edit `config.example.yml` and save it as `config.yml` for develope usage.
2. Then edit `config.example.yml` and save it as `config.deploy.yml` for deploy usage.
3. When deploying, mv `config.deploy.yml` to the same folder with your binary file you've built.

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