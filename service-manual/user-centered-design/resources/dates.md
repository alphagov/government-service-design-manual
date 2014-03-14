---
layout: detailed-guidance
title: Dates
subtitle: 
category: design-and-development-resources
type: resource
audience:
  primary: content-designers, designers
type: guide
status: draft
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: User centered design
    url: /service-manual/user-centered-design
---


## Choose the right date picker

There are lots of ways to ask users to enter dates. 
It depends on the kind of date they're entering.

### Is this a date that most people would expect to have in their heads? 

For example: date of birth, date of their child's birth, another memorable anniversary. 

If so, it's likely to be easier for them to type it into free text fields.

[See our guidance on capturing known dates](#capturing-known-dates).


### Is this a date that they will need to copy from another document? 

For example: start date of a passport, expiry date of a credit card. 

If so, it's likely to be easier for them to copy the exact same format as the format used in the other document.


### Is this a date that they may only remember approximately? 

For example: date of starting a job (if several years ago), date of a doctor's appointment (if a few weeks ago).

If so, make sure that you allow for approximate dates.


### Is this a date that makes most sense relative to today's date? 

For example: booking an appointment in the future. 

If so, provide an optional calendar, and let users enter values like 'today' or 'tomorrow'.


### Is this a date that makes most sense relative to another date? 

For example: departure and return dates from a trip. 

If so, consider offering options like 'a week later' for the later date?

---

## Capturing known dates

Our current recommendation for capturing known dates like date of birth is to use three clearly labelled text boxes.

<div class="example">
  <img src="/service-manual/assets/images/design-patterns/date-of-birth.png" alt="An example of a date of birth picker">
</div>

We've found that some users struggle with long select boxes, which is why we're not recommending you use them.
See [this blog post](https://designnotes.blog.gov.uk/2013/12/05/asking-for-a-date-of-birth/) for a more detailed account of our findings.

Don't use calendar controls for dates of birth. 
They're more appropriate for recent past or future dates, or where the day of the week is relevant.

**Tip:** If you want to trigger the num-pad on iPhones, add a pattern attribute to the input element like this: `pattern="[0-9]*"`

---

## Date formats

Periods of time when used in questions or help text should follow the format in this example: 
8 July to 9 August.

---

## Validating dates

When asking for date ranges make sure the validation forces the user to enter consecutive dates or show an error message.
Include days of the week to help the user with this. 

When asking for past or future dates make sure validation starts from today’s date. 
Show an error message if user input isn’t in the past or future.

---

[Discuss this page on Hackpad](https://designpatterns.hackpad.com/General-dates-vpx6XlVjIbE)

