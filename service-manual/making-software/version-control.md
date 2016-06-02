---
layout: detailed-guidance
title: Version control
subtitle: Make sure the team can work together on code
category: making-software
type: guide
audience:
  primary: developers
  secondary: service-managers, web-ops
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

Every software development project has to use a version control system. Version control allows you to track changes to code over time, which means you can:

* quickly step back to an earlier version where necessary
* annotate your changes with explanatory details to help future developers understand the process

Version control also has tools for seeing what code has been changed, how it's been changed, and who made the changes.

## Commits

When updating the code, make small, discrete 'commits' of changes and group the changes according to their intention. Any changes should be committed with a clear message explaining what changes have happened, and (where appropriate) provide links to any supporting information like development stories, bug reports, or 3rd party documentation.

## Version control systems

GDS use a distributed version control system called [Git](http://git-scm.com/), which is one of the most well-known options.

Using a distributed version control system:

* means that everyone involved in the process has a full copy of the code and its history
* makes it easier for developers to create ‘branches’ in their code to explore new features or approaches without treading on the work of others
* will provide extra resilience when the network is unavailable, eg the developers can continue to work and make small incremental commits, merging their changes back with everyone else's at a later date

## Not just for code

It's a good idea to also use version control for other aspects of your work, not just code. GDS use the same version control tools to manage code and this document. The [Government Digital Strategy](/government/publications/government-digital-strategy) was also produced that way.
