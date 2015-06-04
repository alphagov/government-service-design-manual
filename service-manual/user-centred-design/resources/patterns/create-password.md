---
layout: design-pattern
title: Create a password
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
Help people create secure passwords.



### On this page:

1. [Help users to create good passwords](#section-1)
2. [Help users to enter passwords securely and accurately](#section-2)
3. [Help users who have forgotten their password](#section-3)
4. [Examples](#section-4)

---

<h2 class="heading-36" id="section-1">1. Help users to create good passwords</h2>


### Password constraints

You can force stronger passwords by setting minimal constraints on them relating to length, characters used etc.

Make sure you:

* choose constraints appropriate to your users security needs
* explain to users what those constraints are

Overly strict or confusing constraints on password formation make it harder for people to use your service. 

Some users may give up. Others may regularly forget their password and have to rely on whatever support you’ve provided. Others will end up storing their password in some non-secure place.


### Password strength indicators

We are not currently recommending use of password strength indicators because some evidence indicates that they:

* are distracting to some users
* encourage some users to create minimally secure passwords
* occasionally report easy-to-guess passwords as being strong

---

<h2 class="heading-36" id="section-2">2. Help users to enter passwords securely and accurately</h2>


When users are creating a password they need to know that they’ve entered it correctly, but they also need to know that no-one has seen what they’ve entered.


### Hide passwords by default

Users might be in a public space when entering or creating a password, so hide passwords by default. However, if they can't see what they're typing, it's:

* easier to mistype a password
* harder to meet strict password contraints

You can help by:

* letting users see their password if they want to
* showing the last typed character of their password
* making users enter their password twice and comparing them 

This last approach is only appropriate when creating a password. It should not be implemented in combination with the other two approaches and is not guaranteed to block all mistypes.


### Don’t disable paste

Don’t disable paste on password fields. People may have very good reasons why they want to paste their password -- they're using a password manager for example.

---

<h2 class="heading-36" id="section-3">3. Help users who have forgotten their password</h2>

Passwords that are hard to guess can also be hard to remember. Help users who have forgotten their password to reset it.

### Reset passwords, don’t email them

Email is not a secure channel. Sending passwords by email is a bad idea.
Instead, email a link or code to the address or phone number registered on the account that the user can use to create a new password.


### Avoid password reset questions

These are not particularly effective. The information they ask for is either too obscure and therefore just as hard to remember as a password, or it's easy to find out (mother’s maiden name), or it's unstable (favourite colour).


### Avoid password reminders

Password reminders are a bad idea because they encourage users to reveal information about their password and they don't work for very strong passwords involving random strings of characters.

---

<h2 class="heading-36" id="section-4">4. Examples</h2>

An example with a 'show typing' toggle:

<div class="example">
  <img src="/service-manual/assets/images/design-patterns/create-password.png" alt="An example of a password field with a 'show typing' toggle">
</div>

An example with a 're-type password' field:

<div class="example">
  <img src="/service-manual/assets/images/design-patterns/create-password-2.png" alt="An example of a password field with a 're-type password field">
</div>


---

[Discuss this page on Hackpad](https://designpatterns.hackpad.com/Passwords-and-passphrases-4dSSBUhCYjj)

