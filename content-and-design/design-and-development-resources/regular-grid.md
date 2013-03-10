---
layout: detailed-guidance
title: Grids
subtitle: Quickly lay out lists of items in a grid 
section: design-and-development-resources
type: resource
phases:
  - alpha
  - beta
  - live
status: draft
page_class: grid
---

A mixin for creating a regular grid of elements.

## Arguments

The [mixin](https://github.com/alphagov/government-service-design-manual/blob/master/assets/stylesheets/design-patterns/_regular-grid.scss) accepts two arguments:

`$columns` : The number of columns in the grid, or an array representing the relative width of each column.

`$min-height` : An optional minimum height for grid elements.

This second argument is useful if your grid elements contain varied amounts of content.

Avoid applying an extra border or margin styles to the grid elements themselves as this may break the layout.


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

You can create grids of equally-sized elements by passing in a single value representing the number of elements in a row. 

The following examples are for demonstration purposes only, and not ones we'd ever recommend.

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

You can create grids of unequally-sized elements by passing in an array representing the relative widths of the elements in a row. 

### 5,3 ratio example

        @include regular-grid((5,3));

<ul class="grid example-5">
  <li><p>Item 1</p></li>
  <li><p>Item 2</p></li>
  <li><p>Item 3</p></li>
  <li><p>Item 4</p></li>
</ul>





### 2,2,1 ratio example

        @include regular-grid((2,2,1));

<ul class="grid example-6">
  <li><p>Item 1</p></li>
  <li><p>Item 2</p></li>
  <li><p>Item 3</p></li>
  <li><p>Item 4</p></li>
  <li><p>Item 5</p></li>
  <li><p>Item 6</p></li>
</ul>

### 2,1 ratio example

        @include regular-grid((2,1));

<ul class="grid example-7">
  <li><p>Item 1</p></li>
  <li><p>Item 2</p></li>
  <li><p>Item 3</p></li>
  <li><p>Item 4</p></li>
</ul>


## Cross browser support

* Chrome: All good
* IE 8: Far-right element retains right-margin, so all margins are reduced accordingly
* IE 7: Margins slightly reduced even more to account for IE7 percentage rounding bug
* IE 6: Pattern uses unsupported child selector, so all elements get 100% width
* IE 6: Bullets are showing

## Guidance

The mixin is tag-agnostic, so the elements can be list items, divs, paragraphs etc.
Avoid applying border effects to the item elements as this will throw out the widths.
Instead, style the contents of those elements.

### When to use

* When you want to split part of the page into equal columns
* When you want a grid of equally sized elements - for example, an image gallery





