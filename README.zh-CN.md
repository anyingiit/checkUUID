[English](README.md) · **简体中文**

> 英文版是规范版本。本页与 [README.md](README.md) 不一致时，以英文版为准。

<!-- translation-of: README.md sha256:87d46f864a981bfb -->

<!-- Source: Best-README-Template BLANK_README (Unlicense) — https://github.com/othneildrew/Best-README-Template -->
<a id="readme-top"></a>

# checkUUID

一个 Go 命令行程序，一次运行生成两亿（200,000,000）个 UUID，并报告其中是否出现重复。

[![CI](https://github.com/anyingiit/checkUUID/actions/workflows/ci.yml/badge.svg)](https://github.com/anyingiit/checkUUID/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/anyingiit/checkUUID)](LICENSE)

[报告问题](https://github.com/anyingiit/checkUUID/issues/new?template=bug_report.yml) · [提出需求](https://github.com/anyingiit/checkUUID/issues/new?template=feature_request.yml)

<details>
  <summary>目录</summary>
  <ol>
    <li><a href="#about-the-project">关于本项目</a></li>
    <li><a href="#getting-started">开始使用</a></li>
    <li><a href="#usage">用法</a></li>
    <li><a href="#contributing">参与贡献</a></li>
    <li><a href="#license">许可证</a></li>
    <li><a href="#contact">联系方式</a></li>
  </ol>
</details>

## 关于本项目

checkUUID 是一个用来压力测试 `google/uuid` 包的小型 Go 程序。`main.go`
在一个循环里生成两亿（200,000,000）个 v4 版本的 UUID，并记录每一个已经出现过的
值，以便发现重复；`woker.go` 中定义的一个工作协程（goroutine）会在运行过程中
每完成百分之十就打印一次进度。循环结束后，程序会报告这两亿个 UUID 是否全部唯一，
如果不是，则列出每一个重复出现的 UUID 及其出现次数。

计划中的功能与已知问题，见 [open issues](https://github.com/anyingiit/checkUUID/issues)。

## 开始使用

### 环境要求

- Go 1.19 或更高版本，即 `go.mod` 中声明的版本
- Git，用于克隆本仓库

### 安装

```sh
git clone https://github.com/anyingiit/checkUUID.git
cd checkUUID
go build .
```

## 用法

该程序不接受任何参数，也不需要任何输入；运行后会立即开始固定次数
（200,000,000 次）的重复检测：

```sh
go run .
```

程序会每完成百分之十就打印一次进度，并在生成完全部 200,000,000 个 UUID 后，
打印 `test uuid of 200000000 ok!!!!!` 表示全部唯一，或者列出出现过不止一次的
UUID 及其重复次数。

## 参与贡献

欢迎参与。[CONTRIBUTING.md](CONTRIBUTING.md) 说明如何提交 issue 或 pull request，[CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md) 说明对所有参与者的行为要求。

请不要在公开的 issue 或 pull request 中报告安全问题。[SECURITY.md](SECURITY.md) 说明了私下报告的方式。

## 许可证

以 MIT 许可证分发。详见 [LICENSE](LICENSE)。

## 联系方式

项目地址：[https://github.com/anyingiit/checkUUID](https://github.com/anyingiit/checkUUID)

<p align="right">(<a href="#readme-top">back to top</a>)</p>
