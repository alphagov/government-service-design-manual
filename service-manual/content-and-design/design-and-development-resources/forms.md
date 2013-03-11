---
layout: detailed-guidance
title: Forms
subtitle: Markup, styles and layout for form elements
category: design-and-development-resources
type: resource
audience:
    primary: designer, developer
status: draft
phases:
  - alpha
  - beta
  - live
---

## Styling form elements

Forms should be styled as per the examples on this page. The [GOV.UK forms mixin](https://github.com/alphagov/government-service-design-manual/blob/master/assets/stylesheets/design-patterns/_forms.scss) provides a configurable framework for styling your forms in this way. Use it in your Sass like this:

### Sass

    @import "forms"

    // Default style with top-aligned labels
    .form-1 { 
      @include form 
    }

    // Left aligned labels
    .form-2 { 
      @include form(left)
    }

    // Right aligned labels. Lable width is 9em
    .form-3 { 
      @include form(right, 9em)
    }


Check out the [registration form example](registration-form.html) to see the different layout options in action.


## Basic structure

Wrap each control in an element with a class of 'group'.

### Example

<div class="pattern-example">
    <div class="form-example-1">
      <p class="group">
        <label for="label">Label</label>
        <input id="label" type="text">
      </p>
    </div>
</div>

### HTML

    <p class="group">
      <label for="label">Label</label>
      <input id="label" type="text">
    </p>

## Radios and checkboxes

Wrap your set of options in an 'option group' element.

### Example

<div class="pattern-example">
    <div class="form-example-1">

      <p class="option group">
        <label for="checkbox1"><input id="checkbox1" type="checkbox"> Job offers</label>
        <label for="checkbox2"><input id="checkbox2" type="checkbox"> Networking</label>
        <label for="checkbox3"><input id="checkbox3" type="checkbox"> Business opportunities</label>
      </p>

    </div>
</div> 

### HTML

Use nested inputs...

      <p class="option group">
        <label for="checkbox1">
          <input id="checkbox1" type="checkbox"> Job offers
        </label>
        <label for="checkbox2">
          <input id="checkbox2" type="checkbox"> Networking
        </label>
        <label for="checkbox3">
          <input id="checkbox3" type="checkbox"> Business opportunities
        </label>
      </p>


or use a list...

      <ul class="option group">
        <li>
          <input id="checkbox1" type="checkbox">
          <label for="checkbox1"> Job offers</label>
        </li>
        <li>
          <input id="checkbox2" type="checkbox">
          <label for="checkbox2"> Networking</label>
        </li>
        <li>
          <input id="checkbox3" type="checkbox">
          <label for="checkbox3"> Business opportunities</label>
        </li>
      </ul>


## Aligning controls in a row

You might occasionally need to arrange form controls in a row. To do this, wrap the controls in an
'inline group' element.

### Example

<div class="pattern-example">
    <div class="form-example-1">
        <fieldset>
          <legend><span>Gender</span></legend>
          <p class="inline option group">
            <label for="male"><input id="male" type="radio" name="gender"> Male</label>
            <label for="female"><input id="female" type="radio" name="gender"> Female</label>
          </p>
        </fieldset>
  </div>
</div>

### HTML

        <fieldset>
          <legend><span>Gender</span></legend>
          <p class="inline option group">
            <label for="male">
              <input id="male" type="radio" name="gender"> Male
            </label>
            <label for="female">
              <input id="female" type="radio" name="gender"> Female
            </label>
          </p>
        </fieldset>

## Fieldsets and legends

Use these to break up forms into logical sections

### Example

<div class="pattern-example">
    <div class="form-example-1">

        <fieldset>
          <legend>Name</legend>
          <p class="group">
            <label for="first-name">First name</label>
            <input id="first-name" type="text" class="name">
          </p>
          <p class="group">
            <label for="last-name">Last name</label>
            <input id="last-name" type="text" class="name">
          </p>
        </fieldset>

    </div>
</div>

### HTML

        <fieldset>
          <legend>Name</legend>
          <p class="group">
            <label for="first-name">First name</label>
            <input id="first-name" type="text" class="name">
          </p>
          <p class="group">
            <label for="last-name">Last name</label>
            <input id="last-name" type="text" class="name">
          </p>
        </fieldset>


### Nested fieldsets

There are times when you might want to treat a set of form controls like they were a single, compound control 
(a date-of-birth selector for example). One way to do this is with a nested fieldset. On GOV.UK, when you
nest a fieldset inside another, the legend is styled like a label.

<div class="application-notice info-notice">
<p>Note - if you're planning on doing this with left or right aligned form labels you'll need to wrap your
legend text in a span. Blame inconsistent and generally poor support for legend positioning in browsers.</p>
</div>

### Example

<div class="pattern-example">

    <div id="form-example-2" class="left">

      <fieldset>
        <legend>Your details</legend>
        <p class="group">
          <label for="full-name">Full name</label>
          <input id="full-name" type="text" class="full-name">
        </p>
        <fieldset>
          <legend><span>Date of birth</span></legend>
          <div class="inline group">
            <p class="group">
              <label for="day" class="visuallyhidden">Day</label>
              <select id="day" type="text">
                <option value="Day">Day</option>
                <!-- Other options go here -->
              </select>
            </p>
            <p class="group">
              <label for="month" class="visuallyhidden">Month</label>
              <select id="month" type="text">
                <option value="Month">Month</option>
                <!-- Other options go here -->
              </select>
            </p>
            <p class="group">
              <label for="year" class="visuallyhidden">Year</label>
              <select id="year" type="text">
                <option value="Year">Year</option>
                <!-- Other options go here -->
              </select>
            </p>
          </div>
        </fieldset>
      </fieldset>

  </div>

</div>

### HTML

      <fieldset>
        <legend>Your details</legend>
        <p class="group">
          <label for="full-name">Full name</label>
          <input id="full-name" type="text" class="full-name">
        </p>
        <fieldset>
          <legend><span>Date of birth</span></legend>
          <div class="inline group">
            <p class="group">
              <label for="day" class="visuallyhidden">Day</label>
              <select id="day" type="text">
                <option value="Day">Day</option>
                <!-- Other options go here -->
              </select>
            </p>
            <p class="group">
              <label for="month" class="visuallyhidden">Month</label>
              <select id="month" type="text">
                <option value="Month">Month</option>
                <!-- Other options go here -->
              </select>
            </p>
            <p class="group">
              <label for="year" class="visuallyhidden">Year</label>
              <select id="year" type="text">
                <option value="Year">Year</option>
                <!-- Other options go here -->
              </select>
            </p>
          </div>
        </fieldset>
      </fieldset>



## Hints

Hints can be placed above or below the relevant control.

### Example

<div class="pattern-example">
    <div class="form-example-1">

      <p class="group">
        <label for="telephone">Telephone</label>
        <input id="telephone" type="text">
        <span class="hint">Include your country code</span>
      </p>

      <p class="group">
        <label for="code">Code</label>
        <input id="code" type="text" class="small">
        <span class="inline hint">The three numbers on the back of the card</span>
      </p>

      <p class="group">
        <label for="password">Password</label>
        <span class="hint">Make it at least six characters long</span>
        <input id="password" type="password">
      </p>

    </div>
</div> 

### HTML

      <p class="group">
        <label for="telephone">Telephone</label>
        <input id="telephone" type="text">
        <span class="hint">Include your country code</span>
      </p>

      <p class="group">
        <label for="code">Code</label>
        <input id="code" type="text" class="small">
        <span class="inline hint">The three numbers on the back of the card</span>
      </p>

      <p class="group">
        <label for="password">Password</label>
        <span class="hint">Make it at least six characters long</span>
        <input id="password" type="password">
      </p>


## Hidden labels

Use the 'visuallyhidden' class to hide labels. You need a really good reason to do this though.

### Example

<div class="pattern-example">
    <div class="form-example-1">

      <p class="group">
        <label for="street1">Street</label>
        <input id="street1" type="text" class="street">
      </p>
      <p class="group">
        <label for="street2" class="visuallyhidden">Street line two</label>
        <input id="street2" type="text" class="street">
      </p>
      <p class="group">
        <label for="town">Town/City</label>
        <input id="town" type="text" class="town">
      </p>
      <p class="group">
        <label for="postcode">Postcode</label>
        <input id="postcode" type="text" class="postcode">
      </p>

    </div>
</div> 

### HTML

      <p class="group">
        <label for="street1">Street</label>
        <input id="street1" type="text" class="street">
      </p>
      <p class="group">
        <label for="street2" class="visuallyhidden">Street line two</label>
        <input id="street2" type="text" class="street">
      </p>
      <p class="group">
        <label for="town">Town/City</label>
        <input id="town" type="text" class="town">
      </p>
      <p class="group">
        <label for="postcode">Postcode</label>
        <input id="postcode" type="text" class="postcode">
      </p>


## Buttons

Nest rows of buttons in an 'action group' element.

### Example

<div class="pattern-example">
    <div class="form-example-1">

        <p class="action group">
          <input class="button" type="submit" value="Submit">
          <input class="button-secondary" type="submit" value="Cancel">
        </p>

    </div>
</div> 

### HTML

        <p class="action group">
          <input class="button" type="submit" value="Submit">
          <input class="button-secondary" type="submit" value="Cancel">
        </p>


See the [seperate page on buttons](buttons.html) for detailed guidance on how to style and word them.

## Validation messages

Summarise any validation errors at the top of your page like this:

### Example

<div class="pattern-example">
    <div class="form-example-1">

      <div class="validation-summary">
        <h1>Please check the form</h1>
        <ul>
          <li><a href="#error1">Confirm your email address</a></li>
          <li><a href="#error2">Select at least one area of interest</a></li>
        </ul>
      </div>

    </div>
</div> 

### HTML

      <div class="validation-summary">
        <h1>Please check the form</h1>
        <ul>
          <li><a href="#error1">Confirm your email address</a></li>
          <li><a href="#error2">Select at least one area of interest</a></li>
        </ul>
      </div>


Each link should jump the user down to the corresponding form control. Add a 'validation' class to the control group and insert a 'validation-message' element.

### Example

<div class="pattern-example">
    <div class="form-example-1">

      <p class="validation group">
        <span class="validation-message" id="error1">Confirm your email address</span>
        <label for="email-confirm">Confirm email <abbr title="Mandatory">*</abbr></label>
        <input id="email-confirm" type="text" class="email">
      </p>

    </div>
</div> 

### HTML

      <p class="validation group">
        <span class="validation-message" id="error1">Confirm your email address</span>
        <label for="email-confirm">Confirm email <abbr title="Mandatory">*</abbr></label>
        <input id="email-confirm" type="text" class="email">
      </p>


# Cross browser support

* Chrome: No problems
* IE 8: V-aligned - all good
* IE 8: H-aligned - slightly ragged alignment
* IE 7: Nested fieldsets - legends are hidden in H-aligned forms
* IE 7: Validated option groups - lots of extra left padding
* IE 6: Option groups are aligning horizontally in vertically-aligned forms
* IE 6: Nested fieldsets layout is broken and legends are hidden
* IE 6: Labels with nested controls need 'for' attributes to work

# Putting it all together

To see all the examples above in a single form, check out the [registration form example](registration-form.html).

# Rationale

## Text input fields

* Light grey background: To make them stand out equally on a white page or coloured panel
* Inset border style: By convention people type into 'holes' cut into the interface

## Label positioning

The framework provides support for top, left or right alignment because there are valid cases for the use of all three. The table below (from a [great article on form design](http://uxdesign.smashingmagazine.com/2011/11/08/extensive-guide-web-form-usability/) in Smashing Magazine) outlines the relative advantages of each approach:

|-----------------------------------|--------------|-------------------|---------------|
|                                   | Top          | Right             | Left          |
|-----------------------------------|--------------|-------------------|---------------|
| Speed of completion               | Fastest      |                   | Slowest       |
| Horizontal space required         | Least        |                   | Most          |
| Vertical space required           | Most         |                   | Least         |
| Label text space available        | Most         |                   | Least         |
| Proximity to input                | Closest      |                   | Least close   |
| User eye movement                 | Down         | Down + right      | Down + right  |
| Time to move from label to input  | 50ms         | 240ms             | 500ms         |
| Ideal for...                      | Simple forms | Less simple forms | Complex forms |

## Validation messages

When a form is submitted, any validation messages are summarised at the top of the page.
The messages link down to the part of the form they relate to.
This helps users of assistive technology navigate around the form.

The red bar connects the summary to the messages in the form and aids quick scanning of the form for errors.

<script>
  $(function() {

    // Style toggle for pattern examples
    // Takes the text of the clicked 'option' and assigns it as
    // a class to the element named in the 'data-for' attribute
    $('.class-toggle .option').click(function(){
      $('.class-toggle .option').removeClass('selected');
      $(this).addClass('selected');
      var selectedClass = $(this).text();
      var selectedElement = "#" + $(this).parents('.class-toggle').data("for");
      $(selectedElement).removeClass().addClass(selectedClass);
      return false;
    });



  });
</script>
