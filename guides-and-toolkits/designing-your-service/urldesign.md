---
layout: gsdm
title: URL design
section: guidance
subsection: Designing your service
status: draft
---
    
# URL design
URLs are part of the user interface of a service. They should be designed to be descriptive and bookmarkable, they shouldn't be used to preserve state unless it can be shared. Use HTTP properly. 

## Guidance

URLs are a fundamental building block of any digital service. They are part of its user interface, providing a way for users to navigate the service and to understand how it's structured. For services that need to be indexed in search engines they are also an important piece of data that search indexers take into account for ranking pages.

URLs should be designed for clarity, and according to the practices set out in the HTTP specifications. They should be easy for moderately skilled humans to read as well as for machines to process. They should provide ways to point to the specific pieces of information or entry points for interactions within your service.

HTTP expects URLs to be used in conjunction with verbs (primarily GET and POST) and it is important to get these right in order to build a flexible, effective digital service.

* A GET should be used to retrieve information about a resource from a server. It should not change the state of that resource. A user executing a GET for a URL should get a simple, appropriate response without side effects. A search usually makes sense to be executed via a GET.
* A POST should be used where the state of a resource is expected to change. Submitting a form is usually done with a POST where you expect to store or process that data.

URLs are a commitment and shouldn't be changed without careful thought. If you're changing your URLs then you should redirect from the old ones to the new ones as users may well have linked to or bookmarked the old URLs and will expect a consistent service.

##Why we do this

> The "Smart Answers" on GOV.UK are designed for sharing. Having calculated your maternity or paternity entitlements you may well want to share the outcome with your partner or save it to come back later. The URLs contain all your previous answers and so sending that URL to anyone else will allow them to see exactly the same outcome you saw.

##Further reading

* No Link Left Behind, Paul Downey on [how we saved all the DG and BL URLs](http://digital.cabinetoffice.gov.uk/2012/10/11/no-link-left-behind/)
* Kyle Neath of Github.com talks about [their URL design](http://warpspire.com/posts/url-design/)
* Tim Berners-Lee, the creator of the web is commonly heard saying [Cool URIs Don't Change](http://www.w3.org/Provider/Style/URI.html)
