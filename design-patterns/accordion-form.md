---
layout: design-patterns
title: Accordion form
subtitle: Build to the GOV.UK style 
section: guidance
subsection: Design patterns
type: resource
status: draft
page_class: accordion-form
---

One of a number of ways of breaking a long form up into a series of manageable steps..

In an accordion form, the steps are arranged vertically on a single page, but only one
step is ever 'open' at a time. As each step is completed the next one opens.


<div class="pattern-example">
  <div class="inner">

      <div class="accordion-form">
        <!-- STEP 1 -->
        <section id="step1" class="section">
          <h2 class="accordion-header"> 
            Step one
          </h2>
          <div class="accordion-content">
            <p>[Content for step one goes here]</p>
            <ul class="accordion-actions">
              <li><a href="#step2" class="next-part button">Go to question 2 of 3</a></li>
            </ul> 
          </div>
          <div class="accordion-summary">
            <p>Summary of step one <a href="#" class="edit">edit</a></p>
          </div>
        </section>

        <!-- STEP 2 -->
        <section id="step2" class="section">
          <h2 class="accordion-header"> 
            Step two
          </h2>
          <div class="accordion-content">
            <p>[Content for step two goes here]</p>
            <ul class="accordion-actions">
              <li><a href="#step3" class="next-part button">Go to question 3 of 3</a></li>
            </ul> 
          </div>
          <div class="accordion-summary">
            <p>Summary of step two <a href="#" class="edit">edit</a></p>
          </div>
        </section>

        <!-- STEP 3 -->
        <div id="step3" class="section">
          <h2 class="accordion-header"> 
            Step three
          </h2>
          <div class="accordion-content">
            <p>[Content for step three goes here]</p>
            <ul class="accordion-actions">
              <li><a href="#step4" class="next-part button">Go to ...</a></li>
            </ul> 
          </div>
          <div class="accordion-summary">
            <p>Summary of step three <a href="#" class="edit">edit</a></p>
          </div>
        </div>
      </div>

  </div>
</div>


<script type="text/javascript">
$(function() {

  // ACCORDION FORM

  // Initialise: open first part
  $(".accordion-form .section:first").addClass('current');

  // Move to next part
  $(".next-part").click(function() {
      $(this).closest(".section").removeClass('current').addClass('complete')
      .next(".section").addClass('current');
      return false;
    });

});
</script>


