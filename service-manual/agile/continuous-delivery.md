---
layout: detailed-guidance
title: Continuous delivery
subtitle: Making releases boring
category: agile
type: guide
audience:
  primary: service-managers, web-ops, designers, developers, performance-analysts, tech-archs
  secondary:
status: draft
phases:
  - discovery
  - alpha
  - beta
  - live
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Agile
    url: /service-manual/agile
---

Continuous delivery is about producing regular iterations of your software that’s ready to be released (deployed), though you don’t have to release these iterations to the public.

Producing regular iterations makes it easier for you to:

* add improvements
* fix bugs
* test expectations about your end product

If your software isn’t ready to be used then it’s not creating any real value. Treat it as stock piling up -- and stock is waste.

The [Lean software development philosophy](https://en.wikipedia.org/wiki/Lean_software_development) tells us to:

> Test all iterations of your software, either through public user testing or automated testing (using separate software to perform the testing).

## Deployment
Automate your deployment process so you are forced to fully understand it. Then any issues with moving code from your [version control system](/service-manual/making-software/version-control.html) into production (when it’s gone live) can be dealt with early on.

Automating it early also means that code will be tested and any bugs fixed so that [releases become frequent](/service-manual/making-software/release-strategies.html), low-risk, almost boring events.

Don’t plan production release slots far in advance. You can’t be certain what will be ready in 6 months' time.

Make your release planning simpler -- make sure it’s flexible enough that any feature or update to your software can be deployed when it’s ready. Then, if a feature needs to miss a release slot, it can easily be rescheduled for tomorrow or next week, rather than in 6 months, or more if the next slot is already full.

Another advantage is being able to quickly respond to security patches or similar changes in underlying  libraries or frameworks used by your application. You can quickly make a change and watch the update flow through the various gates in your deployment pipeline, confident that nothing has been broken.

##The deployment pipeline

What happens to code between it being written by a developer and deployed to production is known as the deployment pipeline.

Understand your end-to-end deployment pipeline. Knowing how it works and how each element works together will have implications for:

* [configuration management](/service-manual/making-software/configuration-management.html) (how you maintain consistency with your product’s performance and functionality)
* the automation of your build, deploy, test and release processes

The deployment pipeline has 4 stages:

* commit stage
* shared sandbox environment
* specialist testing environment
* production environment

###The commit stage
When your developer checks into [version control](/service-manual/making-software/version-control) (where all code, including previous versions, are stored), [a range of tests](/service-manual/making-software/testing-in-agile) should be run against the latest version of the code. Any quick, easy-to-identify defects, like [compile errors](https://en.wikipedia.org/wiki/Compilation_error) or [unit test failures](https://en.wikipedia.org/wiki/Unit_testing) will be found at this stage. If the tests pass, the code progresses to the next stage and should be considered for release into production.

###Shared sandbox environment
Send the code to a shared [sandbox](/service-manual/making-software/sandbox-and-staging-servers.html) (testing) environment. This is the first environment where the application is deployed and run. It's also the first stage where it can be visually inspected for [quality](/service-manual/agile/quality.html) by anybody on the team.

Make the sandbox environment as similar to the production (live) version as far as is practical. For example, if production uses Postgres, the sandbox should also use Postgres and not another database like MySQL or SQLite.

The purpose is to find any defects in the code. If a defect is found, stop the version of the code at this stage. If it passes the sandbox environment it can move on to further specialist testing environments.

###Specialist testing environments
You may need to perform additional testing for specialist requirements, like [load and performance testing](/service-manual/operations/load-and-performance-testing.html), [penetration testing](/service-manual/operations/penetration-testing.html), or [accessibility testing](/service-manual/user-centred-design/user-research/accessibility-testing.html). The amount of environments you’ll need will depend on the requirements and conditions of your individual projects.

When you are satisfied with the quality of the code, move it to the live production environment.

###Production environment

Your code is ready to go live if it has passed:

* the commit stage
* the shared sandbox environment
* any necessary specialist testing

Deploy to production the same way as you deploy to any other environment -- use the same scripts, same [configuration management](/service-manual/making-software/configuration-management.html) tooling, and the same version of the code.  This means you’re not releasing code for the first time -- you’re performing a task that’s been validated at each stage throughout the deployment pipeline.
