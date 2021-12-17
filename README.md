# docker-based-monitoring-stack-boilerplate
This is a boilerplate for docker based monitoring stack.

<img src="https://user-images.githubusercontent.com/13291041/146595299-9e21d2b8-3b3d-4931-a739-ea3a8e69fa13.png" alt="drawing" width="400"/>

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
- OS metrics monitoring
  - By using Prometheus, node-exporter, Grafana

# Get started
Copy a `.env.example` as `.env`

Build containers.
`docker compose build`

Then, run containers.
`docker compose up`

|   Application   |             URL             |
| --------------- | --------------------------- |
| app             | http://localhost:8080/      |
| prometheus      | http://localhost:9090/graph |
| node-exporter   | http://localhost:9100/      |
| mysqld-exporter | http://localhost:9104/      |
| grafana         | http://localhost:3000/      |
| kibana          | http://0.0.0.0:5601/        |

![スクリーンショット 2021-12-18 4 04 19](https://user-images.githubusercontent.com/13291041/146595314-d593b9d8-9faa-4275-8c94-6c12cfcbfe36.png)

# Appendix
[gobel-example](https://github.com/bmf-san/gobel-example)
