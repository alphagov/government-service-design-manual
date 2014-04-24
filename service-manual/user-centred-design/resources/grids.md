---
layout: detailed-guidance
title: Grids
subtitle: Quickly lay out lists of items in a grid
category: design-and-development-resources
phases:
  - alpha
  - beta
  - live
status: draft
page_class: grid
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: User centred design
    url: /service-manual/user-centred-design
---

Use this mixin if you need to arrange content in a grid or split part of a page into columns. You might want to do this for an image gallery, product catalogue or home page layout. It's particularly useful if you don't want to explicitly represent rows or columns in the markup.

Don't use this

* when a simple linear layout will work
* when a data table is more appropriate
* for full page layouts (it's not that kind of grid)

## Using the mixin

[The grid mixin](https://github.com/alphagov/government-service-design-manual/blob/master/service-manual/assets/stylesheets/design-patterns/_grid.scss) accepts the following arguments:

`$columns`
  : The number of columns in the grid, or an array representing the relative width of each column.

`$min-height`
  : An optional minimum height for grid elements. Useful if your grid elements contain varied amounts of content.

`$max-rows`
  : Adds IE7 and IE8 support for grids with varied column widths. See below for details.

The mixin is tag-agnostic, so the elements can be list items, divs, paragraphs or anything else.
Avoid applying border effects to the elements as this will make the width calculation incorrect.
Style the contents of those elements instead.

At mobile screen sizes the grid elements switch to being full-width.

## Simple example

<ul class="grid example-1">
  <li><p>Item 1</p></li>
  <li><p>Item 2</p></li>
  <li><p>Item 3</p></li>
  <li><p>Item 4</p></li>
  <li><p>Item 5</p></li>
  <li><p>Item 6</p></li>
  <li><p>Item 7</p></li>
  <li><p>Item 8</p></li>
</ul>

### HTML

    <ul class="grid example-1">
      <li><p>Item 1</p></li>
      <li><p>Item 2</p></li>
      <li><p>Item 3</p></li>
      <li><p>Item 4</p></li>
      <li><p>Item 5</p></li>
      <li><p>Item 6</p></li>
      <li><p>Item 7</p></li>
      <li><p>Item 8</p></li>
    </ul>

### Sass

    @import "grid"

    .example-1{
      @include grid(4);
    }


## Regular grids

You can create grids of equally sized elements by passing in a single value representing the number of elements in a row.

The following examples are for demonstration purposes only and not ones we'd ever recommend.

### Example

<ul class="grid example-4">
  <li><p>Item 1</p></li>
  <li><p>Item 2</p></li>
</ul>

<ul class="grid example-1">
  <li><p>Item 1</p></li>
  <li><p>Item 2</p></li>
  <li><p>Item 3</p></li>
  <li><p>Item 4</p></li>
</ul>

<ul class="grid example-2">
  <li><p>Item 1</p></li>
  <li><p>Item 2</p></li>
  <li><p>Item 3</p></li>
  <li><p>Item 4</p></li>
  <li><p>Item 5</p></li>
</ul>

<ul class="grid example-3">
  <li><p>Item 1</p></li>
  <li><p>Item 2</p></li>
  <li><p>Item 3</p></li>
  <li><p>Item 4</p></li>
  <li><p>Item 5</p></li>
  <li><p>Item 6</p></li>
  <li><p>Item 7</p></li>
  <li><p>Item 8</p></li>
</ul>

## Irregular grids

You can create grids of unequally sized elements by passing in an array representing the relative widths of the elements in a row.

### 5,3 ratio example

    @include grid((5,3), $max-rows: 2);

<ul class="grid example-5">
  <li><p>Item 1</p></li>
  <li><p>Item 2</p></li>
  <li><p>Item 3</p></li>
  <li><p>Item 4</p></li>
</ul>

### 2,2,1 ratio example

    @include grid((2,2,1), $max-rows: 2);

<ul class="grid example-6">
  <li><p>Item 1</p></li>
  <li><p>Item 2</p></li>
  <li><p>Item 3</p></li>
  <li><p>Item 4</p></li>
  <li><p>Item 5</p></li>
  <li><p>Item 6</p></li>
</ul>

### 2,1 ratio example

    @include grid((2,1), $max-rows: 2);

<ul class="grid example-7">
  <li><p>Item 1</p></li>
  <li><p>Item 2</p></li>
  <li><p>Item 3</p></li>
  <li><p>Item 4</p></li>
</ul>


## Browser support

This Sass is supported perfectly by Chrome, Firefox, Safari and Internet Explorer 9 and higher.

In Internet Explorer 7 and 8, regular grids are fine but irregular grids use `first-child` rather
than `nth-child`. For multiple rows you'll need to pass in a `$max-rows` variable representing the
maximum number of rows in the grid.

In Internet Explorer 6, all grid elements take up the entire width of the page.
