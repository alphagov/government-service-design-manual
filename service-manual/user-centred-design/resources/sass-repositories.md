---
layout: detailed-guidance
title: Sass repositories
subtitle: A library of CSS mixins in the Sass format
category: design-and-development-resources
type: resource
phases:
  - alpha
  - beta
  - live
audience:
  primary: designers, developers
status: draft
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: User centered design
    url: /service-manual/user-centered-design
---

Sass lets us share blocks of code and techniques. [GOV.UK](https://www.gov.uk) has a collection of Sass files which are bundled up into a [gem](https://github.com/alphagov/govuk_frontend_toolkit_gem) for easy inclusion into new projects. For projects not written in Ruby note that the gem is just a wrapper for the [GOV.UK Frontend toolkit](https://github.com/alphagov/govuk_frontend_toolkit), a repository containing all the files.

## Categories of Sass mixins

The files in the [gem](https://github.com/alphagov/govuk_frontend_toolkit_gem) can be categorised into three different things:

>1. GOV.UK typography, colours and image assets
>2. Mixins for responsive designs
>3. Mixins for targeting old versions of Internet Explorer
>4. Mixins for cross browser CSS

The first is the key bit which makes [GOV.UK](https://www.gov.uk) projects look the same. There are a collection of pre-defined font-sizes that we use on [GOV.UK](https://www.gov.uk). There is a mixin for each one, for example `heading-26`. These also include a standard amount of whitespace around the text to help with vertical rhythm on the page, spacing things out nicely.

The second is a way to develop sites that can respond to being on different sized displays.

The third is a very easy way of writing IE specific CSS in the middle of our stylesheets without using hacks.

The forth is a way to keep browser specific styles out of our projects. We encapsulate new or non-standardised CSS into mixins. In this way we can easily update all the instances of a new CSS property without having to do a search and replace across all of our projects.

## Responsive Design

It is generally advised to write your markup with a mobile and up attitude. That is, add desktop styles to an otherwise narrow screen stylesheet. In this way you only add styles for desktop and don't reset desktop styles for a mobile device.

## Cross Browser

There are two types of cross browser technique. There are some which are just for encapsulating vendor prefixes. Then there are some for using different methods to achieve a consistent effect. An example of these are:

    .related-box {
      float: left;
      @extend %contain-floats;
      @include border-radius(4px);
    }

The `border-radius` line here is designed to use the different border-radius implementations to create a standard border-radius. The `contain-floats` however, uses a cross-browser technique to ensure that the element wraps all of the floated elements within it. It is not a property that normally exists in CSS but is something we often need to do and don't want to use different techniques everywhere.

## Further reading

For more information check out the Readme on the [GOV.UK Frontend toolkit](https://github.com/alphagov/govuk_frontend_toolkit)
