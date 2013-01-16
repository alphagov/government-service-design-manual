---
layout: gsdm
title: Regular grid
section: guidance
subsection: Design patterns
status: draft
css: /gsdm/css/design-patterns/regular-grid.css
---

A mixin for creating a regular grid of elements. All widths are expressed as percentages of the parent element.

# Arguments

The [mixin](https://github.com/alphagov/prototyping/blob/master/_includes/scss/design-patterns/_regular-grid.scss) accepts two arguments:

`$columns` : The number of columns in the grid

`$min-height` : An optional minimum height for grid elements

This second argument is useful when your grid elements contain varied amounts of content.


## Four column example

<ul class="regular-grid example-1">
  <li><p>Item 1</p></li>
  <li><p>Item 2</p></li>
  <li><p>Item 3</p></li>
  <li><p>Item 4</p></li>
  <li><p>Item 5</p></li>
  <li><p>Item 6</p></li>
  <li><p>Item 7</p></li>
  <li><p>Item 8</p></li>
</ul>

<div class="side-by-side">
  <div>
    <h3>HTML</h3>
<pre><code>&lt;ul class="regular-grid example-1"&gt;
  &lt;li&gt;&lt;p&gt;Item 1&lt;/p&gt;&lt;/li&gt;
  &lt;li&gt;&lt;p&gt;Item 2&lt;/p&gt;&lt;/li&gt;
  &lt;li&gt;&lt;p&gt;Item 3&lt;/p&gt;&lt;/li&gt;
  &lt;li&gt;&lt;p&gt;Item 4&lt;/p&gt;&lt;/li&gt;
  &lt;li&gt;&lt;p&gt;Item 5&lt;/p&gt;&lt;/li&gt;
  &lt;li&gt;&lt;p&gt;Item 6&lt;/p&gt;&lt;/li&gt;
  &lt;li&gt;&lt;p&gt;Item 7&lt;/p&gt;&lt;/li&gt;
  &lt;li&gt;&lt;p&gt;Item 8&lt;/p&gt;&lt;/li&gt;
&lt;/ul&gt;
</code></pre>
  </div>
  <div>
    <h3>Sass</h3>
    <pre><code>.regular-grid p{
    background-color: $highlight-colour;
    border: 1px solid $border-colour;
    text-align: center;
    padding: 1em;
}

.example-1{
  @include regular-grid(4);
}
</code></pre>
  </div>
</div>

## Varying columns

As you can see below, you can vary the number of columns and the margin widths will be maintained. The margin width approximates to the standard GOV.UK margin when the grid is placed in a 'wide' GOV.UK page template.

<ul class="regular-grid example-4">
  <li><p>Item 1</p></li>
  <li><p>Item 2</p></li>
</ul>

<ul class="regular-grid example-1">
  <li><p>Item 1</p></li>
  <li><p>Item 2</p></li>
  <li><p>Item 3</p></li>
  <li><p>Item 4</p></li>
</ul>

<ul class="regular-grid example-2">
  <li><p>Item 1</p></li>
  <li><p>Item 2</p></li>
  <li><p>Item 3</p></li>
  <li><p>Item 4</p></li>
  <li><p>Item 5</p></li>
</ul>

<ul class="regular-grid example-3">
  <li><p>Item 1</p></li>
  <li><p>Item 2</p></li>
  <li><p>Item 3</p></li>
  <li><p>Item 4</p></li>
  <li><p>Item 5</p></li>
  <li><p>Item 6</p></li>
  <li><p>Item 7</p></li>
  <li><p>Item 8</p></li>
</ul>

# Cross browser support

* Chrome: All good
* IE 8: Far-right element retains right-margin, so all margins are reduced accordingly
* IE 7: Margins slightly reduced even more to account for IE7 percentage rounding bug
* IE 6: Pattern uses unsupported child selector, so all elements get 100% width
* IE 6: Bullets are showing

# Guidance

The mixin is tag-agnostic, so the elements can be list items, divs, paragraphs etc.
Avoid applying border effects to the item elements as this will throw out the widths.
Instead, style the contents of those elements.

# When to use

* When you want to split part of the page into equal columns
* When you want a grid of equally sized elements - for example, an image gallery

* * * 

# Discussion

Once the parent element width shrinks below a certain size the layout should probably switch to a reduced number of columns rather than continue to shrink down.





