# nglog-exporter

```
log_format nglog_json '{"scheme":"$scheme","status":"$status","request_time":"$request_time","upstream_status":"$upstream_status","upstream_connect_time":"$upstream_connect_time","upstream_response_time":"$upstream_response_time"}';

access_log syslog:server=127.0.0.1:8888,nohostname,tag= nglog_json;
```

```
$ curl -s http://localhost:9999/metrics | grep ^nglog
nglog_http_requests_total{method="http"} 4.121157e+06
nglog_parse_errors_total 0
nglog_request_time_hist_bucket{le="0.005"} 4.109986e+06
nglog_request_time_hist_bucket{le="0.01"} 4.114444e+06
nglog_request_time_hist_bucket{le="0.025"} 4.116831e+06
nglog_request_time_hist_bucket{le="0.05"} 4.121092e+06
nglog_request_time_hist_bucket{le="0.1"} 4.121157e+06
nglog_request_time_hist_bucket{le="0.25"} 4.121157e+06
nglog_request_time_hist_bucket{le="0.5"} 4.121157e+06
nglog_request_time_hist_bucket{le="1"} 4.121157e+06
nglog_request_time_hist_bucket{le="2.5"} 4.121157e+06
nglog_request_time_hist_bucket{le="5"} 4.121157e+06
nglog_request_time_hist_bucket{le="10"} 4.121157e+06
nglog_request_time_hist_bucket{le="+Inf"} 4.121157e+06
nglog_request_time_hist_sum 4347.117000220557
nglog_request_time_hist_count 4.121157e+06
nglog_status_total{code="200",host="SOMEDOMAIN"} 18751
nglog_status_total{code="200",host="SOMEDOMAIN"} 271
nglog_status_total{code="204",host="SOMEDOMAIN"} 3.855458e+06
nglog_status_total{code="204",host="SOMEDOMAIN"} 237142
nglog_status_total{code="204",host="SOMEDOMAIN"} 9534
nglog_status_total{code="404",host="SOMEDOMAIN"} 1
nglog_upstream_connect_time_hist_bucket{le="0.005"} 4.121151e+06
nglog_upstream_connect_time_hist_bucket{le="0.01"} 4.121155e+06
nglog_upstream_connect_time_hist_bucket{le="0.025"} 4.121155e+06
nglog_upstream_connect_time_hist_bucket{le="0.05"} 4.121155e+06
nglog_upstream_connect_time_hist_bucket{le="0.1"} 4.121155e+06
nglog_upstream_connect_time_hist_bucket{le="0.25"} 4.121155e+06
nglog_upstream_connect_time_hist_bucket{le="0.5"} 4.121155e+06
nglog_upstream_connect_time_hist_bucket{le="1"} 4.121155e+06
nglog_upstream_connect_time_hist_bucket{le="2.5"} 4.121155e+06
nglog_upstream_connect_time_hist_bucket{le="5"} 4.121155e+06
nglog_upstream_connect_time_hist_bucket{le="10"} 4.121155e+06
nglog_upstream_connect_time_hist_bucket{le="+Inf"} 4.121155e+06
nglog_upstream_connect_time_hist_sum 12.491999999999068
nglog_upstream_connect_time_hist_count 4.121155e+06
nglog_upstream_header_time_hist_bucket{le="0.005"} 4.117173e+06
nglog_upstream_header_time_hist_bucket{le="0.01"} 4.120632e+06
nglog_upstream_header_time_hist_bucket{le="0.025"} 4.121157e+06
nglog_upstream_header_time_hist_bucket{le="0.05"} 4.121157e+06
nglog_upstream_header_time_hist_bucket{le="0.1"} 4.121157e+06
nglog_upstream_header_time_hist_bucket{le="0.25"} 4.121157e+06
nglog_upstream_header_time_hist_bucket{le="0.5"} 4.121157e+06
nglog_upstream_header_time_hist_bucket{le="1"} 4.121157e+06
nglog_upstream_header_time_hist_bucket{le="2.5"} 4.121157e+06
nglog_upstream_header_time_hist_bucket{le="5"} 4.121157e+06
nglog_upstream_header_time_hist_bucket{le="10"} 4.121157e+06
nglog_upstream_header_time_hist_bucket{le="+Inf"} 4.121157e+06
nglog_upstream_header_time_hist_sum 3263.7399999507916
nglog_upstream_header_time_hist_count 4.121157e+06
nglog_upstream_response_time_hist_bucket{le="0.005"} 4.117166e+06
nglog_upstream_response_time_hist_bucket{le="0.01"} 4.120632e+06
nglog_upstream_response_time_hist_bucket{le="0.025"} 4.121157e+06
nglog_upstream_response_time_hist_bucket{le="0.05"} 4.121157e+06
nglog_upstream_response_time_hist_bucket{le="0.1"} 4.121157e+06
nglog_upstream_response_time_hist_bucket{le="0.25"} 4.121157e+06
nglog_upstream_response_time_hist_bucket{le="0.5"} 4.121157e+06
nglog_upstream_response_time_hist_bucket{le="1"} 4.121157e+06
nglog_upstream_response_time_hist_bucket{le="2.5"} 4.121157e+06
nglog_upstream_response_time_hist_bucket{le="5"} 4.121157e+06
nglog_upstream_response_time_hist_bucket{le="10"} 4.121157e+06
nglog_upstream_response_time_hist_bucket{le="+Inf"} 4.121157e+06
nglog_upstream_response_time_hist_sum 3263.7959999507907
nglog_upstream_response_time_hist_count 4.121157e+06
nglog_upstream_status_total{code="200",host="SOMEDOMAIN"} 18751
nglog_upstream_status_total{code="200",host="SOMEDOMAIN"} 271
nglog_upstream_status_total{code="204",host="SOMEDOMAIN"} 3.855458e+06
nglog_upstream_status_total{code="204",host="SOMEDOMAIN"} 237142
nglog_upstream_status_total{code="204",host="SOMEDOMAIN"} 9534
nglog_upstream_status_total{code="404",host="SOMEDOMAIN"} 1
```
