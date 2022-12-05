# 3. Build on Kubernetes

Date: 2022-12-05

## Status

Accepted

## Context

There exists a requirement to support the threee major cloud providers.  Additionally, we would like to potentially
work with as many more options as possible, including bare metal environments.  Providing deployment manifests or
tooling for the resources specific to each of these clouds would be impossible.

## Decision

The various systems here will be developed to run on top of Kubernetes.  This gives us a common substrate that
allows for as much of the deployment as possible to be agnostic of the underlying infrastructure environment.

## Consequences

- Supporting the major three cloud vendors is easier due to all having managed kubernetes offerings.
- Additional environments like bare metal are possible if a kubernetes cluster is set up in them.
