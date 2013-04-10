---
layout: detailed-guidance
title: Buttons
subtitle: How to code them, how to word them and when to use them   
category: design-and-development-resources
type: resource
audience:
    primary: designers
    secondary: 
status: draft
phases:
  - alpha
  - beta
  - live
page_class: buttons
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Design and content
    url: /service-manual/design-and-content
---

Buttons should be used to signify actions that the user can perform. Here's how to create buttons in the GOV.UK style, using the [GOV.UK button Sass mixin](https://github.com/alphagov/government-service-design-manual/blob/master/service-manual/assets/stylesheets/design-patterns/_buttons.scss).

## Button styles

The button mixin styles can be applied to links, inputs and button tags, like this:

### Example

<div class="pattern-example">
  <p>
    <button class="button">Button tag</button>
    <a class="button">Link tag</a>
    <input class="button" type="submit" value="Input tag" /> 
  </p>
</div>

### HTML

    <p>
      <button class="button">Button tag</button>
      <a href="#" class="button">Link tag</a>
      <input class="button" type="submit" value="Input tag" /> 
    </p>

### Sass

    @import "buttons"

    .button{
     @include button;
    }

## Button colours

The default button colour is `$green`, but different colours can be assigned. You need a good reason to do this though. 

### Example

<div class="pattern-example">
  <p>
    <a href="#" class="button">Primary action</a> 
    <a href="#" class="button-secondary">Secondary action</a> 
    <a href="#" class="button-warning">Warning action</a>
  </p>
</div>

The button text colour automatically switches from light to dark, depending on the background colour.

### HTML

    <p>
      <a href="#" class="button">Primary action</a> 
      <a href="#" class="button-secondary">Secondary action</a> 
      <a href="#" class="button-warning">Warning action</a>
    </p>

### Sass

    .button{
     @include button;
    }
    .button-secondary{
      @include button($grey-3);
    }
    .button-warning{
     @include button($red);
    }

## Disabling buttons

Use the 'disabled' attribute or class, depending on which kind of element you're styling.

### Example

<div class="pattern-example">
  <p>
    <button class="button" disabled="disabled">Button tag</button>
    <a class="disabled button">Link tag</a>
    <input class="button" disabled type="submit" value="Input tag" /> 
  </p>
</div>

### HTML

    <p>
      <button class="button" disabled="disabled">Button tag</button>
      <a class="disabled button">Link tag</a>
      <input class="button" disabled type="submit" value="Input tag" /> 
    </p>


## Button sizes

Buttons will inherit the font size of their parent elements. Use the standard paragraph text size wherever possible. More complex interfaces may occasionally require smaller buttons.

### Example

<div class="pattern-example">
  <p>
    <a href="#" class="button">Primary action</a> 
    <a href="#" class="button-secondary">Secondary action</a> 
    <a href="#" class="button-warning">Warning action</a>
  </p>
  <div>
    <a href="#" class="button">Primary action</a> 
    <a href="#" class="button-secondary">Secondary action</a> 
    <a href="#" class="button-warning">Warning action</a>
  </div>
  <p>
    <a href="#" class="x-small button">Primary action</a> 
    <a href="#" class="x-small button-secondary">Secondary action</a> 
    <a href="#" class="x-small button-warning">Warning action</a>
  </p>
</div>

## Types of button

### Primary actions
<div class="pattern-example">
  <p>
    <a href="#" class="button">Next step</a>
  </p>
</div>

* Primary actions move the user on to the next part of the transaction
* Avoid having multiple primary actions on a single page

### Secondary actions
<div class="pattern-example">
  <p>
    <a href="#" class="button-secondary">Save</a>
  </p>
</div>

* Secondary actions modify the current view
* They don't move users to the next step
* There can be multiple secondary actions per page
* They should be less prominent than the primary action

### Warning actions
<div class="pattern-example">
  <p>
    <a href="#" class="button-warning">Delete account</a>
  </p>
</div>

* Actions that have irreversable effects should look 'scary'
* Keep them away from the other actions
* Use an alert to check that the user really wants to do this
* Even better, make the action reversible

### Launch button
<div class="pattern-example">
  <p>
    <a href="#" class="button" rel="external" title="Get started on the HMRC website">Get started</a> 
     on the HMRC website
  </p>
</div>

* Use to initiate a transaction
* Let users know if they'll be taken to a different website


## Writing button text

Do use:

* verbs and an active voice
* clear, informative and succinct language
* sentence case

Don't use:

* technical terms - eg. use ‘delete’, not ‘form reset’, use one of the common actions below instead of ‘submit’
* lots of words - eye tracking shows users are less likely to read long labels


### Common actions

Many of these can be used in place of ‘Submit’, which is a technical term to be avoided.

'Start now'
: Use at the beginning of the service

'Next'
: not 'forward', 'go' etc.

'Finish'
: Use at the end of a service before the final confirmation screen 

'Back'
: Not 'previous'

'Sign in' and 'Sign out'
: Not 'log-in' or 'log-out'

'Create an account'
: Not 'register'

'Sign up'
: Only use this for mailing lists


### Compound actions

Sometimes you want a single button to perform more than one action. For example, 'Save and quit'.

It's worth trying to avoid this situation but if you can't, use common sense. If one of the actions is obvious or not important to know, don't mention it.

For example, if a button saves the current state and moves the user to the next screen, don't use 'Save and next', just use 'Next', because users will assume the former. When in doubt, test with real users.




