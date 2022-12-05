# 4. Build with Golang

Date: 2022-12-05

## Status

Accepted

## Context

There is a requirement to use a popular and productive programming language in order to help
encourage community contributions.

## Decision

This project will be primarily implemented using [golang](https://go.dev/).

- Golang is a systems programming language well suited to the task of handling blockchain API requests.
- A popular choice for blockchain projects like [go-ethereum](https://github.com/ethereum/go-ethereum).
- It is a first class language for interacting with kubernetes.

## Consequences

Unless motivation exists otherwise (e.g. Terraform for infrastructure code), the default programming
language used for implementation in this repo will be golang.
