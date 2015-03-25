---
layout: detailed-guidance
title: Security as enabler
subtitle: Using technological change to build secure services
category: technology
type: guide
audience:
  primary: chief-technology-officers
  secondary: tech-archs
status: draft
phases:
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Chief technology officer
    url: /service-manual/technology
---

Protecting information from valid threats to its confidentiality, integrity and availability is an enabler of digital services. Without such protection, digital services would be impossible or unsafe.

Please note: this guide sets out the governing principles for developing efficiently secured digital services. It also busts some common myths around security. There is another guide about [the practical process of security accreditation](/service-manual/making-software/information-security.html).

##Security must be proportionate and justified

Security must be applied intelligently. This means analysing the probable interest in official information from threat sources, establishing their capabilities and methods, and matching proportionate mitigations against these in a traceable manner. Other attempts at security run the risk of over-engineering security controls, or providing an illusion of security by not mitigating the actual risks.

Technology alone is never capable of addressing security and privacy risks: there needs to be a risk model that spans technology, people (their behaviours and culture), and processes.

##User experience and security aren’t exclusive

Bad user experience arising from over-prescriptive use of technology can lead users to circumvent security controls by employing less secure unofficial IT solutions. Users who then retain the official IT systems suffer degraded productivity.

The outcome is that security is not maintained and user experience is unpleasant. However, security and a good user experience do not need to be mutually exclusive.

Modern, intelligently designed security can often be made largely transparent to the user, while also providing the enterprise with the confidence it needs that its information is suitably managed.

Where some degradation of user experience is unavoidable, a risk management analysis must consider the negative impact of users avoiding unpleasant official IT and degraded productivity. Furthermore, a transparent understanding of the mapping between official information, threat sources and capabilities, through to the required mitigations, enables informed risk management when there is a change in environment or appetite for risk.

##Commercial threats, commercial solutions

The vast majority of government’s information will fall into the lowest 'official tier'. This 'official tier' has been created on the basis of a ‘commercial threat model'. This means protection from the type of threats faced by a large company or bank, eg cyber criminals or hacktivists intent on stealing personal or financial information, or disrupting services.

To defend against threats such as these, government will use only the very best security technology sourced from the commercial market, and there will be no need for any bespoke or government-only controls at this level.

A small set of software will perform security enforcing functions, such as firewalling or encrypting data. Such products require assurance that they function as advertised, achieved through the [CESG Commercial Product Assurance (CPA) scheme](https://www.cesg.gov.uk/servicecatalogue/Product-Assurance/CPA/Pages/CPA.aspx). This scheme is lightweight compared to previous schemes, reflecting the commercial grade risks for OFFICIAL information.

##Trust responsible users, audit and verify

It's the intention of the [Civil Service Reform Plan](http://www.civilservice.gov.uk/wp-content/uploads/2012/06/Civil-Service-Reform-Plan-acc-final.pdf) and the new Security Classification Policy that there's greater emphasis on user responsibility, reducing expensive and overbearing technical controls. This requires proper training to assist users in handling sensitive information, and auditing to verify users are acting responsibly.

Users should be trusted to carry out their roles and given the responsibility to do so securely.

Audit and verification of user behaviour should be used to ensure policy compliance instead of preventative measures which add cost and degrade productivity. Such audit and verification should be implemented by services or network infrastructure, away from the end user device.

Government should not invest in security controls to protect users from risks they can protect themselves from. Departments should, however, invest in security controls that help defend individual users against threats that they themselves cannot reasonably defend.

For example, government should not invest in special technology to prevent civil servants working on sensitive information in an open busy public place; users should be able to judge, assess and use appropriate risk mitigation approaches by themselves.

Likewise, civil servants should be able to exercise reasonable judgement about what information is sent to external recipients by email over the public internet. This will lead to reduced technical controls and their associated costs, while also optimising the usability and flexibility of the IT tools for the majority of responsible users.

##New classification policy
The new security policy provides for a simpler and more meaningful approach to denoting the value of information assets and threat model. Together with a set of common controls for OFFICIAL information, it enables a more consistent, standardised, reusable and interoperable approach to securing information assets in government.

There should no longer be situations where information cannot be shared between bodies solely because their interpretation of the same security guidance has led to incompatible controls.

The Cabinet Office Security Classification Policy provides guidance as to which information falls into the 'official' category.

##Procurement

It's important to start thinking about security from the very start of an IT procurement project, as bolting it on later invariably introduces additional cost and delays.

In fact, getting appropriate security into your system needs to start with specifying security requirements correctly in the contract. When we don't adequately articulate security requirements for design and ongoing maintenance of a system, our suppliers price for the level of contractual risk the ambiguity introduces.

Specifying security requirements of the form "must be accreditable for IL3” must become a thing of the past. The risk appetite of departments varies and suppliers not used to working with government will struggle to find out what that means for the design and maintenance of their system -- adding cost and risk to both parties.

Instead, when writing IT contracts, departments should think about what they care about from an information assurance perspective and specify requirements that actually manage their concerns. As well as getting a solution that address the real business needs, this approach also allows industry to innovate to solve the security problem.

##Top 8 security myths

###1. "Security says No!"

It is often said that security is the reason something can't be done. This is very rarely the case in practice -- rather, security is being (mis)used as a handy excuse to not do something.

Good information risk management practices allow organisations to understand the risks they are taking with their information assets, and work out the most effective ways to manage and control those risks -- while not hindering business.

If someone tells you that "Security says No", or that "CESG say No", you should ask for more information to learn what the risks actually are, and what techniques and tools are available to help manage those risks.

###2. Accreditation of government systems is costly, time consuming, and doesn't help secure them.

Accreditation effort should always be proportional to the complexity, threat and impact of a system. It is vital that the effort spent should scale to match the challenge at hand. Can the business accept the risks of undertaking a given activity? Accreditation documentation only needs to contain the information needed to enable this decision.

Clearly, any accreditation activity which just generates documentation that's never read is not adding value, nor helping to secure information. When performed well, accreditation helps risk owners have confidence that all aspects of a system's security have been appropriately considered, and that proper through-life security processes are in place to maintain information security.

###3. Open source software is more / less secure than proprietary code.

Experience has shown that the licensing model for software is not an accurate gauge for the security of the finished item. There are very securely developed open source projects and proprietary products. The opposite is also true. The experience and competence of those developing code is of primary importance.

The same considerations should be given to the security of open source products as are applied to proprietary code. The [Open Source toolkit](https://www.gov.uk/government/publications/open-source-procurement-toolkit) has more information.

###4. "Product X can't be used because it isn't accredited by CESG"

There are 3 problems with this statement.

Firstly, CESG don't accredit products -- they provide assurance in the security properties of products and certify them when they meet various standards.

Secondly, only those products which provide a security enforcing function need to be evaluated and certified by CESG; products like switches or email servers don't need to be CESG-certified.

Finally, CESG certification of a product is a component of risk management -- it doesn't absolve an organisation of their responsibilities to information security. While using certified products should certainly speed up the risk management process, it isn't mandatory.

###5. Restricted systems need to have bespoke security controls.

Restricted systems can be built today using assured commodity security products -- many of which are found natively in the operating systems of modern platforms.

In the past, some of the security controls which were recommended for protecting Restricted IT systems were bespoke variants of commodity products, or were "government specials". The requirements which drove these choices are no longer relevant given advances in commodity IT, and better approaches are available to provide an equivalent level of security, but with enhanced usability and overall cost of ownership.

Foundation grade assured products are appropriate for the protection of these systems, and commodity security products (including open source) can achieve this grade of certification.

###6. "I'm not a target for cyber attack!"

Sometimes it's obvious that attackers are going to be interested in getting hold of your sensitive information -- it has clear value to someone. However, immediate financial reward isn't the only reason government IT systems get attacked.

Attackers may be disgruntled insiders, seeking to disrupt or embarrass an organisation. Attacks may come from protest groups, seeking a platform for their views. Or attacks may seek to use one system as a point to launch attacks on another -- so you might not be the primary target, but might unwittingly be a conduit to them.

Are you specifically being targeted? Maybe not -- but your systems and internet presence probably are being attacked on a daily basis for a wide range of reasons.

###7. Impact Levels define security requirements.

Asking a someone to build an "IL3 system" or to "protect at IL5" is a misuse of the Business Impact Level framework, and will generally cause confusion and an ineffective approach to securing information. Business Impact Levels are intended to help organisations to think about the consequences of certain events, and thus steer where effort should be used when managing information risks.

There is no standard set of technical, procedural or personnel controls for any particular Impact Level, and so specifying requirements in terms of Impact Levels is unclear and inaccurate. Instead, security requirements should be defined in terms of the level of protection that each information asset should be given, and expectations around its through-life protection.

###8. "I've got lots of IL3 records, which aggregate to IL4, so I need to build a confidential back end system"

It's true that if you are building a system which handles many records, the impact of loss of many of them is likely to be higher than loss of one or two. Although Business Impact Levels are often misused as shorthand for protective markings, there is no direct mapping from Impact Level to protective marking.

In the IL3 aggregated to IL4 example, rather than building a system capable of storing confidential records, it's likely to be acceptable to simply add some simple additional controls around the data store, to reduce the risk of a bulk loss of data.
