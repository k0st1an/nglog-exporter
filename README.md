# nglog-exporter

```
curl -s http://localhost:9999/metrics | grep nglog | grep -v ^#
nglog_http_requests_total{method="http"} 601040
nglog_parse_errors_total 0
nglog_request_time_hist_bucket{le="0.005"} 599714
nglog_request_time_hist_bucket{le="0.01"} 600239
nglog_request_time_hist_bucket{le="0.025"} 600414
nglog_request_time_hist_bucket{le="0.05"} 601034
nglog_request_time_hist_bucket{le="0.1"} 601040
nglog_request_time_hist_bucket{le="0.25"} 601040
nglog_request_time_hist_bucket{le="0.5"} 601040
nglog_request_time_hist_bucket{le="1"} 601040
nglog_request_time_hist_bucket{le="2.5"} 601040
nglog_request_time_hist_bucket{le="5"} 601040
nglog_request_time_hist_bucket{le="10"} 601040
nglog_request_time_hist_bucket{le="+Inf"} 601040
nglog_request_time_hist_sum 478.1599999970288
nglog_request_time_hist_count 601040
nglog_status_total{code="200"} 6319
nglog_status_total{code="204"} 594719
nglog_upstream_connect_time_hist_bucket{le="0.005"} 601038
nglog_upstream_connect_time_hist_bucket{le="0.01"} 601038
nglog_upstream_connect_time_hist_bucket{le="0.025"} 601038
nglog_upstream_connect_time_hist_bucket{le="0.05"} 601038
nglog_upstream_connect_time_hist_bucket{le="0.1"} 601038
nglog_upstream_connect_time_hist_bucket{le="0.25"} 601038
nglog_upstream_connect_time_hist_bucket{le="0.5"} 601038
nglog_upstream_connect_time_hist_bucket{le="1"} 601038
nglog_upstream_connect_time_hist_bucket{le="2.5"} 601038
nglog_upstream_connect_time_hist_bucket{le="5"} 601038
nglog_upstream_connect_time_hist_bucket{le="10"} 601038
nglog_upstream_connect_time_hist_bucket{le="+Inf"} 601038
nglog_upstream_connect_time_hist_sum 1.460000000000001
nglog_upstream_connect_time_hist_count 601038
nglog_upstream_status_total{code="200"} 6319
nglog_upstream_status_total{code="204"} 594719
```
