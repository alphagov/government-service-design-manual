README.md

# YAML page meta data schema

Pages in a Jekyll site begin with a section of YAML meta data,
which specifies how the page should be rendered, where it should be linked from and so on.

We've used a mixture of built in and custom tags, as follows:





* `layout` — Page layout. Can be 'default' or 'wide'
* `title` — Page tite. Used at the top of the page, and in links to the page
* `subtitle` — Page subtitle. Used below page title, and optionally in links to the page

* `audience` — Audience for the content. Split into primary and secondary audience. Takes a list of roles (designer, developer etc.)
  * `primary`
  * `secondary`

==================   ======================================================
Command              Description
==================   ======================================================
gh-browse            Manually browse the GitHub API by URL
gh-list-members      List members of a GitHub organisation (may be removed)
gh-login             Login to GitHub, generating an OAuth login token
gh-migrate-project   Migrate an entire project between GitHub instances
gh-migrate-teams     Migrate organisation teams between GitHub instances
gh-org               Interact with organisations 
gh-repo              Interact with repositories
gh-status            Submit commit build status to GitHub
==================   ======================================================



## audience:

  primary:
  secondary:

Who this content is intended for.
Can be one or more of: service-manager, designer, developer, tech-arch, researcher, analyst, content-designer
Used to populate the role-specific landing pages
  If the audience is primary, the page is promoted in the top half of the page
  If the audience is secondary, the page will appear in the 'More guides from the manual' section


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