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
  primary: designers
  secondary: developers
status: draft
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: User-centred design
    url: /service-manual/user-centred-design
---

Sass is a language for creating stylesheets that lets us share blocks of code and techniques. GDS has created a repository called the [GOV.UK Frontend Toolkit][govuk_frontend_toolkit] to simplify the creation of services with a consistent look and feel. The toolkit is [available as a gem][govuk_frontend_toolkit_gem] for easy inclusion in Ruby projects.

## Categories of Sass mixins

The files in [the gem][govuk_frontend_toolkit_gem] can be categorised into four main parts:

1. GOV.UK typography, colours and image assets
2. Mixins for responsive designs
3. Mixins for targeting old versions of Internet Explorer
4. Mixins for cross browser CSS

The first part is the bit that gives all GOV.UK projects a familiar look. There are a collection of pre-defined font sizes that we use on GOV.UK. There is a mixin for each one, for example `heading-26`. These also include a standard amount of whitespace around the text to help with vertical rhythm on the page, spacing things out nicely.

The second is a way to develop sites that are able to respond to different sized displays.

The third is an easy way of writing Internet Explorer specific CSS in our stylesheets without using hacks.

The fourth is a way to keep browser specific styles out of our projects. We encapsulate new or non-standardised CSS into mixins. In this way we can easily update all the instances of a new CSS property without having to do a search and replace across all our projects.

## Responsive design

It is generally advised to write your markup with a mobile first attitude. That is, add desktop styles to an otherwise narrow screen stylesheet. In this way you only add styles for desktop and don't reset desktop styles for a mobile device.

## Cross browser

There are two main types of cross browser CSS that we are concerned with. Firstly using different techniques to achieve a consistent effect and secondly using vendor prefixes to apply consistent behaviour for newer features. For example:

    .related-box {
      float: left;
      @extend %contain-floats;
      @include border-radius(4px);
    }

`@extend %contain-floats` uses a cross-browser technique to ensure that the element wraps all the floated elements within it. It is not a property that normally exists in CSS but is something we often need to do and don't want to use different techniques everywhere. It gives us consistency across our code.

`@include border-radius` is designed to use the different border radius implementations (`-moz-border-radius`, `-webkit-border-radius` etc) to create consistent presentation across different browsers.

## Further reading

The `README.md` file in the [GOV.UK Frontend Toolkit][govuk_frontend_toolkit] has more information.

[govuk_frontend_toolkit]: https://github.com/alphagov/govuk_frontend_toolkit
[govuk_frontend_toolkit_gem]: https://github.com/alphagov/govuk_frontend_toolkit_gem
