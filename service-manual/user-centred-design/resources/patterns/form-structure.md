---
layout: design-pattern
title: Form structure
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
Choose a structure for your forms that most naturally fits the way people need to use them.

### On this page:

1. [Asking questions](#asking-questions)
2. [Structuring forms](#structuring-forms)
3. [Example structures](#example-structures)

---

## 1. Asking questions

You can use a question or statement format. 
For example ‘What is your date of birth?’ or ‘Date of birth’.
Just remember to be consistent across the transaction.

Guidance:

* use questions or statements, but be consistent
* ask for as little information as you can
* keep the questions short and straightforward 
* explain why you're asking a question if it's not obvious

---

## 2. Structuring forms

Ask yourself:

* how many questions are there?
* will users need to answer the questions in a fixed order?
* will they need to complete the transaction in a single session?
* will their answers affect other parts of the transaction?
* will certain answers result in the user having to abandon the transaction?
* will they need to add or remove items from a list, or change the order of items?
* how many parties are involved in the transaction?
* do any parts of the transaction take place offline?
* at what point is the transaction regarded as complete?

---

## 3. Example structures

Thinking about the above will help you decide on an appropriate structure.
Here are some of the more common approaches:

### 1. Single page of questions

All questions are placed on a single page.

{:.left}
![Diagram showing all sections on a page](/service-manual/assets/images/designing-transactions/one-page.png)

**The good:**

* there's only one submit button to press
* a single URL gives access to all form fields
* it doesn't force a fixed order of completion
* you benefit from context of neighbouring sections
* progress is self-evident

**The bad:**

* long forms can be overwhelming and off-putting
* it's less well suited to branching or non-linear flow
* how do you save partial progress?
* can be harder to track analytics like drop-off rates
* making validation errors usable is harder

---

### 2. One question per page

Each question goes on its own page.

![Diagram showing each section on it's own page](/service-manual/assets/images/designing-transactions/wizard.png)

**The good:**

* it's easier to handle branching and dependencies between questions
* it's easier to let the user save progress
* a transaction can feel more manageable
* easier to guide a user through an unfamiliar process
* easier to capture analytics like drop-off rates for each section

**The bad:**

* harder to show progress
* uers have to click more to progress through the questions
* you lose the context of neighbouring questions
* you need to build a seperate page to review and edit questions
* doesn't naturally handle non-linear processes like looping, adding and removing

---

### 3. Accordion form

All questions on a single page, but each new question only appears once the previous section has been completed.

![Diagram showing an accordion form](/service-manual/assets/images/designing-transactions/accordion-3.png)

**The good:**

* can handle branching and dependencies between sections
* users can review and edit previous questions at any time
* can help guide a user through an unfamiliar process
* user still benefits from some surrounding context
* progress is clear

**The bad:**

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
