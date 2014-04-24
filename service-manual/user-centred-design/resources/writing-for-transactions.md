---
layout: detailed-guidance
title: Writing for transactions
subtitle: Writing microcopy and help text for government services
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
    title: User centred design
    url: /service-manual/user-centred-design
---

Microcopy is the term given to the short words or phrases used during transactions,
for things like buttons, form labels, help text, alerts and questions.

### General guidance

* Use sentence case
* Address the user directly
* Follow the style guide
* Don't use 'please' on labels - but give polite, clear, short instructions
* Don't use colons after labels
* Don't add 'your' to labels, for example: 'Your name', 'Your address', unless you have a page with multiple people where you need to differentiate - for example: 'your name', 'partner name'

### Asking for personal details

* 'First name' not 'Christian name'
* 'Middle names'
* 'Last name' not 'Family name'
* 'Age'
* 'Date of birth' (not 'DOB')

## Writing help text

Your first strategy when it comes to help text is to design a service that's
so intuitive it doesn't need any. For this reason it helps to stick to interface design conventions where possible. Avoid innovation for its own sake - the real innovation is an
easy to use government service.

Sometimes though, users need a little help. Here are some ways of providing it.

### Inline help

<div class="pattern-example">
    <div class="form-example-1">

      <p class="group">
        <label for="telephone">Telephone</label>
        <input id="telephone" type="text">
        <span class="hint">Include your country code</span>
      </p>

    </div>
</div>

Use this to provide examples for unfamiliar information requests or formats.


### Progressive disclosure

This refers to help that appear on a page when the user interacts with a link, but remains hidden otherwise. It's useful for delivering important help to some users, without distracting or confusing everyone else. For this reason it's particularly useful for dealing with edge case user scenarios.


### Pop-up help window

This is the nuclear option and should only be considered as a last resort. Popups or lightboxes should ONLY be used for delivering help relating to the concepts or terminology involved in a service or transaction. If you're using one to explain how to use a service then you need to go back and make the interface more intuitive.


## Transaction start pages

The start point for any GOV.UK transaction should be a page on the GOV.UK domain. Users should not be able to jump to a later page in the service via some other means (eg Google).

The design of the start page will be determined by the nature of the service and its audience. All start pages should meet the following goals:

### Describe the service

The page should include the name of the service, expressed as an action if possible ('Renew your passport', Claim for disability allowance', 'Book a driving test' etc). If you need to, include a very brief description of the service.

### Set expectations

People will tend to base their expectations of what a transaction involves on their experience of other digital transactions, for example, online shopping and banking.

If your service has features that do not match these expectations then you should inform your users of this as early on as possible, ideally on the start page.

Users should be told about any financial costs or long waiting periods involved. If they'll be asked to provide relatively obscure information let them know so they can get it ready.

If there are specific eligibility requirements for the service let people know. If the eligibility requirements are complex you should consider using the GOV.UK smart answer format to help people understand whether they're eligible.

It's better to ask a few questions up front (and explain why you're doing this) than to let people invest time and effort in a transaction only to discover part way through that they're not eligible to use it.

## Transaction end pages

The end point of any transaction should be a page on the GOV.UK domain.

These pages should let the user know:

* that they have successfully completed the transaction
* what further actions they need to perform, and when
* what further actions they can expect from the service, and when
* who they should contact with any queries or complaints
* about any related information and services

##Further reading
[Information on designing forms that work](http://www.formsthatwork.com/TheArtOfWritingVeryLittle) on the 'Forms that work' site

