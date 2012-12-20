#Typography
Services should use clear, easy to read type, with consistent styles and a clear hierarchy of information.

##Guidance/Tool
We use the typeface New Transport, cut especially for Government use. This should be included via CSS. We only serve this to modern browsers, via two files: a WOFF file for the majority of browsers and an EOT file for Internet Explorer (although we only serve this to IE8). The font is hinted to display well on all browsers.

We currently only use one weight of New Transport: Light. We are in the process of adding a Bold weight too, but in the meantime differentiation between copy styles should be accomplished using different type sizes. Italics should not be used. The number of different type sizes on a page should be kept to the minimum, and only one typeface/font should be used on each website.

If the service uses numbers in columns or tables, you should change these to use the tabular number version of Transport New. This replaces the standard numbers with new versions that are spaced more evenly. The main noticable difference is a base on the character 1. GDS has used this on the Performance Platform and Trade Tariff.

Text must have enough contrast against the background colour to be readable. This should be tested to conform with our [Accessibility Requrements](/handbook/4/). Generally we use type in #0B0C0C against a white or light grey background. Links should be blue and underlined - see [Colour palettes](/handbook/180/).

Type should be large enough to be easily read. This is generally larger than many current websites: 36px for headlines, 19px for body text. This can be included using default styles in scss from the [Frontend Toolkit](/handbook/???/). These include line height spacing that works across browsers.

TODO: method for services to include GDS-hosted Transport (+ system for GDS to manage/restrict use)

##Code/Templates

See [Colour palettes](/handbook/180/) for the Sass variables for text colours.
See [Shared asset libraries](/handbook/131/) and [SCSS Repositories](/handbook/130/) for easy implementation of GOV.UK typography.

##Why we do this
See [this blog post about why we've chosen Transport](http://digital.cabinetoffice.gov.uk/2012/07/05/a-few-notes-on-typography/).

##Further reading

[Shared asset libraries](/handbook/131/)
[SCSS Repositories](/handbook/130/)
[Colour palettes](/handbook/180/)
[Accessibility requirements](/handbook/4/)
