---
layout: detailed-guidance
title: Choosing technology
subtitle: How to go about choosing what software tools to use
category: making-software
type: guide
audience:
  primary: service-managers, tech-archs
  secondary: developers
status: draft
phases:
  - alpha
  - beta
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Making software
    url: /service-manual/making-software
---

At different points in your project you'll have to choose one technology over others. That might be a programming language, database, operating system or some small tool that helps the development team work more efficiently.

## You can change your mind

The most important consideration is to work on the assumption that most technology choices can change, especially during the early stages of development. 

On day one of your project you simply won't know enough about the domain or the user need to select the _right_ technology. It's OK to make an educated guess at this stage, as long as everyone understands that is what is happening. Then find time to challenge the selection as you learn more about the problem at hand.

Maybe you selected a programming language that you knew would be easier and quicker to prototype in for the early stages of the project, and then moved to another one that was easier to work in large teams with for the final product. Or maybe you started out with an open source product to allow you to get started quickly, before going on to buy a commercial product which provided some required feature (or vice-versa).

## Start with capabilities, not implementations

It's very easy to immediately jump to a specific product when making technology choices. This tends to be based either on past experience or on fashion. 

Try and take a step back and think about the capabilities of the technology you're after. Are you looking for a relational database? Or a document database? A key/value store? Or maybe a graph database? Argue first about the capabilities, rather than any specific implementation.

## Cost

With the growth of open source many technology products (databases, operating systems, programming languages, development tools, etc) are freely available. But a large market still exists for commercial software products. 

When choosing technology make sure you do consider the total cost, as well as any upfront fees. Try and take into consideration costs for things like staff, support and productivity.

## Consider people

Try and involve the whole team in technology choices. That doesn't mean no-one owns the decision making but that you want the development team to buy in to the technology choices made. Technology preferences vary, and technology choice can divide opinion. 

All things being equal, picking technologies that developers and operations staff like will likely result in improved productivity.

## Lock-in

Technology lock-in happens when previous decisions regarding technology limit future decisions, possibly so that only one real choice exists.

For example, if you select a database that only runs on one operating system you have no choice about the operating system you will use. If the costs of that operating system jump you have no simple way of reducing that cost quickly or cheaply.

Over time, and after many decisions, you can find yourself in a situation where all your technology decisions are tightly coupled and you are locked-in to one vendor or one way of doing things. This can have unforeseen financial costs (for example an overnight cost increase) but can also limit how quickly you can iterate on your product in the future, for instance if the ideal technology choice isn't compatible with your current vendor or technology.

Aim to have a clear understanding of the cost or implications of moving away from a technology when you commit to it.

## Build versus buy

Where there is an existing software solution which solves your problem, you should certainly consider using it. For the purposes of making this decision, making use of a piece of open source is effectively buying that piece of software.

You are more likely to buy in software that fulfils a commodity need. Development tools, build tools, utility libraries, databases and monitoring tools are all examples of software where many projects will have the same need, and it makes little sense to reinvent the wheel. For software that serves a strategic need you are far less likely to find a tool which serves and will continue to serve your needs.

Software that requires customisation to fit your needs is best avoided. The ongoing cost of maintaining this customisation is always greater than you think it's going to be.

## Deployment

The services that we write will need to be deployed, [early and frequently](/service-manual/making-software/release-strategies.html). Any components that we choose should be easy to deploy and upgrade as part of an [automated pipeline](/service-manual/agile/continuous-delivery.html).

## Testing

[Frequent, automated testing](/service-manual/making-software/testing-in-agile.html) is essential in agile development. Software that makes testing difficult should be avoided.

## Playing nicely version control

One frequent problem to watch for is tools that don't work well with [version control](/service-manual/making-software/version-control.html). Tools that create large binary files which need to be committed to version control, or whose files change in ways which make merging difficult should be avoided.

## Why we do this

Choosing technology is important, but it's probably not quite as important as you think. What is important are the users of that technology and being able to deliver quality at a sustainable pace and suitable cost. When making technology choices, and importantly as you develop your product and constantly reassess your selections, try and make decisions that:

* maximise developer productivity
* minimise total cost of ownership
* avoid lock-in

## Further reading

[Open Source Considerations](/service-manual/making-software/open-source.html)
