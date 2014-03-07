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

## Date of birth

Our current recommendation for capturing date of birth is to use three clearly labelled text boxes.

<div class="example">
  <div class="ribbon">Recommended</div>
  <form class="form">
    
    <fieldset class="date-of-birth">
      <legend style="font-weight: bold">Date of birth</legend>
      <div class="form-group">
        <label for="day">Day</label>
        <input type="text" pattern="[0-9]*" id="day" class="form-control">
      </div>
      <div class="form-group">
        <label for="month">Month</label>
        <input type="text" id="month" class="form-control">
      </div>
      <div class="form-group form-group-year">
        <label for="year">Year</label>
        <input type="text" pattern="[0-9]*" id="year" class="form-control">
      </div>
    </fieldset>
    <p class="hint">Eg. 21/03/1976</p>
  </form>
</div>

<h3 class="heading-24">Guidance</h3>

A date of birth is personal information, so don't ask for it unless you absolutely have to. For example, you don't need to ask for full date of birth if you only need to know:

* a person's current age
* a persons year of birth
* whether a person falls within a particular age range

Avoid using calendar controls for dates of birth. They're more appropriate for recent past or future dates, or where the day of the week is relevant.

We've found that some users struggle with long select boxes, which is why we're not recommending you use them.
See [this blog post](https://designnotes.blog.gov.uk/2013/12/05/asking-for-a-date-of-birth/) for a more detailed account of our findings.

**Tip:** If you want to trigger the num-pad on iPhones, add a pattern attribute to the input element like this: `pattern="[0-9]*"`


## Other dates

When asking for date ranges make sure the validation forces the user to enter consecutive dates or show an error message. You can add days of the week to help the user with this. 

Example:<br/>
Enter the dates when you were abroad<br/>
From [day] [month] [year] to [day] [month] [year]

Validation example: 09/03/2013 to 12/11/2012<br/>
Error message: Your end date must be after the start date.

When asking for past or future dates make sure validation starts from today’s date. Show an error message if user input isn’t in the past or future. 

Periods of time when used in questions or help text should follow the format in this example: 
8 July to 9 August.

Send any other labels you’re using to [hinrich.von-haaren@digital.cabinet-office.gov.uk](mailto: hinrich.von-haaren@digital.cabinet-office.gov.uk).

