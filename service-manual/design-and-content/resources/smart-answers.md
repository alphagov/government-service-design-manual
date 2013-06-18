---
layout: detailed-guidance
title: Smart answers
subtitle: Building questions to get the right answers from complex content
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

Part of the content strategy of [GOV.UK](https://www.gov.uk) is to hide the complexity & show only what the user needs.

With complicated subjects, if they are new to the user, the need may be unclear as they may not know what to ask. We can build up what they need to ask from values obtained by asking smaller, more focused questions.

For example, according to the regulations on towing vehicles:

<div class="example">
<p>If you have a <strong>category B car</strong> and <strong>your licence has certain entitlements on it</strong> you can tow any trailers as long as their weight isn't over that of the car.</p>
</div>

In the above example the values are

* category B car
* your licence has certain entitlements on it

By asking the category of car the user owns and the entitlements on their licence we get both the answer and an explanation of the question required.

## Starting pages

Each smart answer should have a start page. This should explain what the tool can be used for, any criteria users need to fit to use it and any other information needed to prevent unnecessary journeys.

If the process is longer or more complex than normal, it may be necessary to say how much time is needed to complete it.

## Stages in the process

Each step should be numbered to show where it is in the overall process.

Every question answered should be represented in the user interface (UI). This should show the question, the answer given and links or buttons to allow the user to change that answer.

The UI should allow the user to return to the beginning of the process.

When asking users for information, make sure any restrictions around its format are made clear in the question. For example dates may have be in DD/MM/YYY format or amounts limited to a range of values.

## Presenting the result

Show the resulting answer along with options of how what actions should be taken next.

## Retaining state

As you go through the process it's important that progress so far is retained. Depending on the nature of the process and any restrictions it has, you can do this differently:

## HTTP state

State can be stored in the URL by using HTTP requests/responses using a GET method.

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

## Preventing indexing

Applications of this type should have a meta tag to prevent search engines from indexing their content.

      <meta name="robots" content="noindex">

Because each potential state in a process has a URL it will be indexable, meaning user journeys could begin from any step in the process.

It is important users have a view of the whole process when they arrive at the application. If journeys begin at a page in the process there is risk that all steps will not be visible which could effect the relevance of the result.

## What doesn't match the format

There will be some types of information are not suitable for this type of navigation.

It is important 'Real' content is indexable so if the information that will end up in your results is more than a few simple facts it may not be suitable.

## Further reading

[Blog post on Smart answers](http://digital.cabinetoffice.gov.uk/2012/02/16/smart-answers-are-smart/)

## Examples on GOV.UK

[https://www.gov.uk/student-finance-calculator](https://www.gov.uk/student-finance-calculator)
