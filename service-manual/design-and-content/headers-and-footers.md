---
layout: detailed-guidance
title: Headers and footers
subtitle: How to get a head (and a foot) on GOV.UK
category: design-and-content
type: guide
audience:
  primary: designers
  secondary: service-managers, developers, content-designers
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
    title: Design and content
    url: /service-manual/design-and-content
---

Users will need to know if your service is in [alpha]() or [beta](). Many will not recognise these terms, or know what they mean. Users therefore need to be made aware that a service is not yet 'officially' live, and they therefore to use it at their own risk. That means labelling the service clearly, and explaining what it means for them. 

## Headers
Alpha or beta labels should appear in the header, and link to an about page on your service clearly explaining what these terms mean.

### Service title and navigation
The title of your service should sit in the header, clearly labelling the service to your users. This title should link back to the homepage of your service (which should generally be http://servicename.service.gov.uk/).

[![Example headers showing service title and alpha/beta badges](/service-manual/assets/images/header-footer/not-govuk-standard.png)](/service-manual/assets/images/header-footer/not-govuk-standard.png)

If necessary a navigation menu should live in the header. This should link to the main areas of your service. This menu should not be more than two lines deep at full width and ideally should contain as few links as possible. This will obviously depend on the complexity of your service but if you are finding you have a lot of links here you should think about the user need for each top level link.

[![Example headers showing navigation](/service-manual/assets/images/header-footer/not-govuk-nav.png)](/service-manual/assets/images/header-footer/not-govuk-nav.png)

The links in the header should remain constant no matter the state of the service or the stage in the process a user has reached. Navigation and information that does change throughout the process (eg user account information, shopping basket count, etc) should sit in a separate sub-navigation.

### Search
If your service requires a search function this should sit in the header. The search form should be clearly labelled, indicating that it is only searching your service and not the entire GOV.UK domain.

[![Example headers showing search box](/service-manual/assets/images/header-footer/not-govuk-search.png)](/service-manual/assets/images/header-footer/not-govuk-search.png)

You should try to not complicate the interface by having multiple search boxes on a single page, however, you may need to use in-page filters. These should be clearly and accessibly labeled and given appropriate ARIA attributes. They should also be visually differntiated from the main search box.

### Contextual navigation
Information that changes over the course of the transaction or service should not sit in the header bar. This includes examples such as user account and login information, shopping basket and checkout links. We recommend using a separate sub-navigation level to show log in/out, user account status and settings links. This sub-navigation could also include basket and checkout links if your service contains shopping basket.

### Using the GOV.UK logotype
You can only use the GOV.UK logotype if your service is currently available on the service.gov.uk subdomain. If the service is on a temporary domain then you must use an alternative header that does not show the GOV.UK logotype or crown device.

[![Example headers showing GOV.UK logo](/service-manual/assets/images/header-footer/govuk-standard.png)](/service-manual/assets/images/header-footer/govuk-standard.png)

If your service is availble on the service.gov.uk subdomain you may use the GOV.UK logotype prefixed with an alpha or beta label. You may not use the crown until your service has passed the standard and has been declared live by the [go live panel](/service-manual/digital-by-default/go-live-panel.html). The logotype should link back to [GOV.UK](https://www.gov.uk/)

[![Example headers showing GOV.UK logo with navigation](/service-manual/assets/images/header-footer/govuk-search-nav.png)](/service-manual/assets/images/header-footer/govuk-search-nav.png)

### Using the crown
Once your service has passed the standard and has been declared live you may use the crown.

## Footers
The footer of your service should follow the design patterns used by the GOV.UK footer. It should include links to secondary information for your service including user support and feedback.

The footer should also state the appropriate copyright and licence notices.

## Templates and styles
If need access to the templates and styles please make a request via [our feedback form](/feedback).