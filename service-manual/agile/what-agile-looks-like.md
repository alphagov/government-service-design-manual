---
layout: detailed-guidance
title: What agile looks like
subtitle: Common features of agile projects
category: agile
type: guide
audience:
  primary:
  secondary: service-managers, designers, delivery-managers, developers, tech-archs
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
exclude_from_search: true
---

Agile can be a liberating way of working. It won’t stop you from using existing skills and knowledge, but it will require your team, users and stakeholders to start working together in new ways.

## Understand your users


![Real people will use your product](https://farm8.staticflickr.com/7177/6987029385_410a1c8d26.jpg)

Prioritise features for users over everyone else -- including your big, scary stakeholders -- and ask for their feedback early and often. Really listen to your users. Even when they tell you things you don’t want to hear or disagree with.

If possible, use data from real people that are using your product and let it influence the direction of your project. Constantly put users first.

## “What do you want next Friday? What have we learned last week?”

![A sprint backlog, courtesy of http://www.flickr.com/photos/psd/](https://farm9.staticflickr.com/8043/8100964765_acf4032d09.jpg)

Iterate often. Build something that strives to achieve the most valuable user need and show it to the users, listen to their feedback and improve it. Keep doing this until you have something so useful that they wouldn’t be without it.

It might sound like over-simplifying the complexity of software development and project management, but agile development is all about “what do you want by next Friday?”

The process of producing incremental, production-ready software allows your team to:

* give value to their users and stakeholders regularly
* shorten feedback loops that could be longer if using a waterfall methodology (where you only move on to the next phase when the phase you’re working on is complete)
* think about what features are the next most important to produce
* direct their efforts on creating usable software

Run a [retrospective](/service-manual/agile/running-retrospectives.html) at the end of each delivery cycle or [sprint](/service-manual/agile/features-of-agile.html) to review what worked or what could be improved.

Your team will continue to learn through delivery cycles and improve throughout the project.

## Small, agile teams

![The unit of delivery is the team](https://farm9.staticflickr.com/8374/8451589322_e9f612cf5b.jpg)

Small teams of between 5 to 10 people are often more productive and predictable than larger teams. Forget man-days (the amount of work produced by an average worker in a day) and think about your team as a unit.

A good team includes members with all of the skills necessary to successfully produce software. A fully functioning team has 3 main roles:

* product manager -- (a role probably performed by the [Service manager](/service-manual/the-team/service-manager.html)) responsible for delivering return on investment, usually by creating products that users love
* [delivery manager](/service-manual/the-team/delivery-manager.html) (aka Scrum master or project manager) -- the agile expert responsible for removing blockers (things slowing a team down), they also act as a facilitator at team meetings
* [team members](/service-manual/the-team/index.html) -- self-organising and multi-disciplinary, they produce user stories, carry out the product manager’s vision and are responsible for estimating their output and speed

Encourage your team members to pair up, as working together is beneficial. 2 people working on 1 thing will:

* produce better software solutions
* encourage better quality controls
* spread knowledge across the team

A good team means you’re able to estimate your output, or speed, very effectively and consistently. You can then plan much more accurately.

## Fail fast

![Failing, so fix it!](https://farm8.staticflickr.com/7189/6875228285_9b2409663f.jpg)

Regularly releasing little pieces of code will:

* improve quality
* improve visibility
* reduce cost to market

Agile techniques don’t guarantee success -- you can still fail!

But these techniques do allow you to spot problems earlier on and resolve them. You can resolve issues, and stop issues from happening, by:

* releasing working software to your users regularly -- it allows you to get feedback quickly and hear or see what they think; if the product is wrong you can easily change direction and iterate
* demonstrating value to your sponsor with regular releases -- if your software is rarely released you run the risk of creating a ‘too-big-to-fail’ service that shouldn’t be released, but must be released anyway 
* checking your teams’ progress -- if your teams’ speed is still inconsistent after the initial 4 to 6 sprints, then something needs fixing (possibly unknown complications or poor estimation with timings)
* using test-driven development (writing tests before you develop the features to be tested) to highlight issues with quality early on -- establish the issues, baseline metrics, and monitor throughout the project

Don’t be afraid to fail or experiment. Learn to fail, and create a culture that learns from failure.

##Continuous planning

![Planning session](https://farm9.staticflickr.com/8001/7113823877_80c4dfb613.jpg)

It’s a myth that you don’t plan on agile projects.  The freedom of agile projects does not come free: you have to plan. You just plan differently and continuously.

Base your agile planning on solid, historical data, not theories or opinions. Your plan must continuously demonstrate its accuracy: nobody on your agile project should take it for granted that the plan is workable.

Your teams should plan together, on at least 2 levels:

* release level -- identify and prioritise the features you must have, and would like to have by the deadline
* iteration level -- plan for the next features to implement, in priority order (if features are too large to be estimated or delivered within a single iteration, break them down further)

Review these plans after every sprint and adjust them based on:

* the progress of the previous sprint
* any new facts and requirements

##Common situations to look out for

![Do go here!](https://farm9.staticflickr.com/8424/7503675672_72ff8a1fa9.jpg)

If your team is new to agile, be wary of familiar situations and reactions from having to do things differently. These situations have a bad smell about them and will undermine your project and its chances of success:

* your core team isn’t full-time or is working on multiple projects -- your team is the unit of delivery and you need 100%, so push back on managers and stakeholders if this is happening
* you don’t have a dedicated team area -- sit your team together, preferably in your own room, with space on the walls to draw ideas and stick up cards and post-its
    * rearrange your workspace or use tools in innovative ways to improve your teams’ working environment and increase productivity -- you’ll challenge some longstanding working practices, but this is very important
* there’s no continuous integration/development environment -- if your teams aren’t insisting on this from the beginning you’ve probably got the wrong team:
    * iterative software development is, in many areas, dependent on the ability to continuously deploy and run automated tests
* you have a separate quality assurance (QA) department -- if your team pass software they’ve developed over to a QA department, they’ve got the wrong attitude to delivering production-ready software; embed a quality culture into the team

This is by no means a complete list, but these are most common things to watch out for.

*[QA]: quality assurance
