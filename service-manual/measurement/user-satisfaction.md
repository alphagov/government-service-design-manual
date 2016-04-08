---
layout: detailed-guidance
title: User satisfaction
subtitle: How satisfied are the people who use your service?
category: measurement
type: guide
audience:
  primary: service-managers, performance-analysts
  secondary: designers, content-designers
phases:
  - beta
  - live
status: draft
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Measurement
    url: /service-manual/measurement
exclude_from_search: true
---

Measuring user satisfaction helps you to gauge the overall quality of your service.

Tracking user satisfaction can help make sure the changes you make are improving the service for users. If your user satisfaction falls, then by asking for feedback, you can investigate to find out why.

## How to measure user satisfaction

1. Get a [Feedback page](/service-manual/user-centred-design/resources/patterns/feedback-pages.html) on GOV.UK -- [request a content change](https://support.production.alphagov.co.uk/)
2. Add links from your service to the feedback page (more details below)
3. Use the [Feedback Explorer](https://support.production.alphagov.co.uk/anonymous_feedback/explore) to see feedback on your service

![A Feedback page on GOV.UK](/service-manual/assets/images/feedback-page.png)

## Measure satisfaction across the whole service

More often than not, the end of the transaction isn’t the end of the service. For example, if you claim Carer’s Allowance, the end of the transaction means you’ve finished filling in your claim. You’re still waiting for a decision.

You must prompt users to give feedback at service endpoints. In the example above, this means prompting the user to give feedback when they get their benefit decision. This could be some time after they finish their transaction -- the prompt could be in an email or a letter.

You must also allow users to give feedback from anywhere in the service, in case anything goes wrong.

And you must be able to show that you’re collecting user satisfaction data appropriately at your [Service Assessment](/service-manual/digital-by-default/assessments-at-gds.html).

## Prompt users to give feedback at service endpoints

Add a feedback link to all your service endpoints, for example:

* Your Carer’s Allowance claim has been successful
* You are not eligible for Carer’s Allowance
* Error -- we couldn’t find your National Insurance number

## What pattern to use

We recommend the following pattern to link to the [Feedback page](/service-manual/user-centred-design/resources/patterns/feedback-pages.html):

> [What did you think of this service?]() (takes 30 seconds)

> `<p><a href=“https://www.gov.uk/done/name-of-your-service”>What did you think of this service?</a> (takes 30 seconds)</p>`

If the service endpoint is an email or a letter, use the following pattern:

> What did you think of this service? (takes 2 minutes)

> Visit www.gov.uk/done/name-of-your-service

## Allow users to give feedback from anywhere in the service

Add a ‘Feedback’ link to the footer on every page of your service.

You can also link to the Feedback page from your alpha or beta banner.

This is a courtesy to your users -- if they have something to tell you, they should be able to do so. And it will help make our services better.

## Further reading

* [Discuss this page on Hackpad](https://designpatterns.hackpad.com/User-satisfaction-and-user-feedback-zfk4wUpWNj3)
* Blog post: [How good is your service? How many users give up? (August 2015)](https://designnotes.blog.gov.uk/2015/08/13/how-good-is-your-service-how-many-users-give-up/)
