---
layout: design-pattern
title: Dates
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
Not all dates are the same. Choose a format appropriate for the kind of date you're asking about.


### On this page:

1. [Memorable dates](#memorable-dates)
2. [Copied dates](#copied-dates)
3. [Approximate dates](#approximate-dates)
4. [Relative dates](#relative-dates)
5. [Calendar controls](#calendar-controls)
6. [Formatting dates](#formatting-dates)
7. [Validating dates](#validating-dates)

---

<h2 class="heading-36" id="memorable-dates">1. Memorable dates</h2>


For capturing memorable dates like dates of birth, the simplest option is to provide text fields for users to type the date in.

For example:

<div class="example">
  <img src="/service-manual/assets/images/design-patterns/date-of-birth.png" alt="An example of a date of birth picker">
</div>

Calendar controls are not particularly useful for known dates and [some users struggle with select boxes](https://designnotes.blog.gov.uk/2013/12/05/asking-for-a-date-of-birth/).
We're currently recommending using three fields as it's easier to reliably validate than a single field.

Don't automatically tab between fields, as this can be unexpected behaviour that clashes with normal keyboard controls, and confuses users.

**Tip:** If you want to trigger the num-pad on iPhones, add a pattern attribute to the input element like this: `pattern="[0-9]*"`

---

<h2 class="heading-36" id="copied-dates">2. Copied dates</h2>

If you require a date to be provided EXACTLY as it's given on another document (a passport or credit card for example),
then it's easier for users to copy the date across if the format is the same.

---

<h2 class="heading-36" id="approximate-dates">3. Approximate dates</h2> 

If you don't need an exact date and users are likely to struggle to remember it (for example the date you started your first job),
make sure you allow users to enter approximate dates like 'June 1996'.

---

<h2 class="heading-36" id="relative-dates">4. Relative dates</h2> 

Some dates make most sense relative to today's date or another date.
This is particularly common when booking appointments of some kind.

In these cases it helps if you let users enter or select relative dates like 'tomorrow' or '4 days later'.
If the day of the week is important, show this as well. 

---

<h2 class="heading-36" id="calendar-controls">5. Calendar controls</h2> 

Calendar controls are only appropriate for near past or future dates, where the day of the week is relevant.
Their main use case is for appointment booking.

If you also need to show information like availability, embed the calendar in the page and make it large enough for
the information to be readable.

A calendar control that depends on JavaScript should never be the only input option.

---

<h2 class="heading-36" id="formatting-dates">6. Formatting dates</h2> 

The default date format to use is this one:

* 8 July 2014

Periods of time should be formatted like this:

* 8 July to 9 August


---

<h2 class="heading-36" id="validating-dates">7. Validating dates</h2>

When you're validating dates, remember to check the following:

* **date ranges** - check the dates are consecutive
* **past dates** - check the date is in the past
* **future dates** - check the date is in the future
* **mistyped dates** - if a date is obviously mistyped, warn users

---

[Discuss this page on Hackpad](https://designpatterns.hackpad.com/General-dates-vpx6XlVjIbE)

