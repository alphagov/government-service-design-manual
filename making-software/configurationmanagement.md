---
layout: detailed-guidance
title: Configuration management
subtitle: Manage a team's approach to configuration
section: making-software
type: guide
audience:
  primary: developer, tech-arch
status: draft
phases:
  - alpha
  - beta
  - live
---
    
# Configuration management
Your system is likely much larger than a single application, relying
on other supporting infrastructure components. Even a simple application
probably requires some configuration, to provide database credentials or
a web service endpoint for instance. In order to build robust, scalable
and portable systems this configuration data should be well managed.

## Guidance/Tool

Configuration management tools help with documenting and maintaining the
configuration and dependencies of a software system. Although this could
be done using hand-made software, it's common to use existing tools.
Three examples of existing open source configuration management tools
are CFEngine, Chef and Puppet.

### Infrastructure as code

One approach to managing configuration is to describe the configuration
and the software dependencies in code. This brings with it all the
advantages of programming in general, including:

* Testability
* Reusability
* Executable documentation
* Common and constrained language to describe a problem domain

Once described in code the infrastructure configuration is executed
against the servers, networks and software in question.

### Build for portability

Moving software systems between providers can be difficult and time-
consuming. Even with compatible providers and simpler procurement rules
it's possible to build lockin with simple technical inertia.
Configuration management encourages a deep understanding of the
configuration of the system and this can be used to move software easily
between providers.

### Use the same tools for development and production

A common problem in software systems is seen when code written by a
development team works on their machine or a test environment but not on
the production environment. A common cause of this is differences in
configuration - different versions of software, different types of
database or application server. This can be avoided by using the same
tools for both development and production environments.

## Why we do this

Existing approaches to managing configuration are often manual, process
heavy, slow and error prone. Ultimately people are bad at doing detailed
monotonous tasks. And installing and configuring software across tens or
hundres of servers (if done by hand) is defintely monotonous.

Even if this could be done to provide everything correctly configured on
day 0, over time configuration drifts if not kept in check. One
traditional approach to this problem is to make configuration changes
hard, thereby limiting the number of them. When trying to build agile
and flexible software systems rapid change is needed and manual
processes break down. 

## Further reading

* [Infrastructure as
Code](https://speakerdeck.com/garethr/infrastructure-as-code)
* [CFEngine](http://cfengine.com/)
* [Chef](http://www.opscode.com/chef/)
* [Puppet](http://puppetlabs.com/solutions/configuration-management/)
