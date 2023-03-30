# Go Api2d
[![GoDoc](http://img.shields.io/badge/GoDoc-Reference-blue.svg)](https://godoc.org/github.com/leizhenpeng/go-api2d)
[![Go Report Card](https://goreportcard.com/badge/github.com/leizhenpeng/go-api2d)](https://goreportcard.com/report/github.com/leizhenpeng/go-api2d)
[![codecov](https://codecov.io/gh/leizhenpeng/go-api2d/branch/master/graph/badge.svg?token=bCbIfHLIsW)](https://codecov.io/gh/leizhenpeng/go-api2d)

> **Note**: This library is in beta and may change in backwards-incompatible ways.

本仓库基于 [openai/go-openai](https://github.com/sashabaranov/go-openai),因此兼容[OpenAI API](https://platform.openai.com/)所有接口

同时做出下面优化
- 我们进行了针对[api2d](https://api2d.com/r/186763)自定义参数的优化，使其更加易于使用(如：内容审核服务)
- 增加了流对话接口的token计数服务，方便您在服务端监控token消耗情况



This library provides Go clients for . We support:

* ChatGPT
* GPT-3, GPT-4

Installation:
```
go get github.com/leizhenpeng/go-api2d
```

