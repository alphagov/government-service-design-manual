---
layout: detailed-guidance
title: User-centred design in alpha and beta
subtitle: Combining design and research to create user focused services
category: user-centred-design
type: guide
audience:
  primary: service-managers, designers, performance-analysts, user-researchers, qa
  secondary: developers
status: draft
phases:
  - alpha
  - beta
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: User-centred design
    url: /service-manual/user-centred-design
---

Carry out user research during every stage of your project. Do it continuously through each stage --- don’t leave it as something that happens at the beginning and end of phases.

Doing user research continuously will:

* keep your team concentrating on real user needs
* help teams design products which are prioritised by user needs
* help teams iterate products in response to user feedback

## Designers and researchers are essential

Each team needs to have a designer and a researcher working together, both having an active interest in the design and user insight.

This isn’t about a user researcher ‘testing’ the work of the designer --- having designers and researchers work together means that:

* they’re sharing all the knowledge that’s been gathered about users, so any decisions made on the product design will be influenced by real user insights
* the researcher can provide more information and more regular feedback --- helping the entire team to iterate and prioritise, and create the best possible product
* they’ll continually create experiments that allow them to test and prove whether or not design approaches create products that users find understandable and desirable

Make sure that you have a team member leading on research that sits with the team. They may also do other work (design, write content etc), but are also responsible for running research at least every 2 weeks.


## Researching in design iterations

Don’t spend more than 2 weeks working on design without testing it with real end users.

Each 2-week iteration should comprise of 3 stages:

* [build](#build)
* [measure](#measure)
* [learn](#learn)

![A sketch of the Identity Assurance (IDA) team’s fortnightly iteration plan](/service-manual/assets/images/user-centred-design/process-wheel-small.png)

## Build

With every iteration, create materials to be included in the next round of research. In the alpha and beta phases, these will typically be prototypes that vary in accuracy --- from paper prototypes through to working prototypes in code.

Use [existing GDS code](https://github.com/alphagov) and other open source frameworks as:

* it’s usually quite quick to create prototypes of design concepts in code
* doing so will usually gain more helpful insights from users than paper prototypes

Accept that you’ll design and build a number of prototypes to test with end users, and that these prototypes will be thrown away. This gives the team freedom to explore a range of different concepts and gain a better understanding of what works best for end users. Allow time for these experiments particularly in the early stages of the project.

In the early stages of your project the prototype will probably look unfinished and the design will still require a lot of work. Stick to the ‘test every 2 weeks’ rule, instead of waiting until you’re completely happy with the design or have something that feels finished.

Try to make sure that the prototype is completed at least 24 hours before your testing sessions begin, as it:

* allows time for making tweaks and small changes
* reduces the risk of technical failures in the research sessions

## Measure

There are many ways to create ‘experiments’ that will help you to answer questions about your product, your end users and your design ideas.

Your researcher will help choose the best research methods for answering each question. With their help, you’ll probably end up using a number of different techniques on your project.

Whatever your project, be sure to do fortnightly user research.

### Fortnightly user research (eg ‘Testing Tuesday’)

In these sessions you’re able to:

* have people interact with the design you have been working on (usability testing)
* explore any specific issues you‘d like to gather more information on through a depth interview (one-on-one interviews that produce deep insights on user needs and behaviours)

Carry out research sessions with users every 2 weeks on the same day. Doing user testing at predictable and regular intervals means that your team members will be able to schedule time to watch sessions.

A good approach is to schedule 5 one-on-one interviews for that day, each interview being 1 hour in duration. (Schedule a 6th participant as a ‘spare’ in case some participants don’t attend earlier in the day or you have a poorly recruited participant).

Conduct these sessions in either:

* your own offices
* a research laboratory (get one from a commercial supplier if there’s not one readily available)

Make sure that you have both facilities booked and participants recruited as early in your project as possible for the expected duration of the alpha and beta phases.


#### Before the session

To prepare for these sessions, you need to:

* define your research questions --- what do you want to learn from the round of user testing?
* prepare some theories about your design, eg changing the words on the button will encourage people to read the page more carefully
* decide how you’ll know when these theories have been proved or disproved, eg you’ll know this is true because you’ll observe people taking more time to read the page
* decide the type/demographic of users you want to talk to
  * age
  * location
  * suitability to task
  * other factors
* start recruitment at least a week before testing (GDS recommend that you find participants through a recruitment agency)
* prepare a discussion guide that explains how the sessions should be run and the important research outcomes
* create a participant release form that authorises the recording of the session
* send a schedule for the day to project team members, inviting and strongly encouraging them to attend sessions
* arrange for live streaming of the sessions online via a service such as GoToMeeting, so team members who can’t attend in person can still watch it

Prepare your prototype for the sessions by making sure that it’s:

* ready to be seen by users
* ready to be tested end-to-end
* usable and accessible from the testing venue (eg firewall restrictions, screen resolutions, compatible browsers)


#### During the session

In each user testing session:

* make sure the session is being recorded
* make sure that the live stream of the session is available for external viewers
* keep research theories at the top of your mind so you don’t forget any important topics
* write down notes on post-its of important observations
  * use yellow, super-sticky post-it notes for observations
  * only 1 observation per post it note
  * use a marker pen and write in capital letters (easier to read when writing quickly)
  * if you’re not sure if it’s important, still write it down
* get a written transcript of the session (either during the session or have it transcribed at a later time)

A day of research should be followed by a period of analysis before making any significant design changes. The analysis stage is described in detail below.

### Guerrilla research

You can use the time between more formal testing to conduct some guerrilla-style testing, getting some initial feedback on your revised designs.

Guerrilla testing normally involves taking your prototype into a coffee shop or other public location and finding volunteers who'll give you some quick feedback. Guerrilla testing participants aren't necessarily representative of your target audience, but they can provide a quick sense-check in between formal testing sessions.

### Other research methods

There are many other research methods you can use to supplement your ‘in person’ qualitative research and to address specific research questions. Using a different technique can provide better insight into a particular research question, or validate insights from fortnightly research with a larger audience.

Here‘s a [list of some techniques](/service-manual/user-centred-design/user-research) you can use and information on how and when you should use them.

## Learn

Now it’s time to see what you can learn from your user testing and research. Make the best possible use of this information by:

* [analysing](#analysis)
* [sharing](#share-your-findings)
* [communicating](#communicating)

### Analysis
It’s important to properly debrief after spending a day doing user testing as you’ll have seen a lot of people and a lot of different reactions.

Although it takes time to analyse sessions in detail, doing so means that you’ll:

* have a clear understanding of what you have learned
* know what the implications on design will be
* be able to make sure that important observations are not lost

#### Affinity sorting

To do this correctly, you need to:

* gather all the post-it note observations created during the testing sessions and stick them up on a big wall
* group the observations into similar themes
* create findings --- a statement that summarises the observations --- for each group, and write it down on a post-it note, sticking it on top of the group
* try to address the theories and decide if you have enough information to state a theory is now proven or disproven
* a theory needs further qualitative/quantitative testing

At this stage you’re aiming to create a full set of insights (things that you’ve observed and learned), so don’t start designing solutions just yet. You should allow several hours for this analysis process.

Work with anyone that’s seen the user testing session to analyse observations, as it will help to:

* gather findings
* confirm things that people have seen
* form a consensus on the findings

![Picture of post it notes on wall](/service-manual/assets/images/user-centred-design/post-it-wall.png)

#### Actions

Decide what you’re going to do about each finding, if anything, and when you’ve proven a theory, make sure that it’s recorded and shared with the project team.

If a finding needs further action, determine what is required to address it in the next design iteration. Write this on an orange post-it note and stick it on top of the finding.

#### Prioritisation

When you’ve decided on what actions to take, use a prioritisation method (like [dot voting](https://en.wikipedia.org/wiki/Dotmocracy)) to decide how you’ll spend your efforts for the next iteration.

### Share your findings

Put user testing findings and actions in a place that’s easily accessible to the project team, as it’s likely that you’ll come back to the findings to:

* remember what you discovered
* see how a particular theme or feature has developed over the course of the design iterations

You can document your findings for others to see in many different ways, like using a:

* shared document
* blog
* story wall (eg [Trello](https://trello.com/))
* research catalogue

It’s also useful to keep a record of what you tested with users, eg if you’ve prototyped in code, keep a folder in your repository for each round. You may also like to keep screenshots of the prototype.

GDS intend to create a format for sharing research findings between projects across government. While GDS work on the best way to do this, keep your research findings in an accessible shared format so findings can be shared and learned from in the future. The findings should, wherever possible, be easily understood with no need for an explanation.

Move any actions onto your story wall so that the whole team can see what you’re working on and who's working on it.

#### Video

Use videos to demonstrate findings and store the session videos in a safe place, ideally where they can be found and watched by team members. If you store the videos on a public service (like Vimeo), make sure they are adequately protected.

### Communicating

You need to discuss the project and its progress with others. You can do this by:

* [showcasing/presenting regular updates](#showcase)
* [always publishing the design](#publish-the-design-as-it-stands)
* [having retrospectives](#retrospectives)

#### Showcase

Just as a product team often presents a showcase at the end of each iteration, so should the research and design team. This will make sure the wider project team has an understanding of:

* the research and design work being done
* what has been learnt from research
* how research is changing the design of the service

This is an opportunity for members of the project team to catch up if they haven’t been able to attend user testing sessions, as well as an opportunity to ask questions.

Depending on how your team works, you might put your presentation into the project team’s showcase or you may have a showcase just for research and design.

#### Consider your more distant stakeholders

There may be other people beyond your more immediate project team who are interested in the outcomes of your work. You can communicate and/or work together with these people by:

* inviting them to your showcases
* publishing a blog
* maintaining a shared document
* sending out regular update emails, using screenshots from the prototype to help people see what has been tested

#### Publish the design as it stands

Make it possible for your team to always access the latest version of the prototype, but make sure it’s appropriately protected. Team members may need to check something or use the prototype to talk about the project with their own stakeholders.

Put a message across the top of the prototype to remind people that it’s a work in progress.

#### Retrospectives

It’s important to reflect on your progress at regular intervals. Schedule a retrospective where research and design team members have the chance to discuss the process openly and can suggest changes to the process. Assign owners and due dates to the actions to make sure that change occurs.




