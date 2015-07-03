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
How to structure web forms for GOV.UK services.

### On this page:

1. [Know why you’re asking every question](#section-1)
2. [Design for the most common scenarios first](#section-2)
3. [Start with one thing per page](#section-3)
4. [Examples](#section-4)
5. [Further reading](#section-5)

---

<h2 class="heading-36" id="section-1">1. Know why you’re asking every question</h2>

Before you start, make a list of all the questions you want to ask. For each question, you need to know:

* why does the service need that information?
* who uses the information and what for?
* which users need to provide the information?
* how will you check that the information is accurate?
* how will you keep the information up to date?

We call this list a 'question protocol'; it’s different from the form itself, because it’s about *how* you use the answers.

A question protocol forces you (and your department) to question why you’re asking users for each item of information, and gives you a way of challenging and pushing back if you need to.


---


<h2 class="heading-36" id="section-2">2. Design for the most common scenarios first</h2>

Once you have a question protocol you can start to decide how to order the questions.

Structure the form so most users have a simple, quick path. Use branching questions to hide questions from people who don't need to answer them.

You’ll need to make difficult decisions about which users to prioritise, so make sure you have good data from the business about them.


---


<h2 class="heading-36" id="section-3">3. Start with one thing per page</h2>

Most public-facing government services are used infrequently, by people with a broad range of digital skills and confidence. Because of this, we recommend that you:

* have one thing per page
* only merge pages when you have clear evidence from user research that this would improve the user experience

'One thing' could be:

* one piece of information to understand
* one decision to make
* one question to answer

We’ve found that this approach works for both high and low confidence internet users, and that people are generally very tolerant of clicking through multiple pages.

Having only one thing on a page helps people to:

* understand what they’re being asked to do 
* focus on the specific question and its answer
* find their way through an unfamiliar process
* use the service on a mobile device
* recover easily from form errors

It also helps you to:

* save people’s work automatically as they go
* capture analytics about each question
* handle branching questions and loops


### Tips

‘One question’ doesn’t necessarily mean one form field. For example, date of birth is best captured with three text fields.

For page titles you can use either a question or a statement. For example - ‘What is your date of birth?’ or just ‘Date of birth’.

Use questions or statements consistently to help users get into a rhythm of answering. This lets them focus on the content of the questions rather than their presentation.

Try to ask questions that might result in the user being ineligible for the service as soon as possible, so you don't waste people's time.




---


<h2 class="heading-36" id="section-4">4. Examples</h2>


For a good example of this approach, try the [Register to vote](https://www.gov.uk/register-to-vote) service on GOV.UK (you can get to the last page without submitting any data).

Joe Lanman has written more about the [things we learned designing 'Register to vote'](https://designnotes.blog.gov.uk/2014/07/14/things-we-learnt-designing-register-to-vote/).


---


<h2 class="heading-36" id="section-5">5. Further reading</h2>


[The Question Protocol: How to Make Sure Every Form Field Is Necessary](http://www.uxmatters.com/mt/archives/2010/06/the-question-protocol-how-to-make-sure-every-form-field-is-necessary.php)


---

[Discuss this page on Hackpad](https://designpatterns.hackpad.com/Form-structure-XDwY2wv3lCn)
