---
layout: gsdm
title: Domain names
section: guidance
subsection: Technical Architecture
status: draft
---
    
#Domain names

The domain name for your service is an important consideration for establishing the identity of the service and building trust in it. Government services should follow a consistent pattern in their choice of domain names so we can develop trust in those services.

##Guidance/Tool

The user's experience of government services will always start and end on www.gov.uk. This site will work as a starting point for users' interactions and hand off to the specific service in an appropriate way to be decided on a case-by-case basis.

New services should live at {servicename}.services.gov.uk. The user should interact with the service at that domain name, but www.{servicename}.services.gov.uk should work as a redirect to that name as users are likely to try and type that into their browser.

service managers should make arrangements to manage [DNS](http://en.wikipedia.org/wiki/Domain_Name_System) for that domain and control of it will be delegated by GDS as part of the approval and review process for new services. service managers should also arrange SSL certificates for those domain names and ensure their service is protected by Transport Layer Security (ie. available at https://{servicename}.services.gov.uk)

##Why we do this

Services should live on independent domain names (ie. not all on www.gov.uk) to provide clarity about the service the user is interacting with, and to ensure some separation of the services from a security perspective.

By providing a consistent but distinct set of domain names for services we establish trust in the approach as a whole, but ensure that each service can be independently managed and ring-fenced.

