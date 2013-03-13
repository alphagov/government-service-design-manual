---
layout: detailed-guidance
title: Testing code
subtitle: How to make sure your code does what it's supposed to
category: making-software
type: guide
audience:
  primary: developers, qa
  secondary: web-ops, tech-archs
status: draft
phases:
  - alpha
  - beta
  - live
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Making software
    url: /service-manual/making-software
---
    
We use automated testing to ensure that our code does what is intended, to protect against mis-use of that code, and to provide assurance that iterating that code for better design or new features doesn't break existing behaviour. 

We also add manual testing as an extra check where appropriate.

## Types of testing

Any code written for your service should have a suite of tests operating at two levels:

### Acceptance testing

Requires broad tests that run through high-level functionality end-to-end, making sure that the pieces of the system come together to provide the right service. 

A developer should be able to describe the steps in any acceptance test to the product/service manager in a way that makes sense to them and matches how they expect the service to be used (or abused!)

### Unit testing

Focussed on the specific details of the code ensuring that each discrete unit of code does what is expected of it. They allow the developers to verify that complex calculations are performed correctly, to ensure that code handles bad input properly, and that optimisations to the code don't break its behaviour.

## When to write tests

We aim to write a first set of tests at the start of working on a feature. An acceptance test that describes the end-to-end behaviour ensures that everyone involved understands the objective of a piece of work, and can demonstrate progress through the story at hand. Unit tests can then be written to understand the implementation of the code.

Tests are often described as "happy path" or "sad path". Happy path tests verify that the system can be used as intended, while sad-path tests verify that it handles errors (whether bad input from a user, a vital API being unavailable, or some other issue) gracefully. We start with happy path tests and a few simple sad path tests and then add more sad path tests as our understanding of the code and its dependencies develops.

Tests should also be written whenever a bug is discovered. A test to reproduce the bug should be written before it is fixed, allowing you to verify that the bug has been fixed and ensure that it isn't reintroduced later.

## Test early and often

Developers are expected to run tests regularly, especially before sharing new code, they are verified as part of the code review process, and they are also run regularly in a shared continuous integration system to ensure the whole team has a chance to see how they're performing.
