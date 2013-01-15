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
