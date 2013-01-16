---
layout: gsdm
title: Show/hide link, animated
section: guidance
subsection: Design patterns
status: draft
---

When you click on the link it toggles some content immediately below it. The details and summary tags in HTML5 are specifically intended for this kind of thing, but the current browser behaviour is a little clunky.

This page instead shows the type of animation which could work well.

For example, the following code: 
 
    <section>
     <p class="showhide-trigger">More information about this show/hide design pattern</p>
     <p class="showhide-description">This show/hide implementation reveals the additional information as if masked by a sheet in front. The transition DOES NOT re-flow any text, squash lines together, or reposition the content being revealed.</p>
    </section>

    <p>This is content after the show/hide design pattern. It moves up and down rather than being obscured by the description when it is shown.</p>


Is rendered on GOV.UK like this:

<div class="pattern-example">
  <div class="inner">

    <section>
      <p class="showhide-trigger">More information about this show/hide design pattern</p>
      <p class="showhide-description">This show/hide implementation reveals the additional information as if masked by a sheet in front. The transition DOES NOT re-flow any text, squash lines together, or reposition the content being revealed.</p>
    </section>

  </div>
</div>

* * *

# Dependencies

Include the following in your SCSS to activate the [show/hide mixin](https://github.com/alphagov/prototyping/blob/master/_includes/scss/design-patterns/_show-hide.scss):

    @import "frontend_toolkit";
    @import "design_patterns";
    
    @include details-and-summary;

The optional argument; `highlighted`, creates a highlight effect on the open state 
(you can also pass a colour variable or value if you don't want the default grey background).

So, for example:

    @include details-and-summary(highlighted);

gives you:

<div class="pattern-example show-hide-version-2">
  <div class="inner">

    <details>
      <summary>What happens if I click on this summary tag?</summary>
      <p>The remaining contents of the details tag are revealed inside a highlighted panel</p>
    </details>

  </div>
</div>


# Cross browser support

Cross browser support for these tags is [patchy right now](http://caniuse.com/details). However, there's a [JQuery plugin](https://github.com/mathiasbynens/jquery-details) to bridge the gap,
 which has the added bonus of adding appropriate ARIA roles to the elements. You'll need to add the following to
 your pages:


    <script src="../javascripts/jquery.details.js"></script>
    <script>
      $(function() {
        // Add conditional classname based on support
        $('html').addClass($.fn.details.support ? 'details' : 'no-details');
        // Emulate <details> where necessary and enable open/close event handlers
        $('details').details();
      });
    </script>

* * *

# Guidance

Use this pattern carefully. The objective is to declutter your interface by hiding information
that's only relevant to a small proportion of users. If a majority of users need that information, *don't hide it*.

The wording of the summary tag is critical. It must directly and unambiguously address the audience for whom the hidden information is intended.

One way we've found that's quite effective is to write the summary in the users own voice, as a question.
This forces you to try and pre-empt what it is the user will be asking themselves at that point.

Here's an example to illustrate. Imagine the user has just been asked to submit some documents via a web form:

<div class="pattern-example">
  <div class="inner">

    <details class="animated">
      <summary>What if I can't submit these documents electronically?</summary>
      <p>Don't worry, you can send in copies of your documents by post. We'll give you the address
        and reference number to use once you've paid your licence fee.</p>
    </details>

  </div>
</div>

It goes without saying that if you string a whole bunch of these together you've got yourself an FAQ page.

* * * 

# Rationale

One of the key challenges with this pattern is to help users understand that they won't be taken away from the current page if they click the link.

The 'arrow' style bullet is the closest thing we've got to a convention here. It's the default style for details/summary tags in modern browsers and has precendents in things like expanding tree diagrams, so we've decided to keep it. If we use it consistently across the site then hopefully it will start to become familiar to people.

* * * 

# Discussion

OK, so I opted for details/summary tags plus a polyfill. Anyone violently object? The thing is, if we don't use them we're going to
have to create the pattern using JavaScript and divs anyway.

TO DO: Ask Leonie is she could test drive the plugin and see if it really does what it claims to.


<script src="../javascripts/jquery.details.js"></script>

<script>
  $(function() {
    // Add conditional classname based on support
    $('html').addClass($.fn.details.support ? 'details' : 'no-details');
    // Emulate <details> where necessary and enable open/close event handlers
    $('details').details();
  });
</script>





