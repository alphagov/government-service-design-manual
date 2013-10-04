# Government Service Design Manual

This is the repository for the UK government's guidance and standards for developing digital services and contains both the assets and the content for the site.

The site is built using [Jekyll](http://jekyllrb.com/), and the CSS is [Sassified](http://sass-lang.com).

## Running the app yourself


Depending on your version of Git, you may need to run these commands to bring in the [GOV.UK Frontend Toolkit](https://github.com/alphagov/govuk_frontend_toolkit):

    $ git submodule init
    $ git submodule update

This should import Sass files from the GOV.UK Frontend Toolkit repository into `/service-manual/assets/toolkit/`.

Next, install dependencies using bundler::

    $ bundle install

To start the app::

    $ bundle exec jekyll --server --auto

This will compile the site and make it available on http://localhost:4000/service-manual. Changes will be reloaded without a restart.


## YAML page meta data schema

Pages in a Jekyll site begin with a section of YAML meta data, which specifies how the page should be rendered, where it should be linked from and so on.

The following keys are used throughout:

### Content page meta data

| Key | Values | Description |
|-----|--------|-------------|
| `title:` | Page title | Used at the top of the page, and in links to the page |
| `subtitle:` | Page subtitle | Used below page title, and optionally in links to the page |
| `category` | `agile`, `assisted-digital`, `user-centered-design`, etc | The category the page lives in the browse page |
| `layout:` | `detailed-guidance`, `role-index`, `phases`, etc | Page layout |
| `type:` | `guide`, `resource` | Used to differentiate content types so they can be grouped on index pages |
| `audience: primary:` | `designer, developer, researcher` etc. | Primary audience. Link to page will appear in top half of audience index page |
| `audience: secondary:` | | Secondary audience. Link to page will appear in bottom half of audience index |
| `phases:` | `discovery`, `alpha`, `beta`, `live` | An array of values. Adds links to page header and lists pages in relevant phase page |
| `breadcrumbs:` | An array of objects with `title:` and `url:` values | Adds a breadcrumb trail to the top of the page |

Index page meta data
--------------------

Index pages list out links to content pages, for example, the role specific pages that are built using the ``audience`` keys.

| Key | Description |
|-----|-------------|
| `hero:` |  Takes the value of the `title` key of a page | Use this to select a page for the hero promo on an index page. Any items appearing in the 'hero' slot will be removed from other lists on the page |



