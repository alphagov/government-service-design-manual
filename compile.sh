#!/bin/bash

set -e

DIRECTORY="../design-principles/public"
GUIDANCE_PATH="$DIRECTORY/service-manual"

if [ ! -d "$DIRECTORY" ]; then
  echo "Couldn't find design principles app in $DIRECTORY"
fi

if [ -d $GUIDANCE_PATH ]; then
  echo "Emptying existing service-manual folder"
  rm -rf $GUIDANCE_PATH
  rm -rf ./_site
fi

bundle exec jekyll ./_site
cp -R ./_site/service-manual $GUIDANCE_PATH

echo "*****"
echo "You should now use git to push the changes to the design principles app."
echo "And then schedule a deployment for it."
