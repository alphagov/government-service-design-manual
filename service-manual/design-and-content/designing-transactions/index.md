---
layout: detailed-guidance
title: Designing transactions
subtitle: Desiging useful and usable government services
category: content-and-design
type: guide
audience:
  primary: designer
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
    title: Content and design
    url: /service-manual/content-and-design
---

## What is a transaction?

A transaction is an exchange between two or more parties.

![A transaction is an exchange between two or more parties](transactions.png)

Government transactions typically involve an exchange of information, money, rights or some combination of these. Usually one party will be a citizen or business and the other will be the government.

### An example: applying for a passport

When someone [applies for a passport](https://www.gov.uk/apply-renew-passport), they exchange information and money, in the form of a passport application form and fee, for the right to travel, in the form of a passport.


![Applying for a passport](passport-transaction.png)

## Digital transactions

Most digital transactions take the form of a dialogue between a user and a service, with the service acting as a proxy for the government.

![A transaction is a dialogue](transaction-dialogue.png)

A dialogue is typically made out of a series of questions and answers, mediated via an interface, often a website. Questions are usually rendered as form controls (checkboxes, a text field, date picker, button etc.) but might use something else entirely (a map for example).

How these questions are structured and ordered and how users are allowed to answer (or ask) the questions needs careful consideration.

## How to structure the questions

You need to choose a structure for your transaction that supports the questions you need to ask. It can help to think in terms of levels of structure: group, sub-section, section etc. Try not to worry about how those levels should be represented until you have a broader understanding of what questions you need to ask.

For example, you might choose to have one section per page, or multiple sections. It depends on 



### Option 1: Put all the sections on a single page


#### The good

* Only one submit button to press
* A single URL gives access to all form fields
* Suggests, but doesn't force a fixed order of completion
* Users benefits from context of surrounding form elements
* Progress is self-evident

#### The bad

* Long forms can be overwhelming and off-putting
* Not well suited to branching or non-linear processes
* How do you save partial progress?
* Harder to track analytics like drop-off rates - you'll need to use Javascript


### Option2 : Put each section on it's own page

#### The good

* Easier to handle branching and dependencies between sections
* Easier to let user save progress
* Can make a long transaction feel more manageable to the user
* Can help guide a user through an unfamiliar process
* Easier to capture analytics like drop-off rates for each section 

#### The bad

* Harder for users to orient themselves within the transaction
* It can slow users down as they have to click and load each section
* Users lose the contextual information from neighbouring sections
* Harder for users to review and edit previous sections
* No single URL for users to go and edit their data
* Not a natural fit for non-linear processes like looping, adding and removing

### Option 3: All sections on a page, but only one visible at a time

#### The good

* Can handle branching and dependencies between sections
* Can easily review and edit previous questions