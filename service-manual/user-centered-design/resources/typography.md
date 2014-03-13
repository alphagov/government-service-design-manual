---
layout: detailed-guidance
title: Typography
subtitle: Creating clear, easy to read type for the web
category: design-and-development-resources
type: resource
phases:
  - alpha
  - beta
  - live
audience:
  primary: designers
status: draft
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: User centered design
    url: /service-manual/user-centered-design
---

Services should use clear, easy to read type, with consistent styles and a clear hierarchy of information.


## The typeface

GOV.UK uses the typeface [New Transport](http://en.wikipedia.org/wiki/Transport_(typeface)#New_Transport).

It's been cut especially for government use and is not licenced for use outside of the GOV.UK domain.

If you're making your service publically available on a different domain (a beta version for example), 
use an alternative typeface like Arial.


## Type size

Keep the number of different type sizes on a page to a minimum.
Type should be large enough to be easily read. 
This is generally larger than many current websites: 36px for headlines, 19px for body text. 


## Type style

There are two weights of New Transport: Light and Bold. 
Do not use italics, as it's harder to read.


## Colour contrast

Text must have enough contrast against the background colour to be readable. 
This should be tested to conform with our [Accessibility requirements](/service-manual/user-centered-design/accessibility.html).

Check the [colour palettes](/service-manual/user-centered-design/resources/colour-palettes.html)
page to see our standard colours for text and links.


## Headings

All headings must be sentence case, eg ‘Your tax account’ not ‘Your Tax Account’. However, where a proper noun appears in the heading this is upper case, eg ‘Your National Insurance contributions’. Check the section on capitalisation in the [GOV.UK style guide](https://www.gov.uk/design-principles/style-guide/style-points#style-capitalisation).


## Tabular data

If the service uses numbers in columns or tables, you should use the tabular number version of New Transport. 
This replaces the standard numbers with new versions that have a fixed width. 
The main noticable difference is a base on the character 1. 
GDS has used this on the [Performance platform](/performance) and Trade tariff.


## Implementing the typographic styles

The New Transport typeface is embedded in the CSS provided with the GOV.UK templates.

It's served to browsers via a WOFF file (or an EOT file for Internet Explorer 8). 
The font is hinted to display well on all browsers. 
Older browsers do not receive the typeface.

Sass users can get mixins for the different type sizes and styles (including the tabular version) from the [GOV.UK frontend toolkit](https://github.com/alphagov/govuk_frontend_toolkit). These include line height spacings that works across browsers.


## Further reading

* [Why we've chosen Transport](http://digital.cabinetoffice.gov.uk/2012/07/05/a-few-notes-on-typography/)
* [Shared asset libraries](/service-manual/user-centered-design/resources/shared-asset-libraries.html)
* [Sass Repositories](/service-manual/user-centered-design/resources/sass-repositories.html)
* [Colour palettes](/service-manual/user-centered-design/resources/colour-palettes.html)
* [Accessibility requirements](/service-manual/user-centered-design/accessibility.html)
