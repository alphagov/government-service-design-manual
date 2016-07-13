# How to deploy

The Government Service Design Manual lives at [https://www.gov.uk/service-manual](https://www.gov.uk/service-manual). It is served by [the design principles application](https://github.com/alphagov/design-principles) but that is not the canonical source of the content. The content is managed in [a separate github repo](https://github.com/alphagov/government-service-design-manual).

The content is managed using a [jekyll](http://jekyllrb.com/) app and there's a script that compiles that into HTML, rewrites the URLs appropriately and copies it into the design-principles app for deployment.

That isn't exactly ideal, but allows us to experiment with the format and continue the existing workflow while we decide if this is a useful format, think about what the next iteration of the manual might look like, and generally figure out what's next.

## Steps to deploy

* Make sure you have up to date copies of the [government-service-design-manual](https://github.com/alphagov/government-service-design-manual) and [design-principles](https://github.com/alphagov/design-principles) apps in the same folder (eg ~/govuk)
* Create a version tag in the service manual repo:
  * ```$ git tag``` will list the current tags, look for the most recent tag (we are using the nomenclature `version-X`)
  * create a new tag with the next incremental version number ```$ git tag version-X```
  * push your tag to the remote ```$ git push --tags origin master```
* [Compare your tag](https://github.com/alphagov/government-service-design-manual/compare) with the previously deployed tag to verify the changes 
* Run the compile script, this will run jekyll to compile the source and copy the compiled output to the design-principles app (you will need to have golang installed).
  * ```$ ./compile.sh```
* Review and commit the changes in the design-principles app. In the commit message you should reference the version tag and include a link comparing the previous release and this one on github, (eg: https://github.com/alphagov/government-service-design-manual/compare/version-30...version-31) I don't think there's a short way of linking to a tag in a commit message yet, but I am currently using `#version-X` (as well as the URL).
* Deploy `design-principles` as usual (The buttons in the release app don't work so use Jenkins)