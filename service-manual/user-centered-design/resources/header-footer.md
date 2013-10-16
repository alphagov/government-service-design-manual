---
layout: detailed-guidance
title: Header and footer
subtitle: Guidance on using the GOV.UK header
category: user-centered-design
type: guide
audience:
  primary: designers, developers
  secondary:
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
    title: User centered design
    url: /service-manual/user-centered-design
---

Users will need to know if your service is in [alpha](/service-manual/phases/alpha) or [beta](/service-manual/phases/beta), many will not recognise these terms, or know what they mean – therefore, users need to be made aware that a service is not yet 'officially' live, and they should use it at their own risk. That means labelling the service clearly and explaining what it means for them.

## Headers

Your service must be clearly labelled Alpha or Beta using one of the [three design patterns](#design-patterns) described below. The terms [Alpha](/service-manual/phases/alpha) and [Beta](/service-manual/phases/beta) should link to a page on the service manual which clearly explains what these terms mean.

### Design patterns

So far, we have identified three patterns that fit the situations we have encountered, we may update this list as the exemplar programme matures and we have strong evidence on how they perform.

Each design pattern also has a banner below the title, which should briefly explain about what 'alpha' or 'beta' means linking through to the appropriate page on the service manual.

The text for the banners should be as follows:

> **ALPHA**: This is a prototype, designed to be simpler, clearer and faster. [Find out more](https://www.gov.uk/service-manual/phases/alpha)

> **BETA**: This is a trial service, designed to be simpler, clearer and faster. [Find out more](https://www.gov.uk/service-manual/phases/beta)


In all cases the GOV.UK logotype must link to `www.gov.uk`.

### Pattern 1: No title

This is simply the GOV.UK header as is, without a title. This is relevant if your service is small and straightforward, taking place over only a few screens.

You should also use this on the first page of your service if you are using the title of the service as the main heading the body.

![Pattern 1: No title](/service-manual/assets/images/header-footer/header-pattern-1.png)

### Pattern 2: Title

This is the pattern that will be used in most cases. Your service will not need any internal navigation, but will need a title to identify itself.

The service title should link to the service start page unless there is a reason why this would not work.

![Pattern 2: Title](/service-manual/assets/images/header-footer/header-pattern-2.png)

### Pattern 3: Title and navigation

If you service needs to show navigation, then this should sit below the service title. You shouldn’t need to include a ‘Home’ link as the service title should link to the service start page.

![Pattern 3: Title and navigation](/service-manual/assets/images/header-footer/header-pattern-3.png)

### Search

If your service requires a search function this should sit in the header. The search form should be clearly labelled, indicating that it is only searching your service and not the entire GOV.UK domain.

You should try to not complicate the interface by having multiple search boxes on a single page, however, you may need to use in-page filters. These should be clearly and accessibly labeled and given appropriate ARIA attributes. They should also be visually differentiated from the main search box.

### Using the GOV.UK logotype

You can only use the GOV.UK logotype if your service is [currently available on GOV.UK](/service-manual/user-centered-design/what-should-service-look-like). If the service is on a temporary domain then you must use an alternative header that does not show the GOV.UK logotype or crown device. The logotype must link back to [GOV.UK](https://www.gov.uk/).

## Footers

The footer of your service should follow the design patterns used by the GOV.UK footer. It should include links to secondary information for your service including user support and feedback.

The footer should also state the appropriate copyright and licence notices.

## Templates and styles

If need access to the templates and styles please make a request via [our feedback form](/feedback).

