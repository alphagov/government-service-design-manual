---
layout: design-pattern
title: Start pages
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
Start pages are how people access transactions on GOV.UK. All services on GOV.UK must have a start page.


Start pages consist of:

1. the name of the service
2. a short introduction
3. a call to action button
4. any additional information

<div class="example">
  <div class="inner-block">

    <img src="/service-manual/assets/images/design-patterns/start-page.png" alt="Example of a start page" />

  </div>
</div>


### 1. Service name

The name of the service should tell people what it lets them do and who it's for.

Avoid acronyms and noun phrases. The best service names start with a verb, like "Pay your car tax".



### 2. Short introduction

If you need to, include a few lines of SEO friendly introductory text.



### 3. Call to action button

This button text should contain a clear call to action to start the service.

If the service is not hosted on GOV.UK then this should be made clear.


### 4. Additional information

Tell people about:

* information or documents they will need to use the service
* anything that will help them decide whether to use the service (eg how much it costs, how long it takes)
* other ways that people can access the service

There’s no need to mention something that most users will already know or assume.

Don't try to get across complex eligibility criteria. Instead, deal with these inside the service itself.


## Examples on GOV.UK


* [Register to vote start page](https://www.gov.uk/register-to-vote)


## Implementing start pages

There's a [start page template](http://govuk-prototype-kit.herokuapp.com/examples/start-page) in the [GOV.UK Prototype Kit](https://github.com/alphagov/govuk_prototype_kit). This will help you test your service end-to-end during the alpha phase.

When you're close to your beta launch, [contact the GOV.UK content team](https://support.production.alphagov.co.uk/new_feature_request/new) to ask for a start page.

You should tell them:

+ what the service is and who it's for
+ what user needs it addresses
+ eligibility criteria and restrictions
+ what documents or information the user will need to use the service
+ a preview of the service (if available)
+ other methods of applying eg phone and post
+ a demo URL (with a username and password if necessary)
+ when the service is expected to pass its beta assessment

Factor in time it'll take for the content to be written and fact-checked -
the start page can only be published after both:

+ the service passes its [beta assessment](https://www.gov.uk/service-manual/phases/beta.html) or self-certification
+ the service goes live

---

[Discuss this page on Hackpad](https://designpatterns.hackpad.com/Transaction-start-pages-8fitVQYufJX)