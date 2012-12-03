# Hosting
The software running your service will need servers to run on. This
guide aims to help with deciding how you host your applications and the
things to think about if selecting a vendor.

## Guidance/Tool

The recommended approach is to involve a small cross-functional group of
people to quickly access different options, shortlist suppliers,
interview  and finally make a descision. This group should include people with a
knowledge of the available Procurement options and acceptable costs but
must include people with a hands-on technical understanding of the service and
underlying software.

It is important to keep good notes from any interviews or deliberation
sessions and to access different suppliers equally. A scoring matric can
help here.

It is worth stating that it's very common to use multiple suppliers.
This may be due to them offering different but compatible services or
potentially for additional redundancy. This can be technically
challenging but for larger projects can provide extra resiliance.

### Types of provider

There are a number of different approaches taken by suppliers of hosting
services, which can make comparing offerings difficult. The following is
intended only as a brief introduction.

This advice is complicated because many service providers redefine the
meanings for marketing reasons. In particular Infrastructure as a
service and Platform as a service are marketable at the moment and often
used incorrectly. Alwats look into the details of the services being
offered.

#### Ownership

For particularly large projects with very specific requirements you may
decide that purchasing hardware and even running a dedicated data centre
is suitable. The costs and timescales involved here are very high and
this is unlikley to be the best option in most cases.

#### Co-location

Many providers offer co-location services which is where you purchase
your own hardware to put into a managed data centre. This provides a
great deal of flexibility but can introduce lead times and other
physical constrains. It also requires a wide range of technical
specialist skills. 

#### Shared or managed hosting

Lots of service providers have a shared or managed hosting option. This
tends to mean renting specific virtual or physical machines for fixed
periods of time. Different suppliers other different management
services, some just manage the underlying machine while others will
support the operating system and even specific applications running on
the machines.

#### Infrastrucure as a service

In the last several years Infrastructure as a service has become a
common approach to managing hosting and infrastructure requirements.
This tends to involve a capability to rapidly add or remove capacity,
often in minutes, and to be billed only for what is used. This provides
a great deal of flexibility and the ability to hold costs down but also
requires a degree of technical skill to manage well. 

#### Platform as a service

Similar to Infrastructure as a service above, Platform as a service
offerings tend to allow for quickly adding or removing capacity and fine
grained pay on demand pricing. The difference is is that you are
abstracted away from the underlying infrastructure completely. The unit
here is the running application, not a virtual or physical machine.
Using a Platform as a service places a number of constrains on the
software architecture but can move the support burden for parts of the
stack onto the supplier.

## Why we do this

Making a decision about your hosting supplier involves weighing up a
wide range of different components, including:

* Technical requirements of your software applications
* Future capacity requirements
* Customer support
* Relisance and redundancy
* Mandated Government information security requirements
* Government network connectivity
* Cost of service
* Flexibility and on-demand billing
* Different operating models (physical, co-location, infrastructure as a service,
  platform as a service)

This is unfortunately a technical field with many options. Seemingly
similar services can have wildly different architectures or different
cost models can result in large differences in total cost of ownership.
It is recommended to involve technical colleagues or trusted third
parties in any discussions and descision.

## Further reading

* [Why IaaS?](http://digital.cabinetoffice.gov.uk/2012/09/25/why-iaas/)
  a blog post about why the Government Digital Service chose an
  Infrastructure as a service for GOV.UK
* [G-Cloud](http://gcloud.civilservice.gov.uk/) the procurement
  framework intended for purchasing Infrastructure as a Service,
  Platform as a Service and Software as a Service products
