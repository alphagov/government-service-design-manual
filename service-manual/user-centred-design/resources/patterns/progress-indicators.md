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
Help people understand where they are in a transaction and give them the confidence to continue.


### On this page:

1. [Start without a progress indicator](#section-1)
2. [If you do use one, keep it simple](#progress-bar)
3. [Avoid complex progress indicators](#section-3)

---

<h2 class="heading-36" id="section-1">1. Start without a progress indicator</h2>

Test your service first without any progress indicators at all. It may be simple enough that you don't need them. If it isn't, then at least you'll discover the point at which people start to struggle.

It's often the order, type or number of questions that causes issues, so try improving these first.

If people still have problems then try adding a simple progress indicator.

---

<h2 class="heading-36" id="section-2">2. If you do use one, keep it simple</h2>

If you do need a progress indicator, just add the step or question number at the top of the page, like this:

<div class="example">
  <div class="inner-block">

    <img src="/service-manual/assets/images/design-patterns/step-indicator.png" alt="Example of a step indicator" />

  </div>
</div>

Only include the total number if you can do so reliably.

This approach is compact, accessible and usually enough to give people the confidence to continue.



---

<h2 class="heading-36" id="section-3">3. Avoid complex progress indicators</h2>

<div class="example">
  <div class="inner-block">

    <img src="/service-manual/assets/images/design-patterns/progress-bar.png" alt="Don't use horizontal progress bars" />

  </div>
</div>

We recommend against these because:

* they're rarely used
* they take up valuable space
* they don't scale well on small screens
* they can distract and confuse some people
* it's hard to write good labels for the steps
* it's hard to handle conditional questions

Complex progress indicators have been safely removed from a number of services on GOV.UK. For example, the Carer's Allowance team removed a 12-step horizontal progress indicator from their live service, with no effect on completion rates or times.


# Progress indicators as navigation

We occasionally see progress indicators (usually vertical ones) used as navigation. These suffer from many of the issues listed above.

If people really need to be able to jump around the transaction, try meeting this need on a separate page.


---

<h2 class="heading-36" id="section-4">Further reading</h2>

* [Do Less: Problems as shared spaces](https://designnotes.blog.gov.uk/2014/07/07/do-less-problems-as-shared-spaces/). Ben Holliday talks about removing the progress indicator from the Carer's Allowance service.


---

[Discuss this page on Hackpad](https://designpatterns.hackpad.com/Progress-indicators-3AOrLoia9Us)

