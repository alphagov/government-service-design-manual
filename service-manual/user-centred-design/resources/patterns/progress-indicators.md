---
layout: design-pattern
title: Progress indicators
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
Use these to reassure users that they're making progress and give an indication 
of how much further there is to go.


### On this page:

1. [Step indicator](#step-indicator)
2. [Progress bar](#progress-bar)
3. [Summary menu](#summary-menu)

---

<h2 class="heading-36" id="step-indicator">1. Step indicator</h2>

Tell the user what step they're on.

<div class="example">
  <div class="ribbon">Recommended</div>
  <div class="inner-block">
    
    <img src="/service-manual/assets/images/design-patterns/step-indicator.png" alt="Example of a step indicator" />
    
  </div>
</div>

**The good:**

* simple, compact and accessible

**The bad:**

* doesn't give details of future steps

As this is the simplest approach we recommend you try it before any others.
It's often all users need to reassure them that they're making progress.


---

<h2 class="heading-36" id="progress-bar">2. Progress bar</h2>

Show the steps horizontally across the top of the page, with the current step highlighted.

<div class="example">
  <div class="inner-block">
    
    <img src="/service-manual/assets/images/design-patterns/progress-bar.png" alt="Example of a progress bar" />
    
  </div>
</div>


**The good:**

* provides an overiew of the entire process
* steps can be used as navigation


**The bad:**

* takes up a lot of space, particularly on small screen devices
* adds visual noise which may distract from the important content


**Guidance:**

* don't assume you'll need one of these
* only use with a small, fixed sequence of steps
* number the steps to reinforce their sequential nature
* don't rely on progress bars for navigation - you'll still need to provide 'Back' and 'Next' links on each screen

---

<h2 class="heading-36" id="summary-menu">3. Summary menu</h2>

[This pattern is still under discussion](https://designpatterns.hackpad.com/Progress-indicators-3AOrLoia9Us)

Provide a vertical list of links to each section. Let users complete the sections in any order.

<div class="example">
  <div class="inner-block">
  
    <img src="/service-manual/assets/images/design-patterns/summary-menu.png" alt="Example of a summary menu" />

  </div>
</div>

**The good:**

* section titles can be longer
* room for more sections
* users can complete sections in an order that suits them
* users can jump around and preview sections before completing them


**The bad:**

* takes up a lot of space
* requires more effort to implement
* some usability challenges - for example, letting users know they've finished all the sections


---

[Discuss this page on Hackpad](https://designpatterns.hackpad.com/Progress-indicators-3AOrLoia9Us)

