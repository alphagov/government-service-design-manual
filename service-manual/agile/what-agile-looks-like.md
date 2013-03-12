---
layout: detailed-guidance
title: What agile looks like
subtitle: Common features of agile projects
category: agile
type: guide
audience:
  primary: service-manager
  secondary: designer, developer, tech-arch, analyst, researcher
status: draft
phases:
  - discovery
  - alpha
  - beta
  - live
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Agile
    url: /service-manual/agile
---

{:.intro}
Agile is a liberating way of working.  It does not preclude the use of existing skills and knowledge. But it does require teams, users and stakeholders to adopt new ways of working together.  

{:.intro}
This short guide lists a few of the behaviours common to agile projects that support successful delivery and learning. 

## Understand your users 


![Real people will use your product](http://farm8.staticflickr.com/7177/6987029385_410a1c8d26.jpg)


Prioritise features for them over everyone else -- including your big, scary stakeholders, and seek their feedback early and often.  Really listen to them.  Even when they tell you things you don’t want to hear or disagree with.  If possible, use data from real people using your product to influence the direction of the project. Your focus on the user should be relentless.

## “What do you want next Friday? What have we learned last week?”


![A sprint backlog, coutesty of http://www.flickr.com/photos/psd/](http://farm9.staticflickr.com/8043/8100964765_acf4032d09.jpg)


Iterate often. Build something focused on the next most valuable user need and show it to them; listen to their feedback and improve it. Keep doing this until you have something so useful that they would not be without it. 

It perhaps sounds like over-simplifying the complexity of software development and project management, but at its heart this is what agile development is all about:  “What do you want next Friday?”

The process of delivering incremental, production-ready software allows a team to deliver value to their users and stakeholders regularly.  It shortens the feedback loops that might otherwise have been longer using a waterfall methodology.  An iterative delivery cycle also forces the team to think about what the most important features are to deliver next and focuses the mind on useable software.

At the end of each delivery cycle, or sprint, teams should run a [retrospective](/service-manual/agile/running-retrospectives.html) to review ‘what worked, what could be improved’ in the next sprint.  

The software and the team continue to learn through delivery and iterate and improve throughout the project.

## Small, agile teams

![The unit of delivery is the team](http://farm9.staticflickr.com/8374/8451589322_e9f612cf5b.jpg)

Small teams of between five to ten people are often more productive and predictable than larger teams. Forget man-days and think about the team as the unit of delivery.  

A good team includes members with all of the skills necessary to successfully deliver software. A fully-functioning team has three key roles embedded into them, usually full-time:

*Product Manager* - responsible for delivering return on investment, usually by creating products that users love.  The team delivers the Product Manager’s vision.

*Delivery Manager* (a.k.a. Scrum master or Project Manager) - is the agile expert that is responsible for removing blockers (things slowing a team down).  They also usually act as a facilitator at team get togethers.

*Team member* - Self organising, multi-disciplinary team that delivers prioritised user stories. Responsible for estimation.

You help each other and work together toward delivering your sprint goals.  It’s common to encourage team members to pair. It sounds counter-intuitive to have two people work on one thing, but this is not so.  Working together closely produces better software solutions, promotes better quality controls and spreads knowledge across the team.

A good team can estimate their output, or velocity, very effectively and consistently.  This allows for much more accurate planning.

## Fail fast

![Failing, so fix it!](http://farm8.staticflickr.com/7189/6875228285_9b2409663f.jpg)

Releasing little pieces of code often improves quality and visibility and reduces cost to market, but using agile techniques does not guarantee success. You can still fail!  What agile methodologies do allow you to do is to spot problems earlier and resolve them.  

Here’s a few examples of how:

>* releasing working software to your users often allows you to get feedback quickly and hear or see what they think.  If the product is wrong you can easily change direction and iterate.
>* if your software is rarely released to production you are not demonstrating value to your sponsor. You run the risk of creating a “too-big-to-fail” service that isn't fit for public consumption but must be releaesed anyway. That means another press headline!  Ship! Ship! Ship!
>* if your teams’ velocity is consistently volatile, beyond the initial 4-6 sprints, then this is indicative of something that needs fixing.  Perhaps there is hidden complexity or poor estimation.
>* Test Driven Development (writing tests in code before you develop the features) has a wealth of metrics that highlight quality issues early.  Establish what these are early on, baseline and monitor throughout the project.

Don’t be afraid to fail or experiment.  Learn to fail, and create a culture that learns from failure.

##Continuous Planning

![Planning session](http://farm9.staticflickr.com/8001/7113823877_80c4dfb613.jpg)

It’s a myth that you don’t plan on agile projects.  The freedom of agile projects does not come free: you have to plan.  You just plan differently and continuously.

Agile planning is based as much as possible on solid, historical data, not speculation. The plan must continuously demonstrate its accuracy: nobody on an agile project will take it for granted that the plan is workable.

Typically teams plan together, usually on at least two levels:

>* at the release level, they identify and prioritise the features we must have, and would like to have by the deadline.  
>* at the iteration level, they plan for the next features to implement, in priority order. If features are too large to be estimated or delivered within a single iteration, they break them down further.

These plans are usually reviewed after every sprint and adjusted based on “the weather yesterday”, new facts and requirements that will inevitably be uncovered along the way.

##Bad smells

![Do go here!](http://farm9.staticflickr.com/8424/7503675672_72ff8a1fa9.jpg)

Teams new to agile should be wary of these familiar situations and reactions to doing things differently.  They have a bad smell about them and undermine your project and its chances of success.

>* *Your team is not full time*. If your core team of product manager, scrum master, and key members of your multi-disciplinary team are not on the project full-time and spread over many projects then expect difficulties.  The team is the unit of delivery and you need focus.  Push back on managers and stakeholders if this is happening.

>* *You don't have a dedicated team area*. Your team should be sat together, preferably in your own room, with space on the walls to draw ideas and stick up cards and post-its.  As the project gets going, consciously 'hack the environment' to create a working environment conducive to team working.  You might upset a few people and challenge some long-standing working practices. But this is so, so important, and really should not be a big ask.

>* *There's no continuous integration environment*. Start right: with a continuous development environment.  If your teams are not insisting on this from the outset then you've probably got the wrong team.  So much about iterative software development is contingent on the ability to continiously deploy and run automated tests as you do.

>* *You have a separate QA department*. If your team's attitude to quality is to throw the software they've developed over the wall to a QA department, then they've not got the right attitude to delivering production-ready software.  You need to embed a quality culture into the team.

This is by no means an exhaustive list, but these are most common things to watch out for.
