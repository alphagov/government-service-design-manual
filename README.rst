Government Service Design Manual
================================

This is the repository for the UK Government's guidance and standards for developing digital services and contains both the assets and the content for the site.

The site is built using [Jekyll](http://jekyllrb.com/), and the CSS is [Sassified](http://sass-lang.com).

YAML page meta data schema
==========================

Pages in a Jekyll site begin with a section of YAML meta data, which specifies how the page should be rendered, where it should be linked from and so on.

The following keys are used throughout:

Content page meta data
----------------------

====================   ====================================   =========================================================================================================================================================================================================================================================
Key                    Values                                 Description
====================   ====================================   =========================================================================================================================================================================================================================================================
layout:                default, wide                          Page layout
title:                                                        Page tite. Used at the top of the page, and in links to the page
subtitle:                                                     Page subtitle. Used below page title, and optionally in links to the page
status:                draft, experimental etc.               Triggers a label in the top-right of the content area
type:                  guide, resource                        Used to differentiate content types so they can be grouped on index pages
section:               home, dbd, about                       Added as class to the body tag. Used to set global navigation and section-specific styles
audience: primary:     designer, developer, researcher etc.   Primary audience. Link to page will appear in top half of audience index page
audience: secondary:                                          Secondary audience. Link to page will appear in bottom half of audience index page
assets:                local                                  Use this on a page if you need to work temporarily offline. The page will reference locally stored copies of the main template files instead of the ones in the GOV.UK preview environment. Note that the local versions will most likely be out of date.
====================   ====================================   ==========================================================================================================================================================================================================================================================


Index page meta data
--------------------

Index pages are pages that list out links to content pages, for example, the role specific pages that are built using the ``audience`` keys.

====================   ===============================================   ===================================================================================================================================================
Key                    Values                                            Description
====================   ===============================================   ===================================================================================================================================================
hero:                  Takes the value of the ``title`` key of a page    Use this to select a page for the hero promo on an index page. Any items appearing in the 'hero' slot will be removed from other lists on the page.
====================   ===============================================   ===================================================================================================================================================


Tables of contents
=================

For longer pages you may want to add a table of contents.

Just add this line after your intro paragraph:

{% include _toc.md %}