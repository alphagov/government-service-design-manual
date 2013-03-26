---
layout: detailed-guidance
title: Continuous Delivery
subtitle: Making releases boring
category: agile
type: guide
audience:
  primary: service-managers, designers, developers, performance-analysts, user-researchers, content-designers
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

Continuous Delivery builds upon the iterative, adaptive process of delivering valuable software and testing assumptions about product direction. From the Lean philosophy, software that is not in production isn't delivering value, it should be treated as inventory piling up, and inventory is waste.

Agile development practices normally look to have a product in a deployable state at the end of each iteration. It does not necessarily follow that the product should be deployed at the end of each iteration. The important thing is the capability to do so.

Understanding your end-to-end deployment pipeline has massive benefits for addressing configuration management and the automation of your build, deploy, test and release processes. By automating your deployment process, you are forced to fully understand it and resolve early on any issues with getting code from your version control system into production. Automating the deployment early also means that it will be tested and bugs ironed out, so that releases become frequent, low-risk, almost boring events.

If one treats long-term planning as an attempt to predict the future, then attempting to predict what features will be available for a production release slot in 6 months time is a somewhat redundant exercise. If instead, your process allows for deploying anything when it is ready, then release planning becomes a lot simpler. If a project needs to miss a release slot, then it can always be rescheduled for tomorrow or next week, rather than in 6 months, or perhaps in a year if the next slot is already full up!

Another advantage is being able to quickly respond to security patches or similar changes in underlying libraries / frameworks used by your application.
