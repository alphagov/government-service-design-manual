---
layout: detailed-guidance-2
title: Forms
format: 'Design pattern'
category: design-and-development-resources
type: resource
audience:
    primary: designers, developers
    secondary:
status: draft
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Design resources
    url: /service-manual/designers
---


{{ page.content | toc_list }}


---

## Form groups

A form usually consists of one or more groups of elements, which we call form groups.
The diagram below shows a typical form group, with the different elements labelled.

<div class="example">
  <img src="/service-manual/assets/images/design-patterns/form-group.png" alt="A form group">
</div>

More complex forms may also be broken into sections or spread across multiple pages.
[Find out more about structuring forms](question-pages.html).

---

## Form labels


### Give your form elements labels

All form elements other than buttons should have a label.
Each form group should have a form group label.
More complex form groups may also need individual form control labels.

the HTML elements you use to mark up labels will depend on the structure of your form.


### Associate your labels

A HTML label must be associated with it's control using the 'for' attribute or by nesting the control inside the label.


### Don't hide labels

Don't hide labels unless the surrounding context makes them unnecessary. Never use `display: none` as this will stop the label being read out by screenreaders. Instead, use the 'visuallyhidden' class provided with the GOV.UK templates.


### Write clear label text

* Use sentence case
* Give polite, clear, short instructions
* Don't use 'please'
* Don't use colons at the end


### Indicate optional fields in their label

Mark any optional fields in the labels with '(optional)'.
This avoids repetitive '*' symbols, which then have to be explained.

If you need to clarify that fields are mandatory, add text at the top of the form saying 'all fields are required unless stated otherwise'.


### Position labels near their controls

We recommend labels be positioned above their controls, however there maybe valid cases for other label alignments.

The table below (from an [article on form design](http://uxdesign.smashingmagazine.com/2011/11/08/extensive-guide-web-form-usability/) in Smashing Magazine) outlines the relative advantages of each approach:

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

Note that the position of an element's label does not effect how screenreaders announce it to users.


---

## Text fields

Text fields should be sized according to their content. 
Doing this gives people an visual cue as to the kind of content they're being asked for. For example, a postcode field should be shorter than an email address field.

Avoid using placeholder text inside text fields. It confuses some people and the text disappears once a user starts typing into the field. If extra help is needed, use [hint text](#form-hints).


---

## Checkboxes and radio buttons

Radio buttons let users choose a single option from a list. 
Check boxes let users select on/off toggles or multiple selections.


### Make it clear whether it's multiple or single choice

Research indicates many users do not know the difference between radios and checkboxes. They will decide whether it's multiple or single choice based on the question and answer text instead, so word these carefully.


### Take care when pre-selecting options

Pre-selecting options can help users, but it needs to be done with care.

Do pre-select options if:

* you're playing back something the user has already told you
* there's a good reason to steer users towards a particular answer
* there's a strong common case bias towards a particular answer

Do not pre-select options if:

* there's a legal requirement for the user to make a choice
* you don't want to bias users towards making a particular choice


### Don't forget 'none of the above'

For radio buttons, if 'none of the above' is a valid response remember to represent it explicitly. A set of radio buttons cannot be reset to an unselected state once it's been interacted with.


### Highlight the clickable region

Enlarging the clickable region and making it visible can make these controls easier to use,
especially on touch devices.

<div class="example">
  <img src="/service-manual/assets/images/design-patterns/block-radio.png" alt="Radio buttons with an enlarged and visible clickable region">
</div>

Some people are unaware that they can click the label to select these controls. They try to click the control itself, which is difficult as they're very small.


---

## Drop-down lists (select boxes)

Try to find an alternative to these.

Research has found that some people struggle when using these.
Test with your users and use alternatives where possible.

### Short lists

For short lists, radio buttons are often preferable. They make all the options immediately visible which can help with comprehension.

### Long lists

For long lists, consider a free-text field with an auto-suggest feature, or try breaking the list down into shorter lists.


---

## Form hints

Place form hints immediately above their control, unless this makes the form harder to understand.
In those cases, place the hint below the control.
On GOV.UK form hints are coloured grey (Sass variable: $secondary-text-colour).

Use the 'aria-describedby' attribute to associate form hints with the controls they describe.

[Find out more about ways of providing help for users](help-text.html).


---

## Buttons

<div class="example">
  <img src="/service-manual/assets/images/design-patterns/buttons.png" alt="Primary, disabled and secondary button styles">
</div>


### Buttons are for actions

Use buttons to represent actions.
Don't use multiple buttons to represent choices.
Instead, use radio buttons or similar.


### Use buttons sparingly

Buttons are powerful, but only when used sparingly - one button per form is ideal.


### Make buttons obvious

Position buttons in the user's line of sight.
The bottom left of a form, aligned with the controls is good.


### Make buttons describe their action

Keep button text clear and concise. Avoid jargon and words like 'submit'.
Use verbs and an active voice - buttons should describe the thing they do.
As elsewhere on GOV.UK, use sentence case.

Here are some common examples:

'Start now'
: Use at the beginning of the service

'Next'
: Not 'forward', 'go' etc.

'Finish'
: Use at the end of a service before the final confirmation screen

'Back'
: Not 'previous'

'Sign in' and 'Sign out'
: Not 'log-in' or 'log-out'

'Sign up'
: Use this for mailing lists

'Create an account'
: Not 'register'


---

## Validation messages

Summarise validation errors at the top of your page.
Each error should link the user down to the corresponding form control.
The form control should repeat the error and be styled similarly.
Error messages should be actionable - the user should be told how to fix the error.

---

## Coding forms

Use the [HTML5 specification](http://www.w3.org/TR/html51/) and [progressive enhancement](/service-manual/making-software/progressive-enhancement.html) when coding forms. This will ensure that they're fully accessible across the widest range of browsers and devices.

Use the HTML form elements wherever possible. Avoid inventing your own versions.

---

[Discuss this page on Hackpad](https://designpatterns.hackpad.com/GOV.UK-design-patterns-0eUk1OdHvql#:h=Form-elements)

