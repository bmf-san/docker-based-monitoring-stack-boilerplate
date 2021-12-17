# docker-based-monitoring-stack-boilerplate
This is a boilerplate for docker based monitoring stack.

# Monitoring stacks
- Elasticsearch
- Fluentd
- Kibana
- Grafana
- Prometheus
- node-exporter
- cadvisor

# What you can do with this boilerplate?
- Application log monitoring
  - By using Elasticsearch, Fluentd, Kibana
- Container metrics monitoring
  - By using Prometheus, cadvisor, Grafana
- System metrics monitoring
  - By using Prometheus, node-exporter, Grafana

# Get started
Copy a `.env.example` as `.env`

Build containers.
`docker compose build`

Then, run containers.
`docker compose up`

# Appendix
[gobel-example]](https://github.com/bmf-san/gobel-example)