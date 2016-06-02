---
layout: detailed-guidance
title: Managing software dependencies
subtitle: Making use of third party libraries and frameworks
category: making-software
type: guide
audience:
  primary: developers, web-ops
  secondary: tech-archs, designers
phases:
  - alpha
  - beta
  - live
breadcrumbs:
  -
    title: Home
    url: /service-manual
  -
    title: Making software
    url: /service-manual/making-software
exclude_from_search: true
---

Most software projects, especially those using open source technologies, will
use third party code. This approach brings with it many advantages and some
challenges.

## Not reinventing the wheel

The amount of software available to reuse in your project is huge. This may be open
source libraries or frameworks, or it may be commercial software appliances or
applications, or it may be code already written by your organisation. Ultimately very
very few software development projects will write absolutely everything from scratch.

Where possible and sensible, aiming to reuse existing code makes sense and will
reduce the time taken to develop the software you are tasked with building. Areas
that are general, and not specific to your domain, are good examples of where a
third party library is likely to exist. For instance for managing:

* time and dates
* money
* routing HTTP requests
* rendering HTML
* maths and statistics

## Strategies

There are a number of different strategies for reusing code, some of which are
described below. In many cases a single project may use more than one of these, but
in all cases it is worth doing so consciously.

### Vendor dependencies

One approach to managing dependencies is simply to copy the code into your project.
If you do take this approach be aware that not all open source licenses allow for this,
while others will force you to license your code in the same way as the dependency.

In general this approach is likely to be the last resort. Vendoring (as it is called)
breaks the link with the third party code, meaning that tracking upstream changes becomes
a manual (and likely slow, inconsistent and time consuming) job. Vendored code
tends to become out of date quickly, ideally this should be avoided for more
structured approaches.

### Trust but verify

A common approach, especially amongst the open source community, is to pull in
third party dependencies automatically using a dependency management tool
-- either at runtime, deploy time or compile time. The assurance that this third
party code is working is provided by extensive testing, generally of an automated
nature.

This should include testing of security characteristics. For example if a library
used to generate a web form introduces a SQL injection vulnerability, then your acceptance
tests should fail. It is important under this model that a suitable level of verification
is done, rather than just trusting any third party dependency completely. It's also important
to have trust in the specific code and version you are using, rather than just the library
or framework in general. What is suitable will vary between services depending on the risk
profile.

### Fork dependencies

Sometimes simply trusting a third party code repository does not provide the level
of assurance required. In these cases, one approach is to fork, or copy, the third
party code into a separate repository managed by you. This decouples updates
to the upstream code from updates to your organisations copy of that code, but importantly
maintains the link with the original. Your fork of the code can be updated
periodically or based on required new functionality. The process for updating
your fork may involve code reviews, virus scanning, checks against known vulnerabilities,
etc. It would also be possible to modify the code at this point for your own internal
use.

The downside of this approach compared with the trust but verify approach above is
the time taken to maintain the fork and to pull in changes from upstream. Not
keeping up to date with upstream at all increases these costs and has the same issues
as vendoring described above.

## Different language communities are different

It is worth noting that different programming language communities tend to view
dependencies differently. It is very common for instance within the Ruby or JavaScript
communities to find tiny libraries that do one trivial thing well. This tends to lead
to projects with hundreds of dependencies very quickly. For instance the popular
[Ruby on Rails](http://rubyonrails.org/) framework has 31 external dependencies at
the time of writing.

Within other communities it is less common to have so many independent dependencies,
but it is worth noting that this tends to mean either a more comprehensive
standard library (eg Python) or larger dependencies (eg Java).

Ultimately the problem of dependency management exists within all programming
communities, but the differences in tooling and approach are worth noting.

## Dependency management tools

Most programming languages will have a small number of popular dependency management
tools. The following list has a number of examples from a variety of languages.

* [Maven](http://maven.apache.org/) (Java)
* [Bundler](http://bundler.io/) (Ruby)
* [Pip](https://pypi.python.org/pypi/pip) (Python)
* [sbt](http://www.scala-sbt.org/) (Scala)
* [Leiningen](http://leiningen.org/) (Clojure)
* [Composer](https://getcomposer.org/) (PHP)
* [Bower](http://bower.io/) (JavaScript, CSS)

## Managing risks in third party code

By including third party dependencies in your application you are potentially including
code with known or unknown vulnerabilities. Accepting this risk is important in
order to mitigate it. In some cases code reviews of third party code may be required,
in others good testing practices can be enough.

### Tooling

Commercial tools do exist for some languages which aim to alert you to vulnerabilities
in third party dependencies and can help with building trust in code.

* [Gemnasium](https://gemnasium.com/)
* [Sonatype CLM](http://www.sonatype.com/clm/overview)

These are often based on publicly available vulnerability disclosure lists so it is
also possible to track those individually.

### Mailing lists

Most operating systems, vendors and some open source projects  will also have mailing lists
to advise about security vulnerabilities in third party code. You should always be
monitoring these lists where they exist for software you use. Some common examples include:

* [RedHat](https://access.redhat.com/security/updates/advisory/)
* [Ubuntu](https://lists.ubuntu.com/mailman/listinfo/ubuntu-security-announce)
* [Microsoft](https://technet.microsoft.com/en-us/security/dd252948.aspx)
* [CERT-UK](https://www.cert.gov.uk/register-for-alerts/)

The above is a small example of what is available. Having these emails come into an
internal mailing list is also advisable, rather than rely on a single individual to
alert everyone else.

## Further reading

* [Configuration management](/service-manual/making-software/configuration-management)
