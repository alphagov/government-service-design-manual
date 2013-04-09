---
layout: detailed-guidance
title: Continuous Delivery
subtitle: Making releases boring
category: agile
type: guide
audience:
  primary: service-managers, designers, developers, performance-analysts, user-researchers, content-designers, tech-archs
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

Continuous Delivery builds upon the iterative, adaptive process of delivering valuable software and testing assumptions about product direction. From the [Lean philosophy][], software that is not in production isn't delivering value, it should be treated as inventory piling up, and inventory is waste.

[Agile][] development practices normally look to have a product in a deployable state at the end of each iteration. It does not necessarily follow that the product should be deployed at the end of each iteration. The important thing is the capability to do so.

Understanding your end-to-end deployment pipeline has massive benefits for addressing [configuration management][] and the automation of your build, deploy, test and release processes. By automating your deployment process, you are forced to fully understand it and resolve early on any issues with getting code from your version control system into production. Automating the deployment early also means that it will be tested and bugs ironed out, so that [releases become frequent, low-risk][release regularly], almost boring events.

If you treat long-term planning as an attempt to predict the future, then attempting to predict what features will be available for a production release slot in 6 months time is a somewhat redundant exercise. If instead, your process allows for deploying anything when it is ready, then release planning becomes a lot simpler. If a project needs to miss a release slot, then it can always be rescheduled for tomorrow or next week, rather than in 6 months, or perhaps in a year if the next slot is already full up!

Another advantage is being able to quickly respond to security patches or similar changes in underlying libraries / frameworks used by your application. You can quickly make a change and watch the update flow through the various gates in your deployment pipeline, confident that nothing has been broken.

## The deployment pipeline

What happens to code between it being written by a developer, and
deployed to production? We refer to this process as the **deployment
pipeline**.

### The commit stage

Whenever a developer checks into [version control][], a
[suite of tests][testing in agile] is run against the latest version
of the code. At this stage, any quick, easy-to-identify defects such
as compile errors or unit test failures are caught. If the tests pass,
the code progresses to the next stage.

### Shared sandbox environment

The code is deployed to a shared [sandbox][] environment, where
everyone involved in the project can observe it. The sandbox should be
similar to production as far as is practical: for example, if
production uses Postgres, the sandbox should also use Postgres and not
another database such as MySQL or sqlite.

Every commit is considered a potential candidate to be released into
production. The sandbox environment is the first environment where the
application is deployed and run. This is the first stage where it can
be visually inspected for [quality][] by anybody on the team. The
purpose is to identify any defect which means the application should
not be deployed to production. If such a defect is found, this version
of the code stops here; otherwise, it can proceed to further
specialist testing environments.

### Specialist testing environments

There may be a need for other testing environments, to enable testing
for specialist requirements such as [load and performance testing][],
[penetration testing][], or [accessibility testing][]. How many
environments are needed will depend on the requirements and
constraints of individual projects.

If code is determined to be of satisfactory quality, it can now
proceed to the live production environment.

### Production environment

Once code has passed the commit stage, been deployed into the shared
sandbox environment, had any necessary specialist testing run on it,
it is considered suitable to go live. Deploying to production should
be done in the same way as deploying to any other environment -- using
the same scripts, same [configuration management][] tooling, and the
same version of the code. This ensures that when code is released to
production, you are not doing it for the first time; you are instead
performing an operation which has been validated at each stage
throughout the deployment pipeline.

[Agile]: /service-manual/agile
[Lean philosophy]: http://en.wikipedia.org/wiki/Lean_software_development
[accessibility testing]: /service-manual/making-software/accessibility-testing.html
[configuration management]: /service-manual/making-software/configuration-management.html
[development environment]: /service-manual/making-software/development-environment.html
[load and performance testing]: /service-manual/operations/load-and-performance-testing.html
[penetration testing]: /service-manual/operations/penetration-testing.html
[quality]: /service-manual/agile/quality.html
[release regularly]: /service-manual/making-software/release-strategies.html
[sandbox]: /service-manual/making-software/sandbox-and-staging-servers.html
[testing in agile]: /service-manual/making-software/testing-in-agile.html
[version control]: /service-manual/making-software/version-control.html
