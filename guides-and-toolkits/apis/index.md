---
layout: gsdm
title: Consuming/Using APIs
section: guidance
subsection: APIs
status: draft
---
    
#Consuming/Using APIs
Don't do everything yourself (you can't). Sometimes the functionality your service needs will be provided by other parts of your organisation, other government departments or by reliable third parties via APIs.

##Guidance
Most modern digital services are built on top of a wide range of APIs. This allows each part of the service to focus on its core responsibility rather than constantly reinventing the wheel.

### Code Integration
When consuming APIs you should be careful to keep the integration with your code clean and distinct. This is important to ensure that you can swap between providers or update to new versions of an API without making substantial changes to your core code. We encourage the use of adapter code that is entirely focussed on interfacing with the system and mapping code that will provide the linkage between your code's domain model and the concepts and services provided by the API.

### Testing
You should consider carefully how you intend to test your integration with the service. In day to day development you will want to be able to test your code without making (computationally or potentially financially) costly calls out to third party services so you should come up with a way of providing mock versions of those APIs. For full system tests, however, you will want to test the full flow including the third-party service so an automated mechanism should be built for that.

Many of the GOV.UK publishing applications send emails to provide alerts for content designers. When running tests we don't want to send lots of fake emails so we swap the normal email adapter for one that logs the emails it would have sent. This lets us test our code is doing the right thing without depending on external services.

Our "data insights" code involves significant interactions with Google Analytics. It wouldn't be practical to test this by sending events to google, waiting for them to be processed, and then reviewing the results. Our developers therefore built a mock service that can be run alongside tests and provides a dummy version of google's api that lets us check the right data is being sent.

Our publishing systems make use of a single sign-on service. In most of our tests the interaction with that service are mocked so the applications' tests can be run in isolation, but we also have a suite of "smoke tests" that run in our preview environment and use dummy accounts to ensure that the full authentication and authorisation flow is working.

The Licence Application Tool integrates with a number of third-party payment services. It makes use of test accounts with those services to verify it is able to communicate with them and is sending the right data to complete payments.

### Service Agreements and Resilience
By depending on a third party API you could very easily be tying your service's availability to that of the third party. In some cases that may be acceptable, but often you will want to ensure that you have a fallback plan in place.

The details of that fallback will vary according to your service. It may be that you will need to offer the user the opportunity to use an alternative service, or queue the action to take place later. That could be an automated queue with software that monitors it and retries transactions, or it could be a manual queue where someone follows up to collect further details.

You should be clear with your users about what is happening. If a third party payment provider isn't available you might queue the transaction to try again later. That will mean you can't offer users the same guarantee that their payment will be processed correctly and you should tell them so.

##Why we do this
This is the reasoning for the decision. You might want to link to relevant blog posts or legislation here.

##Further reading
Link within the Manual, or to other posts, that will help people to work on the tool.

#Producing APIs
Martha Lane Fox's report report called for government to act as a "wholesaler as well as the retail shop front for services and content by mandating the development and opening up of Application Programme Interfaces [APIs](http://en.wikipedia.org/wiki/Application_programming_interface) to third parties". This section is a set of guiding principles for exposing a Digital service as an API.

##Guidance

###Build an API by building with the API
When building an API there is a danger of building the wrong thing in the wrong way for the wrong people.
This is a particular risk in the absence of a developer community driving [the needs](https://www.gov.uk/designprinciples#first) of an API.

The simplest way to ensure your API is useful and consumable is to build a Web site using your own API.

Building a Web site forces a need to model data around bookmarkable resources; becoming a consumer of your own user not only validates your API,
but leads to services being exposed [on The Web](http://www.w3.org/TR/webarch/).

### Just use The Web
Consider an API to be a part of a Web site.
Provide links to machine-friendly formats from human readable pages,
and enable agents to easily construct URLs which link to human-friendly representations of pages.

Use standard formats for content, and follow established Web patterns for authentication.

Building a service to enjoy mass adoption and support from a wide disparate community of developers and programming environments whilst being able to reach a world-wide audience is difficult.
Whilst proprietary and open technologies abound for machine-to-machine communication, none combine the interoperability, reach and ability to scale to compete with The Web.

Standards are powerful agreements, and nowhere are agreements more quickly established and adopted than on The Web.
Using HTTP and URIs, the core technologies of The Web, together with emergent standards such as [JSON](http://www.json.org/) and [OAuth](http://en.wikipedia.org/wiki/OAuth) changes a Website from [a retail shop window into a wholesaler](http://www.cabinetoffice.gov.uk/resource-library/directgov-2010-and-beyond-revolution-not-evolution), meeting our design principle to [Build digital services, not websites](https://www.gov.uk/designprinciples#eighth).

### Give each thing a bookmarkable URI
Expose data as a set of resources, offering a [clean URL](http://en.wikipedia.org/wiki/Clean_URL) for each thing, and each collection of things.

Only use [query strings](http://en.wikipedia.org/wiki/Query_string) for URLs with unordered parameters, such as options to search pages.

Consider creating URIs for different granularity of resources. For example, `/members.json` could return a list of names, whilst `/members.json?detail=full` could return detailed information about each member in a list.

These principles enable network effects which arise through linking and allow information published beyond the Web, sent in alerts email, SMS, XMPP and other messages to link back to the canonical content, on The Web.

### Use HTTP methods as Tim intended
Ensure all [HTTP GET](http://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol#Request_methods) requests are [safe](http://www.w3.org/2001/tag/doc/whenToUseGet.html), and actions which change state are conducted using a [POST](http://en.wikipedia.org/wiki/POST_\(HTTP\)), PUT or DELETE method.

Use PUT and DELETE which are commonly blocked by firewalls, intranet proxies, hotel Wifi and mobile operators with caution; always offer a POST alternative.

Avoid HTTP methods which are not well defined, such as PATCH.

### Representations are for the consumer
Offer content for each thing as a human-readable HTML, with links to content in alternative machine-readable representations:

- [JSON](http://en.wikipedia.org/wiki/JSON) for convenient processing in most programming languages
- [JSONP](http://en.wikipedia.org/wiki/JSONP) and JSON with [CORS](http://en.wikipedia.org/wiki/Cross-origin_resource_sharing) for client-side JavaScript
- [CSV](http://en.wikipedia.org/wiki/Comma-separated_values) for importing into spreadsheets
- [Atom](http://en.wikipedia.org/wiki/Atom_\(standard\)) for [feeds](http://en.wikipedia.org/wiki/Web_feed).

Where possible, also offer other formats most suited to a specific domain, such as:

- [iCalendar](http://en.wikipedia.org/wiki/ICalendar) for events
- [vCard](http://en.wikipedia.org/wiki/VCard) for name and addresses
- [KML](http://en.wikipedia.org/wiki/Keyhole_Markup_Language) and [geoRSS](http://en.wikipedia.org/wiki/GeoRSS) for geographical data
- [m3u](http://en.wikipedia.org/wiki/.m3u) for playlists

Include hyperlinks to alternative representations as [link headers](http://www.w3.org/TR/html51/document-metadata.html#the-link-element) as well as in content.

_Consider also encoding meta-data inside HTML content using semantic markup: [Microformats](http://microformats.org/). [RDFa](http://en.wikipedia.org/wiki/Rdfa) or [schema.org](http://schema.org/)._

The representations supported by an API for input will vary depending upon the complexity of the action, but where possible should include `application/x-www-form-urlencoded` to allow the construction of simple [POST](http://en.wikipedia.org/wiki/POST_\(HTTP\)) forms.

### Names reinforce conventions
Use names for fields, formats and path segments in a URI path consistently across your API.
Establish conventions others may easily follow, and anticipate.
Where possible, reuse names widely used elsewhere on The Web, as with the [Microformats naming policy](http://microformats.org/wiki/naming-principles).

### Document by discovery … and example
Building a Website which exposes the data through links, and services through HTML forms encourages exploration and leads to [discovery through hypertext](http://roy.gbiv.com/untangled/2008/rest-apis-must-be-hypertext-driven).
Provide documentation for your API using examples.
Collect how people are using your API, especially link to any open source projects for projects, wrappers and programming language libraries.
Provide simple ways to experiment, as with [The Guardian API explorer](http://explorer.content.guardianapis.com/).

### Explicitly set expectations
Be clear in Web pages and other documentation as to the security, availability, rate-limiting, expected responsiveness of the platform and the provenance of data, so consumers may plan their commitment to using your API.

### Be public by default
Lower the barriers to others using your data; don't demand registration or API keys for public data.

Open data increases the number of people able to use your data and service, and leads to feedback loops where consumers become motivated to resolve issues at source, feeding back issues and correction to your service and the data within.

Where content is sensitive, or requires authentication, use encryption (HTTPS) and a standard authentication such as [Basic Authentication](http://en.wikipedia.org/wiki/Basic_access_authentication) or [OAuth](http://en.wikipedia.org/wiki/OAuth), depending upon the sensitivity of your content.

### Communicate breaking changes
Practice service evolution:

- Build for [forwards compatibility](http://en.wikipedia.org/wiki/Forward_compatibility) by gracefully handling content that is unexpected. The [robustness principle](http://en.wikipedia.org/wiki/Robustness_principle) &mdash; 
Postel's Law explains the ability for The Web and Internet to evolve, though you shouldn't ignore protocol errors, corrupted, or invalidly formatted content.
- Preserve [backwards compatibility](http://en.wikipedia.org/wiki/Backward_compatibility) with existing consumers of your API, by sending expected content,
Eschew changes to the semantics of content, e.g. don't change a `title` field from to the prefix for a name to the person's job title.

Where a revolutionary change is unavoidable, communicate a breaking change by changing the URL.
When changing URIs, continue to honour old consumers, possibly use a [redirection](http://digital.cabinetoffice.gov.uk/2012/10/11/no-link-left-behind/).
[Cool URIs don't change](http://www.w3.org/Provider/Style/URI.html).

##Further reading
This guide is an outline of the alpha [GDS API Design Principles](https://github.com/alphagov/api-design-principles).
The [API Craft Group](http://groups.google.com/group/api-craft) is a reasonably active public forum for discussing publishing APIs.

