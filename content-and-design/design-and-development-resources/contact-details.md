---
layout: design-patterns
title: Contact details
subtitle: Build to the GOV.UK style 
category: design-and-development-resources
type: resource
status: draft
phases:
  - alpha
  - beta
  - live
page_class: contact-details
---

Contact details for an organisation or individual.

Remember to add the relevant [hCard](http://microformats.org/wiki/hcard) classes.

## Example

<div class="pattern-example">
  <div class="inner">

    <div class="address vcard">
      <div class="adr org fn">
        <p>
              Office for Judicial Complaints
          <br>Steel House
          <br>11 Tothill Street
          <br>London
          <br>SW1H 9LJ
        </p>
      </div>
    </div>

  </div>
</div>

## Code

    <div class="address vcard">
      <div class="adr org fn">
        <p>
              Office for Judicial Complaints
          <br>Steel House
          <br>11 Tothill Street
          <br>London
          <br>SW1H 9LJ
        </p>
      </div>
    </div>  

* * * 

# Discussion

Q: Does the envelope design work for telephone numbers, email addresses etc.?
<br> No - lose the envelope. BT

Q: What's our policy on publishing email addresses?




