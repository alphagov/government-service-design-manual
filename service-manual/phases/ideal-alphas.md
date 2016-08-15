---
layout: phases
title: Ideal alphas
subtitle: What an alpha project can look like
phase: alpha
category: guidance
status: draft
phases:
  - alpha
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Phases of service design
    url: /service-manual/phases
  -
    title: Ideal Alphas
    url: /service-manual/phases/ideal-alphas
exclude_from_search: true
---


{:.intro}
An [alpha](/service-manual/phases/alpha) is a process that happens at the beginning of a project. It allows us to:

* explore the major risks to the project
* discover whether the project is workable
* get some ideas about the cost of the project

It is primarily defined by two activities:

* identifying risks
* prototyping solutions

This document covers how to go about executing an alpha project.

Following this guidance will help you create a successful alpha project. You should bear in mind that these guidelines are flexible and are intended to direct you towards achieving your project goals simply and quickly.

## Preparing

### Identifying goals

Before planning an alpha, the team must work out what the goals of the alpha are.

For example, some of your goals might be:

* attempting to prove that you can replace an existing digital service
* prove that you can replace a paper service with a digital service
* create a new digital service

The department should communicate their goals in the request for proposals.

The team will probably develop their own additional goals as well.


### Identifying risks

The reason for running an alpha is to identify risks to the wider project early.

This means it should identify a project's biggest risks and explore those risks to understand them better.

In most government projects, the risks fall into 3 main categories:

- design
- business process
- technical risks

The design risks cover the issues that come with running a user-centred design process within an iterative model. We find that it often takes several attempts to get the service proposition correct for the citizen. You'll need to think about things like:

* how easy is it to use
* will the user get through it correctly the first time
* what do errors look like
* how will I analyse the user research
* how do I build prototypes quickly

The business process risks include:

* how the department will integrate with the new process
* do they have sufficient fraud checks
* are they set up to handle support calls
* how to operate the service

The technical risks tend to be about integrating into the existing systems. For example:

* what sort of ongoing connections will be necessary
* are there special hosting arrangements, eg high security hosting

In our experience, there are always technical risks to all projects. Most government departments do have a history of solving technical risks and have entire teams dedicated to them. Business process and design risks are far less well understood, and so we encourage alphas to focus on exploring those risks with prototyping over full service integrations.


### Identifying the team

The well-rounded team will need to have various skills, including:

- design (service, UX, content as needed)
- user research (user testing)
- prototype development
- service integration development
- delivery management
- business analysis

The team will probably be bigger than this, and the service may supply its own team members who will normally be learning during this process, and thus probably won't be contributing significantly.

The team will work together in an agile manner, doing iterations of either 1 or 2 weeks. For a short project we'd normally recommend 1 week iterations where possible.

The team must have the core skills to:

* build a wireframe or prototype user journey in 1 week
* host it on the internet
* conduct some user research
* analyse the findings

The Business Analyst and Delivery Manager roles are there to ensure that the findings can be fed back into the team and the wider organisation.


### Identifying process

Each alpha will be different, because the risks and goals are different.

But we find that there are common aspects in how good alphas are run.

We want to maximise the amount of learning that is happening in the alpha. So we tend to expect alphas to be short, about 8 weeks in length.

We expect alphas to normally consist of:

- an inception
- a series of iterations of design, development and test
- either:
  - alpha termination
  - alpha to Beta transition

## Executing

### Inception

We believe that alphas should start with an inception process. An inception is designed to bring the team together and share the knowledge about the risks and goals identified.

We tend to recommend a short inception -- between 4 to 7 days seems to work well.

The inception should be run by the supplier team, probably by the Delivery Manager. It will look at a variety of business, technical and user aspects of the project. Each supplier has a different format for these, but we see fairly common outputs including:

- a shared understanding of the service
- User Personas (who will use the system)
- a clarified current business process (where applicable)
- a clear set of goals and tradeoffs
- some vision exercises
- a clear understanding of the existing technical estate
- a prioritised backlog of work for first iteration

The facilitator will of course give up time to various specialists to run individual sessions. The aims include:

* the entire team get to know one another
* beginning to form a team identity
* beginning to form a team understanding of the business itself

The inception may well cover introductions to agile itself, if the project team need it, and should also have a demonstration and retrospective at the end of it.


### Iterations

Alphas are often about identifying competing hypotheses about the user journey or business process. We expect to produce many prototypes to test the hypothesis, and run tests in a short amount of time so we can learn, change and retest. These feedback loops are crucial.

At this early a stage of a project, it is rare to be able to choose a single flow or design. Because of this, we encourage the building and testing of many prototypes. This allows you to quickly get feedback on the best experience for the users and to test competing ideas.

We expect that a first draft of a user flow will have issues. You should seek to revisit the design, research and test process on each flow during many iterations.

The iterative process we would recommend should be a straight up agile process combined with UX and user research.

We would encourage the following format:

- Day 1
  - Retrospective (1 hour)
  - Iteration Planning (2 hours)
  - Start iteration
- Day 2
  - User research on previous iterations prototypes
  - Analysis of test results
- Day 5
  - Demo

This structure might move around a bit, but you will want to do the user research early in the iteration to give plenty of time to feed into the next iteration. This enables fast turnaround on the stories.

Working at this pace will result in:

- user research in iteration 1
- stories completed in iteration 2
- further research on updated prototypes in iteration 3

This assumes fairly high definition prototypes. These may have backend systems that mock out service integrations, but which contain actual working logic.

If you need to iterate faster (as is likely at first), you can start with much lower fidelity prototypes (like paper). This can let you test various flows and allow much faster iteration of research and prototype development.

We would expect that you will do a demonstration of the service journey so far. This happens at the end of each iteration, and highlights what the team has learnt and what the team is planning on doing next.

We've often found that aligning the iterations to Monday to Friday is not a best practice. It can cause logistical issues:

* the first and last day of the iteration are the most critical for the entire team
* Mondays/Fridays have the most absence

Starting the iteration on a Wednesday, with User testing on Thursday and demo the following Tuesday has worked well in the past.

### Ending the alpha

The value of the alpha process is that the team can use it to identify early the biggest risks to the Beta part of the program.

It is important to remember that the most valuable goals from the alpha process are the increased domain understanding and the team forming behaviours.
If the code produced during the alpha is not high-enough quality, then it's fine to throw it away and rebuild. For example, the team may elect not to use test-driven development during this stage. Those activities represent investment in quality. Remember that a primary goal during the alpha is learning. We might not be ready to make the investment in creating production-ready systems. In particular, that investment might be premature if the team learns they are not solving the right problem.

Another output of the alpha can be the early decision not to continue into Beta. This represents a successful alpha since it reduces wasted time and money. It may well be the case that at the end of the alpha you go back to the discovery outputs and start a new alpha with a different focus or go back and perform another discovery.

However, assuming a successful alpha project, then early in the alpha process, the Delivery Manager and the Service Manager will start to work on the business proposal for the Beta program.

The reason for starting so early on the Beta program proposal is that often the team would like to continue without a break into the beta program to maintain the momentum and so need to allow for the time to gain approval for further spending

The last demo of the alpha should have the attendance of the project board or senior management team from the service. At that demo, you are able to show what the alpha has achieved and show the feasibility of the beta program.

Assuming that the status reports to project board and the show and tell have got approval, then as the project draws to an end, the team will want to have a final retrospective, and produce a backlog of epic stories for the Beta program.

{:.phase-nav}
* **[Next phase: beta](/service-manual/phases/beta.html)**
* [Previous phase: discovery](/service-manual/phases/discovery.html)
