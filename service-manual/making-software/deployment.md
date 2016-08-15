---
layout: detailed-guidance
title: Deploying software
subtitle: Principles to ensure frequent, low-risk deployments
category: making-software
type: guide
audience:
  primary: tech-archs, web-ops
  secondary: developers, performance-analysts, qa
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
# Principles for software deployment

We've identified some common principles for software deployment
which we've applied in a number of different projects, with
different technology stacks and needs. These principles underpin a
software deployment process which meets user needs. Those principles
are:

 * little and often
 * quality software
 * optimise for cycle time
 * repeatable, auditable deployments
 * zero downtime deployments

## Little and often ##

Deploying software should be a low-risk activity. By deploying
software frequently and in small increments, the risk is reduced in a
number of ways. See
[Regular Releases Reduce Risk](https://gds.blog.gov.uk/2012/11/02/regular-releases-reduce-risk/)
from the GDS blog for more on this.

Deploying software frequently makes life better for the product
managers in your organisation. Frequent deployments allow the product
managers to get things right in a timely fashion: both fixing bugs and
releasing new features. 

Roo Reynolds, GOV.UK mainstream product
manager, said that "Deploying once a week would be frighteningly
slow." 

The GOV.UK site design has changed radically 4 times since
its public release in October 2012. This was enabled in part by
frequent releases enabling rapid gathering of feedback and responding
to change.

## Quality software ##

The software that you deploy to production should be of a consistently
high quality. The user impact of bugs is obvious; less obvious is that
the earlier you identify bugs, the easier and cheaper they are to fix.

The deployment itself should not be a risky process. By the time a
version of the software is deployed to production, you should have
confidence that it will work smoothly and seamlessly.

## Optimise for cycle time ##

How long does it take from a developer making a code change to that change
hitting production? The shorter this time is, the faster a product can
respond to change. The quicker you can release the next iteration, the
faster you will converge on an ideal solution.

## Repeatable, auditable deployments ##

You should know at any moment what version of your service is running
in each environment. When a deployment hits production, you should be
able to trace the changes that it introduced all the way back to the
initial code commits in the source code repositories which went into
that deploy.

Combined with small, frequent releases, if any problem does hit
production, you will be able to immediately narrow down the cause to a
small number of commits. 

Rolling back to a previous version is less
onerous as less of the system has changed. And "rolling forward" --
with a code change to fix the production issue -- is achievable
because the deployment process is automated and the lead time is
short.

An additional benefit of having a repeatable deployment process is
that scaling and recovering from failure become easy. Suppose you want
to add more application servers to host a particular application,
either to respond to higher demand, or to replace failed
instances. 

Once you have provisioned the required machines, you can
just re-run your deployment process on the new machines to deploy the
software. Without a repeatable deployment process, adding machines
becomes manual and error-prone.

## Zero downtime deployments

Many deployment processes incur a downtime cost. The more frequently
you deploy, the more downtime you will experience as a result of
deployment. This may be acceptable depending on the needs of your particular project,  or you may need to consider how your deployment process needs to change to achieve zero downtime.

This isn't a one-dimensional problem. Achieving zero downtime for read
operations is easier than doing so for write operations.

Whether or not your project has a business need for zero downtime
deployments, it's worth considering the tools and processes which
make it possible, as the constraints of zero downtime deployments can
result in better engineering practices generally.

# Techniques and tools to achieve these ideals #

## Single artefact ##

An antipattern in deployment processes is building a different
application artefact for each environment. Examples of this might
include controlling the presence of debugging symbols in the binary or
variations in the use of optimisation flags. The problem is that the
testing you do of code artefacts in preview environments may not be
applicable to the artefact you deploy to production.

The better alternative is to build a single artefact which gets
deployed to all environments. With the same code running in each
environment, you can deploy to production safe in the knowledge that
this code has been tested in every other environment and has not been
found wanting.

Note that the exact nature of an artefact is intentionally vague. It
may be

* a .jar file for JVM languages
* for languages without compilation artefacts it may even be a tag in the source control
system
* an entire virtual machine image with the application pre-deployed.

## Ordered environments ##

You should have multiple environments to deploy to. At the very least,
you will have a development environment running the latest version of
the software, and a production environment being used by live
users. You may also have other environments dedicated to exploratory
testing, user testing, performance testing or a staging environment
prior to production.

The environments should be ordered so that a version of the
software cannot be deployed to a later environment before it's been
deployed and tested in an earlier stage. That way, the software cannot
be deployed to production without having been tested in every previous
environment first.

This does not need to be a strict linear ordering. Some sets of tests
may be run in parallel -- such as user acceptance testing and
performance testing. However there is very often one single production
environment which is later than all others, and one single entry point
which precedes all others.

## Repeatable deployments of infrastructure configuration ##

One of the principles of good deployment is repeatable
deployments. This does not just apply to application
code. Applications don't run in a vacuum, and often have particular
requirements of the underlying system in which they run. The
configuration of that system should be automated and repeatable.

There are two main issues: ensuring that new builds of machines are
repeatable, and ensuring that once built, machines do not suffer from
*configuration drift*, in which small manual configuration changes are
made over time, resulting in a system which is not in a reproducible
state.

Scripting the configuration of a new machine is not a difficult
process. It will always start from a known state and can have a
number of tasks to install packages, put configuration files in place
and start services until the machine is in a good state.

Managing configuration drift is more tricky, as to counteract
manual changes to configuration, your system must be robust enough to
take a machine from an unknown state to a known state. There are a
number of tools for managing your infrastructure configuration,
such as [CFEngine][], [Chef][], and [Puppet][]. Each of these is
designed such that they can be run repeatedly on a machine to
alleviate configuration drift.

### Deploying configuration management code ###

If you're using one of these tools, you need to provide a means to
deploy new versions of the infrastructure code. There are two means of
distributing infrastructure code:

**Using a server** (Chef Server, Puppetmaster) In this kind of system,
  you'll have a central server that distributes configuration code
  to each machine in your environment. Deploying new versions of code
  requires only deploying to the server, that will then distribute it
  to the clients.

**Serverless** (Chef Solo, Masterless Puppet) Here, you'll need to
  distribute the configuration code to each individual node and ensure
  that each node runs the code.

### Avoiding configuration management code ###

An alternative strategy to avoid configuration drift is to use the
[immutable server][] pattern, in which once a machine is configured it's
never touched again. In order to deploy a new version of the
software, an entirely new machine is provisioned and the old one
discarded. Configuration drift is avoided because servers have short
lifespans and are frequently replaced by new instances. This is a
natural fit in virtualised environments and where the application
artefact is a virtual machine template with the app pre-deployed, but
can also be achieved using containerisation technology such as [lxc][].

## Repeatable deployments of code ###

There are a number of options for deploying your code:

 * construct your artefacts as operating system packages (.debs or
   .rpms) and install using your infrastructure configuration
   management tool from a local package repository (apt or yum)
 * use a push-based system to deploy such as [fabric][],
   [capistrano][], or similar
 * create a new [immutable server][] for each deployment

You should think about how you'll discover hosts that you deploy
to. In a simple scenario, your deployment script may have a hard-coded
list of application servers that it deploys to. 

In this situation, there's a risk that the hard-coded list of servers drifts to differ
from the number of servers which actually exist in reality. This
risk grows more likely with larger and more dynamic
infrastructures. 

There are more involved host discovery mechanisms,
such as internal DNS, Zookeeper, or using a message-queue based system
such as MCollective.

## Management of environment-variable configuration ##

Since you should be deploying the same artefact to each environment,
both for infrastructure configuration management code and for
application code, you'll inevitably find a need to inject
configuration which varies between environments, such as URLs of
dependent services.

For application configuration, your deployment mechanism should
provide a way of injecting environment-specific configuration files
into each environment.

For infrastructure configuration, your infrastructure tool should
provide some means of achieving this. For example, Puppet 3 provides
[Hiera][], a hierarchical datastore for managing these values.

### Secrets ###

Extra care must be taken when managing secrets such as database
passwords or SSL keys. You want to ensure:

 * at a coarse-grained level, secrets cannot be accessed outside of
   the environment which uses them
 * at a fine-grained level, secrets are known only by those machines
   in an environment which need to know them.

For example, in a three-tier app with database, application and web
servers, the database server does not need to know the  SSL (secure sockets layer) private
keys for the site, nor does the web server need to know the database credentials.

If you are using hiera, then [hiera-gpg][] provides a solution to this
problem. It allows the injection of values from [GPG][]-encrypted
files. Only those with an appropriate private key can access the
contents. By creating a GPG key for each host in an environment, you
can decide on a host-by-host basis which host can access which sets of
secrets.

If you are using chef, then [chef data bags][] provide a similar solution.

## Zero downtime deployments

In projects which have high availability requirements, the process of
deploying small code changes to production frequently may incur an
unacceptable loss of service, if each deployment results in a short
period of downtime. Therefore, it's important to consider what
engineering is necessary to enable deployments which do not result in
any downtime at all.

This is all subject to what your definition of downtime is.
Maintaining uptime for read-based operations is relatively simple: a
caching layer which can serve from stale can hide the absence of
application servers; a database is easy to migrate from one master to
another if it is placed into read-only mode first.

Maintaining uptime for write-based operations is trickier, and
requires up-front thought and design. If you know that you'll have
high uptime requirements for write-based or transactional operations,
you'll need to consider how that will affect your architecture and
infrastructure.

### Database migrations ###

As applications evolve over time, so do the requirements that they
place on their databases. Database migration scripts are short pieces
of database code which transform the database in some way for the
benefit of the application.

To achieve zero downtime deployments, you should decouple
application deployments from database migrations. If you're
performing zero downtime deployments, you'll necessarily end up with
multiple different versions of the application running
concurrently. Conversely, the application will need to be tolerant to
the eventuality of a database migration script running concurrently
within the application lifetime.

Note that database migrations should be subject to the same rigorous
deployment pipeline as application code. They should be deployed to
testing environments first, and only go to production once they have
been applied and verified against all other environments.

### Service dependencies ###

Services which depend on one another via an application programming interface (API) can experience similar
deployment problems as applications which depend on a database. For
example, a frontend application which communicates with a backend
application over an API of some sort. 

Once again, the answer is to decouple deployments of the applications to make sure that the frontend
application is tolerant to additions to the backend API, and that
similarly the backend API can add functionality without disrupting the frontend application's operation.

### Making writes asynchronous ###

Another method of avoiding failures during deployments is to make
write operations asynchronous by posting them to a message queue. That
way, when the backend system which consumes from the queue is disabled
during a deployment, the frontend does not start seeing errors;
rather, it just sees an increase in the time taken to see a write
reflected in further read operations.

## Smoke tests ##

Once you have deployed your application, you should determine whether
the application is working as expected. If it's not working, the
deployment can be cancelled or rolled back. The test used to determine
this is often referred to as a "smoke test".

A good smoke test is simple and fast, and exercises not just the
application but also all of its essential dependencies. For example,
if an application needs a database to be present to operate
effectively, the smoke test should exercise an application code path
which will fail if the database is not present or returns an error.

If and when the smoke test fails, you should know what your response
will be. The simplest option is to manually roll back to a previous
version of the application -- which should be easy enough if you have
a versioned artefact repository to draw the application from.

A solution with more sophistication may automatically detect the smoke
test failure and cancel the deployment or roll back to the previous
version.

An ideal solution would not even add the new version of the
application to the production [load balancer](https://en.wikipedia.org/wiki/Load_balancer) until it has been smoke
tested and verified good. If the application fails the smoke test, it
is simply discarded; no rollback is necessary, and no interruption in
service happens. This works particularly well with the immutable
server pattern.

## Emergency deployments ##

From time to time, there may come a situation where you wish to deploy
to production right now. This may be due to a published security
vulnerability in a library you are using, or because a bug has hit
production which has broken the system for a number of users.

It may be the case that you subvert your usual deployment pipeline
to fix things, then back-port the change you made in production
(or "hotfix") to your development environment and push it through the
normal deployment process once the crisis is over. Should this be the
case, then your cycle time is too long. In the ensuing post-mortem
analysis of what went wrong, you should ask questions about why the
deployment pipeline was not streamlined enough to accommodate a rapid
deployment of a fix to production.

[CFEngine]: http://cfengine.com/
[Chef]: https://www.chef.io/chef/
[Puppet]: https://puppetlabs.com/
[immutable server]: http://martinfowler.com/bliki/ImmutableServer.html
[lxc]: https://linuxcontainers.org/
[fabric]: http://docs.fabfile.org/en/1.6/
[capistrano]: https://github.com/capistrano/capistrano
[Hiera]: http://projects.puppetlabs.com/projects/hiera
[hiera-gpg]: https://github.com/crayfishx/hiera-gpg
[GPG]: https://www.gnupg.org/
[chef data bags]: https://docs.chef.io/data_bags.html

*[API]: application programming interface 
