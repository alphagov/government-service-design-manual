---
layout: gsdm
title: Progress indicator
section: guidance
subsection: Design patterns
status: draft
css: /gsdm/css/design-patterns/progress-indicator.css
---

Use this near the top of multi-page transactions to help users understand how
many steps there are, which step they are on and what's involved in each step.


### Default example
<div class="pattern-example">
  <div class="inner">

    <nav role="navigation" class="progress-indicator group">
      <span class="visuallyhidden step-count">Step 1 of 4</span>
      <ol class="group">
        <li class="active">Current step</li>
        <li>Step two</li>
        <li>Step three</li>
        <li>Step four</li>
      </ol>
    </nav>

  </div>
</div>

### Code


    <nav role="navigation" class="progress-indicator group">
      <span class="visuallyhidden step-count">Step 1 of 4</span>
      <ol class="group">
        <li class="active">Current step</li>
        <li>Step two</li>
        <li>Step three</li>
        <li>Step four</li>
      </ol>
    </nav>



# Dependencies

Include the following in your SCSS to pull in styles for this pattern:

    @import "frontend_toolkit";
    @import "design_patterns";
    
    @include progress-indicator;

The [progress-indicator mixin](https://github.com/alphagov/prototyping/blob/master/_includes/scss/design-patterns/_progress-indicator.scss) accepts three optional arguments:

`$active-colour` : The colour of the current step

`$border-colour` : The border colour

`$background-colour` : The background colour


So for example, this code:

    @include progress-indicator($green, $green, $white);

produces the following:

<div class="alt pattern-example">
  <div class="inner">

    <nav role="navigation" class="progress-indicator group">
      <span class="visuallyhidden step-count">Step 1 of 4</span>
      <ol class="group">
        <li class="active">Current step</li>
        <li>Step two</li>
        <li>Step three</li>
        <li>Step four</li>
      </ol>
    </nav>

  </div>
</div>

Note that the text colour will switch automatically from dark to light, depending on the background colours chosen.

* * *

# Rationale

* This pattern is best used to represent progress through a fixed number of steps
* The steps are numbered to indicate the fixed order in which they must be completed
* The arrow motif is used to emphasise the goal of progressing through the steps
* Steps are *not* links:
  * Their affordance isn't strong enough to rely on them as links
  * In many cases we won't be able to let users jump back and forth through steps anyway

* * * 

# Discussion

Q: We're wrapping this pattern in a <nav> tag. Is this right? There are no links so you can't actually
navigate to anywhere.

Q: We need to consider a small-screen version - how about we hide the steps and just show the step count?

