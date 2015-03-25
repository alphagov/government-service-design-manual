# Service manual search

This directory contains a layout, view and includes for
search in the service manual.

The GOV.UK template in `layouts/` is duplicated from the
template for the Jekyll site with a minor change:

- `{% include %}` tags have been changed from
  `{% include layouts/_page_title.html %}` to
  `{% include 'layouts/page_title' %}`. This is due
  to an unresolvable difference in the way Jekyll implements
  Liquid includes

These includes reference files in `_search/includes/layouts/`,
which are duplicated from the main Jekyll `_includes` directory.
