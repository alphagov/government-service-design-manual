---
layout: detailed-guidance
title: Configuration management
subtitle: Manage a team's approach to software configuration
category: making-software
type: guide
audience:
  primary: web-ops
  secondary: developers, tech-archs
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
exclude_from_search: true
---

Configuration management is about managing how the pieces of your software and/or service work together.

It's important to get it right as your system is likely to be much larger than a single application, relying on other supporting infrastructure components. Even a simple application probably requires some configuration, like providing database credentials or a web service endpoint.

Managing your configuration data well will mean you can build a stable, scalable and portable system.

## Management tools

Configuration management tools will help with documenting and maintaining the configuration and dependencies of your software system. Although you can do this using hand-made software, it’s common to use existing tools.

Three examples of existing open source configuration management tools are:

* [CFEngine](http://cfengine.com/)
* [Chef](https://www.chef.io/chef/)
* [Puppet](https://puppetlabs.com/)

### Infrastructure as code

One approach to managing configuration is to describe in code the configuration and software dependencies (what your software needs to perform its function).

This brings with it all the advantages of programming in general, including:

* testability
* reusability
* executable documentation
* common and constrained language to describe a problem domain

Once written in code, the infrastructure configuration is then run against the servers, networks and software in question.

### Build for portability

Moving software systems between providers can be difficult and time-consuming. Even with compatible providers and simpler procurement rules, it’s possible to lock yourself in through inaction alone.

Configuration management encourages a deep understanding of the configuration of your system. Use this understand to move software easily between providers, if needed.


### Use the same tools for development and production

A common problem in software systems is when code written by a development team works on their machine or a test environment, but not on the production environment. A common cause of this is differences in configuration:

* different versions of software
* different types of database
* different application servers

Avoid this by using the same tools for your development and production environments.

## Why GDS do this

Existing methods to managing configuration are often manual, process-heavy, slow and error prone.

Ultimately, people are bad at carrying out detailed monotonous tasks - and installing and configuring software across tens or hundreds of servers (if done by hand) is definitely monotonous.

Even if this could be done to provide everything correctly configured on day 0, configuration will drift if not kept in check. A traditional solution to this problem is to make it harder to change configuration, which then limits the number of changes that can made.

Overall, the ability to change rapidly is essential when trying to build agile and flexible software systems, which means breaking down those restrictive manual processes.


## Further reading

* [Infrastructure as Code](https://speakerdeck.com/garethr/infrastructure-as-code)
* [Managing dependencies](/service-manual/making-software/dependency-management)
