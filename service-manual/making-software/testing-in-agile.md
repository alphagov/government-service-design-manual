---
layout: detailed-guidance
title: Testing in an agile environment
subtitle: How to get testing right
category: agile
type: guide
audience:
  primary: qa, developers
  secondary: service-managers, web-ops, tech-archs
theme: agile
status: draft
phases:
  - alpha
  - beta
  - live
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Making software
    url: /service-manual/making-software
exclude_from_search: true
---

The basics of any testing approach still apply in the agile world, but the aim of testing can be quite different.

It's important to recognise you are testing in the first place to:

* build the best quality system you can
* make sure it does what the customer requires
* complete it at a cost that everyone agrees can be afforded (cost being money, business change, risk etc)

Too often the aim of testing is to validate what has been produced and nothing else. Your testing should be more about these 7 concepts:

* [building quality in](#building-quality-in)
* [everyone is responsible for quality](#ownership-of-quality)
* [fast feedback](#fast-feedback)
* [tests are an asset of the product](#tests-are-an-asset)
* [faster delivery into production](#faster-delivery)
* [clear and consistent view of testing](#clear-view)
* [optimise value](#optimise-value)

## Building quality in
{: id="building-quality-in"}

Use the vast majority of your efforts building quality in at the start, and test the quality throughout the project -- not at the end when it’s too late.

Define acceptance criteria standards for your user stories. You can either do this when you first write your user stories or later on in the form of acceptance tests when you start development work. Testing should confirm what you already know and understand to be true, so there should be no surprises in the latter stages.

## Everyone is responsible for quality
{: id="ownership-of-quality"}

Service quality isn’t just a testing issue. The quality of a system is defined by the people who create it.

Your team should be able to see a problem in the quality of your system. Every person on the project should be taking action to increase quality and fix issues.

## Fast feedback
{: id="fast-feedback"}

A successful agile project relies on fast feedback loops. Getting feedback and getting it fast means you can actually be agile and change when you need to change.

Your testing should be about giving and getting that fast feedback -- at a time when it’s useful. [Automated code testing techniques](/service-manual/making-software/code-testing) have their place and you can use them, but don’t make them the centre of your approach to testing.

## Tests are an asset of the product
{: id="tests-are-an-asset"}

When you build a test, build it to a standard that makes it suitable to use for testing multiple times throughout your project.

It takes a lot of effort to do testing correctly, so don’t make it a throwaway exercise that has to start from scratch each time there’s a new release, or a new project.

Write your automated tests with the same care and accuracy as production code.

## Faster delivery into production
{: id="faster-delivery"}

Testing is necessary and valuable to the programme, but the time it takes for your code to go live once written is time wasted.

Use testing to get the fastest possible confirmation that everything is as you expect -- or that it isn’t and needs to be fixed.

It doesn’t always need to be exhaustive at every level, but it does need to be relevant to the current situation. Your team needs to agree on the necessary testing at each level, based on the preferences of the product owner and the likelihood of risk in the application itself.

## Clear and consistent view of testing
{: id="clear-view"}

Everybody in your programme needs to understand and agree on:

* the approach to testing
* where they fit in
* what they’re required to do

## Optimise value
{: id="optimise-value"}

When done well, testing will inform you of the best way forward. It’ll also give you the best value in terms of effort used in various functional or nonfunctional areas and will:

* help make the tough decisions
* direct development based on the risk of each user story
* help with the order of priority, based on the complexity of the system

## Types of testing

The most noticeable difference with testing in an agile project is that the majority of your test efforts will be centred on automated tests. 

These tests run in continuous integration, which means that they form part of your codebase -- every time you make a change to your code, your tests are automatically run. You’ll get immediate feedback on the quality of your code and it’ll help to prevent bugs happening at a later stage (when they’re expensive and complicated to resolve.)

### Code testing
Read the guidance about [testing code](/service-manual/making-software/code-testing.html).

### Exploratory testing
Read the guidance about [exploratory testing](/service-manual/making-software/exploratory-testing.html).

### Load and performance testing
Read the guidance about [load and performance testing](/service-manual/operations/load-and-performance-testing.html)

### Penetration testing
Read the guidance about [penetration testing](/service-manual/operations/penetration-testing.html)

### Accessibility testing
Read the guidance about [accessibility testing](/service-manual/user-centred-design/user-research/accessibility-testing.html)

### Crowdsourced testing
Crowdsourced testing doesn’t use a set group of people to carry out testing (known as outsource testing). It uses different people from different places in different jobs. It’s a good way of speeding up your manual testing and/or covering more ground.

There are external companies who provide this as a service, but GDS do it internally by:

* putting out a call for as many volunteers as possible across the organisation to put aside a few hours on a particular day for testing
* giving some guidance on what needs to be tested
* providing somewhere to log bugs
* sometimes motivating testers by creating a ‘leaderboard’ that shows who’s tested the most things

Examples of where GDS used this effectively include:

* the pre-release testing of additional devices/browsers
* the detailed checking of hundreds of pieces of content on GOV.UK against old content on DirectGov and BusinessLink

### Test your ideas
Don’t forget, don’t just test the product itself -- test your ideas.   For information on how to do this read [the guidance about user research](/service-manual/user-centred-design/user-research/index.html).
