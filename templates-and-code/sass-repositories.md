---
layout: default
title: Sass repositories
subtitle: A library of CSS mixins in the Sass format
section: guidance
audience:
  primary: designer, developer
type: resource
status: draft
---
    
Sass lets us share blocks of code and techniques. [GOV.UK](www.gov.uk) has a collection of Sass files which are bundled up into a gem that enable us to quickly build pages that conform to our styles.

## Categories of Sass mixins

Within the `govuk_frontend_toolkit` gem we have a collection of Sass which is designed to enable you to easily add GOV.UK styles. They can be categorised into three different things:

>1. GOV.UK Typography and Colours
>2. Mixins for responsive designs
>3. Mixins for cross browser CSS

The first is the key bit which makes [GOV.UK](www.gov.uk) projects look the same. There are a collection of pre-defined font-sizes that we use on [GOV.UK](www.gov.uk). There is a mixin for each one, for example `heading-26`. These also include a standard amount of whitespace around the text to help with vertical rhythm on the page, spacing things out nicely.

The second is a way to develop responsive sites while not giving older browsers a bad experience. That is to still deliver a desktop stylesheet to older versions of IE which don't understand media queries, and serve the correct media queries to modern devices. As a side effect of this approach we also have a very easy way of writing IE specific CSS in the middle of our stylesheets without using hacks.

The third is a way to keep browser specific styles out of our projects. We encapsulate new or non-standardised CSS into mixins. In this way we can easily update all the instances of a new CSS property without having to do a search and replace across all of our projects.

For a full list of the different Sass techniques we have available look at the [readme on the `govuk_frontend_toolkit` gem](https://github.com/alphagov/govuk_frontend_toolkit).

## Typography and Colours

Some of the typography mixins we have available are: `heading-36`, `heading-24`, `copy-19`. They can be used as such:

    h1 {
      @include heading-36;
      @include text-colour;
    }
    p {
      @include copy-19;
      
      a {
        @include link-colour;
      }
    }

## Responsive Design

It is generally advised to write your markup with a mobile and up attitude. That is, add desktop styles to an otherwise narrow screen stylesheet. In this way you only add styles for desktop and don't reset desktop styles for a mobile device. This can be done as such:

    .navigation {
      background: $panel-colour;
      
      @include media(desktop){
        float: left;
        width: 33%;
      }
    }
    .content {
      padding-top: 30px;
      
      @include media(desktop){
        float: left;
        width: 33%;
      }
    }
    
## Cross Browser

There are two types of cross browser technique. There are some which are just for encapsulating vendor prefixes. Then there are some for using different methods to achieve a consistent effect. An example of these are:

    .related-box {
      float: left;
      @extend %contain-floats;
      @include border-radius(4px);
    }
    
The `border-radius` line here is designed to use the different border-radius implementations to create a standard border-radius. The `contain-floats` however, uses a cross-browser techinque to ensure that the element wraps all of the floated elements within it. It is not a property that normally exists in CSS but is something we often need to do and don't want to use different techniques everywhere.

## Further reading

For more information check out the Readme on the [`govuk_frontend_toolkit` gem](https://github.com/alphagov/govuk_frontend_toolkit)
