---
layout: detailed-guidance
title: Hosting
subtitle: Deciding where your service will live
type: guide
status: draft
audience: 
  primary: tech-archs, web-ops
  secondary: service-managers, delivery-managers, developers
category: operations
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Operations
    url: /service-manual/operations
exclude_from_search: true
---

You need servers to run your software. The information here will help you decide how you host your applications and the things to consider when selecting a vendor.

## Decide how you'll host your application

To begin with, involve a small cross-functional group of people who will quickly:

* assess different options
* shortlist suppliers
* interview suppliers
* make a decision

Make sure the people in your group have:

* a hands-on technical understanding of the service and underlying software -- these people are essential
* knowledge of the available procurement options
* knowledge of acceptable costs

It’s important that you keep good notes from any interviews or deliberation sessions and assess different suppliers equally. One way to do this is to use a scoring matrix.

Know that it’s very common to use multiple suppliers. This can be because they offer different but compatible services, or additional redundancy. It can be technically challenging to do this, but can provide extra resilience for larger projects.


## Types of suppliers

Suppliers offer a number of different approaches to hosting services, which can make it difficult to compare them. The following information is intended as only a brief introduction of each type, covering:

* ownership
* co-location
* shared or managed hosting
* infrastructure as a service
* platform as a service

This advice is complicated because many service providers redefine the meanings of hosting terms for marketing reasons, especially:

* infrastructure as a service
* platform as a service

These terms are marketable at the moment and often used incorrectly, so always look into the details of the services being offered.

### Ownership

If you have a large project with very specific requirements, you may decide that purchasing hardware and even running a dedicated data centre is suitable.

The costs and timescales involved here are very high and in most cases it’s unlikely to be your best option.

### Co-location

Many providers offer co-location services -- you purchase your own hardware and rent space in a managed data centre. 

This provides a great deal of flexibility, but can introduce lead times and other physical constraints. It also requires a wide range of technical specialist skills.

### Shared or managed hosting

Lots of service providers have a shared or managed hosting option. This tends to mean renting specific virtual or physical machines for fixed periods of time.

Different suppliers offer different management services; some just manage the underlying machine while others will support the operating system and even specific applications running on the machines.

### Infrastructure as a Service

In the last several years this has become a common approach to managing hosting and infrastructure requirements. It tends to involve a capability to rapidly add or remove capacity, often in minutes, and to be billed only for what is used. This provides a great deal of flexibility and the ability to keep costs down, but also requires a degree of technical skill to manage well.

### Platform as a Service

Similar to infrastructure as a service, platform as a service offerings tend to allow for quickly adding or removing capacity and precise ‘pay on demand’ pricing.

The difference is that you are considered completely separate from the underlying infrastructure. The unit here is the running application, not a virtual or physical machine.

Using a platform as a service places a number of restrictions on the software architecture, but can move the support burden for parts of the stack onto the supplier.

## Make your decisions carefully

Before making a decision about your hosting supplier, weigh up a wide range of different factors, including:

* the technical requirements of your software applications
* your future capacity requirements
* user support
* reliance and redundancy
* mandated government information security requirements
* government network connectivity
* the cost of service
* flexibility and on-demand billing
* different supplier operating models

Unfortunately this is a technical field with many options. Seemingly similar services can have wildly different structures or different cost models that can result in large differences in your total cost of ownership (cost of the project). 

It’s important to involve technical colleagues or trusted third parties in any discussions and/or decisions.

## Further information

* [Why IaaS?](https://gds.blog.gov.uk/2012/09/25/why-iaas/) -- a blog post about why the Government Digital Service chose an infrastructure as a service for GOV.UK
* [G-Cloud](/how-to-use-cloudstore) -- the procurement framework intended for purchasing infrastructure as a service, platform as a service and software as a service products
