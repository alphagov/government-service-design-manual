---
layout: detailed-guidance
title: Choosing appropriate formats
subtitle: Help your users by providing content in a format they can use
category: design-and-content
type: guide
audience:
  primary: designers, developers, tech-archs, user-researchers, service-managers, content-designers
  secondary: delivery-managers, performance-analysts
status: draft
phases:
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

Almost all content relating to the policies or publications of government departments should live on [GOV.UK](https://www.gov.uk/). Where exceptions to this rule are required, content and data should be provided in formats that appropriately reflect their purpose and intended audience.

## Choose the format to fit the content

You should publish documents in formats that reflect the format of the information they contain, and the uses to which they will likely be put.

- For tabular data, use [CSV](https://en.wikipedia.org/wiki/Comma-separated_values) or a similar "structured data" format (see also [JSON](https://en.wikipedia.org/wiki/JSON) and [XML](https://en.wikipedia.org/wiki/XML)). Do not publish structured data in unstructured formats such as PDF.

- For written reports, consider the accessibility of your report to your readers. PDF can be an excellent display format, but without additional effort it can be inappropriate for users of screenreading software. You might use a more accessible format such as HTML, provide accessible alternate formats for the document, or invest effort in [accessibility tagging your PDFs](/service-manual/design-and-content/resources/creating-accessible-PDFs.html).

If you are publishing tabular information (financial reports, statistical data, etc.) then your users may well wish to process this data programmatically. By publishing in a "machine-readable" format such as CSV, you are helping them do so with minimal additional effort. Similarly, by making available accessible versions of textual reports, you are ensuring that you do not exclude users who have every right to receive your publications.

In summary, consider your users, and the uses to which they will put your published data and content. If in doubt, treat the native format of the web, HTML, as a good fallback option. Web browsers are available on all platforms and devices, and web pages tend to be both passably accessible and machine-readable.

## Don't assume your users can read proprietary formats

Wherever possible, publish in accessible, patent-free, [open formats](https://en.wikipedia.org/wiki/Open_format), for which software is widely available on a variety of platforms. If publishing in proprietary formats, you should always make a non-proprietary alternative available.

- For textual reports, provide HTML, plain text (.txt), or PDF rather than formats that require proprietary software to view, such as Word documents (.doc).

- For tabular data, provide CSV or [TSV](https://en.wikipedia.org/wiki/Tab-separated_values) rather than Excel spreadsheets (.xls).

Again, if in doubt, you should treat the native format of the web, HTML, as your best default option.

*[CSV]: Comma-separated values
*[TSV]: Tab-separated values
*[PDF]: Portable Document Format
*[JSON]: JavaScript Object Notation
*[XML]: Extensible Markup Language
