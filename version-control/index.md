---
layout: detailed-guidance
title: Version control
subtitle: Ensure the team can collaborate on code
section: guidance
type: guide
audience:
  primary: service-manager
  secondary: tech-arch, developer, content-designer
status: draft
---

All software development projects must use a version control system. Version control allows you to track changes to code over time, meaning that you can quickly step back to an earlier version where necessary and you can annotate your changes with explanatory details to help future developers understand the process. Version control will also provide tools to audit who has made changes to the code and what has changed.

## Commits

Those updating the code should make small, discrete 'commits' of changes that are grouped according to their intention. They should be committed with a clear message explaining what the intention of the change was and (where appropriate) providing links to any supporting information such as development stories, bug reports, or third-party documentation.

## Version control systems

We prefer to use distributed version control systems (we use git which is the highest profile option), which mean that everyone involved in the process has a full copy of the code and of its history. This makes it easier for developers to create 'branches' in their code to explore new features or approaches without treading on the toes of those working on different aspects of the service. 

It also provides extra resilience - if the network is unavailable the developers can continue to work and make small incremental commits, merging their changes back with everyone else's at a later date.

## Not just code

It's a good idea to also use version control for other aspects of your work, not just code. We use the same version control tools to manage this document as we do our code, and the Government Digital Strategy was also produced that way.
