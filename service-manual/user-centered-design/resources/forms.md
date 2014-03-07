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
    title: User centered design
    url: /service-manual/user-centered-design
---

## General guidance


### Writing forms

* Use sentence case
* Address the user directly
* Follow the style guide
* Don’t invent proper nouns, eg don’t say ‘Check your Identity Profile’


### Coding forms

Use the [HTML5 specification](http://www.w3.org/TR/html51/) and [progressive enhancement](/service-manual/making-software/progressive-enhancement.html) when coding forms. This will ensure that they're fully accessible across the widest range of browsers and devices.

Avoid changing the default appearance or behaviour of form elements without good reason. Users will be familiar with the defaults and familiarity aids comprehension.


## Form labels

All form elements other than buttons should have a label. A label must be associated with it's control using the 'for' attribute or by nesting the control inside the label.

### Positioning labels

By default, labels should be positioned above their control, but there are valid cases for left or right alignment of labels too. The table below (from an [article on form design](http://uxdesign.smashingmagazine.com/2011/11/08/extensive-guide-web-form-usability/) in Smashing Magazine) outlines the relative advantages of each approach:

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

The position of an element's label does not effect how screenreaders announce it to users.


### Hiding labels

Only hide labels if people can still infer how to use the control from it's context. Don't use `display: none` as this will stop the label being read out by screenreaders.


### Writing label text

* Don't use 'please', but give polite, clear, short instructions
* Don't use colons after labels
 


### Indicating mandatory and optional fields

You should be asking users for the minimum amount of information you need to run your service, 
so nearly all your fields should be mandatory.

Rather than pepper the interface with repetative '*' symbols, which then have to be explained, 
simply mark the optional fields in the labels with '(optional)'.

If you need to, add text at the top of the form saying 'all fields are required unless stated otherwise'.


## Text fields

Text fields should be sized according to their content. For example, a postcode field should be shorter than an email address field.
Doing this gives people an extra cue as to the kind of content they're being asked for.

Avoid using placeholder text inside text fields. It confuses some people and the text dissapears once they start typing into the field.


## Checkboxes and radio buttons

Use radio buttons to let users choose a single option from a list. Use checkboxes for on/off toggles or for multiple selections.

Research indicates many users do not know the difference between radios and checkboxes. They will decide whether it's multiple or single choice based on the question and answer text instead, so word these carefully.

### Pre-selecting options

Do not pre-select options if:

* there's a legal requirement for the user to make a choice
* you don't want to bias users towards making a particular choice

Do pre-select options if:

* we already know the answer because it was given previously
* there's a good reason to steer users towards a particular answer
* there's a strong common case bias towards a particular answer

For radio buttons, if 'none of the above' is a valid response remember to represent it explicitly. A set of radio buttons cannot be reset to an unselected state once it's been interacted with.

### Highlighting the clickable region

Some people are unaware that they can click the label to select these controls. They try to click the control itself, which is difficult as they're very small. To help these people you can enlarge the clickable region and make it visible.


## Select boxes

Research is showing that some people struggle to use select boxes. Test with your users and use alternatives where possible.

For short lists, radio buttons are often preferable. They make all the options immediately visible which can help with comprehension.

For long lists, consider a free-text field with an auto-suggest feature, or try breaking the list down into shorter lists.


## Help text

First try to design a service that's so intuitive it doesn't need help text.
You shouldn't need help text to explain the interface - it should be simple enough to understand without explanation.
Help text should be contextual - it should relate to the user's current situation.
Here are some different ways of providing help text.


### Form hints

Form hints are short pieces of text positioned next to a form control.
Use form hints to:

* explain where to find the information being asked for
* give an example to show the format of the information
* explain what will be done with the information

By default, form hints should be placed immediately above the associated control. However, you can position it below the control if it aids overall comprehension. On GOV.UK form hints are coloured grey (Sass variable: $secondary-text-colour).


### Hidden help text

If the help text is only relevant to a minority of users you can keep your interface less cluttered by hiding the text behind a summary which the user has to click on. Be careful not to hide important content from people or force them to interact needlessly with the page.

Remember - if most people will need the text, don't hide it.

Summary text tends to work better when it's targetted, rather than the generic 'Help' or 'Info'. For example:

* 'I can't find my National Insurance number'
* 'What's this?'


### Detailed help

Some transactions are complex and some users will dive in to a transaction without the required knowledge to complete it.

You need to support these users by allowing them to access more detailed guidance from within the transaction,
without them losing their place. If you need to open a new browser window remember to tell people you're going to do this.


## Buttons

[ Primary button, secondary button, disabled button ]

Use buttons to represent actions. Use them sparingly - one button per form is ideal.

Position buttons in the user's line of sight. The bottom left of a form, aligned with the controls is good.

### Button text

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


## Validation messages

Summarise validation errors at the top of your page.
Each error should link the user down to the corresponding form control.
The form control should repeat the error and be styled similarly.
Error messages should be actionable - the user should be told how to fix it.


