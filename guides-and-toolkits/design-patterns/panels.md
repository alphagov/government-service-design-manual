---
layout: gsdm
title: Panel styles
section: guidance
subsection: Design patterns
status: draft
css: /gsdm/css/design-patterns/panels.css
---

One way to visually emphasise or highlight content is to put it in a panel.
We've created a mixin to make it easier to create panels that match the GOV.UK style.

## Default panel style

<div class="pattern-example">
  <div class="inner">
    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
    tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
    quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
    consequat.</p>

    <div class="default-panel">
      <p>Default panel style</p>
    </div>

        <p>Duis aute irure dolor in reprehenderit in voluptate velit esse
    cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
    proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>

  </div>

 <div class="side-by-side">
  <div>
    <h2>HTML</h2>
<pre><code>&lt;div class="default-panel"&gt;
  &lt;p&gt;Default panel style&lt;/p&gt;
&lt;/div&gt;
</code></pre>
  </div>
  <div>
    <h2>Sass</h2>
<pre><code>.default-panel {
  @include panel; 
}
</code></pre>
  </div>
</div>


</div>

## Parameters

The [mixin](https://github.com/alphagov/prototyping/blob/master/_includes/scss/design-patterns/_panels.scss) accepts the following arguments:

`$colour`          Background colour. Default: Light grey

`$width`           Panel width. Default: 100% (or auto if floated)

`$float`           Left, right or none. Default: none

`$icon`            Pass an image URL. Positioned top-right. Default: none

`$outdent`         True or false. Default: false

Panels are outdented by default so that their contents are vertically aligned with the surrounding contents.
If you need to you can turn outdenting off by setting $outdent to false.

## Floated panels

<div class="pattern-example">
  <div class="inner">

    <h2>Lorem ipsum</h2>

    <div class="right-panel">
      <pre>.right-panel{
  @include panel($float: right);
}</pre>
    </div>

    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
    tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
    quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
    consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
    cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
    proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>

  </div>
</div>

<div class="pattern-example">
  <div class="inner">

    <h2>Lorem ipsum</h2>

    <div class="panel-4">
      <p>50% wide.</p>
      <p>Floated left.</p>
    </div>

    <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod
    tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,
    quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo
    consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse
    cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non
    proident, sunt in culpa qui officia deserunt mollit anim id est laborum.</p>

  </div>
</div>

## Panels with icons

You can pass an image URL to the mixin and it will display in the top-right of the panel. Use this to clarify the function of the panel contents.

<div class="pattern-example">
  <div class="inner">
    <div class="info-panel">
      <p>Info panel</p>
    </div>

    <div class="download-panel">
      <p>Download panel</p>
    </div>

    <div class="calendar-panel">
      <p>Calendar panel</p>
    </div>

  </div>
</div>


## Existing panel styles

The following panel classes are available from any page. 

<div class="pattern-example">
  <div class="inner">

    <div class="application-notice">
      <p>class="application-notice"</p>
    </div>

    <div class="application-notice info-notice">
      <p>class="application-notice info-notice"</p>
    </div>

    <div class="application-notice help-notice">
      <p>class="application-notice help-notice"</p>
    </div>

    <div class="advisory minor">
      <p>class="advisory minor"</p>
    </div>

    <div class="form-download">
      <p>class="form-download"</p>
    </div>

    <div class="subscribe">
      <p>class="subscribe"</p>
    </div>

    <div class="intro">
      <p>class="intro"</p>
    </div>

    <div class="contact">
      <p>class="contact"</p>
    </div>

  </div>
</div>


* * * 

# Discussion

Q: It feels like the CSS for these needs some rationalising.
For example, why do some of the styles need two classes to work properly?

The help notice icon feels more like a warning or alert. Could we just use the info notice instead
and repurpose the help one as a warning one?



