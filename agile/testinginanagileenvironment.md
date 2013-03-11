---
layout: detailed-guidance
title: Testing in an agile environment
subtitle: What testing your service might look like
section: agile
type: guide
audience: 
  primary: service-manager, developer, researcher
  secondary: designer, analyst
theme: agile
status: draft
phases:
  - alpha
  - beta
  - live
---

Fundamentally, the basics of any testing approach still apply in the Agile world. However, the focus of testing can be quite different.  

It is important to recognize why we are testing in the first place, and that is to build the best quality system we can, that does what the customer requires, at a cost that everyone agrees we can afford  (cost being money, business change, risk etc.).  Too often, the focus of testing is to validate what has been produced and that alone, when in actuality it should be more about the following 7 concepts:

## Building quality in

You cannot test quality at the end, which is the mistake many teams make. Expend the vast majority of your effort building quality in at the start instead. 

When you write user stories, ensure that you define acceptance criteria against which you can test. Once past that point, testing should primarily be a verification of what we already know and understand to be true, there should be no surprises in the latter stages.

## Everyone is responsible for quality

Service quality isn’t just a testing issue. The quality of a system is defined by the people who create it. If there is a problem in the quality of the system being produced, then it should be evident to everyone involved, and every person on the project should be taking action to increase quality and fix issues.

## Fast Feedback

Agile is reliant on fast feedback loops, so that we can actually be agile and change when we need to change. Testing should be about giving that fast feedback, at the time when it is useful. Agile test techniques (e.g. Behaviour driven development, Acceptance test driven development) have their place and we may well use them, but they are not the primary focus of the test approach.

## Tests are an asset of the product

By this we mean that testing is built with reuse in mind. It takes a lot of effort to do testing correctly, we don’t want it to be a throwaway exercise that has to start from scratch each time there is a new release, or a new project. Automated tests need to be written with the same care and rigour as production code.

## Faster delivery into production

While testing is necessary and valuable to the programme as a whole, any time it takes for that code to go live once written is essentially wasted time. Testing should be a tool that is used to get the fastest possible confirmation that it is as expected, or that it isn’t and needs to be reworked.

Testing doesn’t always need to be exhaustive at every level, it needs to be applicable to the situation at hand. It should be the job of the team to agree on the necessary testing at each level, based on the appetite for risk of the product owner and the likelihood of risk in the application.

## Clear and consistent view of testing

Everybody in the programme needs to understand and agree the approach to testing, and everybody needs to understand where they fit in and what they are required to do.

## Optimise value

Testing, done well, will inform the the best way forward and get the best “bang for buck” in terms of effort expended in various functional or non-functional areas.  It will help make the tough decisions, and drive the development effort based on the risk of each choice of story.  It will help the prioritization based on the understanding of the complexity of the system.

## Types of testing

The most noticeable difference with testing in an Agile world is that the majority of your  test effort will be focussed on automated tests.   
These tests run in Continuous Integration (C.I.) which means that they form part of your code base and every time you make a change to your code, your tests are automatically run. This gives you immediate feedback on the quality of your code and helps prevent bugs being found at a later stage when they are expensive and complicated to resolve.

### Code Testing
Read the guidance about [testing code](/making-software/code-testing.html).  


### Exploratory Testing
Exploratory testing is a term commonly used to describe unscripted manual testing whereby the tester uses his or her knowledge, experience and intuition to navigate through the software and identify bugs. A scripted test can only ever test a predetermined outcome. Exploratory testing aims to find and test the less obvious outcomes. A good way to think about it is that automated tests prevent bugs whereas exploratory tests find them.  

Exploratory testing is normally time-boxed and has a specific purpose - e.g. 'I will spend x hours exploring y and z aspects of the system (though my explorations may take me elsewhere and it may take more or less time, I’ll use my judgement as I go)'. 

The term does not imply that the tester has not prepared for the testing. They will have given the testing some detailed planning in advance, thinking for instance about the specific aspects they want to explore, and any data or other system set-up that they will need.  

Automation may still play a part - not to run the tests themselves but, for instance, to set up the data or to get a set of transactions into predetermined states.

In a team where you have a one or more dedicated ‘quality analysts’ or ‘testers’ this type of testing will normally be part of their role. In a developer-only team time will need to be put aside for them the developers themselves to do this type of testing. As a developer has been deeply involved in writing the code, it is sometime difficult for them to step far enough enough from the system to see paths through the system that they hadn’t previously envisaged. To help with this it’s ideal if they can be assigned to exploratory testing for a full day to allow the appropriate amount of context switching. It is also preferable if they are exploring parts of the system that they have been less involved in developing.

When a manual test uncovers a defect, it is important to always add an automated test to catch it going forward and hence prevent any reoccurrence.

Read [Cem Kaner on exploratory testing](http://www.kaner.com/pdfs/QAIExploring.pdf)

### Load & Performance Testing
Read the guidance about [load & performance testing](/operations/load-and-performance-testing.html)

### Penetration Testing
Read the guidance about [penetration testing](/operations/penetration-testing.html)

### Accessibility Testing
Read the guidance about [accessibility testing](/making-software/accessibility-testing.html)

### Crowd Sourced Testing
Crowd sourced testing is a good way of speeding up your manual testing and/or achieving better coverage.  

There are external companies who provide this as a service, but at GDS we do it internally. We simply put out a call for as many volunteers as possible across the organisation to put aside a few hours on a particular day for testing. We then give them some guidance on what we need testing and a central place for them to log bugs and hey presto! Sometimes we incentivise people by creating a ‘leader board’ showing who has tested the most things.  

Examples of where we have used this effectively include for pre-release testing of additional devices/browsers and for the detailed checking of hundreds of pieces of content on GOV.UK back to old content on DirectGov and BusinessLink.

### Test Your Ideas
Don’t forget, don’t just test the product itself - test your ideas.   For information on how to do this read [the guidance about user research](/users/introduction-to-user-research.html).