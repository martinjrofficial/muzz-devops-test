# Devops Interview Task

This repository is used as part of the interview process for a DevOps Engineer at Muzz.

You'll need the latest [Go runtime installed](https://go.dev/dl/) to get started.




1. **Metric Types**: I used different metric types from the Prometheus client library based on their suitability for the data being tracked:
   - `Counter`: Used for `muzz_echo_requests_total` and `muzz_echo_requests_failed_total` to track cumulative counts.
   - `Histogram`: Used for `muzz_echo_request_duration_seconds` to capture request duration distributions and calculate quantiles.
   - `SummaryVec`: Used for `muzz_echo_request_duration_seconds` to track request latency percentiles with configurable quantile objectives.
   - `CounterVec`: Used for `muzz_echo_requests_total` to track request counts with additional dimensions (method and status).

2. **Latency Percentiles**: For the `muzz_echo_request_duration_seconds` summary metric, i've chosen to track the 50th, 90th, and 99th percentiles. These percentiles are commonly used for setting SLOs and provide a good overview of the latency distribution.

3. **Simulated Errors**: In the `server.go` file, ive added a simulated random error for demonstration purposes. This is not part of the actual service implementation but helps showcase the usage of the `muzz_echo_requests_failed_total` and `muzz_echo_requests_total` metrics with the "error" status label.

4. **Metric Naming**: I followed the Prometheus naming conventions for metrics, using a `<app_name>_<metric_name>_<metric_type>` format. The `muzz` prefix represents the application name, followed by the metric name and type.

5. **Metric Labeling**: I  added labels to the `muzz_echo_requests_total` metric to capture additional dimensions (method and status). This decision was made to provide more granular insights and enable future extensibility if more gRPC methods are added to the service.

6. **Metric Exposition**: I exposed the metrics over an HTTP endpoint (`/metrics`) using the Prometheus client library's `promhttp.Handler()`. This allows Prometheus to scrape the metrics from the gRPC server.






In this implementation, i've added several metrics to gain better visibility into the performance, errors, and latencies of the gRPC service. Here's a breakdown of the added metrics and the rationale behind them:

## Metrics

1. **`muzz_echo_requests_failed_total`**
   - Description: Total number of failed Echo requests.
   - Importance: Tracking failed requests is crucial for monitoring the service's reliability and identifying potential issues. This metric allows you to measure the error rate and set alerts if it exceeds a certain threshold.

2. **`muzz_echo_request_duration_seconds`**
   - Description: Summary of Echo request durations in seconds, with percentile buckets at 50th, 90th, and 99th percentiles.
   - Importance: Request latency is a key performance indicator for any service. By tracking latency percentiles, you can understand the distribution of request durations and set appropriate service-level objectives (SLOs) for different percentiles. This metric helps identify potential performance bottlenecks or regressions.

3. **`muzz_echo_requests_total`**
   - Description: Total number of Echo requests, with labels for the request method and response status (success or error).
   - Importance: This metric provides a more detailed view of request counts by separating successful and failed requests. Additionally, the "method" label allows for future extensibility, should more gRPC methods be added to the service. This level of granularity helps in analyzing traffic patterns and correlating request counts with other metrics.

## Rationale

The added metrics provide comprehensive visibility into the gRPC service's health, performance, and reliability. By tracking request counts, error rates, and latency distributions, you can:

1. **Monitor Service Reliability**: The `muzz_echo_requests_failed_total` metric allows you to track the error rate and set alerts if it exceeds an acceptable threshold, ensuring that the service is functioning as expected.

2. **Measure Performance**: The `muzz_echo_request_duration_seconds` metric, especially the percentile buckets, provides insights into the service's performance characteristics. You can set SLOs based on the desired latency percentiles and monitor for regressions or performance degradations.

3. **Analyze Traffic Patterns**: The `muzz_echo_requests_total` metric, with its dimensions for method and status, allows you to analyze traffic patterns and correlate request counts with other metrics. This can help identify potential bottlenecks or areas for optimization.

4. **Future Extensibility**: By including the "method" label in the `muzz_echo_requests_total` metric, the implementation is prepared for future additions of new gRPC methods. This ensures that the metrics remain relevant and valuable as the service evolves.

Overall, these metrics provide a solid foundation for monitoring, troubleshooting, and optimizing the gRPC service. They enable data-driven decision-making and help ensure that the service meets performance and reliability requirements.

Run make first:

make install_tools
make generate_protos


To run the server run:

```bash
go run server/main.go
```

To run the client run:

```bash
go run client/main.go
```


This  client will  send 100 concurrent Echo requests to the server. You should see the metrics being updated in the Grafana dashboard as the requests are processed.

To visualize the metrics, you can use Prometheus and Grafana. Here's an example of how to set up a Grafana dashboard to visualize the metrics:

Install and run Prometheus. You can use the official Prometheus Docker image:

A prometheus yml config provided in the repo


docker run -d --name=prometheus -p 9090:9090 -v <path-to-prometheus-config>:/etc/prometheus/prometheus.yml prom/prometheus

Install and run Grafana. You can use the official Grafana Docker image:

docker run -d --name=grafana -p 3000:3000 grafana/grafana

Access the Grafana UI at http://localhost:3000 (default username/password is admin/admin).

Add a new data source in Grafana and configure it to connect to your Prometheus instance (e.g., http://<prometheus-ip>:9090).

Create a new dashboard and add panels to visualize the metrics  defined  (muzz_echo_requests_total and muzz_echo_request_duration_seconds).

You should now be able to see the metrics being collected and visualized on the Grafana dashboard.