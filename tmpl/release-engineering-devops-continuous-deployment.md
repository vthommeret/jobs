# Release Engineering - DevOps - Continuous Delivery

{{ template "about" }}

## The Job

* This job is in the engineering group.
* This is an individual contributor role (not people management)
* Located in Los Angeles, although exceptional candidates a few time
  zones away will also be considered.
* Suitable for mid- and senior- level engineers

## Problems

“Continuous Delivery” isn’t a trendy buzzword for us -- it’s a competitive
advantage.   We strive to have a "culture of shipping" where developers are
responsible for getting their code to production and then responsible for how
it performs.  However, to do that, they need great tools.  Our current system
allows us to deploy the entire application stack through a web interface in
under 30 seconds.  That's "ok", but as our system grows and improves, so must our tooling.
Here are some problems we are investigating and need help with: 

* Reliability - there are a number of external dependencies in the deploy process.
How can we minimize impact if they fail (or not use them at all)?
* Delivery - moving artifacts from point A to point B.   DSH? Or configuration
management tools (e.g. Chef, Puppet, et al), or special purpose tools such as
Fabric/Invoke/Capistrano?
* Containers: How does switching to a container-based architecture (e.g. Docker)
change our deployment strategy?
* Pros and cons of “in-place updates” versus “build from scratch and swap”
* Integrity - how can we monitor “what we say we deploy actually is what got
deployed”
* Accuracy - Our system is somewhat coarse. On deploy, everything gets updated,
which has much to be desired.  We would like to have the deploy process be more fine
grained.
* Monitorability - Once code is running, how can we make sure it is monitored
but also that the right people are notified when it is not.  How can we make
tools to better empower developers to understand how their code runs in
production.

We'd love hear your stories from the battlefield.  And better yet, working on
our battlefield!

## Solutions

* Strong linux and bash fundamentals.  Even in 2016, bash is essential
  glue.
* Experience with multiple scripted (php, ruby, python) or compiled languages (java, et al).  The more the merrier.  Our current tools are mostly written in Go ("golang").  However we gladly
  welcome people with different backgrounds.  Most of the team was not familiar
  with Go when they started but were productive very quickly.
* Built or a major contributor to in-house or open-source deploy tools or
  CI/CD pipelines (e.g. Jenkins, CircleCI, Travis, etc)
* Or an engineer who has used CI/CD and wants them to be better
* We are slowly migrating to docker-based deploys. If you have docker
  experience already, great!  If not, this is a great opportunity to do so!

{{ template "apply" }}
