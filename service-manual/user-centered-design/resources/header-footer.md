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

Your service must be clearly labelled Alpha or Beta using one of the [three design patterns](#design-patterns) described below. 

### Design patterns

So far, we have identified three patterns that fit the situations we have encountered, we may update this list as the exemplar programme matures and we have strong evidence on how they perform.

Each design pattern also has a banner below the title, which should briefly explain the stage of the project -  'alpha' or 'beta' - and that it is a trial service.

The text for the banners should be as follows:

> **ALPHA**: This is a prototype

> **BETA**: This is a trial service

Services that aren't a transaction funnel may optionally choose to link to a feedback survey and description of alpha or beta:

> **ALPHA**: This is a prototype -- your [feedback]() will help us to improve it. [Find out more](/service-manual/phases/alpha)

> **BETA**: This is a trial service -- your [feedback]() will help us to improve it. [Find out more](/service-manual/phases/beta)

Services should only add these links when it is acceptable for a user to be taken away from their existing flow.

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

If you service needs to show navigation, then this should sit below the service title. The service title should link to the service start page.

![Pattern 3: Title and navigation](/service-manual/assets/images/header-footer/header-pattern-3.png)

In this case "Your Tax account" is not a transactional funnel, so it can also include links to feedback and information.

### Search

If your service requires a search function this should sit in the header. The search form should be clearly labelled, indicating that it is only searching your service and not the entire GOV.UK domain.

You should try to not complicate the interface by having multiple search boxes on a single page, however, you may need to use in-page filters. These should be clearly and accessibly labeled and given appropriate ARIA attributes. They should also be visually differentiated from the main search box.

### Soliciting end user feedback

End user feedback can provide data that is invaluable for developing and improving the subsequent iterations of your service. You should solicit feedback by providing a link to your service's feedback mechanism from the 'alpha'/'beta' banner. The link could:

* open a blank email in the user's email client, using the [`mailto` URL scheme](http://www.ietf.org/rfc/rfc2368.txt)
* toggle the appearance of an embedded web form
* take the user to a dedicated feedback web form on another page within your service

### Using the GOV.UK logotype

You can only use the GOV.UK logotype if your service is [currently available on GOV.UK](/service-manual/user-centered-design/service-look-and-feel). If the service is on a temporary domain then you must use an alternative header that does not show the GOV.UK logotype or crown device. The logotype must link back to [GOV.UK](https://www.gov.uk/).

## Footers

The footer of your service should follow the design patterns used by the GOV.UK footer. It should include links to secondary information for your service that would allow users to contact you directly or leave feedback about your service (such as email addresses, phone numbers or contact forms).

The footer should also state the appropriate copyright and licence notices.

## Templates and styles

If you need access to the templates and styles, please make a request via [the GOV.UK support form](/support/internal).

