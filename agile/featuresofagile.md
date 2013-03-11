---
layout: detailed-guidance
title: Features of agile
subtitle: Sprints, stand-ups and other regular meetings
section: agile
type: guide
audience:
  primary: service-manager
  secondary: designer, developer, tech-arch, analyst, researcher
status: draft
phases:
  - alpha
  - beta
  - live
---

Here are some of the common features of the agile development methods we've used at GDS.

## Sprints

Sprints are the rhythmic foundation of an agile process that uses [Scrum](http://en.wikipedia.org/wiki/Scrum_(development)). They usually last 1-2 weeks.

Factors in deciding the length of your sprints include:

* how long is the project? If it’s only a couple of months then shorter sprints give you greater flexibility and visibility of what’s going on and an opportunity to iterate and adaptively plan
* shorter sprints allow teams that are new to Scrum get used to the regular meeting formats and familiarise themselves with the process. For those teams that a familiar with Scrum then moving to longer sprint cycles reduces the overhead of Sprint Planning, Review and [Retrospectives](/agile/runningretrospectives.html)
* if you choose to release to production only at the end of each sprint (it’s optional, but perfectly OK), then how often do you want to ship production ready code?

Teams commonly start with a Sprint 0, which is used to get the development and working environments setup, agree some of the design principles (technical, product, interaction, content) and prepare the backlog for the first sprint.

Note: There are other agile methodologies that do not rely on sprinting such as [Kanban](http://en.wikipedia.org/wiki/Kanban_(development)). However it’s common for teams to start with Scrum and sprint.

## Standard meetings

### Daily Stand Up

A daily meeting with all of the team. It should take no more than 15 minutes and typically takes less time. It’s best if you do it standing up, in a semi-circle in front of the project wall. This helps keep it short and allows participants to point at story cards on the wall to keep things on topic.

The normal format is each person answers the following 3 questions:

* what I worked on/delivered yesterday
* what I am working on today (and help I might need)
* what’s blocking me (i.e. stopping me delivering a story card)

(It sometimes helps to have an object that you (gently) throw randomly to the a person in the stand-up to signify they should speak next. It keeps people on their toes, the randomness stops it feeling too repetitive and allows the last person that speaks to choose the person they wish to hear from next. At GDS we use cuddly toys or a piece of fruit. It’s a bit of fun. You don’t have to this, it’s just something to experiment with.)

It’s OK to ask a waffler to wind up. If people try to solve issues within the stand up then team members can call time on the conversation and decide to convene a huddle after the stand-up to discuss in a smaller group. I think it’s a sign of a good, constructive stand up if you spin up a couple of ad-hoc huddles afterwards and I believe it shows people are getting something out of the meeting.

### Sprint planning

Sprint planning is done at the start of each sprint. It requires stories to have been written in advance with acceptance criteria.

It’s the Product Owners job to read out the stories and explain the acceptance criteria in priority order. It is the team's’ job to understand the story and acceptance criteria and agree the number of stories they will commit to within each sprint and agree the tasks needed to complete it.

There is a good description of this meeting on the [Agile Learning Labs website](http://www.agilelearninglabs.com/resources/scrum-introduction/) and how it has two parts: the “what we will do?” and the “how we will do it?”

With bigger teams this can be a hard meeting to get right. Some people want to dig deep and question every story, others want to keep moving and don’t want to go into detail. Persevere! it’s an important part of the sprint and like Marmite, love or hate it, it’s good for you.

### Story time

Some teams choose to to write or refine their stories at a single, set time within their sprint cycle, others choose to do it over a couple of sessions. It’s up to you, but don’t miss it. It’s vital that you have good stories and there has been constructively discussed with relevant team members, subject matter experts and stakeholders in good time, ahead of sprint planning.

Make sure stories are well formed (i.e. don’t skip the “So that...” part because it’s hard), have a good, sensible list of acceptance criteria that supplement your teams’ standing criteria for ‘definition of done’ and are estimated. If stories are too big then split them into smaller stories. They stand more chance of delivering shippable code.

### Sprint Review

There needs to be a moment at the end of the sprint where the team get to demonstrate the work they have delivered. The whole team needs to feel on the hook for delivery.

### Sprint Retrospective

These meetings are important opportunities to iterate the team's working process. [Anna Shipman from GDS has written detailed guidance about how to run retrospectives](runningretrospectives.html).
