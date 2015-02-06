---
layout: detailed-guidance
title: Security testing
subtitle: What is security testing and how to use it
category: making-software
type: guide
audience:
  primary: developers, qa
  secondary: web-ops, tech-archs
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
  	title: Making software
  	url: /service-manual/making-software
---

Security Testing intends to find flaws and vulnerabilities in an IT system, as well as the people and processes that surround the system. The practice includes:

* formal security testing for live service accreditation, eg [penetration testing](https://www.gov.uk/service-manual/operations/penetration-testing.html), which aggressively attacks a system to check its defences, and [DDOS](http://en.wikipedia.org/wiki/Denial-of-service_attack) (Distributed Denial of Service) durability testing, which tries to make a service unavailable 
* system feature testing, eg quality assurance testing of password functions


## Finding the right skills

People with the appropriate skills, tools and accreditation can carry out security testing. They can be developers, security specialists, software testers or quality analysts. Specialist suppliers offering security testing services can be found on the Government [Digital Marketplace](https://www.gov.uk/digital-marketplace).
 
The test and quality analyst team must give appropriate coverage to internal and cyber security features on areas such as user administration, password functionality, role based access control and auditing. This testing of security features is different to formal security testing, which may be out of scope for the quality assurance team.


## Distinction between 'formal security testing' and typical assurance testing of security features:

<table>
  <thead>
    <tr>
      <th scope="col"></th><th scope="col">Formal security Testing</th><th scope="col">Security related assurance</th>
    </tr>
  </thead>
  <tr>
    <td>Objectives </td><td>Test your system meets the required level of security to be deployed as a live service, with regard to software, hardware, users, process, etc... Testing is concerned with whether your service is running a secure platform, rather than whether software acceptance criteria are met.</td><td>Test your system features meet your security acceptance criteria (from a functional and non-functional perspective). Features are defined by user stories or requirements, and you’ll test these alongside other stories/requirements.</td>
  </tr>
  <tr>
    <td>Examples of areas tested </td><td>Penetration testing and DDOS.</td><td>User administration, login and password functions, and role based access controls (RBAC).</td>
  </tr>
  <tr>
    <td>Skills needed  </td><td>CESG accredited specialised teams. These teams are typically provided as a commodity service to projects that require them.</td><td>Quality analysts or testers, who work as part of an agile team. They’ll test security features alongside all other in-scope product features using the same toolset.</td>
  </tr>
  <tr>
    <td>Outputs </td><td>Identified areas of risk with recommendations for potential corrective actions. Department security units use testing reports to make decisions regarding ‘approval to operate’ status.</td><td>Identified defects in security features. Development teams use test reports to determine whether acceptance criteria are met, and whether user stories are ‘done’.</td>
  </tr>
</table>


## Timeline for security testing

Security must be taken into account throughout the development process, with regular tests and reviews. Time should be spent on testing areas of a system that present the most risk to your department or agency.
 
The amount of security testing required will depend on your system and the data it uses. Systems handling sensitive information are usually subject to formal assessments carried out by the Tiger Scheme or the CREST approved IT health check (ITHC) service, which is accredited by the Communications Electronics Security Group (CESG). 
 
The frequency of ITHCs will depend on their duration and cost, as well as your number of development cycles.
 
Plan your approach to security accreditation at the start of your project so that everyone is aware of the needs and requirements that have to be fulfilled.

## Further reading

* The Open Web Application Security Project ([OWASP](https://www.owasp.org)) has a [Testing Guide](https://www.owasp.org/images/5/52/OWASP_Testing_Guide_v4.pdf) that offers details on the different ways to test for security flaws. The guide is the consensus of leading industry experts.

* The CESG website has more information on the ITHC scheme. There are more details on the CREST and Tiger Scheme on their [websites](http://www.crest-approved.org/ www.tigerscheme.org).



