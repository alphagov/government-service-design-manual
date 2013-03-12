---
layout: detailed-guidance
title: Shared asset libraries
subtitle: How gov.uk uses shared asset libraries, and where we provide our code
category: design-and-development-resources
type: resource
phases:
  - alpha
  - beta
  - live
audience:
  primary: designer, developer
  secondary: tech-arch
status: draft
---

Shared asset libraries are helpful for using the same frontend and
branding on multiple services.

## Why we do this

When building services around patterns and consistency, it's important
to share your frontend assets so that they can be easily reused as
required.

There's are two additional benefits to this approach:

- Single serve assets are
  [recommended by Yahoo! as a performance best practice](http://developer.yahoo.com/performance/rules.html#num_http)
- All used libraries are kept at known versions, which guarantees
  compatibility and reduces the risk of security vulnerabilities
  through external server compromise.

## Where to find our code

The templates on [GOV.UK](https://www.gov.uk) are constantly changing as we react to user
feedback and evolving best practice, so the best place to find them is
on our open source frontend libraries - more explanation of these to
follow:

- [Static](https://github.com/alphagov/static)
- [Frontend](https://github.com/alphagov/frontend)
- [govuk_frontend_toolkit](https://github.com/alphagov/govuk_frontend_toolkit)

### Static

The [static repository](https://github.com/alphagov/static) contains
our
[wrapper templates](https://github.com/alphagov/static/tree/master/app/views/root),
our
[basic CSS](https://github.com/alphagov/static/tree/master/app/assets/stylesheets),
and our
[basic Javascript](https://github.com/alphagov/static/tree/master/app/assets/javascripts). Anything
added to this repository is built to be used across the entire gov.uk
website and is used to provide a consistent look and feel.

### Frontend

The [frontend repository](https://github.com/alphagov/frontend)
contains the wrapper templates and views for many of the various pages
used across gov.uk. It gives a good indication as to how to structure
HTML and assets together.

### govuk_frontend_toolkit

Please see
[Sass repositories](/service-manual/ass-repositories.html) for
more information on this repository.


## Linking to shared assets

The template code contains direct links to CSS and JavaScript assets hosted on
the [GOV.UK](https://www.gov.uk) domain. We recommend that you leave these links as is whilst you
develop your service (as opposed to linking to your own copies of these files).
That way, you'll always be using the latest version of the assets and when they
change you can identify and resolve any conflicts immediately.

Once the service is ready for production we'll need to decide whether you should
continue to link to the assets in this way or whether you should now use your
own copies of them. Please contact GDS to discuss this at the appropriate time.

## Keeping shared assets up-to-date

[GOV.UK](https://www.gov.uk) is continuously being improved, which means that template and asset code
changes regularly. All services on GOV.UK are expected to keep their templates
and assets up-to-date. How you do this will depend on how you implement the
templates and where your service is hosted. Please contact the GDS team to
discuss options.
