# norm

An orm library support nsql for Golang.

[![go report card](https://goreportcard.com/badge/github.com/zhihu/norm "go report card")](https://goreportcard.com/report/github.com/zhihu/norm)
[![test status](https://github.com/zhihu/norm/workflows/tests/badge.svg?branch=master "test status")](https://github.com/zhihu/norm/actions)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/norm.io/norm?tab=doc)

## Overview

* Build insert nsql by struct/map (Support vertex, edge).
* Parse nebula execute result to struct/map.
* Easy use.
* Easy mock for unit test.

**roadmap**
1. Session pool. For detail please see [dialector](/docs/dialector.adoc)
    * Plan: before 2021/07/20
2. Support more types in insert/execute function.
    * Types: time.Time
3. Support batch insert, query list.
4. Chainable api. For detail please see [chainable api](/docs/chainable_api.adoc)

**Maybe Support**
- [ ] Statistic hooks. Insert/Query count and latency.
- [ ] Fix fields Order when build insert nsql. (now norm use map store keys, and in go range map is out-of-order.)

**Need improve**
- [ ] Betchmark.
- [ ] Unit tests.
- [ ] Documents.

## Getting Started
install: `go get github.com/zhihu/norm`

use example: please go [use example](/examples/toddle/main.go)

## Contributing guidelines
* [code of conduct](/CODE_OF_CONDUCT.md)
* [行为规范 中文版](/CODE_OF_CONDUCT_CN.md)

## License

© Zhihu, 2021~time.Now

Released under the [MIT License](/LICENSE)

_copy and paste from gorm_