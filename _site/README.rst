

YAML page meta data schema
==========================

Pages in a Jekyll site begin with a section of YAML meta data,
which specifies how the page should be rendered, where it should be linked from and so on.

We've used a mixture of built in and custom tags, as follows:





* `layout` — Page layout. Can be 'default' or 'wide'
* `title` — Page tite. Used at the top of the page, and in links to the page
* `subtitle` — Page subtitle. Used below page title, and optionally in links to the page

* `audience` — Audience for the content. Split into primary and secondary audience. Takes a list of roles (designer, developer etc.)
  * `primary`
  * `secondary`


====================   ======================================================
Tag                    Options
====================   ======================================================
layout:                Page layout
title:                 Page tite. Used at the top of the page, and in links to the page
subtitle:              Page subtitle. Used below page title, and optionally in links to the page
audience: primary:     Primary audience. Link to page will appear in top half of audience index page
audience: secondary:   Secondary audience. Link to page will appear in bottom half of audience index page
====================   ======================================================



## status:

Options: draft, experimental, mandatory…
Triggers a flag at the top-right of the content area


## type:

What kind of content is on the page
Options:  guide   Something to be read (words, pictures...)
    resource    Something to be used (code, web assets...)
Used: To group links on landing pages

## hero:

Used on index pages to determine what page should be promoted in the 'hero' slot
Takes the page title of a page.
Any items appearing in the 'hero' slot will be removed from other lists on the page.

## section:

Added as class to the body tag. Used for global navigation and section-specific styles

## assets: local

Use this on a page if you need to work temporarily offline.
The page will reference locally stored copies of the main template files instead of the ones in
the GOV.UK preview environment. Note that the local versions will most likely be out of date.


## Table of contents

For longer pages you may want to add a table of contents.

Just add this line after your intro paragraph:

{ nomarkdown } {% include _toc.md %} { endnomarkdown }