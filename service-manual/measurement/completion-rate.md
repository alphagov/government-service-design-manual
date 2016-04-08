---
layout: detailed-guidance
title: Completion rate
subtitle: How many people manage to use the service successfully?
status: draft
category: measurement
type: guide
audience:
  primary:
  secondary: service-managers, performance-analysts
phases:
  - beta
  - live
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Measurement
    url: /service-manual/measurement
exclude_from_search: true
---

Completion rate shows how many people manage to use the service successfully.

Once your service goes live, you’ll find steps that cause users to give up and bugs that make the service impossible to complete. As you fix these issues, your completion rate will rise.

## How to measure completion rate

1. Count the number of completed transactions
2. Divide it by the number of started transactions, as a percentage

## What is a completed transaction?

A transaction is completed when **the service did what it was supposed to do**.

Often this will be because the user was successful in achieving their desired outcome, for example:

* A user claimed Carer’s Allowance
* A user paid their vehicle tax

Sometimes, even though **the service did what it was supposed to do**, the user won’t be successful. These still count as completed transactions. For example:

* A user couldn’t claim Carer’s Allowance because they weren’t eligible
* A user couldn’t pay their vehicle tax because they have already paid

In all of these cases, the user will reach a [transaction end page](https://designpatterns.hackpad.com/Transaction-end-pages-xkOPGx6R1iM).

You should be able list all your transaction end pages at your [Service Assessment](https://www.gov.uk/service-manual/digital-by-default/assessments-at-gds.html).

## What is a failed transaction?

There are two reasons why transactions are counted as failed:

* The user gave up before reaching a transaction end page
* The user reached an error page

For example:

* A user couldn’t pay their vehicle tax because of a database error
* A user abandoned trying to submit their Self Assessment tax return

You should be able list all your error pages at your [Service Assessment](https://www.gov.uk/service-manual/digital-by-default/assessments-at-gds.html).

## How to set up web analytics

Use web analytics to measure your completion rate:

* in Google Analytics -- use [destination goals](https://support.google.com/analytics/answer/1116091) and funnels
* in Piwik -- use [goals](http://piwik.org/docs/tracking-goals-web-analytics/)
* in WebTrends -- use [scenario analysis](http://help.webtrends.com/en/analytics10/#configuring_scenario_analysis.html)

Make sure each step in the service has a unique URL so you can identify them in your analytics package.

Your goals/scenarios should always start from the GOV.UK start page. Don’t make it possible for users to bypass your service start page via links or search engine results. Users who try to access another page in the service should be sent back to your start page, unless they are resuming a previously saved transaction.

Your goals/scenarios should end with the endpoints where **the service did what it was supposed to do.** Define a separate goal/scenario for each one, then combine the completion scores to give an overall completion rate.

**Your completion rate should never be greater than 100%.**

## Saving and resuming progress

Some services allow users to save a transaction mid-flow and resume another time.

Match saved transactions with resumed transactions so they can be treated as one continuous process.

To do this, make separate goals/scenarios for users who are resuming using the transaction. These don’t count as starts, but do count as completions if they complete.

Apply a time limit to saved transactions after which, if they haven’t been resumed, they are classed as failed.

## Non-digital channels and steps

Some services have both digital and non-digital channels or steps. For example, when granting a lasting power of attorney, users can start and finish the transaction online, but are required to print, sign and post a form in the middle of the process.

In these situations treat the second online part of the transaction as a user continuing on from the first part. Where it’s not possible to match the two, treat the single digital parts of the service as separate tasks, with the completion rate measured for each.

It may not be practical to directly measure completion rate for non-digital channels or steps. For something like a paper form, usability testing will be more effective.

## Further reading

Blog post: [How good is your service? How many users give up? (August 2015)](https://designnotes.blog.gov.uk/2015/08/13/how-good-is-your-service-how-many-users-give-up/)
