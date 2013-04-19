---
layout: detailed-guidance
title: Choosing technology
subtitle: How to go about choosing what software tools to use
category: making-software
type: guide
audience:
  primary: service-managers, tech-archs
  secondary: developers
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

At different points in your project you’ll have to choose one technology over others. That might be a programming language, database, library, operating system or some small tool that helps the development team work more efficiently. In some cases you may also have to choose between proprietary and open source software.

> Please note, this is for guidance purposes only, and should not be taken as legal advice.

## You can change your mind
The most important consideration is to work on the assumption that most technology choices can change, especially during the early stages of development.

On day one of your project you simply won’t know enough about the domain or the user need to select the right technology. It’s OK to make an educated guess at this stage, as long as everyone understands that is what is happening. Then find time to challenge the selection as you learn more about the problem at hand.

Maybe you selected a programming language that you knew would be easier and quicker to prototype in for the early stages of the project, and then moved to another one that was easier for large teams to use for the final product. Or maybe you started out with an open source product to allow you to get started quickly, before going on to buy a commercial (proprietary) product which provided some required feature (or vice-versa).

## Start with capabilities, not implementations
It’s very easy to immediately jump to a specific product when making technology choices. This tends to be based either on past experience or on fashion.

Try and take a step back and think about the capabilities of the technology you’re after. Are you looking for a relational database? Or a document database? A key/value store? Or maybe a graph database? Argue first about the capabilities, rather than any specific implementation.

Once you have defined a set of capabilities, work out how you will test solutions against those capabilities. You should aim to test any software that appears appropriate in an environment that is very similar to your production environment, ie at production scale under real-world conditions (including worst-case user load).

This will allow you to make sure you understand its specific characteristics, confirm that that its claimed capabilities are genuine, and identify any trade-offs or potential issues to consider with a given piece of software such as conceptual fit with your domain model vs. performance under heavy load.

## Cost
When choosing technology make sure you consider the total cost, as well as any upfront fees. Try and take into consideration costs for things like staff, support or licensing costs (where applicable), the productivity of ongoing service improvement work, and any exit costs (especially around migration of data to a future replacement system) that might be caused by the use of non-open standard formats or protocol extensions. Ensure that you understand the cost implications of any unusually high user loads in production systems.

## Consider people
Try and involve the whole team in technology choices. That doesn’t mean no-one owns the decision making but that you want the development team (including web ops people) to share their insights into the options available and how they will affect the overall system, and thereby get their 'buy in' to the technology choices made. Technology preferences vary, and technology choice can divide opinion.

All things being equal, picking technologies that developers and operations staff like will typically result in improved productivity.

## Testing and deployment
The services that we write will need to be deployed, [early and frequently](/service-manual/making-software/release-strategies.html). Any components that we choose should be easy to deploy and upgrade as part of an [automated pipeline](/service-manual/agile/continuous-delivery.html). [Frequent, automated testing](/service-manual/making-software/code-testing.html) is essential in agile development. Software that makes testing difficult should be avoided.

## Lock-in
Technology lock-in happens when previous decisions regarding technology limit future decisions, possibly so that only one real choice exists.

For example, if you select a database that only runs on one operating system you have no choice about the operating system you will use. If the costs of that operating system jump you have no simple way of reducing that cost quickly or cheaply.

Over time, and after many decisions, you can find yourself in a situation where all your technology decisions are tightly coupled and you are locked-in to one vendor, or one way of doing things. This can have unforeseen financial costs (for example an overnight cost increase) but can also limit how quickly you can iterate on your product in the future, for instance if the ideal technology choice isn’t compatible with your current vendor or technology.

Aim to have a clear understanding of the cost or implications of moving away from a technology when you commit to it, and avoid technology lock-in whenever possible.  This may sometimes mean choosing not to use certain features of a particular piece of technology, such as non-standard extensions to protocols, APIs or programming languages, in order to [avoid lock-in](/service-manual/making-software/open-standards-and-licensing.html) and preserve the ability to move to another technology at a later date.

In general, you should avoid making long-term commitments to any particular technology, product or supplier until you fully understand the problem you're trying to solve - and even then, you should ensure you maximise your future development options and avoid technology lock-in if at all possible.

## Make new software versus use existing software
Where there is an existing software solution which solves your problem, you should certainly consider using it.  You are more likely to use existing software than to make new software when you have a commodity need - or if you're facing a niche problem that's peculiar to government and has already been solved by another part of the government - or indeed another government - and [released as open source software](https://github.com/alphagov/). Development tools, build tools, utility libraries, databases and monitoring tools are all examples of software where many projects will have the same need, and it makes little sense to reinvent the wheel.

For software that serves a rare or niche need you are less likely to find a tool which serves and will continue to serve your needs, and in this situation you will likely have to create new software.  Software that is developed to meet the needs of the government - whether it’s developed by government employees, contractors or by a supplier - should be shared wherever possible under a permissive, GPL-compatible open source licence (eg [MIT](http://opensource.org/licenses/MIT)/[X11](http://directory.fsf.org/wiki/License:X11) or [3-clause BSD](http://directory.fsf.org/wiki/License:BSD_3Clause)) so that it can be widely used and improved. This allows the software to be used and improved by anyone in the world who has a similar need.  It’s important that other governments in particular to have the opportunity to re-use the software you've created, because everyone deserves to have [digital services so good that people prefer to use them](/service-manual/digital-by-default).

For example [GovSpeak](https://github.com/alphagov/govspeak) and [unicorn herder](https://github.com/alphagov/unicornherder) are small components which [developed as a part of GOV.UK](http://digital.cabinetoffice.gov.uk/govuk-launch-colophon/), are used by several different organisations, and have received a number of [public contributions](https://github.com/alphagov/unicornherder/commits/master).

You should always share software that has been written by the government and/or its suppliers (this includes source code and documentation) unless doing so would:

a) create an unacceptable risk to the security of our systems or processes that cannot be mitigated with reasonable efforts
b) contravene existing contractual arrangements
c) directly threaten our national security

In practice, sharing usually means uploading the source code and documentation to a public source code repository, keeping it updated with subsequent changes that you make (or accept from other people), and putting in place [appropriate information security assurance to reduce and mitigate the risk of an exploit appearing in publicly-viewable software](/service-manual/making-software/information-security.html).

> In some cases you might want to store the 'master' version of your software on an internal source code control system and replicate the latest version to a public repository.

## Reasons not to share software
Sometimes it's not possible to share software that was developed for the government by a third party because the third party retains ownership of the 'intellectual property' (IP) embodied in that software. Contracts that allow third parties to retain ownership of IP in software that's been developed for the government, and/or that restrict the government's ability to share this software under a permissive, GPL-compatible open source licence, should be avoided.

On other occasions the team may take a risk-based decision not to share some of the software they have created, for instance to avoid exposing in public the details of a particular risk-assessment algorithm or process. It's good engineering practice in any case to encapsulate software, and so often in this situation a large portion of the software for a given system or service can still be shared in public.

The default assumption should be in favour of coding in the open and sharing software widely, but if you have serious concerns about sharing source code in public, then CESG can provide you with advice based on your specific situation. See also the joint Cabinet Office / CESG statement on [Open Source Software and Security](https://www.gov.uk/government/uploads/system/uploads/attachment_data/file/78967/OSS_Toolkit_Security_Note_v1.0.pdf).

## Level playing field
With the growth of free/open source software, many high quality technology products (databases, operating systems, programming languages, development tools, etc) are freely available for government and its suppliers to use and improve. But a large market still exists for commercial software products, and the availability of open source software doesn't automatically mean that you can't choose a proprietary technology if it meets our needs.

However, it remains the policy of the government that, where there is no significant overall cost difference between open and non-open source products that fulfil minimum and essential capabilities, open source will be selected on the basis of its inherent flexibility. For more information on how to ensure a level playing field between proprietary and open source software, please see the [Open Source Procurement Toolkit](https://www.gov.uk/government/publications/open-source-procurement-toolkit) and the section of this manual that covers [use of open standards](/service-manual/making-software/open-standards-and-licensing.html).

> You should ensure that any decision you make to use existing software, whether open source or proprietary, doesn't stop you from sharing any new software that you create (or your suppliers create for you) under a permissive, GPL-compatible open source licence.

## Coding in the open
A successful open source project will garner contributions from a large number of sources, both inside and outside of a single organisation. Allow developers time to review contributions, and answer issues and discussion raised by others using the software.

Larger open source projects often evolve an extension model to enable others to continue to use the service in a variety of often unexpected and possibly undesirable ways whilst keeping the core project coherent under the editorship of a small, trusted group of [committers](http://en.wikipedia.org/wiki/Committer).

Ensure developers have the ability to install and experiment with open source software, have environments to easily publish prototype services on The Web, have convenient access to a wide variety of network connected devices for testing Web sites, and have unrestricted access to collaboration tools such as [GitHub](https://github.com), [Stack Overflow](http://stackoverflow.com/) and [IRC](http://en.wikipedia.org/wiki/Internet_Relay_Chat).

Take every opportunity to contribute back to open source projects you use. Contributions may be in the form of source code, patches, bug reports, feature requests, sponsorship of developers and support staff, engaging in community discussion groups, and giving public attribution to projects. Cite the open source code you use, as in the [GOV.UK colophon](http://digital.cabinetoffice.gov.uk/govuk-launch-colophon/) - you can read more about this approach on the [GDS blog entry about coding in the open](http://digital.cabinetoffice.gov.uk/2012/10/12/coding-in-the-open/).

> Keys, passwords and other secrets must always be stored safely and securely away from source code [following Kerckhoffs’ principle](http://en.wikipedia.org/wiki/Kerckhoffs%27_principle).  This separation of project code from deployed instances of a project is good development practice regardless of whether or not the software itself is shared in public.

> For projects with a high impact level, particularly with a small number of participating developers, it is advisable to have a private space to discuss security issues and develop a patch rather than risk flagging a vulnerability before a fix has been deployed.

## Why we do this
Choosing technology is important, but it’s probably not quite as important as you think. What is important are the users of that technology and being able to deliver quality at a sustainable pace and suitable cost. When making technology choices, and importantly as you develop your product and constantly reassess your selections, try and make decisions that:
* maximise developer productivity
* minimise total cost of ownership
* avoid lock-in
* make it easy for the government to share software that it creates

## Further reading
[Open standards considerations](/service-manual/making-software/open-standards-and-licensing.html)
