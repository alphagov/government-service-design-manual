---
layout: detailed-guidance
title: Shared asset libraries
subtitle: How GOV.UK uses shared asset libraries, and where we provide our code
category: design-and-development-resources
type: resource
phases:
  - alpha
  - beta
  - live
audience:
  primary: designers, developers
  secondary: tech-arch
status: draft
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: User-centred design
    url: /service-manual/user-centred-design
---

Shared asset libraries are helpful for using the same frontend and
branding on multiple services.

## Why we do this

When building services around patterns and consistency, it's important
to share your frontend assets so that they can be easily reused as
required.

There are two more benefits to this approach:

- [Yahoo recommends minimising HTTP requests](https://developer.yahoo.com/performance/rules.html#num_http)
  to improve performance
- All used libraries are kept at known versions, which guarantees
  compatibility and reduces the risk of security vulnerabilities
  through external server compromise

## Where to find our code

The templates on GOV.UK are constantly changing as we react to user
feedback and evolving best practice, so the best place to find them is
by looking at the code we publish:

- [govuk_template](https://github.com/alphagov/govuk_template)
- [govuk_frontend_toolkit](https://github.com/alphagov/govuk_frontend_toolkit)

### govuk_template

The GOV.UK template is a project that provides the GOV.UK header and footer, as well as associated
assets. It generates a variety of output formats and you can extend it with more depending on
the language your service is written in.

GOV.UK's [static][] consumes the Ruby version of the template application, which in turn
provides shared resources like footer links across the various frontend apps that run GOV.UK.

The Performance Platform's [Node.js application][spotlight] uses the [Mustache][mustache] version of the template.

[static]: https://github.com/alphagov/static/blob/release_1994/Gemfile#L27-30
[spotlight]: https://github.com/alphagov/spotlight
[mustache]: https://mustache.github.io/

You should be able to become a consumer of the template in exactly the same way these two projects
are, by adding it to your application's dependencies.

### govuk_frontend_toolkit

The service manual article on
[Sass repositories](/service-manual/user-centred-design/resources/sass-repositories) has
more information on this repository.


GDS is continuously improving GOV.UK, which means that template and asset code changes regularly.

All services on GOV.UK are expected to keep their templates and assets updated. How you do this will
depend on how your frontend is implemented, but if you include the template and the toolkit as
dependencies in your application it should be relatively easy to update as GDS publishes new
versions. Please contact the GDS team for help or advice.
