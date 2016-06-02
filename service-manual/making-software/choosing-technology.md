---
layout: detailed-guidance
title: Choosing technology
subtitle: How to go about choosing what software tools to use
category: making-software
type: guide
audience:
  primary: service-managers, tech-archs, developers
  secondary: web-ops, chief-technology-officers
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
exclude_from_search: true
---

This is for guidance purposes only, so don’t take it as legal advice.

At different points in your project you’ll have to choose one technology over others, which could be a choice between:

* programming languages
* databases
* libraries
* operating systems
* small tools that help the development team work more efficiently
* proprietary (commercial) and open source software, where relevant

## You can change your mind

Assume that most of your technology choices can or will change, especially during the early stages of development, eg:

* selecting a programming language that’s easier and quicker to prototype in, moving to another for the final product that’s easier for large teams to use
* using an open source product to get started quickly, going on to buy a commercial (proprietary) product for required features (or vice-versa)

On day 1 of your project you simply won’t know enough about the domain or the user need to select the right technology. It’s OK to make an educated guess at this stage, as long as everyone understands this.

Be sure to find the time to challenge your selection as you learn more about the problem at hand.

## Start with capabilities
When making choices on what technology to use, it’s very easy to immediately jump to a specific product. This tends to be based either on past experience or on current trends.

1. Take a step back and think about the capabilities of the technology you're after -- eg are you looking for a relational database or a document database? A key/value store or a graph database?
2. Discuss the capabilities you want, rather than any specific implementation.
3. Once you’ve defined a set of capabilities, work out how you will test potential software solutions against those capabilities.
4. Test software in an environment that's very similar to your production environment, ie at production scale under real conditions (including worst-case user load).

Carrying out this process will make sure that you:

* understand the specific characteristics of software
* identify any trade-offs or potential issues with a piece of software, eg the ideal fit with your domain model vs. performance under heavy load
* confirm that its claimed capabilities are genuine

## Cost

When choosing technology, make sure you consider the total cost as well as any upfront fees.

Take into consideration costs for things like:

* staff
* support costs
* licensing costs (where applicable)
* the productivity of ongoing service improvement work
* any exit costs (especially around migration of data to a future replacement system) that might be caused by the use of non-open standard formats or protocol extensions

Make sure that you understand the cost implications of any unusually high user loads in production systems.

## Involve everyone
Involve the whole team in technology choices. That doesn't mean no one is in charge of making decisions, but that you want the development team (including web ops people) to share their insight on:

* the options available
* how it will affect the overall system

Expect technology preferences to vary, and for opinions to be divided when choosing technology.

All things being equal, picking technologies that developers and operations staff like will typically result in improved productivity.


## Testing and deployment

You’ll make these processes a lot easier if you:

* deploy your service [early and often](/service-manual/making-software/release-strategies.html)
* choose component that are easy to deploy and upgrade as part of an [automated pipeline](/service-manual/agile/continuous-delivery.html)
* avoid software that makes testing difficult - [frequent, automated testing](/service-manual/making-software/code-testing.html) is essential in agile development

## Lock-in
Technology lock-in happens when your previous decisions limit future decisions, possibly leaving you with only 1 option.

For example, if you select a database that runs on only 1 operating system, you have no choice about the operating system you’ll use. If the costs of that operating system jump, you have no simple way of reducing that cost quickly or cheaply.

Being locked in to one vendor, or one way of doing things, can create difficulties like:

* unforeseen financial costs, eg an overnight increase in cost
* limiting how quickly you can iterate on your product in the future, eg if the ideal technology choice isn’t compatible with your current vendor or technology

Over time, and after many decisions, you can find yourself in a situation where all your technology decisions are tightly coupled and you are locked-in to one vendor, or one way of doing things. This can have unforeseen financial costs (for example an overnight cost increase). It might also limit how quickly you can iterate on your product in the future (if, for instance, the ideal technology choice isn’t compatible with your current vendor or technology).

When you decide on a certain technology, be sure to have a clear understanding of the cost or implications of moving away from it, and [avoid lock-in](/service-manual/making-software/open-standards-and-licensing.html) whenever possible. This could mean choosing not to use certain features (eg non-standard extensions to protocols, APIs or programming languages), so you have the ability to move to another technology as and when.

In general, avoid making long-term commitments to any particular technology, product or supplier until you fully understand the problem you’re trying to solve. Even then, make sure you maximise your future development options and avoid technology lock-in if at all possible.

## New vs existing software

If there’s an existing software solution that solves your problem, definitely think about using it.

If you’re using software to satisfy a rare or niche need, you’re less likely to find a tool that will serve and continue to serve those needs. In these situations, expect that you’ll have to create new software.

However, there are many situations where it makes more sense to use existing software like:

* when you have a commodity need, because someone has probably already developed a solution to the common problem
* facing a niche problem that’s peculiar to government, which has already been solved by another part of, or another, government and [released as open source software](https://github.com/gds-operations/)

In most cases, many projects will have the same need so it makes little sense to reinvent the wheel for certain types of software, eg:

* development tools
* build tools
* utility libraries
* databases
* monitoring tools

## Sharing software

Share software that’s developed to meet the needs of government wherever possible under a permissive, GPL-compatible open source licence (eg [MIT](http://opensource.org/licenses/MIT)/[X11](http://directory.fsf.org/wiki/License:X11) or [3-clause BSD](http://directory.fsf.org/wiki/License:BSD_3Clause)) unless doing so would:

* create an unacceptable risk to the security of systems or processes that cannot be mitigated with reasonable efforts
* break existing contractual arrangements
* directly threaten national security

Sharing software means it can be widely used and improved by anyone in the world who has a similar need. It’s important that other governments in particular have the opportunity to reuse the software you’ve created, because everyone deserves to have [digital services so good that people prefer to use them](/service-manual/digital-by-default).

For example [GovSpeak](https://github.com/alphagov/govspeak) and [unicorn herder](https://github.com/gds-operations/unicornherder) are small components that were [developed as a part of GOV.UK](https://gds.blog.gov.uk/govuk-launch-colophon/). They're now used by several different organisations, and have received a number of [public contributions](https://github.com/gds-operations/unicornherder/commits/master).

In practice, sharing usually means:

* uploading the source code and documentation to a public source code repository
* keeping it updated with subsequent changes that you make (or accept from other people)
* putting in place [appropriate information security assurance to reduce and mitigate the risk of an exploit appearing in publicly-viewable software](/service-manual/making-software/information-security.html)

> Store the 'master' version of your software on an internal source code control system and replicate the latest version to a public repository.

### Reasons not to share software

Sometimes it’s not possible to share software that was developed for the government by a third party because the third party retains ownership of the ‘intellectual property’ (IP) embodied in that software.

So, when working with third parties, Avoid contracts that:

* allow third parties to retain ownership of IP in software that’s been developed for government
* restrict government’s ability to share software under a permissive, GPL-compatible open source licence

On other occasions a team may take a risk-based decision not to share some of the software they‘ve created, eg to avoid exposing the details of a particular risk-assessment algorithm or process to the public.

It’s good engineering practice in any case to encapsulate software, and so often in this situation a large portion of the software for a given system or service can still be shared in public.

### CESG

The default assumption should be in favour of coding in the open and sharing software widely. But if you have serious concerns about sharing source code in public, then CESG can provide you with advice based on your specific situation. See also the joint Cabinet Office / CESG statement on [Open Source Software and Security](/government/uploads/system/uploads/attachment_data/file/78967/OSS_Toolkit_Security_Note_v1.0.pdf).

## Commercial or open source?

With the growth of free/open source software, many high quality technology products (databases, operating systems, programming languages, development tools, etc) are freely available for government and its suppliers to use and improve.

Yet a large market still exists for commercial software products, and the availability of open source software doesn’t automatically mean that you can’t choose a proprietary technology if it meets your needs.

Still, the policy of government is that open source will be selected on the basis of its inherent flexibility. This is norm when there is no significant overall difference to cost between open and non-open source products that fulfil minimum and essential capabilities.

For more information on how to make sure there’s a level playing field between proprietary and open source software, see the:

* [open source procurement toolkit](/government/publications/open-source-procurement-toolkit)
* section of this manual covering the [use of open standards](/service-manual/making-software/open-standards-and-licensing.html)

> Make sure that any decision to use existing software, whether open source or proprietary, doesn't stop you from sharing new software created under a permissive, GPL-compatible open source licence.

## Coding in the open

A successful open source project will collect contributions from a large number of sources, both inside and outside of a single organisation. Allow developers time to review contributions, and answer issues and discussions raised by others using the software.

Make sure developers have:

* the ability to install and experiment with open source software
* environments to easily publish prototype services on the web
* convenient access to a wide variety of network connected devices for testing websites
* unrestricted access to collaboration tools like [GitHub](https://github.com), [Stack Overflow](https://stackoverflow.com/) and [IRC](https://en.wikipedia.org/wiki/Internet_Relay_Chat)

Larger open source projects often develop an extension model that makes it possible for others to continue to use the service in a variety of often unexpected and possibly undesirable ways whilst keeping the core project coherent under the editorship of a small, trusted group of [committers](https://en.wikipedia.org/wiki/Committer).

Take every opportunity to contribute back to the open source projects you use, which can be in the form of:

* source code
* patches
* bug reports
* feature requests
* sponsorship of developers and support staff
* engaging in community discussion groups
* giving public attribution to projects

Cite the open source code you use, as in the [GOV.UK colophon](https://gds.blog.gov.uk/govuk-launch-colophon/) -- you can read more about this approach on the [GDS blog entry about coding in the open](https://gds.blog.gov.uk/2012/10/12/coding-in-the-open/).

## Security

Keys, passwords and other secrets must always be stored safely and securely away from source code [following Kerckhoffs’s principle](https://en.wikipedia.org/wiki/Kerckhoffs%27s_principle). This separation of project code from deployed instances of a project is good development practice regardless of whether or not the software itself is shared in public.

It's advisable for large, significant projects to have a private space to discuss security issues and develop a patch. This removes the risk of flagging a vulnerability before a fix has been deployed. This is especially advisable if there's a small number of participating developers.

## Why GDS do this
Choosing technology is important, but it’s probably not quite as important as you think. What's important are the users of that technology and being able to produce quality at a sustainable pace and suitable cost. When making technology choices, and importantly as you develop your product and constantly reassess your selections, try and make decisions that:

* maximise developer productivity
* minimise total cost of ownership
* avoid lock-in
* make it easy for the government to share software that it creates

## Further reading
[Open standards considerations](/service-manual/making-software/open-standards-and-licensing.html)

*[IP]: intellectual property
*[GPL]: General Public License
*[CESG]: Communications-Electronics Security Group
