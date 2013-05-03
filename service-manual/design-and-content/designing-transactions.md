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

It's an exchange between two or more parties. Government transactions typically involve an exchange of information, money, rights, goods or some combination of these. Usually one party will be a citizen or business and the other will be the government.

For example, when someone applies for a passport they exchange information and money for the right to travel.

!['Diagram showing 'applying for a passport' transaction'](designing-transactions/passport-transaction.png)

A service is made up of a collection of related or connected transactions.

Most digital transactions take the form of a dialogue between a user and a system, which acts on behalf of the other party. It's the system's interface which determines in large part the quality of the overall user experience.

## Choosing the right interface

The dialogue between a user and a system can be realised in different ways. It might resemble a verbal conversation, with the system asking questions and the user answering them. It might be more visual, with the user querying the system by interacting with a map or diagram.

Either way, you'll need to decide what kind of interface best fits your transaction and how to structure it.

## How to structure the transaction

You should choose a structure for your transaction that most naturally fits the ways your users are going to want to use it.

Ask yourself:

* Will they want to move through the transaction in a fixed order, or one of their choosing?
* Will they be able to complete the transaction in a single go?
* Will their answers affect other parts of the transaction?
* Will they want to go back and review or change answers to previous questions?
* Will they need to add or remove items from a list, or change the order of things?
* How many parties are involved in the transaction?
* At what point is the transaction regarded as complete?

How you answer these questions will help you decide how to structure the transaction. It can help to think in terms of levels: sections, subsections, groups etc. Try not to worry about how those levels should be represented in the interface until you have a broader understanding of the overall structure.

For example, on a web site, you might choose to have one section per page, or multiple sections. It depends on how people are going to want to use it.

## Example structures

### Option 1: Single page

All sections are positioned on a single page.

{:.left}
![Diagram showing all sections on a page](designing-transactions/one-page.png)

#### **The good**

* There's only one submit button to press
* A single URL gives access to all form fields
* It doesn't force a fixed order of completion
* You benefit from context of neighbouring sections
* Progress is self-evident

#### **The bad**

* Long forms can be overwhelming and off-putting
* It's less well suited to branching or non-linear flow
* How do you save partial progress?
* Can be harder to track analytics like drop-off rates

--- 

### Option2 : Wizard

Each section goes on it's own page.

![Diagram showing each section on it's own page](designing-transactions/wizard.png)

#### **The good**

* It's easier to handle branching and dependencies between sections
* It's easier to let the user save progress
* A long transaction can feel more manageable
* Easier to guide a user through an unfamiliar process
* Easier to capture analytics like drop-off rates for each section 

#### **The bad**

* Can be harder for users to see where they are within the transaction
* It can slow users down as they have to click and load each section
* You lose the contextual information from neighbouring sections
* Harder for users to review and edit previous sections
* There's no single place for users to go back and edit their data
* Not a natural fit for non-linear processes like looping, adding and removing

---

### Option 3: Accordion form

All sections on a single page, but each new section only appears once the previous section has been completed.

![Diagram showing an accordion form](designing-transactions/accordion-3.png)

#### **The good**

* Can handle branching and dependencies between sections
* Can easily review and edit previous questions
* Can help guide a user through an unfamiliar process
* User still benefits from some surrounding context
* Progress is clear

#### **The bad**

* Implementation and interface is more complex

Done well, option 3 is a hybrid of the other two that has benefits of both the other options.
Within this hybrid option there are still some important design decisions to make. For example:

* Will future questions be shown in any way or will you only see the questions you've answered?
* What happens if you go back and edit a previous question?
    * Does the current question stay open or closed?
    * How do you get back to the current question once you've edited a previous one?
    * Do you lose all your answers to questions that follow the one you go back to edit?





## Saving progress

If the average time to complete a transaction is more than you can reasonably expect your users to spend in a single session, then you'll need to provide a way for them to save their progress.

The same goes if the session is likely to be interrupted for some reason. For example, if the user is suddenly asked for information which they might not have immediately to hand (a way to mitigate against this is to warn users up front if they're going to be asked for that kind of information).

Saving progress does not necessarily mean you require user accounts, logins, email validation etc. For simpler transactions that don't store personal information you might be able to store the data in the URL itself. The user then simply has to bookmark that URL.

For more complex transactions that don't store personal data you might be able to offer users a unique, and hard-to-guess URL that they can use to get back to their session.








