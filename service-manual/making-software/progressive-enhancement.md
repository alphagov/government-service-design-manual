---
layout: detailed-guidance
title: Progressive enhancement
subtitle: How to create pages that work regardless of browser capability
category: making-software
type: guide
audience:
  primary: designers, developers
  secondary:
status: draft
phases:
  - beta
  - live
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Making software
    url: /service-manual/making-software
---

When creating web pages, the only part of it that you can rely upon working is the HTML (and even that can fail, but without it there is no web page and everything else becomes moot). The attitude towards building for the web with this in mind is called [progressive enhancement](https://en.wikipedia.org/wiki/Progressive_enhancement).

This means in essence that each extra layer (images, styling, behaviour, video, audio) of the page should be seen as optional. If you build pages with the idea that parts other than HTML are optional, you will create a better and stronger web page.

## First, just make it work

Any new page or feature should first be made to work with HTML only. No images, no CSS, no JavaScript, nothing but HTML. Interactive elements are only those capable of being implemented with [forms](https://en.wikipedia.org/wiki/Form_(HTML)) and server-side processing.

This gives a baseline experience which will work in practically every browser, allowing your site to work for as many people and devices as possible, including older, legacy browsers and devices.

## Second, make it work better

From this baseline, the extra layers can be added. The page can have images added, advanced styling can be applied. Interactions can be made smoother and faster without the need to refresh the entire page. Validation of submitted data can be performed before it hits the network. Charts and data tables can be turned into visualisations with interactive elements.

However, it's important that each addition is seen as just that — an addition. Something extra that the more modern browsers are capable of doing *on top* of an already accessible and usable experience with as much of the content and interactivity as possible available already.

## It isn't about "JavaScript off"

A common misunderstanding is that designing sites to work without CSS, JavaScript or anything else is that this is not about a conscious choice made by the person visiting the web page that we can ignore. Nor can we treat this as an error or mistake, and it isn't a case of "just fix your browser".

There are many scenarios where the extra layers can fail to load, including temporary network errors and DNS lookup failures. The server that the resource is found on could be overloaded or down, and fail to respond in time or at all.

On top of failures there's deliberate filtering of the internet. Many large institutions (eg banks and financial institutions, some government departments) have corporate firewalls that block, remove or alter content from the internet. Mobile network providers have been known to resample images and otherwise alter content in order to make the load times faster and reduce bandwidth consumed. Also some antivirus and personal firewall software will alter and/or block content.

And, of course, there are people who do turn off features in their browsers deliberately. We should respect their decision and still provide them with a usable and useful service.

## It isn't only about JavaScript

Another common misunderstanding is that progressive enhancement is only about JavaScript. Anything more than HTML is an enhancement and should be treated as such.

Adding an image to a page requires suitable alternative text (and may also need a longer description) for people with visual impairment, and simply for when the image fails to load.

Styling cannot be the only method by which information is shared. "Text in bold" is not enough when you can't differentiate bold from normal text. "Red items are required" is not enough when you don't have the colours applied, cannot see the colours or cannot distinguish them because of colour blindness.

Video and audio without transcripts and/or subtitles are unsuitable for hearing impaired people. Information appearing in a video in a visual-only form (eg a question then answered by the people on screen in an interview) is unsuitable for people with visual impairments.

Making an interactive element that requires a mouse (eg a hover effect, a drag and drop operation) is unusable by someone who only uses a keyboard or a touch-only device such as a smartphone or tablet.

## Further reading

* [Progressive enhancement](https://en.wikipedia.org/wiki/Progressive_enhancement) on Wikipedia
* [Understanding progressive enhancement](http://alistapart.com/article/understandingprogressiveenhancement) on A List Apart
* [For a Future-Friendly Web](http://alistapart.com/article/for-a-future-friendly-web) on A List Apart
* [Cutting the mustard](http://responsivenews.co.uk/post/18948466399/cutting-the-mustard) on the BBC News responsive design blog
