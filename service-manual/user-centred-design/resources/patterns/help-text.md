---
layout: design-pattern
title: Help text
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
The best way to help users is by making the simplest, clearest service you can.
Sometimes though you'll need to provide extra help.

### On this page:

1. [Writing help text](#writing-help-text)
2. [Inline help text](#inline-help-text)
3. [Hidden help text](#hidden-help-text)
4. [Detailed help](#detailed-contextual-help)

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

<h2 class="heading-36" id="hidden-help-text">3. Hidden help text</h2>

A short link that expands into more detailed help when clicked on.

<div class="example">
    <img src="/service-manual/assets/images/design-patterns/details-tag.gif" alt="Example of expanding help text">
</div> 

Use this to make your page easier to scan, but don't hide text if a majority of users will need it.


---

<h2 class="heading-36" id="detailed-contextual-help">4. Detailed help</h2>

[This pattern is still under discussion](https://designpatterns.hackpad.com/Contextual-help-XqnDcGgTBKQ)

If your users have to make difficult or complex decisions that can't be supported by the above approaches you might need to provide them with more detailed help.

Make sure that users:

* can access the help before and during the transaction
* can get back to the transaction without losing their place
* are shown help that's relevant to their current situation

Don't use detailed help to explain how to use the interface - it should be simple enough to use without this.



---

[Discuss this page on Hackpad](https://designpatterns.hackpad.com/Contextual-help-XqnDcGgTBKQ)
