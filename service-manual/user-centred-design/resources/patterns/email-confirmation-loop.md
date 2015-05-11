---
layout: design-pattern
title: Email confirmation loops
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
If email is an integral part of your service then consider confirming whether users have access to the
email address they give you.


### On this page:

1. [Consider alternatives](#section-1)
2. [Confirmed email addresses don't prove identity](#section-2)
3. [Set an expiry time on the email](#section-3)
4. [Blocking versus non-blocking loops](#section-4)
5. [Activating your account](#section-5)

---

<h2 class="heading-36" id="section-1">1. Consider alternatives</h2>

Email confirmation loops can be a very disruptive user experience, forcing people to wait and then switch from a website to email and back again. They may not have access to their email at that point in time.

Confirmation loops tend to increase drop-off rates, because some users struggle to complete the loop. They don't notice the email, or it goes in to their spam folder, or they give up on the service before the email arrives.

You decision should be based in part on what information you intend to send by email. Consider just [helping users enter their email correctly](email-addresses) instead.

---

<h2 class="heading-36" id="section-2">2. Confirmed email addresses don't prove identity</h2>

An email confirmation loop verifies that the address supplied by the user is a real one, but doesn't verify that they alone own that address, or that they check it regularly. 

Some users may supply temporary addresses as a way of getting around confirmation loops. Other people may share their account with someone else, or the account may have been hacked.

---

<h2 class="heading-36" id="section-3">3. Set an expiry time on the email</h2>

For added security you can set an expiry date on the email, so that the link can't be used after a certain time period.
If the user attempts to use an expired link then they should be told this and given an opportunity to re-send a new link.

---

<h2 class="heading-36" id="section-4">4. Blocking versus non-blocking loops</h2>

There are two versions of the email confirmation loop. In the ‘blocking’ version the user cannot use the service until they have completed the loop. In the ‘non-blocking’ version they can continue to use the service, but will be reminded regularly that they need to confirm their email. Some functionality may not be available until they have done this.

Blocking loops have a slightly simpler flow, but if the users can't complete the loop they're unable to use the service at all. It's especially important that the emails are sent instantly.

Non-blocking loops require more careful design, and you can't guarantee that all users emails are confirmed, but there's no risk that the loop will stop people from accessing your service.

Some services start users in a non-blocking loop initially and then transition to a blocking loop.


---

<h2 class="heading-36" id="section-5">5. Activating your account</h2>

Use the phrase 'Activate your account' to describe the process of confirming an email address. The activate your account page should be shown immediately after the user provides their email address, or if they try to sign in before confirming their email.

It should:

* tell them where the activation email has been sent
* explain that they need to click the link in the email to proceed
* let them resend the activation email

For blocking loops this should be the only page the user sees if they try to sign in before activating their account.

For non-blocking loops, if a user signs in before activating their account, you should:

* let them use the service
* remind them that they need to activate their account
* tell them where the activation email has been sent
* let them re-send the activation email

When a user clicks on the link in the activation email take them to a page that confirms that they have activated their account. You may or may not require them to sign in at this stage, depending on where in the flow the activate your account screen appears.



---

[Discuss this page on Hackpad](https://designpatterns.hackpad.com/Email-confirmation-loops-rJlwfm9N3QA)

