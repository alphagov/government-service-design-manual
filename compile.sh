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
fi

bundle exec jekyll $GUIDANCE_PATH
find $GUIDANCE_PATH -name "*.html" -exec sed -i -e 's#href="/#href="/service-manual/#g' -e 's#src="/#src="/service-manual/#g' {} \;
find $GUIDANCE_PATH -name "*.css" -exec sed -i -e 's#url(/#url(/service-manual/#g' -e 's#url("/#url("/service-manual/#g' {} \;

echo "*****"
echo "You should now use git to push the changes to the design principles app."
echo "And then schedule a deployment for it."
