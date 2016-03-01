---
layout: detailed-guidance
title: Writing user stories
subtitle: How to write a useful user story
category: agile
type: guide
audience:
  primary:
  secondary: service-managers, designers, delivery-managers, developers
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

A user story briefly explains:

* the person using the service (actor)
* what the user needs the service for (narrative)
* why the user needs it (goal)

User stories are an essential part of the agile toolkit. They’re a way of organising your work into manageable chunks that create tangible value, and can be discussed and prioritised independently.

User stories are a technique which relies on other agile practices, including continuous delivery, and face-to-face communication with user representatives. It's not sufficient to simply 'set' a user representative: that user representative needs to be in the same place as the team and available to the team a sufficient amount of time.

User stories can be added to a product backlog at any point in the sprint cycle by any person in the team. It's up to the product owner to coordinate and prioritise them, and to select stories for each sprint at the start of each sprint cycle.

Discuss the stories ahead of sprint planning with:

* relevant team members
* subject matter experts
* stakeholders

## User story cards

{:.left}
![A user story index card](https://farm9.staticflickr.com/8372/8358344190_f48b88c254_n.jpg)

A user story is represented through a story card that has a title and a few lines of text.

Your story cards can be virtual, as well as actual cards. On a large product/service keep your stories in a digital format, and then turn them into physical cards as part of sprint planning.

When writing a user story, make sure the story is well-formed. Don’t skip the part explaining why there’s a need for a service just because it can be difficult.

You might want to include a list of acceptance criteria. These should be a reminder for things to test or check which may have come up during conversation, but they should not be used as a way of defining the scope of a story.

If stories are too big then split them into smaller stories as they stand more chance of producing shippable (ready to use) code.

### Structure
Story cards follow a standard structure:

* title
* actor
* narrative
* goal

They don’t capture every detail, but you should have a more in-depth discussion about a user story at the appropriate time.

{:.hide-headers}
|-------------------------------------------------------------------------------------------
| Name    |
|-------------------------------------------------------------------------------------------
| **Actor:**     | As a      | journalist
| **Narrative:** | I want to | see contact information relating to the news article I am reading
| **Goal:**      | So that   | I can get directly in touch with the press office about it
|-------------------------------------------------------------------------------------------

## Working outside-in using goals

Building useful software systems is hard, so you need to make sure you’re solving the right problem.

Agile methodologies emphasise an outside-in approach -- what outcome is your service user trying to achieve? If you dive into the solution without a good understanding of your users, you risk solving the wrong problem, or coming up with solutions which don’t really work for your users in the real world. That’s why the most important part of a user story is the goal.

###Goal

Use this to help you decide whether the story is “done” or delivered, ie does the work meet the goal of the user?

When writing stories with your development team, always start by thinking about and discussing your users’ goals:

* why do they want to use your service?
* what are they trying to achieve?
* what need has motivated them to seek out your service?
* in what context do they use it -- at home/work/on a mobile phone/whilst caring for a child?
* how often do they use it?

Suzanne and James Robertson have excellent advice on this in the book *Mastering the Requirements Process* (3rd edition).

###Actor
Being specific about the actor will help you to break down interactions into logical chunks.

Sometimes the actor will be a user of your service, or the actor will be an administrator, technician or manager in your organisation.

Make sure you already have a good understanding of your users from your initial project work or existing research. If not, take the time to develop that understanding.

###Narrative
Use this as a reminder of the main interaction that needs to be addressed as part of the user experience. Remember that the story card does not need to spell out every detail.

<a name="promise-of-a-conversation"></a>

## Taking a rain check -- the promise of a conversation

[Agile teams prefer face-to-face conversation over detailed documentation](http://agilemanifesto.org/principles.html).

Face-to-face is:

* faster
* more accurate than written documentation
* allows developers to build up a detailed mental model of the user goals, workflows, constraints, and the many issues which must be taken into account when building a software system

The story card is just a placeholder, a promise to have a conversation when the time is right. Use the story card and some brief initial conversations to estimate the amount of time a story needs to be completed and then put it into an agile backlog.

When development work actually starts, consult the users or user representative to fill out the story details. A user story is the sum of all of these conversations, sketches and whiteboard diagrams -- not just the card. You don’t need to write down or archive your conversations, but translate them directly into automated tests and working code.

Using user stories in this way allows you to avoid ‘analysis paralysis’, the painful condition of trying to guess the details of some far-future goal.

## Acceptance criteria in user stories

Acceptance criteria can be used to determine when a story is done. These follow a conversation between developer and user or user representative and are considered before coding begins.

Only include these on the back of the story card if the team finds them useful for recording user assumptions they might later forget. Sometimes writing acceptance criteria on the story card is useful where the user or user representative is not immediately available but is no substitute at all for face-to-face conversation.

## Stories only work in an agile team

The success of a user story is dependent on regular face-to-face communication between developers and users or user representatives.

Your service manager and other user representatives need to be available to developers every day, and have enough time to think through and answer questions. Don’t underestimate how time-consuming this work can be!

## How to get user stories

Stories can come from many places, but the most common sources include:

* story writing workshops -- more commonly at the start of a project, the development team and stakeholders will get together to write stories
* user interviews with real users -- ideally, you will set up a user panel which the development team have ongoing access to
* user representatives embedded within your team -- this may include the service manager or product owner
* observation -- watch real users using your service

See [Chapter 4 of User Stories Applied](http://www.mountaingoatsoftware.com/books/user-stories-applied).

## Further reading

1. Mastering the Requirements Process, 3rd Ed, Suzanne Robertson & James Robertson, 2012
3. Twelve Principles of the Agile Manifesto -- [http://agilemanifesto.org/principles.html](http://agilemanifesto.org/principles.html)
4. User Stories Applied, Mike Cohn, 2004. Free chapter on "[Writing User Stories](http://www.mountaingoatsoftware.com/system/asset/file/259/User-Stories-Applied-Mike-Cohn.pdf)" available at [http://www.mountaingoatsoftware.com/books/user-stories-applied](http://www.mountaingoatsoftware.com/books/user-stories-applied)
