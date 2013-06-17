---
layout: detailed-guidance
title: Testing in an agile environment
subtitle: How to get testing right
category: agile
type: guide
audience: 
  primary: qa
  secondary: service-managers, web-ops, developers, tech-archs
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
---

The basics of any testing approach still apply in the Agile world, but the aim of testing can be quite different.

It is important to recognise why you are testing in the first place:

* to build the best quality system you can
* to make sure it does what the customer requires
* that it’s done at a cost that everyone agrees can be afforded (cost being money, business change, risk etc.)

Too often the aim of testing is to validate what has been produced and nothing else. Your testing should be more about the following 7 concepts:

* building quality in
* everyone is responsible for quality
* fast feedback
* tests are an asset of the product
* faster delivery into production
* clear and constant view of testing
* optimise value


## Building quality in

Use the vast majority of your efforts building quality in at the start, and test the quality throughout the project - not at the end when it’s too late.

Define acceptance criteria standards for your user stories. You can either do this when you first write your user stories or later on in the form of acceptance tests when you start development work. Testing should confirm what you already know and understand to be true, so there should be no surprises in the latter stages.

## Everyone is responsible for quality

Service quality isn’t just a testing issue. The quality of a system is defined by the people who create it. 

Your team should be able to see a problem in the quality of your system. Every person on the project should be taking action to increase quality and fix issues.

## Fast feedback

A successful agile project relies on fast feedback loops. Getting feedback and getting it fast means you can actually be agile and change when you need to change.

Your testing should be about giving and getting that fast feedback - at a time when it’s useful. Agile test techniques (eg behaviour driven development, acceptance test driven development) have their place and you can use them, but don’t make them the centre of your approach to testing.

## Tests are an asset of the product

When you build a test, build it to a standard that makes it suitable to use for testing multiple times throughout your project.

It takes a lot of effort to do testing correctly, so don’t make it a throwaway exercise that has to start from scratch each time there’s a new release, or a new project.

Write your automated tests with the same care and accuracy as production code.

## Faster delivery into production

Testing is necessary and valuable to the programme, but the time it takes for your code to go live once written is time wasted.

Use testing to get the fastest possible confirmation that everything is as you expect - or that it isn’t and needs to be fixed.

It doesn’t always need to be exhaustive at every level, but it does need to be relevant to the current situation. Your team needs to agree on the necessary testing at each level, based on the preferences of the product owner and the likelihood of risk in the application itself.

## Clear and consistent view of testing

Everybody in your programme needs to understand and agree on:

* the approach to testing
* where they fit in
* what they’re required to do

## Optimise value

When done well, testing will inform you of the best way forward. It’ll also give you the best value in terms of effort used in various functional or nonfunctional areas and will:

* help make the tough decisions
* direct development based on the risk of each user story
* help with the order of priority, based on the complexity of the system

## Types of testing

The most noticeable difference with testing in an agile project is that the majority of your test efforts will be centred on automated tests. 
These tests run in [continuous integration (C.I.)](/service-manual/agile/continuous-delivery.html), which means that they form part of your codebase - every time you make a change to your code, your tests are automatically run. You’ll get immediate feedback on the quality of your code and it’ll help to prevent bugs happening at a later stage (when they’re expensive and complicated to resolve.)

### Code Testing
Read the guidance about [testing code](/service-manual/making-software/code-testing.html).  

### Exploratory Testing
Exploratory  testing is a term commonly used to describe unscripted manual testing. This is where the tester uses his or her knowledge, experience and intuition to go through through the software and identify bugs.

The difference between unscripted and scripted tests is that a scripted test can only ever test a predetermined outcome. Exploratory testing (unscripted) doesn’t mean you don’t prepare for the testing.

Exploratory testing is:

* always planned in advance, in detail, including:
    * the specific aspects you want to explore
    * any need for extra data or a different system set-up to perform the testing
* used to find and test the less obvious outcomes (automated tests prevent bugs whereas exploratory tests find them)
* normally timeboxed (runs for a fixed period of time)
* has a specific purpose - eg ‘I will spend x hours exploring y and z aspects of the system (though this may take me elsewhere and it may take more or less time, so I’ll use my judgement as I go)’
* not always without automation – automation can be used to set up the data or to get a set of transactions (just not to run the tests)

The quality analysts or testers in your team carry out this type of testing. If your team is made up of only developers, you’ll need to put time aside for the developers to do the testing, but bear in mind that:

* as developers are deeply involved in writing the code, it’s sometime difficult for them to see ways through the system that they hadn’t previously thought of
* it’s ideal for developers to be assigned to exploratory testing for a full day to allow the appropriate amount of context switching (the way a central processing unit changes from one task to another without the tasks conflicting)
* it’s preferable if they explore parts of the system that they haven’t really developed

When a manual test finds a defect, it’s important to always add an automated test to stop it from happening again.

Read [Cem Kaner on exploratory testing](http://www.kaner.com/pdfs/QAIExploring.pdf)

### Load & Performance Testing
Read the guidance about [load & performance testing](/service-manual/operations/load-and-performance-testing.html)

### Penetration Testing
Read the guidance about [penetration testing](/service-manual/operations/penetration-testing.html)

### Accessibility Testing
Read the guidance about [accessibility testing](/service-manual/making-software/accessibility-testing.html)

### Crowd Sourced Testing
Crowdsourced testing doesn’t use a set group of people to carry out testing (known as outsource testing). It uses different people from different places in different jobs. It’s a good way of speeding up your manual testing and/or covering more ground.

There are external companies who provide this as a service, but GDS do it internally by:

* putting out a call for as many volunteers as possible across the organisation to put aside a few hours on a particular day for testing
* giving some guidance on what needs to be tested
* providing somewhere to log bugs
* sometimes motivating testers by creating a ‘leaderboard’ that shows who’s tested the most things

Examples of where GDS used this effectively include:

* the pre-release testing of additional devices/browsers
* the detailed checking of hundreds of pieces of content on GOV.UK against old content on DirectGov and BusinessLink

### Test Your Ideas
Don’t forget, don’t just test the product itself - test your ideas.   For information on how to do this read [the guidance about user research](/service-manual/users/introduction-to-user-research.html).
