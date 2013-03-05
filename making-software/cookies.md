---
layout: detailed-guidance
title: Cookies
subtitle: When to use them, and when to tell users about them
section: guidance
type: guide
audience:
    primary: developer, tech-arch
status: draft
---


This short guide tells you what to keep in mind when including cookies into your services, and how and why we notify users about cookies on [GOV.UK](https://www.gov.uk).

##Using Cookies

You should minimise the use of cookies throughout services. Store as little information as you require for as short a time as necessary to deliver a good service to users.

If your service requires new cookies to be set then you need to ensure that they can be explained simply and clearly, in a way that the majority of users can understand.

You will need to notify users whenever a specific action sets a cookie. This will look like this:

"[sets a cookie](https://www.gov.uk/support/cookies#)"

That link will take users to the cookies page, and the appropriate description of the cookie being set.


##Cookie warnings

Whenever a user visits [GOV.UK](https://www.gov.uk) for the first time we notify them on our use of cookies using this message:

"GOV.UK uses cookies to make the site simpler. [Find out more about cookies](https://www.gov.uk/support/cookies)

On [the cookies page at GOV.UK](https://www.gov.uk/support/cookies) you can see a breakdown of the kinds of cookie used throughout the site, followed by an explanation of each cookie’s purpose. For example:

###Storing your approximate location

Some pages on this site provide you with information based on your location. So when you enter your postcode, we save your approximate location to a cookie. This means you don’t have to keep typing your postcode in, and means we’ll always point you towards services that are closest to you e.g. the nearest UK online centre.

<table>
    <tr>
        <td>Name</td>
        <td>Purpose</td>
        <td>Expires</td>
    </tr>
    <tr>
        <td>govukgeo</td>
        <td>Saves details about your location so that services which require this information are consistent</td>
        <td>4 months</td>
    </tr>
</table>

##Cookies on existing transactions

The [GOV.UK](https://www.gov.uk) cookies policy does not cover 3rd party transactions. Unfortunately, most existing services were built before the updated EU Privacy Directive came into force, so weren't built with it in mind (hence not having cookie policies).

New and redesigned services will be expected to notify users that cookies are being set beforehand, using a 'sets a cookie' line of text near the action that triggers the cookie (this text should be linked to the appropriate part of the cookie information page).  You can see an example of how that works on GOV.UK [here](https://www.gov.uk/dvla-offices).

Where a user can't be notified before the cookie is set (i.e. if it's set when the user first visits the service using implied consent), the service should use a banner similar to that on [GOV.UK](https://www.gov.uk) to inform people that this is the case and link them to the further information.

##Why we do this
This policy is in line with EU Law and was devised following extensive research during the beta of [GOV.UK](https://www.gov.uk). It follows the [The Privacy and Electronic Communications (EC Directive) (Amendment) Regulations 2011](http://www.legislation.gov.uk/uksi/2011/1208/contents/made) (PECR). You can read more about how we decided this in [this blog post by GDS Developer Dafydd Vaughan](http://digital.cabinetoffice.gov.uk/2012/01/12/cookies-on-the-beta/).


##Further reading
[The ICO's latest guidance](http://www.ico.gov.uk/for_organisations/privacy_and_electronic_communications/the_guide/cookies.aspx)

This [blog post by GDS Developer Dafydd Vaughan](http://digital.cabinetoffice.gov.uk/2012/01/12/cookies-on-the-beta/) explains how cookies were used on the beta version of [GOV.UK](https://www.gov.uk).

