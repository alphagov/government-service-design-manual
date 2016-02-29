---
layout: detailed-guidance
title: Features of agile
subtitle: Sprints, stand-ups and other regular meetings
category: agile
type: guide
audience:
  primary: developers
  secondary: service-managers, designers, delivery-managers, web-ops, tech-archs, performance-analysts, qa, content-designers
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
    title: Agile
    url: /service-manual/agile
exclude_from_search: true
---

Some common features of the agile development methods we’ve used at GDS.

##Sprints

Sprints are part of an agile process that uses [Scrum](https://en.wikipedia.org/wiki/Scrum_(development)), a framework for teams developing a product.

Sprints are about planning what you and your team will achieve and when you will achieve it by. Each member of your team should have a task to complete within each sprint.

Each sprint is of equal length and usually lasts between 1 to 2 weeks, but you can use longer or shorter sprints in your project.

When deciding the length of your sprints, consider:

* the length of the project -- use shorter sprints for shorter projects, giving you greater visibility of what’s going on, the opportunity to iterate and the flexibility to adaptively plan
* team members that are new to the Scrum process:
    * shorter sprints means those that are new to Scrum get used to the regular meeting formats and processes faster
* how often you release code to production  - if it’s only released at the end of each sprint decide how often do you want to ship production-ready code

Start the process with a sprint 0. This sprint is about preparing your team for the first sprint. Use a sprint 0 to:

* get the development and working environments set up
* agree some of the design principles (technical, product, interaction, content)
* prepare the product backlog for the first sprint

While there are other agile methodologies that don’t rely on sprinting, such
as [Kanban](https://en.wikipedia.org/wiki/Kanban_(development)), it’s common
for teams to start with Scrum.

## Standard meetings

Agile processes have 4 different types of regular meetings:

* daily stand-ups
* sprint planning
* sprint reviews
* retrospectives

### Daily stand-up

Have daily meetings with your team lasting no more than 15 minutes. It’s best if you do it standing up, in a semicircle, in front of your project wall. This will help you to keep it short and allows your team members to point at user story cards on the wall to keep things on topic.

Each member of your team should answer:

* what I worked on/produced yesterday
* what I am working on today (and help I might need)
* what’s blocking me (ie stopping me from finishing a user story card)

Ask people to keep it brief and don’t be afraid to remind them of this. If people try to solve issues during the stand-up, stop the conversation and arrange a huddle after the stand-up to discuss in a smaller group.

### Sprint planning

Sprint planning is done at the start of each sprint. It requires user stories to have been written in advance with acceptance criteria.

The product owner should read out the stories and explain the acceptance criteria in order of priority. It’s the job of the team to:

* understand the story and acceptance criteria
* agree on the number of user stories they’ll aim to achieve within each sprint
* agree on the tasks needed to complete it

A good description of sprint planning is on the [Agile Learning Labs website](http://www.agilelearninglabs.com/resources/scrum-introduction/), explaining its 2 parts:

* what we will do?
* how we will do it?

This meeting can be hard to get right with large teams. Some people want to
dig deep and question every story; others want to keep moving and don’t want
to go into detail. Persevere!

### Sprint review

This is a chance for team members to demonstrate the work they’ve produced during the sprint. You can include stakeholders in this meeting and also use it to inform stakeholders which user stories haven’t been completed and why.

### Retrospectives
These meetings are important opportunities to iterate the team’s working process. [Learn how to run retrospectives](/service-manual/agile/running-retrospectives.html).

### User stories
User stories are essential to the development of your service and the user experience, so it’s important that you know how to write them well. [Learn how to write a user story](/service-manual/agile/writing-user-stories.html).

### Testing in agile
Testing in agile is about constantly improving the quality of your service. [Learn about Agile testing methods](/service-manual/making-software/testing-in-agile.html).
