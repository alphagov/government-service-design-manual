---
layout: detailed-guidance
title: Incident Management
subtitle: Available, stable and secure even under pressure
type: guide
status: draft
audience: 
  primary: service-managers, tech-archs, web-ops
  secondary: service-managers, delivery-managers, developers
category: operations
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Operations
    url: /service-manual/operations
---

# Incident Management

Digital Services so good that people prefer to use them need to be available, stable and secure at all times. That means that quick, calm, and effective management of incidents is an essential component of managing your service.

Incidents may be of many kinds from the exposure of a possible security vulnerability to having your service taken offline in an attack, from a minor data replication issue in a database to a major infrastructure failure. In all cases it's vital that you:
* communicate proportionately and clearly to your users if there's any impact on their experience of the service or the security of their personal data
* respond calmly and thoroughly to establish and resolve the root cause of issues
* fulfil your responsibilities across government, such as providing information to GovCert and to your senior management so that government as a whole can learn from these experiences and improve its services as a result

## Preparation

Preparation for incidents is vital. That starts with doing whatever we can to avoid them and runs through to rehearsing how we will respond in various scenarios.

### Understand the risks
A realistic understanding of risks, threats and the general landscape your service is operating within will help you prepare.
In all of these aspects proportionality is vital and your service team and stakeholders needs to make a call about risks, failure scenarios, and what you might need to do in different situations. For example, if your service fails out of hours but is unlikely to have significant impact on users is it okay to fall back to a friendly error page and avoid tricky late night operations, or is it vital that it's always available? There will be a significant cost to the latter, but some services are vital and will need support staff with detailed technical knowledge to be available 24/7/365.

### Test your service
Effective testing of software for its correctness, performance and security will significantly reduce the chances of an incident. Ensuring that you've got appropriate flexibility in your infrastructure and that it's designed for resilience will help you avoid risks from individual components failing. 
Prior to launch and at regular intervals you should test what does happen when individual or multiple components fail. If a server fails are others able to pick up the slack? If a user's operation can't be completed are they presented with a useful explanation?
Online video provider Netflix have become well known in the web development community for their "Simian Army":https://github.com/Netflix/SimianArmy a collection of tools including the "Chaos Monkey" that they can deploy to randomly switch off parts of their infrastructure. This is run in their development and production environments and ensures that they are building systems that degrade gracefully (or just keep working) even when components fail. Read more about Chaos Monkey on the "Netflix TechBlog":http://techblog.netflix.com/2012/07/chaos-monkey-released-into-wild.html
For more on testing see https://www.gov.uk/service-manual/agile/quality

### Have a process
If an incident occurs there will be many things that may need to happen:
* The impact on users will need to be understood
* Appropriate technical steps will be needed to fix, lock-down or otherwise change the service
* You may need to co-ordinate responses with suppliers such as your infrastructure provider
* Senior Stakeholders will need to be informed
* Decisions will need to be taken rapidly

It is easy for even a small incident to generate confusion, and it's vital to have a framework for how you'll respond. That should cover:
* What kinds of incidents do you have and how do you understand their severity? Is this a small technical hitch that you can fix in the next release or is it a serious security incident that needs your Senior Information Risk Officer's immediate attention?
* Who needs to be informed about each type of incident and how will they be informed? We use email lists for different types of  incidents that include senior technical and management people.
* What do we know? There should be a clear register of what is known and what the questions are, next steps to take and who's doing what. Whether that's a shared online document or a list on the wall, everyone working on the incident should have a clear point of reference.
* Who is in charge? Lots of decisions are needed in particular incidents and lots of information may be generated.  An individual with appropriate skills to understand all elements of the problem should be in charge and able to make the immediate decisions.
* Who's managing communications? Clear communication channels with users, stakeholders and decision makers are vital but it's easy to generate distracting noise. An individual should be in charge of running those communications. This may be the same person who's in charge of decision making, but that will depend on the scale of the situation. If it is a different person they should be reporting to the person making the decisions.
* How do you record incidents and follow-up to ensure that lessons are learned and next steps taken?

### Rehearse Failures
Automated testing can take you so far, but it won't tell you how your team responds under pressure. Consider running "Game Day" exercises where you simulate failures without the team's knowledge and test how your response mechanisms work.

## After The Fact

Any incidents you deal with will be an opportunity to learn and will reveal work that could be done to further mitigate risks.

We've found it very useful to have a standard format for incident reports that are stored in a known location so anyone on the team can refer to them. These reports are live documents with an initial version generated as soon as is practical during or immediately after an incident, and a more complete version produced as the dust settles. They cover:
* What happened? What was the timeline? Who did what?
* The user impact
* The root cause
* Actions required in response, from technical work to organisational changes
* Any other relevant notes

It's also important to make sure that all relevant people are aware of the incident report, that there's an opportunity for questions to be asked, and that the actions are clearly owned. 

We hold postmortem meetings and usually invite the SIRO, the product owner, the technical owner of the system and the team members involved in the incident. This is where the incident report is agreed and actions are assigned. The postmortem should happen as soon as practical after the incident so that memories are fresh and actions can be taken quickly. A timely effective meeting is more important than getting everyone there, so there's some flexibility in attendance.
