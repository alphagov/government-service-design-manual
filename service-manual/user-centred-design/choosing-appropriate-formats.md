---
layout: detailed-guidance
title: Choosing appropriate formats
subtitle: Help your users by providing content in a format they can use
category: user-centred-design
type: guide
audience:
  primary: designers, service-managers
  secondary: delivery-managers, performance-analysts, developers
status: draft
phases:
  - beta
  - live
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: User-centred design
    url: /service-manual/user-centred-design
---

Almost all content relating to the policies or publications of government departments should live on [GOV.UK](https://www.gov.uk). Where exceptions to this rule are required, content and data should be provided in formats that appropriately reflect their purpose and intended audience.

## Choose the format to fit the content

You should publish documents in file formats that reflect the nature of the information they contain, and the uses to which they will likely be put.

- For written reports, the native format of the web (HTML) should be your default option. PDF can be an excellent display format, but without additional effort it can be inappropriate for users of screenreading software. It's faster and easier to make accessible HTML that's suitable for every platform and device. If you must publish PDFs, you should provide accessible alternate formats for the document, and invest effort in [accessibility tagging your PDFs](/service-manual/user-centred-design/resources/creating-accessible-PDFs.html).

- For data, use [CSV](https://en.wikipedia.org/wiki/Comma-separated_values) or a similar 'structured data' format (see also [JSON](https://en.wikipedia.org/wiki/JSON) and [XML](https://en.wikipedia.org/wiki/XML)). Don't publish structured data in unstructured formats such as PDF.

- If you're regularly publishing data (financial reports, statistical data etc.) then your users may well wish to process this data programmatically, and it becomes especially important that your data is 'machine-readable'. PDFs, Word documents and the like are not suitable formats for data publication. In addition, you should consider making your data available through an API (application programming interface) if this will simplify your users' interactions with your publications. For more information on APIs, and for more detailed technical guides on publishing data, please see [our guidance on APIs and formats](/service-manual/making-software/apis.html#representations-are-for-the-consumer).

- If you're publishing a written report that contains statistical tables, provide the tables alongside or in addition to your report in suitable data formats.

In summary, consider your users, and the uses to which they'll put your published data and content. If in doubt, treat the native format of the web, HTML, as a good fallback option. Web browsers are available on all platforms and devices, and web pages tend to be both passably accessible and machine-readable.

## Don't assume your users can read proprietary formats

Wherever possible, publish in accessible, patent-free, [open formats](https://en.wikipedia.org/wiki/Open_format), for which software is widely available on a variety of platforms. If publishing in proprietary formats, you should always make a non-proprietary alternative available.

- For textual reports, provide HTML, plain text (.txt), or PDF rather than formats that require proprietary software to view, such as Word documents (.doc/.docx).

- For tabular data, provide CSV or [TSV](https://en.wikipedia.org/wiki/Tab-separated_values) rather than Excel spreadsheets (.xls/.xlsx).

- For other structured data, see our guidance on [representations for the
  consumer](/service-manual/making-software/apis.html#representations-are-for-the-consumer). Wherever possible, choose an open format over a proprietary one.

Again, if in doubt, you should treat the native format of the web, HTML, as your best default option.

*[APIs]: application programming interfaces
*[API]: Application Programming Interface
*[CSV]: Comma-separated values
*[TSV]: Tab-separated values
*[PDF]: Portable Document Format
*[JSON]: JavaScript Object Notation
*[XML]: Extensible Markup Language
