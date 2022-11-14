# 2. Manage tool dependencies using asdf

Date: 2022-11-14

## Status

Accepted

## Context

This project will have many dependencies (programming language, utilities, etc...).
Managing dependencies across environments (developer machines, CI) is often a difficult
task which is why many independent version managers (nvn, pyenv) exist.

Dealing with multiple version managers is painful, and not all tools will even have one,
with usually varying default install instructions.

Having a single cross-platform version manager that can install all tool dependencies
with a simple command is beneficial to the development of this project.

## Decision

We will use [asdf](https://asdf-vm.com/) to manage all tool dependencies whenever possible.

## Consequences

Developers can easily install all necessary tool dependencies when working on this project and
stay in version sync as we update them. CI and automation environments can also use `asdf` for
installation to run with an equivalent environment.
