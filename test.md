# Release Engineering - DevOps - Continuous deployment



“Continuous Delivery” isn’t a trendy buzzword for us -- it’s a competitive
advantage.   We strive to have a "culture of shipping" where developers are
responsible for getting their code to production and then responsible for how
it performs.  However, to do that, they need great tools.  Our current system
allows us to deploy the entire application stack through a web interface in
under 30 seconds.  That's "ok", but as our system grows and improves, so must our tooling.
Here are some issues you might be looking at:

* Reliability - there are number of external dependencies in the deploy process.
How can we minimize impact if they fail (or not use them at all)
* Delivery - moving artifacts from point A to point B.   DSH? Or configuration
management tools (e.g. Chef, Puppet, et al), or special purpose tools such as
Fabric/Invoke?
* Containers: How does switching to a container-based architecture (e.g. Docker)
change our deployment strategy?
* Pros and cons of “in-place updates” versus “build from scratch and swap”
* Integrity - how can we monitor “what we say we deploy, actually is what got
deployed”
* Accuracy - Our system is somewhat coarse.. On deploy everything gets updated
which has much to be desired.  We like to have the deploy process be more fine
grain
* Monitorability - Once code is running, how can we make sure it is monitored
but also that the right people are notified when it is not.  How can we make
tools to better empower developers to understand how their code runs in
production.

We’d love hear your stories from the battlefield.

