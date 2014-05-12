---
layout: design-pattern
title: Progress indicators
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

When people interact with multi-page transactions they can sometimes become frustrated, 
disorientated or think they've finished when they haven't. 
This can cause them to abandon the transaction before they've completed it.

Add a progress indicator to reassure users that progress is being made and give an indication 
of how much further there is to go. Don't assume that you'll need a progress bar though.


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


**Guidance:**

* only use with a small, fixed sequence of steps
* number the steps to reinforce their sequential nature
* don't rely on progress bars for navigation - you'll still need to provide 'Back' and 'Next' links on each screen

---

<h2 class="heading-36" id="summary-menu">3. Summary menu</h2>

<div class="example">
  <div class="inner-block">
  
    <img src="/service-manual/assets/images/design-patterns/summary-menu.png" alt="Example of a summary menu" />

  </div>
</div>

Provide a vertical list of links to each section, on the left of the page, which can be completed in any order.



**The good:**

* section titles can be longer
* room for more sections
* users can complete sections in an order that suits them
* potential for steps to be partially completed


**The bad:**

* requires a lot of space
* the challenge is letting users know when they have completed all the steps


**Guidance:**

A summary menu is similar to a to-do list. 
It also has some of the advantages of paper forms.
Users can jump around and preview sections,
fill out the ones they can do immediately and then return to the more difficult ones.

To take advantage of this though you'll need to provide a way for users to save and return to their work.

All of this requires effort, so only implement this if you're sure your users need these features.

---

## Conclusion

Go with the simplest approach that addresses your user's needs. 
Don't assume that you'll definitely need a progress bar.

---

[Discuss this page on Hackpad](https://designpatterns.hackpad.com/Progress-indicators-3AOrLoia9Us)

