---
layout: design-pattern
title: Check your answers pages
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
Use these at the end of a transaction to play back to users the information they have provided.

Seeing the information reassures people that the service is working and gives them an opportunity to check for errors and reflect on their answers.

<div class="example">
  <a href="http://govuk-prototype-kit.herokuapp.com/examples/check-your-answers-page">
    <img src="/service-manual/assets/images/design-patterns/check-your-answers-page.png" alt="An example of a check your answers page">
  </a>
</div>

# Guidance

Make it clear that there is a task to perform on this page and that the transaction is not complete until they confirm that the information is correct.

Provide links back to the pages containing information they can change. The user should be returned to the ‘Check your answers’ page once they’ve updated the information (and any other questions triggered by their updated response).

Don’t just write the questions and answers out as they were originally given. Reword them for this particular format and context. For example, the individual elements of an address do not need labelling, and long questions can often be rewritten as shorter statements.

Make it clear what will happen when the user clicks the button at the bottom of the screen. You don’t need an explicit declaration checkbox as long as it’s made clear what will happen.


# Implementation

There's a [coded example of this page](http://govuk-prototype-kit.herokuapp.com/examples/check-your-answers-page) in the [GOV.UK prototyping kit](https://github.com/alphagov/govuk_prototype_kit).


# Accessibility

Add hidden text to the ‘Change’ links, so that they make sense when read out of content by screen readers.


# Research

You can help by researching the following questions:

- what effect does the wording, styling and position of the ‘Change’ links have on their usage?
- for long, multi-section transactions, do we need one of these pages at the end of each section?

[Update the Hackpad with your findings](https://designpatterns.hackpad.com/Check-Your-Answers-page-2DSpTH9J0wU)

