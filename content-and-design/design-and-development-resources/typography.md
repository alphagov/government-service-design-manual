---
layout: detailed-guidance
title: Typography
subtitle: Creating clear, easy to read type for the web
section: guidance
audience:
  primary: designer
type: resource
status: draft
---

Services should use clear, easy to read type, with consistent styles and a clear hierarchy of information.

##Guidance
We use the typeface New Transport, cut especially for Government use. This should be included via CSS. We only serve this to modern browsers, via two files: a WOFF file for the majority of browsers and an EOT file for Internet Explorer (although we only serve this to IE8). The font is hinted to display well on all browsers.

We currently use two weights of New Transport: Light and Bold. Italics should not be used. The number of different type sizes on a page should be kept to the minimum, and only one typeface/font should be used on each website.

If the service uses numbers in columns or tables, you should change these to use the tabular number version of Transport New. This replaces the standard numbers with new versions that are spaced more evenly. The main noticable difference is a base on the character 1. GDS has used this on the Performance Platform and Trade Tariff.

Text must have enough contrast against the background colour to be readable. This should be tested to conform with our [Accessibility requirements](/content-and-design/accessibility.html). Generally we use type in #0B0C0C against a white or light grey background. Links should be blue and underlined - see [Colour palettes](/content-and-design/design-and-development-resources/colour-palettes.html).

Type should be large enough to be easily read. This is generally larger than many current websites: 36px for headlines, 19px for body text. This can be included using default styles in scss from the [Frontend Toolkit](/content-and-design/design-and-development-resources/sass-repositories.html). These include line height spacing that works across browsers.

TODO: method for services to include GDS-hosted Transport (+ system for GDS to manage/restrict use)

##Code/Templates

See [Colour palettes](/content-and-design/design-and-development-resources/colour-palettes.html) for the Sass variables for text colours.
See [Shared asset libraries](/content-and-design/design-and-development-resources/shared-asset-libraries.html) and [Sass Repositories](/content-and-design/design-and-development-resources/sass-repositories.html) for easy implementation of GOV.UK typography.

##Why we do this
See [this blog post about why we've chosen Transport](http://digital.cabinetoffice.gov.uk/2012/07/05/a-few-notes-on-typography/).

##Further reading

[Shared asset libraries](/content-and-design/design-and-development-resources/shared-asset-libraries.html)

[Sass Repositories](/content-and-design/design-and-development-resources/sass-repositories.html)

[Colour palettes](/content-and-design/design-and-development-resources/colour-palettes.html)

[Accessibility requirements](/content-and-design/accessibility.html)
