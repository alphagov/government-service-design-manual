---
layout: design-pattern
title: Addresses
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: User-centred design
    url: /service-manual/user-centred-design
  -
    title: Design patterns
    url: /service-manual/user-centred-design/resources/patterns
---

{:.intro}
If you need to ask for an address on a form, consider where the addresses will be from and what you need to do with the data.

### On this page:

1. [Free text box](#free-text-box)
2. [Multiple fields](#multiple-fields)
3. [Address finder](#address-finder)

---

<h2 class="heading-36" id="free-text-box">1. Free text box</h2>

A single, multi-line text box where users write out the address in full.

<div class="example">
  <img src="/service-manual/assets/images/design-patterns/full-address.png" alt="An example of a single multi-line address field">
</div>

**The good:**

* it can handle any possible address format
* people can copy and paste addresses from the clipboard
* people don't have to work out which part of the address goes in which field

**The bad:**

* parsing addresses for sub-parts (region, street etc.) is hard, and impossible to do with 100% accuracy
* many legacy back-end systems expect multi-field addresses

### Guidance

Use when you're expecting a very broad range of address formats and you don't need to use specific sub-parts of the address.


---

<h2 class="heading-36" id="multiple-fields">2. Multiple fields</h2>

The address is broken down into multiple fields. Here's an example that works for simple UK addresses:

<div class="example">
  <img src="/service-manual/assets/images/design-patterns/multi-line-address.png" alt="An example of multiple address fields">
</div>

**The good:**

* you can easily extract the parts of an address and do things with them
* you can give help for or validate each part of the address separately
* works well with browsers that have auto-complete enabled


**The bad:**

* hard to find a single format that works for a broad range of regions
* no guarantee that people will use the fields as you intended
* can't easily paste addresses from the clipboard

### Guidance

Only use multiple address fields when you know which regions the addresses will come from and can find a format that supports them all.

**UK addresses:**

* 'postcode' is one word
* let people enter postcodes with or without spaces
* avoid making individual fields mandatory (but warn users if they don't fill in any fields)
* Royal Mail do not need a county as long as the town and postcode are included
* include a county field though -- it lets people who don't know the postcode give a valid address
* make the field lengths appropriate -- it helps people understand the form


---

<h2 class="heading-36" id="address-finder">3. Address finder</h2>

Sometimes referred to as 'postcode lookup'. An address finder lets users specify a UK address by inputting their postcode (and optionally street name or number) and selecting the address from a list.

Here's how this pattern was implemented on the [Lasting Power of Attorney](https://lastingpowerofattorney.service.gov.uk/) service.

<div class="example">
  <img src="/service-manual/assets/images/design-patterns/postcode-lookup.gif" alt="An example of an address finder">
</div>

**The good:**

* people entering UK addresses don't have to enter as much information
* reduces the chance of mis-typed UK addresses

**The bad:**

* requires greater effort to implement

### Guidance

* make it clear that the address finder only works for UK addresses
* provide a manual option for people with international adresses or addresses that are missing or badly formed in the Royal Mail database


---

[Discuss this page on Hackpad](https://designpatterns.hackpad.com/Addresses-CgrMSGRAhRc)
