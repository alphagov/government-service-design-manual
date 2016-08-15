---
layout: phases
title: Retirement phase
subtitle: Learn what to do when your service is retired or superseded
phase: retirement
category: guidance
status: draft
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Phases of service design
    url: /service-manual/phases
exclude_from_search: true
---

{:.intro}
Even the best services may eventually reach retirement. Changes in policy may mean that the service is no longer offered or new understanding may mean that those user needs are better provided through a different service.

Whatever the reason, the retirement of digital services should be handled with the same degree of care as their creation, concentrating on user needs.

## User needs

Your service will have been built to serve user needs. It's vital to understand how those needs are to be addressed once your service has been retired, whether they are deemed to no longer exist, are no longer served by government, or will in future be served as part of another service.

If the needs will no longer be met by government but will now be met by the private or voluntary sectors then it's important to for these organisations to be adequately prepared. This means that the appropriate online user journey can be developed.

Similarly, if the needs are to be served as part of another government service or services then you should identify which services they are and work with those services' teams. This means the new teams can learn from your experience and you can understand how you'll support your users to make the transition.

## Next stages

### Let your users know

The vast majority of users of your service will begin and end their journeys via GOV.UK. As soon as you know that your service is likely to be retired you should contact the GOV.UK team to make sure that those journeys are amended and appropriate information is supplied.

The GOV.UK team will need to know why the service is being retired, and how those user needs are to be served in future so that they can provide the appropriate information, advice and links to users.

For that majority of users who begin and end their journeys on GOV.UK the most important thing is to ensure that GOV.UK is updated. There will be some users, however, who access your service directly whether via links in emails, bookmarks on their computer or remembering your URLs. It's important to prepare them for the change and lay out clearly what it will mean to them.

Your planning should aim to produce the minimum possible disruption for users, but it will still be a significant change for them. Details of what the change is, why it's being made, what they will need to do, and what will happen to their data should be made easily available.

Users who access your service via an application programming interface (API) will need time to update their software to use the replacement service's APIs or to make other relevant adjustments. You should reach out to your API users as early as possible and remember that they may have significant lead times for making and distributing changes. Changes to the service online will also need to be seamlessly tied in with messaging to offline users who are receiving the service through assisted digital channels.

### Plan to redirect traffic

Once GOV.UK is updated the vast majority of users will begin to be directed to the new service. Some users will still try to access the service at its current (now retired) home. You should have a plan for redirecting those requests to the appropriate new service, or to provide clear information about the service that has been retired in perpetuity.

> When we retired the DirectGov, BusinessLink and departmental websites, we invested a lot of effort in mapping content on those sites to GOV.UK content. We redirected as many of the old sites' URLs as we could to corresponding GOV.UK URLs, and provided 'Gone' pages with links to the National Archive where the content was not being transferred. We wrote about that in ["No Link Left Behind"](https://gds.blog.gov.uk/2012/10/11/no-link-left-behind/).

### Make sure your subdomain continues to work

If you operate a `service.gov.uk` subdomain, please read [your technical responsibilities for ending your service](/service-manual/operations/operating-servicegovuk-subdomains#lifecycle-of-service-subdomains).

## Data

In running your service you will have accrued a large amount of data about the service and its users. You should already have policies in place to manage that data responsibly, including details of how long it will be retained. Those policies will continue to apply, and you will need to ensure that there is support in place to maintain them.

Where data is being transferred to a new service owner, that should be done in accordance with your existing data protection policies, and communicated clearly to your users.

{:.phase-nav}
* [Previous phase: live](/service-manual/phases/live.html)

*[APIs]: application programming interfaces
*[API]: application programming interface
