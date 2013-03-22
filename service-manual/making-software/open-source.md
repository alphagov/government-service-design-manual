---
layout: detailed-guidance
title: Open source
subtitle: When to use open source
category: making-software
type: guide
audience:
  primary: service-managers, developers, tech-archs
  secondary: web-ops 
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

This section presents architectural practice and considerations for using, publishing and contributing to free and open source software ([FOSS](http://en.wikipedia.org/wiki/Free_and_open-source_software)).

## When to use open source

Use open source software in preference to [proprietary or closed source](http://en.wikipedia.org/wiki/Proprietary_software) alternatives, in particular for operating systems, networking software, Web servers, databases and programming languages.

Problems which are rare, or specific to a domain may be best answered by using [software as a service](http://en.wikipedia.org/wiki/Software_as_a_service), or by installing proprietary software.

In such cases, take care to mitigate the risk of lock-in to a single supplier by ensuring [open standards](https://www.gov.uk/government/uploads/system/uploads/attachment_data/file/78892/Open-Standards-Principles-FINAL.pdf) are available, in particular for data portability, and interfaces used for integration with other systems.

Where possible use DNS addresses you own for services, and demand open formats for the import and export of your data.

For unique needs and common problems which have yet to be solved well elsewhere, develop software by [coding in the open](http://digital.cabinetoffice.gov.uk/2012/10/12/coding-in-the-open/) and publish under an open source licence ([Legal processes/Open standards and licensing]()).

Whenever possible construct software in the form of small components and utilities, re-usable both inside and outside of your organisation. Keep infrastructure code and secrets, including passwords and deployment configuration and scripts, separate and privately from publicly visible source code.

A successful open source project will garner contributions from a large number of sources, both inside and outside of a single organisation. Allow developers time to review contributions, and answer issues and discussion raised by others using the software.

Larger open source projects often evolve an extension model to enable others to continue to use the service in a variety of often unexpected and possibly undesirable ways whilst keeping the core project coherent under the editorship of a small, trusted group of [committers](http://en.wikipedia.org/wiki/Committer).

## Developing with open source

Provide developers with ready access to open source development tools with which they will be familiar so they may be productive immediately.

Ensure developers have the ability to install and experiment with open source software, have environments to easily publish prototype services on The Web, have convenient access to a wide variety of network connected devices for testing Web sites, and have unrestricted access to collaboration tools such as [GitHub](https://github.com), [Stack Overflow](http://stackoverflow.com/) and [IRC](http://en.wikipedia.org/wiki/Internet_Relay_Chat).

Take every opportunity to contribute back to open source projects you use. Contributions may be in the form of source code, patches, bug reports, feature requests, sponsorship of developers and support staff, engaging in community discussion groups, and giving public attribution to projects. Cite the open source code you use, as in the [GOV.UK colophon](http://digital.cabinetoffice.gov.uk/govuk-launch-colophon/).

## Why we do this
Free and open source software has a number of architectural benefits over closed source and proprietary alternatives and is the basis of our [tenth design principle](https://www.gov.uk/designprinciples#tenth) &mdash; _Make things open: it makes things better_.

Coding in the open lowers the barriers to collaboration with others inside and outside of your organisation, increasing the speed of new developers being productive on a project. Modern developers are usually more familiar with open source tools and ways of working than with proprietary products.

Freedom at the point of use means open source software may be downloaded and assessed by developers without payment, prior agreement, a need to sign non-disclosure documentation, needing to waiver rights, or enter into aporia agreements on behalf of themselves or their organisation.

### Avoid vendor lock-in

A large part of the total cost of ownership of a project are incurred at the end, overcoming barriers to exit, especially when moving data to a new system. Using closed source software increases the risk of [lock-in](http://en.wikipedia.org/wiki/Vendor_lock-in) to a single supplier. This is a risk, even when building with [open standards](http://consultation.cabinetoffice.gov.uk/openstandards/) due to the possibility of inadvertently using proprietary features, increasing exit costs.
Access to the source code enables support to be sourced from a number of independent suppliers, enabling better prioritisation of bug-fixes, allowing systems to be maintained outside of vendor product life-cycles, and mitigating the risk of [planned obsolescence](http://en.wikipedia.org/wiki/Planned_obsolescence) or [abandonment](http://en.wikipedia.org/wiki/Abandonware).

## Security and open source

Keys, passwords and other secrets need to be stored safely and securely away from source code following [Kerckhoffs' principle](http://en.wikipedia.org/wiki/Kerckhoffs%27_principle). Whilst it is possible to publish a project initially built in private, coding in the open means confronting this separation of concerns early during development. 

This separation of project code from deployed instances of a project is good development practice, and using open source enables developers to easily fork and experiment with multiple development, and operations to quickly spin-up multiple test and integration environments without encountering limits of licensing.

Although [Linus's Law](http://en.wikipedia.org/wiki/Linus%27s_law) "given enough eyeballs, all bugs are shallow" is contentious, transparent access to the code increases the difficulty of hiding a [back door](http://en.wikipedia.org/wiki/Backdoor_(computing\)) or [Trojan horse](http://en.wikipedia.org/wiki/Trojan_horse_(computing\)) inside a FOSS product.

It is argued a system is most vulnerable after a potential attack is discovered, but before a fix has been deployed. A number of [metrics and models](http://en.wikipedia.org/wiki/Open_source_software_security#Metrics_and_Models) attest to the quicker response to security issues in open source products when compared to closed source equivalents.

For projects with a high impact levels, particularly with a small number of participating developers, it is advisable to have a private space to discuss security issues and develop a patch rather than risk flagging a vulnerability before a fix has been deployed.

## Better products and services

Software is expensive to maintain, and can be seen as a liability, yet the value of software increases the more people use it. Publishing as open source increases the likelihood of code-reuse. For example [GovSpeak](https://github.com/alphagov/govspeak) and [unicorn herder](https://github.com/alphagov/unicornherder) are small components which developed as a part of GOV.UK, are used by several different organisations, and have received a number of public contributions.

Open source encourages re-use and benefits from [network effects](http://en.wikipedia.org/wiki/Network_effect), and as a result has developed tools, services and processes for coordinating massively distributed development teams spread across multiple organisations.

Finally, by lowering the barriers for reuse, building for extensibility, encouraging forking, mutation and experimentation, open source fosters a culture of [open innovation](http://en.wikipedia.org/wiki/Open_innovation), leading to better products and services.

## Further reading
* The legal obligations for using open source software are outlined the section [open standards and licensing](/service-manual/making-software/open-standards-and-licencing.html).
* Official note from [CESG](http://www.cesg.gov.uk/) on [the security of open source](https://www.gov.uk/government/uploads/system/uploads/attachment_data/file/78967/OSS_Toolkit_Security_Note_v1.0.pdf).
