---
layout: gsdm
title: Show/hide link
section: guidance
subsection: Design patterns
status: draft
css: ../css/design-patterns/show-hide.css
---

When you click on the link it toggles some content immediately below it. The details and summary tags in HTML5 are specifically intended for this kind of thing.

For example, the following code: 

    <details>
      <summary>What happens if I click on a summary tag?</summary>
      <p>The remaining contents of the details tag are revealed</p>
    </details>

Is rendered on GOV.UK like this:

<div class="pattern-example">
  <div class="inner">

    <details>
      <summary>What happens if I click on a summary tag?</summary>
      <p>The remaining contents of the details tag are revealed</p>
    </details>

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

This pattern has been tested in IE6-8, FF and Chrome. It's functional in all, and no content is obscured (except intentionally).
However, the following minor issues were observed:

* IE 8: Underline extends below bullet
* IE 7: No bullet
* IE 6: No bullet or highlighted colour

* * *

# Guidance

Use this pattern carefully. The objective is to declutter your interface by hiding information
that's only relevant to a small proportion of users. If a majority of users need that information, *don't hide it*.

The wording of the summary tag is critical:

1.) It must directly and unambiguously address the audience for whom the hidden information is intended.
One way we've found that's quite effective is to write the summary in the users own voice, as a question.  
This forces you to try and pre-empt what it is the user will be asking themselves at that point.  
Here's an example to illustrate. Imagine the user has just been asked to submit some documents via a web form:  
It goes without saying that if you string a whole bunch of these together you've got yourself an FAQ page.

<div class="pattern-example">
  <div class="inner">

    <details class="animated">
      <summary>What if I can't submit these documents electronically?</summary>
      <p>Don't worry, you can send in copies of your documents by post. We'll give you the address
        and reference number to use once you've paid your licence fee.</p>
    </details>

  </div>
</div>

2.) The first few times it is used in a transaction, it should reassure users that clicking it will not take them away from the page (and lose the data they've input so far). Once users are familiar with this pattern it becomes less important to reassure them.   
One way of doing this is to start with the word 'Show'. 

Here's an example to illustrate:  

<div class="pattern-example">
  <div class="inner">

    <details class="animated">
      <summary>Show me where to find my driving licence number.</summary>
      <p>Something</p>
    </details>

  </div>
</div>

There's also an animated version of this pattern which additionally swaps the word 'show' for 'hide' when the description is open.


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





