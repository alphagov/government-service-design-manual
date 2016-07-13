[![Build Status](https://travis-ci.org/alphagov/government-service-design-manual.svg?branch=master)](https://travis-ci.org/alphagov/government-service-design-manual)

# Government Service Design Manual

This is the repository for the UK government's guidance and standards for developing digital services and contains both the assets and the content for the site.

The site is built using [Jekyll](http://jekyllrb.com/), and the CSS is [Sassified](http://sass-lang.com). Some validation of the content happens using [Go](http://golang.org/), so you will [need to have Go installed](http://golang.org/doc/install) if you wish to use the Makefile.

## Deploying

[GDS-specific instructions are for how to deploy this content](https://github.com/alphagov/government-service-design-manual/blob/master/docs/how-to-deploy.md).

## Running the app yourself

Depending on your version of Git, you may need to run these commands to bring in the
[GOV.UK Frontend Toolkit](https://github.com/alphagov/govuk_frontend_toolkit):

    $ git submodule init
    $ git submodule update

This should import Sass files from the GOV.UK Frontend Toolkit repository into
`/service-manual/assets/toolkit/`.

Next, install dependencies using bundler::

    $ bundle install

You need to create the content::

    $ bundle exec jekyll build

To start the app::

    $ bundle exec ruby app.rb

This will compile the site and make it available at
[http://localhost:4567/](http://localhost:4567/).

In a separate terminal, you'll need to run:

    $ bundle exec jekyll build --watch

This will monitor the service manual and rebuild it whenever it's changed.

## Running on Heroku

This application should run fine on Heroku. You will need to have [set yourself up with a Heroku account and the command-line tools](https://devcenter.heroku.com/articles/quickstart). Then create an application on Heroku.

```
> heroku create <app-name>
```

Make your change locally, commit and deploy them to Heroku as needed.
If the code you're deploying is not in master, then you'll need to
make sure you specify your local branch to push to master. Otherwise
it will just deploy your local master (and probably not work as
expected).

```
> git push heroku master
```

Then open https://app-name.herokuapp.com/service-manual/ to see your changes.

### Click-through version

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://heroku.com/deploy)

## YAML page meta data schema

Pages in a Jekyll site begin with a section of YAML meta data, which specifies how the page should be rendered, where it should be linked from and so on.

The following keys are used throughout:

### Content page meta data

| Key | Values | Description |
|-----|--------|-------------|
| `title:` | Page title | Used at the top of the page, and in links to the page |
| `subtitle:` | Page subtitle | Used below page title, and optionally in links to the page |
| `category` | `agile`, `assisted-digital`, `user-centred-design`, etc | The category the page lives in the browse page |
| `layout:` | `detailed-guidance`, `role-index`, `phases`, etc | Page layout |
| `type:` | `guide`, `resource` | Used to differentiate content types so they can be grouped on index pages |
| `audience: primary:` | `designer, developer, researcher` etc. | Primary audience. Link to page will appear in top half of audience index page |
| `audience: secondary:` | | Secondary audience. Link to page will appear in bottom half of audience index |
| `phases:` | `discovery`, `alpha`, `beta`, `live` | An array of values. Adds links to page header and lists pages in relevant phase page |
| `breadcrumbs:` | An array of objects with `title:` and `url:` values | Adds a breadcrumb trail to the top of the page |
| `exclude_from_search:` | `true` | Exclude the page from search. The default behaviour is to include the page in search. |

Index page meta data
--------------------

Index pages list out links to content pages, for example, the role specific pages that are built using the ``audience`` keys.

| Key | Description |
|-----|-------------|
| `hero:` |  Takes the value of the `title` key of a page | Use this to select a page for the hero promo on an index page. Any items appearing in the 'hero' slot will be removed from other lists on the page |



