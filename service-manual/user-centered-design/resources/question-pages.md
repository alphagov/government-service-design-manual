---
layout: detailed-guidance
title: Question pages
subtitle: Choosing an appropriate structure for your transactions
category: user-centered-design
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
    title: User centered design
    url: /service-manual/user-centered-design
---

Typically transactions will contain one or more pages of questions, where users enter information.


## Writing good questions

Ask for as little information as you can.
Keep the questions short and straightforward. 
Explain why your asking a question if it's not obvious.

Your questions don’t have to be questions, you can also use statements, eg ‘What is your date of birth?’ or ‘Date of birth’.
Just remember to be consistent across the transaction.


## Structuring the questions

It's important to choose a structure for the questions that most naturally fits the ways your users are going to want to use your service.

## Things to think about

* how many questions are there?
* will users need to answer the questions in a fixed order?
* will they need to complete the transaction in a single session?
* will their answers affect other parts of the transaction?
* will certain answers result in the user having to abandon the transaction?
* will they need to review and edit their answers?
* will they need to add or remove items from a list, or change the order of items?
* how many parties are involved in the transaction?
* do any parts of the transaction take place offline?
* at what point is the transaction regarded as complete?


## Example structures

Thinking about the above will help you decide on an appropriate structure.
Here are some of the more common approaches:

### 1. Single page of questions

All questions are placed on a single page.

{:.left}
![Diagram showing all sections on a page](/service-manual/assets/images/designing-transactions/one-page.png)

#### The good

* There's only one submit button to press
* A single URL gives access to all form fields
* It doesn't force a fixed order of completion
* You benefit from context of neighbouring sections
* Progress is self-evident

#### The bad

* Long forms can be overwhelming and off-putting
* It's less well suited to branching or non-linear flow
* How do you save partial progress?
* Can be harder to track analytics like drop-off rates

---

### 2. One question per page

Each question goes on its own page.

![Diagram showing each section on it's own page](/service-manual/assets/images/designing-transactions/wizard.png)

#### The good

* It's easier to handle branching and dependencies between questions
* It's easier to let the user save progress
* A long transaction can feel more manageable
* Easier to guide a user through an unfamiliar process
* Easier to capture analytics like drop-off rates for each section

#### The bad

* Can be harder for users to see where they are within the transaction
* It can slow users down as they have to click and load each section
* You lose the contextual information from neighbouring sections
* Harder for users to review and edit previous sections
* There's no single place for users to go back and edit their data
* Not a natural fit for non-linear processes like looping, adding and removing

---

### 3. Accordion form

All questions on a single page, but each new question only appears once the previous section has been completed.

![Diagram showing an accordion form](/service-manual/assets/images/designing-transactions/accordion-3.png)

#### The good

* Can handle branching and dependencies between sections
* Can easily review and edit previous questions
* Can help guide a user through an unfamiliar process
* User still benefits from some surrounding context
* Progress is clear

#### The bad

* Implementation and interface is more complex

Done well, option 3 is a hybrid of the other two that has benefits of both the other options.
Within this hybrid option there are still some important design decisions to make. For example:

* Will future questions be shown in any way or will you only see the questions you've answered?
* What happens if you go back and edit a previous question?
    * Does the current question stay open or closed?
    * How do you get back to the current question once you've edited a previous one?
    * Do you lose all your answers to questions that follow the one you go back to edit?

### 4. Hybrid

For more complicated transactions, some combination of the other options might be your best bet.

![Diagram showing a hybrid transaction](/service-manual/assets/images/designing-transactions/hybrid.png)

Again, done well this can give you the benefits of both the single page and wizard approaches. It also allows you to create a sense of rhythm to the overall flow, which can help people to understand when they have moved into a different part of the transaction, and break up the monotony of filling in forms.

As always, these design decisions must have a strong, user-centred rationale behind them.

---

[Discuss this page on Hackpad](https://designpatterns.hackpad.com/Question-pages-ZztvLlQ7VDV)
