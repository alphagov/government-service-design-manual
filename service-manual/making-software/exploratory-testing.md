---
layout: detailed-guidance
title: Exploratory testing
subtitle: What is exploratory testing and how to use it
category: making-software
type: guide
audience:
  primary: developers, qa 
  secondary: web-ops, tech-archs
phases:
  - alpha
  - beta
  - live
status: draft
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Making software
    url: /service-manual/making-software
exclude_from_search: true
---

The goal of exploratory testing is to explore a system as a user would, possibly finding new defects as you go.

There are no instructions to follow because you’re not looking into anything specific. For this reason exploratory testing is sometimes referred to as unscripted testing.

Unlike following a script to test a predetermined outcome, exploratory testing relies on a tester’s knowledge, experience and intuition to fully check a system.

## The value of exploratory testing over other testing methods

Exploratory testing:

* detects subtle or complex bugs in a system (that are not detected in targeted testing)

* provides user-oriented feedback to developers and business analysts

Exploratory testing aims to find new and undiscovered problems. This contrasts with other more prescribed methods of testing, such as test automation, which aims to show scripted tests can complete successfully without defects.

## Approach to exploratory testing

Practitioners tend to avoid describing exploratory testing as ‘ad hoc’ as this implies the testing is unplanned, unstructured and unfocused. Exploratory testing is a disciplined approach and requires specific techniques.

State a goal for each testing session and document the session as it progresses, particularly when you identify defects.

## Team requirements

Exploratory testing is ideally performed by experienced testers as they tend to have:

* analytical skills

* an ability to behave like users

* a knack of finding defects

If you’re an experienced tester, you’ll easily adapt to exploratory testing without prior knowledge of the field. On-the-job-training, mentoring or reading about the discipline can help you prepare for your first test.

You can carry out exploratory testing in isolation, allowing yourself to focus on what the system does rather than what it ought to do.

Alternatively, you can often achieve good results working alongside developers and business analysts, who can offer insights into system functionality or intended user journeys.

## When to do exploratory testing

Exploratory testing works best on a system that holds enough functionality for you to interact with in a meaningful way.

The system needs to be relatively stable and mature, but you can prioritise your development activities so you can do exploratory testing as early as possible, eg build functionality in slices so you can test from start to finish as early as possible.

## Techniques & outputs

Apply the techniques of test charters, mind maps and timeboxing to achieve best practice during exploratory testing. These techniques will put some general restrictions around your testing, helping you judge whether it’s a success.

## Test charters

Test charters give exploratory testing sessions a mission without being prescriptive, eg a charter won’t detail what problems you’re looking for.

Your charter can outline:

* what parts of the system you’ll be looking at (also known as your mission scope)

* your testing goals, eg make sure a system is secure enough for people to use

* details of the team doing the exploratory testing

* the date and times of the testing

* the test environment, eg the software you’ll use to do the test

* mock data you’ll need for testing, eg a mock username or password

As the exploratory testing session progresses, update the charter to log any observations or issues you’ve come across and note down ideas for further testing.

## Mind maps

Mind maps follow a similar approach to test charters but display the information graphically.

Mind-maps are seen as appropriate for exploratory testing as they reflect the non-linear nature of the session.

## Timeboxing

You can limit your exploratory test session to a set timescale, a practice also known as timeboxing. This can help you focus on the project’s specific goals and scope, rather than drift into unfocused exploration.

More on these techniques can be found in an article posted online by Jonathan Bach:[ http://www.satisfice.com/articles/sbtm.pdf](http://www.satisfice.com/articles/sbtm.pdf)

## Collecting evidence

Document all your queries and observations during the exploratory testing session. Capture the evidence as notes, screenshots and log files to help the testing be replicated or further investigated.

An exploratory test that discovers a bug can be developed into a test scenario and automated.

When you need to formally document the outcome of a session, your report may include:

* a charter

* features tested

* notes on how testing was conducted

* notes on any bugs found

* list of issues (questions and concerns about the product or project that have arisen during testing)

* extra materials used to support testing

You can also report how much of the exploratory testing session you spent on each different activity, including:

* creating and executing tests

* investigating and reporting bugs

* setting up the session

## Tools for exploratory testing

There are many different tools you can use for exploratory testing, eg fully automated video capture and logging tools, planning tools, or walls full of process diagrams and feature descriptions.

However, the only tools you need are a pen and some paper. The better tools are nice to have, but don’t put off your testing if you don’t have access to them.

