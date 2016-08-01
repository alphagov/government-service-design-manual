---
layout: detailed-guidance
title: Vulnerability and penetration testing
subtitle: Identify insecurities in your service
category: operations
type: guide
audience:
  primary: web-ops
  secondary: developers, tech-archs
status: draft
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Operations
    url: /service-manual/operations
exclude_from_search: true
---

Making sure your web-based systems and applications are secure requires more than just good design and development.

Sometimes referred to as pen testing, vulnerability and penetration testing is the act of analysing and testing a service for security problems. This is often a specialist activity done via a third party.

View security testing as an ongoing activity in your project, and not as a final check.

## Involve the right people

Security is important to a product and its audience. It’s essential
that non-technical audiences understand vulnerability testing reports
and the risks they identify, as well developers.

Security isn’t just about whether a product secure or not -- it’s also about reducing the risks of a wide range of issues.

Get an independent body to help find potential security problems in your system/application before it’s released to the public, as this helps to identify its vulnerabilities.

### Talk to CESG

[CESG](https://www.cesg.gov.uk/) is the National Technical Authority for Information Assurance. Based at Cheltenham, they provide both standards and advice for information security.

If your project is fairly large, contact CESG as early as possible. They can provide guidance and expertise on potential problems and help you to make sure the right things are tested.

### Use in-house expertise as well as external services

Test for security issues throughout the development of your service, as well as regular testing when your service is live. 

Having third parties do this testing is a good way of introducing genuine experts and getting a different view on something. However it’s also important to make sure that your team building the software know that security is their responsibility, and not something that is outsourced.

## Not just software

When testing for vulnerability, look at the whole system and not just the software involved, such as:

* physical security (where is the equipment housed and how secure is it?)
* the interaction between an online system and a call centre

An obvious example would be physical security (where is the equipment housed
and how secure is it?) but a more interesting example is often the interplay
between an online system and a call centre.

From using information available from a call centre it’s possible to exploit the software system in some way. For instance, getting a call-centre team to change an email address on record for someone else, and then using a forgotten password facility, which relies on the email address being trusted, to access the account. It can and [does happen](http://www.emptyage.com/post/28679875595/yes-i-was-hacked-hard).

## Automate where possible

Some exploratory manual testing is always a good idea when looking for vulnerabilities, but as this is time-consuming and expensive, also have some level of automation.

This may take the form of tests written or tools used specifically to test the security of a feature, eg:

* [static analysis](https://en.wikipedia.org/wiki/Static_program_analysis)
* [input fuzz testing](https://en.wikipedia.org/wiki/Fuzz_testing)

## Why GDS do this

The web is a hostile environment, and the nature of government services means they can be targets for a wide range of different threats -- from financially motivated criminals and online activists up to nation states. Even where personal or sensitive information is not at risk, the smallest issues can damage the reputation of government.

Exploits of web application tend to follow a relatively small number of common patterns. Automated and manual testing, as well as awareness of these common problems, can have a drastic effect on the security of your system.

[The Government Accreditation processes](https://www.cesg.gov.uk/articles/pan-government-accreditation-service-pga) mandate some form of vulnerability testing, often working with [CHECK approved suppliers](https://www.cesg.gov.uk/scheme/penetration-testing). View this as the minimum amount of effort required.

## Further reading

* [OWASP Top 10](https://www.owasp.org/index.php/OWASP_Top_10)
