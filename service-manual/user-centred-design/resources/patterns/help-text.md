---
layout: design-pattern
title: Help text
breadcrumbs:
  -
    title: Home
    url: /service-manual/index.html
  -
    title: User centred design
    url: /service-manual/user-centred-design/index.html
  -
    title: Design patterns
    url: /service-manual/user-centred-design/resources/patterns/index.html
---

{:.intro}
The best way to help users is by making the simplest, clearest service you can.
Sometimes though you'll need to provide extra help.

### On this page:

1. [Writing help text](#writing-help-text)
2. [Inline help text](#inline-help-text)
3. [Expanding help text](#expanding-help-text)
4. [Detailed help](#detailed-help)

---

<h2 class="heading-36" id="writing-help-text">1. Writing help text</h2>

Keep it short and concise and speak to the user in their language. 
Users are unlikely to read anything longer than 3 lines. 

Focus on the action required and make it relevant to the user's current situation.
Don’t give background information like ‘This used to be called X but in 2008 it was changed to Y’.


### Do this:

<div class="example">  
    <img src="/service-manual/assets/images/design-patterns/help-text-good.png" alt="Example of good help text">
</div> 

### Don't do this:

<div class="example">
    <img src="/service-manual/assets/images/design-patterns/help-text-bad.png" alt="Example of bad help text">
</div> 
 

### Do explain:

* legal jargon
* unfamiliar concepts
* where to find obscure information
* what format the information should be given in
* what you do with personal information

### Don't explain:

The interface. If you have to do that you haven't made it simple enough.
If you find yourself writing things like 'Click on the green button at the bottom of the screen' you know the interface needs work.

---

<h2 class="heading-36" id="inline-help-text">2. Inline help text</h2>

Short, clear text positioned immediately next to the part of the page it relates to.
Use this for help that's relevant to the majority of users.

<div class="example">
    <img src="/service-manual/assets/images/design-patterns/help-text-inline.png" alt="Example of inline help text">
</div> 

Inline help for form fields should be positioned between the label and the field so it's read by sighted and screen reader
users before they get to the field itself.

---

<h2 class="heading-36" id="expanding-help-text">3. Expanding help text</h2>

A short link that expands into more detailed help when clicked on.

<div class="example">
    <img src="/service-manual/assets/images/design-patterns/details-tag.gif" alt="Example of expanding help text">
</div> 

Used carefully this is a good way of keeping the interface free from clutter.

Only use this for help that's only relevant to a minority of users.
Remember - if most people need the text, don't hide it.

Try link text that expresses the immediate user need rather than the generic 'Help' or 'Info'. For example 'I can't find my National Insurance number', or 'What's this?'.

---

<h2 class="heading-36" id="detailed-contextual-help">4. Detailed help</h2>

Very occasionally you'll need to give users access to more detailed supporting content from within a transaction.

One example on GOV.UK is the [Lasting Power of Attorney](https://lastingpowerofattorney.service.gov.uk) service.
Because the subject matter is complicated and the decisions the user is making are very serious the service needs to support users who are learning about the subject as they go through the forms. 

In this case a modal dialogue is opened at the relevant page. The user can read as much or as little of the help content they need, dismiss the dialogue and then continue with the form.

<div class="example">
    <img src="/service-manual/assets/images/design-patterns/detailed-help.gif" alt="Example of detailed help text">
</div> 


---

[Discuss this page on Hackpad](https://designpatterns.hackpad.com/Contextual-help-XqnDcGgTBKQ)
