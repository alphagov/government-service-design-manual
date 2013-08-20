---
layout: detailed-guidance
title: Forms
subtitle: Markup, styles and layout for form elements
category: design-and-development-resources
type: resource
audience:
    primary: designers, developers
    secondary:
status: draft
phases:
  - alpha
  - beta
  - live
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Design and content
    url: /service-manual/design-and-content
---

## Writing the HTML for forms

Like other components of web pages, forms should be created following the principles of [progressive enhancement](/service-manual/making-software/progressive-enhancement.html).

Browsers have default styling for forms. This is usually shared with the styling of the operating system user interface (UI), making it familiar to users. Ensure that any styling you add does not remove any of the native, highly accessible functionality offered by these defaults.

The [HTML5 specification](http://www.w3.org/TR/html51/) should be consulted for guidance on creating the HTML. This is more important than with other HTML elements as some types of user will depend on proper use of the language. For example it is important each form element has a label describing it otherwise screenreaders will not be able to identify it properly.

## User interface patterns

In the forms created so far on [GOV.UK](https://www.gov.uk) certain patterns of display have emerged as solutions to common UI problems. What follows is a guide to those patterns.

## Aligning controls in a column

The base pattern for grouping controls is in a column with one row for each control and its label.

<div class="pattern-example">
    <div class="form-example-1">

      <p class="option group">
        <label for="checkbox1"><input id="checkbox1" type="checkbox"> Job offers</label>
        <label for="checkbox2"><input id="checkbox2" type="checkbox"> Networking</label>
        <label for="checkbox3"><input id="checkbox3" type="checkbox"> Business opportunities</label>
      </p>

    </div>
</div>

## Aligning controls in a row

You might occasionally need to arrange your group of controls in a row, for instance if they have short labels and there are only a few options.

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

## Pre-checked radios and checkboxes

You may want to pre-check radios if:

* we already know the answer because it was given previously
* there is a good business reason to steer users towards a particular answer, for instance 'Contact me by email' may be preferable to 'Contact me by phone' to help manage callcentre workload
* there is a strong ‘common case’ bias towards a particular answer

Do not pre-check radios if:

* selecting none is a valid option (to be avoided for radios only, as they can’t be unchecked)
* we want an unbiased opinion without leading the user
* there is a legal requirement for the user to make a choice

## Wrapping controls in a label tag

In the examples above the controls are wrapped in their associated label tag.

The default behaviour of clicking a label will move focus to its associated form element. By wrapping a form element in its label you increase the size of the area users need to click to interact with that element to whatever size you make the label.

## Label positioning

There are valid cases for top, left or right alignment of labels. The position of an element's label does not effect how screenreaders announce it to users.

The table below (from a [great article on form design](http://uxdesign.smashingmagazine.com/2011/11/08/extensive-guide-web-form-usability/) in Smashing Magazine) outlines the relative advantages of each approach:

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

## Hidden labels

All form elements other than submission buttons should have a label. For situations where a label would not fit into the visual interface the label should be hidden from view.

Hiding labels should be done with care as by doing so you are taking information away from the user.

Labels should never be hidden by using `display: none` as this will also remove them from the [document flow](http://www.w3.org/TR/html51/dom.html#flow-content) meaning they will not be recognised by screenreaders.

In the example below the second input box has a hidden label. It is associated with the first input box visually by its position. For non-visual users it is important the label is present to provide this information.

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

## Styling text input fields

When styling form elements we should try to achieve the same goals as the default styling through our own CSS.

A few examples of how to achieve this for input fields:

* Light grey background: To make them stand out equally on a white page or coloured panel
* Inset border style: By convention people type into 'holes' cut into the interface

## Fieldsets and legends

As explained in their [HTML5 specification section](http://www.w3.org/TR/html51/forms.html#the-fieldset-element) fieldsets are used to break up forms into logical sections.

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

### Nested fieldsets

There are times when you might want to treat a set of form controls like they were a single, compound control
(a date-of-birth selector for example). One way to do this is with a nested fieldset.

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

## Hints

Hints for help with interactions can be placed above or below the relevant control.

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

## Buttons

By default buttons should be horizontally left-aligned beneath the form inputs (not necessarily left aligned with the labels, and not right aligned on the page). The primary action should be the first button in the group.

<div class="pattern-example">
        <div class="form-example-1">
        <p class="group">
          <label for="full-name">Full name</label>
          <input id="full-name" type="text" class="full-name">
        </p>
        <p class="action group">
          <input class="button" type="submit" value="Next">
          <input class="button-secondary" type="submit" value="Cancel">
        </p>

    </div>
</div>

See the [separate page on buttons](/service-manual/user-centered-design/resources/buttons.html) for more detailed guidance.

## Validation messages

Summarise any validation errors at the top of your page like this:

### Example

<div class="pattern-example">
    <div class="form-example-1">

      <div class="validation-summary">
        <h1>Please check the form</h1>
        <ul>
          <li><a href="#error1">Re-type your email address</a></li>
          <li><a href="#error2">Select at least one area of interest</a></li>
        </ul>
      </div>

    </div>
</div>

Each link should jump the user down to the corresponding form control.

The red bar visually connects the summary to the messages in the form and aids quick scanning of the form for errors.

# Examples in this page

To see all the examples above in a single form, check out the [registration form example](/service-manual/user-centered-design/resources/registration-form.html). The CSS for those styles is derived from this [SASS file](https://github.com/alphagov/government-service-design-manual/blob/master/service-manual/assets/stylesheets/design-patterns/_forms.scss).

# Examples on [GOV.UK](https://www.gov.uk)

Examples of forms currently in use across [GOV.UK](https://www.gov.uk) are:

* [Contact page](https://www.gov.uk/feedback/contact)
* [Freedom of information request form](https://www.gov.uk/feedback/foi)
* [Search results page](https://www.gov.uk/search?q=tax)
* [404 page](https://www.gov.uk/notapage)
* [Report a problem form](https://www.gov.uk/vat) (click the 'Report a problem' link to view)
* [Licence finder](https://www.gov.uk/licence-finder/sectors)
* [Trade Tariff](https://www.gov.uk/trade-tariff/sections)

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
