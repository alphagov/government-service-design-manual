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
Use these to play back to users the information they have provided, before they submit it.

Status: [Research required](#research)

<br>

<div class="example">
  <a href="http://govuk-prototype-kit.herokuapp.com/examples/check-your-answers-page">
    <img src="/service-manual/assets/images/design-patterns/check-your-answers-page.png" alt="An example of a check your answers page">
  </a>
</div>

The two main benefits of summary pages are:

- **increased completion rates** -- users are reassured that the correct data has been captured so are less likely
to drop out of the transaction before completing it
- **reduced error rates** -- users are given a second chance to spot and correct errors before submitting the data


---

# Code

There's a [coded example of this page](http://govuk-prototype-kit.herokuapp.com/examples/check-your-answers-page) in the [GOV.UK prototyping kit](https://github.com/alphagov/govuk_prototype_kit).


---

# Guidance

Make it clear that there is a task to perform on this page and that the transaction is not complete until they confirm that the information is correct.

Provide links back to the pages containing information they can change. The user should be returned to the ‘Check your answers’ page once they’ve updated the information (and any other questions triggered by their updated response).

Don’t just write the questions and answers out as they were originally given. Reword them for this particular format and context. For example, the individual elements of an address do not need labelling, and long questions can often be rewritten as shorter statements.

The submit button should fully state the action it performs, eg ‘Change your tax details’ or ‘Send your claim form’. You don’t need an explicit declaration checkbox as long as it’s made clear what will happen.


---

# Accessibility

Add hidden text to the ‘Change’ links, so that they make sense when read out of content by screen readers.


---

# Security

These pages play back information entered by your users. In some cases you might need to obfuscate parts of the data to reduce the risk that it’s read by the wrong people.

Each service is different. Work with your security team to make sure that there’s an acceptable level of risk.


---

# Research

You can help improve this pattern by researching the following:

- what effect does the wording, styling and position of the ‘Change’ links have on their usage?
- for long, multi-section transactions, do people need one of these pages at the end of each section, or just at the end of the whole transaction?

[Update the Hackpad with your findings](https://designpatterns.hackpad.com/Check-Your-Answers-page-2DSpTH9J0wU)

