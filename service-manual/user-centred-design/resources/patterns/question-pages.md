---
layout: design-pattern
title: Question pages
breadcrumbs:
  -
    title: Home
    url: /service-manual/index.html
  -
    title: User-centred design
    url: /service-manual/user-centred-design/index.html
  -
    title: Design patterns
    url: /service-manual/user-centred-design/resources/patterns/index.html
---

{:.intro}
Split complex forms into individual question pages.

Question pages consist of:

1. back link
2. page title
3. form elements
4. continue button

<br>

<div class="example">
  <a href="http://govuk-prototype-kit.herokuapp.com/examples/question-page">
    <img src="/service-manual/assets/images/design-patterns/question-page.png" alt="An example of a question page">
  </a>
</div>

There's a [coded example of this page](http://govuk-prototype-kit.herokuapp.com/examples/question-page) in the [GOV.UK prototyping kit](https://github.com/alphagov/govuk_prototype_kit).

---

# Guidance

For general advice on how to structure complex, multi-page forms read our [form structure guide](form-structure).

### 1. Back link

Some users don't trust browser back buttons when they're entering data (often for good reason). 
Back links reassure users that it's possible to go back and change previous answers.

It's positioned at the top of the page so users don't fill out the form, click back and lose data.


### 2. Page title

Can be in question or statement format as long as you're consistent throughout the transaction.

The question format has the advantage that it's more explicit about what the user should do.

For more complex forms you may want to include a simple [progress indicator](https://www.gov.uk/service-manual/user-centred-design/resources/patterns/progress-indicators.html) above the title.


### 3. Form elements

Follow the [GOV.UK form element styles](http://govuk-elements.herokuapp.com/form-elements/). 
For public facing services we recommend starting with [one thing per page](https://designnotes.blog.gov.uk/2015/07/03/one-thing-per-page/).


### 4. Continue button

* use 'Continue', not 'Next'
* align the button to the left so it's not missed by users 


---

# Accessibility

All HTML forms should contain at least one fieldset and therefore also one legend.
For simple question pages, the page title is often enough. You can hide the legend using the 'visuallyhidden' class. It will still be read out by screen reader users.




