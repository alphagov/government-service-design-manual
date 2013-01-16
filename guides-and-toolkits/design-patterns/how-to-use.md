---
layout: gsdm
title: Design pattern library
section: guidance
subsection: Design patterns
status: draft
---

## How to use the patterns in this library


First, you'll need to be using the GOV.UK [Prototyping app](https://github.com/alphagov/prototyping).

The CSS for most patterns is stored as a Sass mixin, so you'll need to import the following at the top of your
 .scss files:

    @import "frontend_toolkit";
    @import "design_patterns";

The first line imports the mixins from the GOV.UK front end toolkit.

The second line imports the mixins from the pattern library.
