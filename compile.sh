#!/bin/bash

set -e

echo "Updating submodules to ensure toolkit up to date..."
git submodule update --init
echo "Updated."

DESIGN_PRINCIPLES_ROOT="../design-principles"
DIRECTORY="$DESIGN_PRINCIPLES_ROOT/public"
GUIDANCE_PATH="$DIRECTORY/service-manual"
SEARCH_CONTENT_PATH="$DESIGN_PRINCIPLES_ROOT/db/index/"

if [ ! -d "$DIRECTORY" ]; then
  echo "Couldn't find design principles app in $DIRECTORY"
fi

if [ -d $GUIDANCE_PATH ]; then
  echo "Emptying existing service-manual folder"
  rm -rf $GUIDANCE_PATH
  rm -rf ./_site
fi

bundle exec jekyll ./_site

# Jekyll will fail to compile the Sass files if the frontend toolkit isn't checked out
# however Jekyll doesn't exit with a helpful error, so the build doesn't fail.
function check_css_compiled {
    test -f $1 || echo "ERROR: $1 not compiled, Sass compilation above probably failed." && exit 1
}
check_css_compiled ./_site/service-manual/assets/stylesheets/main.css
check_css_compiled ./_site/service-manual/assets/stylesheets/print.css
check_css_compiled ./_site/service-manual/assets/stylesheets/main-ie6.css
check_css_compiled ./_site/service-manual/assets/stylesheets/main-ie7.css
check_css_compiled ./_site/service-manual/assets/stylesheets/main-ie8.css

cp -R ./_site/service-manual $GUIDANCE_PATH
mv ./.search-index.json $SEARCH_CONTENT_PATH/service-manual.json

echo "*****"
echo "You should now use git to push the changes to the design principles app."
echo "And then schedule a deployment for it."
