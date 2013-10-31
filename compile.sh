#!/bin/bash

set -e

if [ $(find . -name '*.orig' -type f | grep -c .) -ne 0 ]; then
  echo "You seem to have some git merge artifacts on the filesystem. Aborting..."
  exit 1
fi

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

echo "Copying the compiled manual to $GUIDANCE_PATH"

cp -R ./_site/service-manual $GUIDANCE_PATH
mv ./.search-index.json $SEARCH_CONTENT_PATH/service-manual.json

echo "************************************************"
echo "** design-principles app updated successfully **"
echo "************************************************"
echo "You should now push the design-principles changes to GitHub"
echo "and then schedule a deployment for it"
