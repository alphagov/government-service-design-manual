---
layout: detailed-guidance
title: Creating accessible PDFs
subtitle: How to make PDFs that everyone can use
category: user-centred-design
type: resource
audience:
  secondary: designers, developers, content-designers
status: draft
phases:
  - discovery
  - alpha
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

The best way to create an accessible PDF is to create an accessible source document. Well structured Microsoft Word documents make good source documents for conversion to PDF. When a source document is converted into PDF it is tagged. The PDF tag tree reflects the structure of the document, and it’s this structure that assistive technologies like screen readers use to navigate the document.

## In Microsoft Word

Use the styles and features available in Word to format your content and give it structure. This will make it easier to convert your source document into PDF because it lays the groundwork for the PDF tag tree.

### Use headings

Use the heading styles in Word to create a logical document structure. Don’t increase the size of text or make it bold to create the appearance of headings. Treat your document like a book: It should have one title (level one heading) and multiple chapters (level two headings). Within each chapter there may be multiple sections (level three headings) and sub sections (level four headings).

### Use lists

Use the list styles in Word to group together related items. If the items follow a specific sequence, use a numbered list instead. Don’t use punctuation or other markers to create the illusion of a list.

### Create a table of contents

If your document is longer than a few pages, use Word to automatically create a table of contents based on your heading structure. Don’t use lists and links to manually create a table of contents.

### Use readable body text

Use left aligned text (unless the language of your document is read right to left). Don’t use justified text in your document.

Choose a sans serif font and use the styles in Word to set it as the default, with a minimum size of 12pt. If you need to include footnotes or other text of a smaller size, increase the size of the body text to 14pt rather than reducing the size of any text below 12pt.

Don’t use chunks of italicised or capitalised text, and don’t underline text unless it’s a link.

### Use good colour contrast

Use foreground/background colours for text that have a good contrast ratio. The 4.5:1 ratio recommended by the [Web Content Accessibility Guidelines 2.0](http://www.w3.org/TR/WCAG/) is a good minimum.

Don’t use colour or shape as the only way to identify something in your document. Use text labels or descriptions instead.

### Use data tables

Use tables with column headings to display data. Don’t use tables to make cosmetic changes to the layout of the document.

### Provide text descriptions

Use Word to add text descriptions to all important images in the document. Make sure the text description includes all the information contained within, or conveyed by, the image.

## In Adobe Acrobat

Use Adobe Acrobat Pro to convert your Word document into PDF. Use the Convert to PDF option under the Adobe menu in Microsoft Word to do this. This will ensure that Acrobat picks up the accessibility you have built into your source document.

### Set the document language

Set the language of the document. Go to File > Properties > Advanced and select a language from the Language menu. If the PDF is written in Welsh, type CY into the box.

### Check the tag tree

All content must be tagged, marked as an artefact (background content), or removed from the tag tree. Use the Tags panel to review and edit the tag tree. If the PDF was converted from a well structured Word document, the tag tree should require little editing.

### Check the tab order

If the PDF contains form fields, use Advanced > Accessibility > Touch up reading order to check they can be navigated with the tab key in a logical order. If the tab order needs improving, use the Order panel to drag and drop the fields into the correct order.

### Check the reading order

Use the Tags panel to review and edit the reading order of the PDF. Don’t rely on the visual order of the PDF. The reading order is based on the structure of the PDF tag tree, which may not match the visual content order.

### Check the reflow order

Use View > Zoom > Reflow then check that the PDF still has a logical reading order. Note: It can sometimes be difficult to guarantee a logical reflow order for PDFs with complex content.

### Check text descriptions

Go to Advanced > Accessibility > Touch up reading order and check that all images have text descriptions. If the text descriptions were present in the source Word document and the Convert to PDF option was used, the text descriptions should already be present in the PDF.

### Remove empty tags

Remove empty tags from the tag tree. Use the Tags panel to highlight and delete any empty tags from the tag tree.

### Set decorative content

Tag decorative content elements as artefacts. Use Advanced > Accessibility > Touch up reading order to select a decorative element, and use the Background button to make the element an artefact.

### Check data tables

Use the Tags panel to check the structure of data tables. The `<table>`, `<tr>` and `<td>` tags should be used to give data tables the proper structure.

### Check active links

Use the Tags panel to check that links are active. Active links should be tagged with the `<link>` tag.

### Check high contrast

Use File > Preference > Accessibility to set a high contrast colour scheme, and check the PDF remains readable. It may not be possible to make high contrast mode work in all PDFs, in which case you should be prepared to make a high contrast version available on request.

### Display document title

Display the document title instead of the file name. Go to File > Properties > Initial view and select Document title from the Show drop down box.

## Before publication

Once all the above steps have been taken, the PDF should be checked before it is published.

### Full Adobe accessibility check

Go to Advanced > Accessibility and select Full check. The PDF should pass the full check for WCAG 2.0 Level AA without any warnings.

### Quick screen reader check

Ask a screen reader user to read through the PDF. If no-one is available to do this, use one of the following options instead.

### Use NVDA

[Non Visual Desktop Access (NVDA)](http://www.nvaccess.org/) is a free open source screen reader for Windows. It can be installed to the desktop or run from a portable USB thumb drive.

With NVDA running, open the PDF and use the following commands to check the PDF:

* from the top of the PDF (with the Numlock off), use Numpad 0 + Numpad 2 to read the PDF from top to bottom and check the reading order
* use the tab key (repeatedly) to move through the PDF and check the tab order
* use the h key (repeatedly) to move through the PDF and check the heading structure
* use the g key (repeatedly) to move through the PDF and check for text descriptions

These commands will also work with the JAWS screen reader from Freedom Scientific.

### Use VoiceOver

VoiceOver is the integrated screen reader with Mac OS X and all iOS devices. In Mac OS X turn VoiceOver on (or off again) using Command + F5.

With VoiceOver running open the PDF and use the following commands to check the PDF:

* from the top of the PDF use a double finger down swipe, or Control + Option + A to read the PDF from top to bottom and check the reading order
* use the tab key (repeatedly) to move through the PDF and check the tab order

VoiceOver does not provide shortcut keys for navigating PDFs by headings or graphics.

## Acknowledgements

Thanks to the Department for Work and Pensions (DWP) for enabling us to incorporate its accessibility best practice guidance into this document.
