# 5. Kubebuilder Operator

Date: 2022-12-05

## Status

Accepted

## Context

It has been [decided](./0003-build-on-kubernetes.md) to build on Kubernetes as a common deployment target.
We want to simplify and standardize that deployment as much as possible to ease both development and usage.

## Decision

We will implement the [operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)
via a [kubebuilder] (https://book.kubebuilder.io/) project. 

The change that we're proposing or have agreed to implement.

## Consequences

- An operator lets us provide easy installation and deployment of InfraKit resources in a kubernetes cluster
via higher level Custom Resource Definitions (CRDs).
- Kubebuilder provides scaffolding and boilerplate to help implementing such an operator.
