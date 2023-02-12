# Go-LB

This is a toy loadbalancer that I built to better understand how Loadbalancers and Service Discovery works.

### Features

- RoundRobin Load Distribution
- Registration and Deregistration of backends through Service Discovery
- Cancellation propogation and timeouts(Configurable duration)
- Heartbeat and auto-deregistration during backend disconnection
- Supports all HTTP methods
- Handles requests concurrently from single and multiple clients

TODO:

- [ ] URL pattern matched balancing
- [ ] Configurable strategy
- [ ] Hot reload config changes
- [ ] Export metrics about healthy servers
- [ ] CLI to add and remove backends and change strategy in realtime
