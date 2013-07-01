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
    
GDS use automated testing to:

* make sure that code does what it's supposed to
* make sure that code is protected against malicious attacks
* provide assurance that iterating the code for better design or new features won't introduce bugs

GDS also add manual testing as an extra check, where appropriate.

## Types of testing

Any code written for your service should have a set of tests operating at 2 levels:

* [acceptance testing](#acceptance-testing)
* [unit testing](#unit-testing)

### Acceptance testing

This requires broad tests that run through high-level functionality end-to-end, making sure that the pieces of the system come together to provide the right service.

A developer in your team should be able to describe the steps in any acceptance test to the product/service manager in a way that makes sense to them and matches how they expect the service to be used (or abused!)


### Unit testing

This concentrates on the specific details of the code, making sure that each separate unit of code does what is expected of it by: 

* allowing developers to verify that complex calculations are performed correctly
* making sure that code handles bad input from users properly 
* making sure that optimisations to the code won’t break its behaviour

## When to write tests

Write tests whenever a bug is discovered. Then write a test to reproduce the bug before it’s fixed, so you can confirm that the bug has been fixed and make sure that it isn’t reintroduced later.

GDS aim to write a first set of tests at the start of working on a feature. Acceptance tests that describe the expected end-to-end behaviour of code make sure that: 

* everyone involved understands the objective of a piece of work
*	progress can be demonstrated through the current user story 

These tests are often described as ‘happy path’ or ‘sad path’:

* happy path tests confirm that the system can be used as intended
*	sad path tests confirm that it handles errors (whether bad input from a user, a vital API being unavailable, or some other issue) smoothly

GDS start with happy path tests and a few simple sad path tests, and then add more sad path tests as understanding of the code and its dependencies develops.


## Test early and often

Developers are expected to run tests regularly, especially before sharing new code. Examine the code in tests as part of the code review process, and regularly run tests in a shared continuous integration system. This gives the whole team a chance to see how they’re performing.
