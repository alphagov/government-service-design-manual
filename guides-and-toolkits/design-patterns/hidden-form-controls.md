---
layout: gsdm
title: Hidden form controls
section: guidance
subsection: Design patterns
status: draft
css: /gsdm/css/design-patterns/hidden-form-controls.css
---

This pattern uses the details and summary tags to reveal content when a link or form label is clicked.

## Example 1: Links

<div class="pattern-example">
  <div class="inner">
    <div class="form-example-1">

        <p class="group">
          <label for="label">Label</label>
          <input id="label" type="text">
        </p>                                      
        <details>
          <summary class="action group">Extra options</summary>
          <p class="group">
            <label for="label2">Label</label>
            <input id="label2" type="text">
          </p>
        </details>

    </div>
  </div>
</div>

## Example 2: Checkboxes

<div class="pattern-example">
  <div class="inner">
    <div class="form-example-1">

      <ul class="hidden-controls option group">
        <li>                                      
          <details>
            <summary>
              <label for="details"><input id="details" type="checkbox"  value="value"> I'd like to provide a detailed report...</label>
            </summary>
            <label class="visuallyhidden">Enter text here</label>
            <textarea class="full-size" placeholder="Enter text here"></textarea>
          </details>
        </li>
      </ul>

    </div>
  </div>
      <pre><code>  &lt;ul class="hidden-controls option group"&gt;
    &lt;li&gt;                                      
      &lt;details&gt;
        &lt;summary&gt;
          &lt;label&gt;&lt;input type="checkbox"  value="value"&gt; I'd like to provide a detailed report...&lt;/label&gt;
        &lt;/summary&gt;
        &lt;label class="visuallyhidden"&gt;Please provide details&lt;/label&gt;
        &lt;textarea class="full-size" placeholder="Enter text here"&gt;&lt;/textarea&gt;
      &lt;/details&gt;
    &lt;/li&gt;
  &lt;/ul&gt;
</code></pre>
</div>

## Example 3: Radio buttons

<div class="pattern-example">
  <div class="inner">
    <div class="form-example-1">

      <ul class="hidden-controls option group">
        <li>
            <label for="optionsRadios1">
              <input type="radio" name="optionsRadios" id="optionsRadios1" value="option1">
              All decisions must be made jointly
            </label>
        </li>
        <li>
            <label for="optionsRadios2">
              <input type="radio" name="optionsRadios" id="optionsRadios2" value="option2">
              All decisions can be made jointly or severally
            </label>
        </li>
        <li>                                      
          <details>
            <summary>
              <label for="optionsRadios3"><input type="radio" name="optionsRadios" id="optionsRadios3" value="option3"> It depends on the decision…</label>
            </summary>
            <label class="visuallyhidden">Please provide details</label>
            <textarea class="full-size" placeholder="Please provide details"></textarea>
          </details>
        </li>
      </ul>
    </div>
  </div>
</div>

## Dependencies

You'll need to import 'forms.scss' and 'show-hide.scss' and include the 'hidden-form-controls' and 'details-and-summary' mixins.

You'll also need the relevant javascripts on your page.


## Cross browser

* FF, Safari, IE 6-8: With radios, once a hidden control has been opened it stays open,
even if another radio is subsequently checked

* All browsers: If you have more than one hidden control in a set of radio buttons it gets confused.

<script src="../javascripts/jquery.details.js"></script>
<script>
  $(function() {
    // Add conditional classname based on support
    $('html').addClass($.fn.details.support ? 'details' : 'no-details');
    // Emulate <details> where necessary and enable open/close event handlers
    $('details').details();


    {% include javascripts/design-patterns/_hidden_form_controls.js %}


  });
</script>


