---
layout: detailed-guidance
title: Question and answer
subtitle: A pattern for identifying specific information
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
    title: Design and content
    url: /service-manual/design-and-content
---

## User needs

Hide the complexity of information & show only what is asked for by the user.

## Introduction

Explain what the tool can be used for, any criteria users need to fit to use it and any other information needed to prevent unnecessary journeys.

If the process is longer or more complex than normal, it may be necessary to say how much time is needed to complete it.

## Stages in the process

Each step should be numbered to show where it is in the overall process.

Every question answered should be represented in the user interface (UI). This should show the question, the answer given and links or buttons to allow the user to change that answer.

The UI should allow the user to return to the beginning of the process.

When asking users for information, make sure any restrictions around its format are made clear in the question. For example dates may have be in DD/MM/YYY format or amounts limited to a range of values.

## Presenting the result

Show the result along with options of how what actions should be taken next.

## Retaining state

As you go through the process it's important that progress so far is retained. Depending on the nature of the process and any restrictions it has, you can do this differently:

## HTTP state

HTTP requests/responses with a GET method store state in the URL. 

Clean URLs can be achieved through HTTP redirects so responses are stripped of their query string and represent the state in their path instead.

For example

      https://www.gov.uk/student-finance-calculator/y/?response=2012-2013

should be redirected to

      https://www.gov.uk/student-finance-calculator/y/2012-2013

## Advantages of HTTP state

Users can navigate between steps by using the back and forward controls.

Each step of the process is able to be bookmarked by browsers so users can return to a process at any point.

Clean URLs are more easily remembered and written.

Browsers will retain the values entered in form fields for any page with a unique URL. By maintaining clean URLs for each step of the process, the values entered for each step will be preserved.

## Disadvantages of HTTP state

Once you navigate to a previous step all subsequent steps will be lost.

The complexity of answers can be limited by the need to maintain clean URLs.

## Examples on GOV.UK

[https://www.gov.uk/student-finance-calculator](https://www.gov.uk/student-finance-calculator)
