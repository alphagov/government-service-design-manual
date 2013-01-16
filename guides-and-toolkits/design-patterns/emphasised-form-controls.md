---
layout: gsdm
title: Emphasised form controls
section: guidance
subsection: Design patterns
status: draft
css: /gsdm/css/design-patterns/emphasised-form-controls.css
---

The pattern below gives selected radio buttons and checkboxes a highlight colour and makes their clickable area larger. Use this to emphasise important decisions.

<div class="pattern-example">
  <div class="inner">
    <div class="form-example-1">

      <ul class="emphasised option group">
        <li>
          <label for="optionsRadios1">
            <input type="radio" name="LPAtypeRadios" id="optionsRadios1" value="option1">
            Create an LPA for property and financial affairs
          </label>
        </li>
        <li>
          <label for="optionsRadios2">
            <input type="radio" name="LPAtypeRadios" id="optionsRadios2" value="option2">
            Create an LPA for health and welfare
          </label>
        </li>
      </ul>

    </div>
  </div>
</div>

      <ul class="emphasised option group">
        <li>
          <label id="optionsRadios1">
            <input type="radio" name="LPAtypeRadios" id="optionsRadios1" value="option1">
            Create an LPA for property and financial affairs
          </label>
        </li>
        <li>
          <label id="optionsRadios2">
            <input type="radio" name="LPAtypeRadios" id="optionsRadios2" value="option2">
            Create an LPA for health and welfare
          </label>
        </li>
      </ul>


<script type="text/javascript">
$(function() {
  {% include javascripts/design-patterns/_emphasised_form_controls.js %}
});
</script>


