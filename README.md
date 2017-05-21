# docker-swarm-probe
Test project to see what we can unearth from Docker regarding Docker Swarm services.

## Blue-Green Deployment Detection

### Assumptions/Requirements Viktor Farcic

* Detect the color of the current release color (e.g. blue)
* Deploy a new release a the other color (e.g. green)
* Publish the new release under a "special" address (e.g. through DFP) so that it can be tested.
* Replace the new release in the proxy config.
* (Optionally) revert to the previous color.

### Assumptions/Requirements Joost van der Griendt

* Deployment tool used for a release should be the same
* Should detect version of deployment tool (for naming schemes)
* Must use naming scheme for blue-green
* Either Blue or Green must be the default
* A change from Blue to Green or vice versa must not impact services using Docker Service DNS for accessing services
* Should be able to deploy a whole stack
* Must expose events log of actions taken
