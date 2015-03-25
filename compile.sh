#!/bin/bash

set -e

if [ $(find . -name '*.orig' -type f | grep -c .) -ne 0 ]; then
  echo "You seem to have some git merge artifacts on the filesystem. Aborting..."
  find . -name '*.orig' -type f
  exit 1
fi

bundle

echo "Linting"
make lint

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

bundle exec jekyll build

# Jekyll will fail to compile the Sass files if the frontend toolkit isn't checked out
# however Jekyll doesn't exit with a helpful error, so the build doesn't fail.
echo "Checking CSS compiled"

function check_css_compiled {
    if [ -f $1 ]; then
        echo " OK: Found CSS file $1";
    else
        echo " ERROR: $1 not compiled, Sass compilation above probably failed.";
        exit 1;
    fi
}
check_css_compiled ./_site/service-manual/assets/stylesheets/main.css
check_css_compiled ./_site/service-manual/assets/stylesheets/print.css
check_css_compiled ./_site/service-manual/assets/stylesheets/main-ie6.css
check_css_compiled ./_site/service-manual/assets/stylesheets/main-ie7.css
check_css_compiled ./_site/service-manual/assets/stylesheets/main-ie8.css

if [ -n "$TRAVIS" ]; then
  exit 0
fi

echo "Copying the compiled manual to $GUIDANCE_PATH"

cp -R ./_site/service-manual $GUIDANCE_PATH
mv ./.search-index.json $SEARCH_CONTENT_PATH/service-manual.json

echo "************************************************"
echo "** design-principles app updated successfully **"
echo "************************************************"
echo "You should now push the design-principles changes to GitHub"
echo "and then schedule a deployment for it"
