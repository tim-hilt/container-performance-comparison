# Container Performance Testing

## Objective

I wanted to compare the impact of containerizing an application to different levels, namely

- dockerized
- Inside of a Kubernetes cluster

to the native performance.

This repository implements a simple endpoint returning a string using the go standard-library and the gin http-framework.

## Approach

- Implement endpoint that returns a plain string
- Start application with different methods
- use `wrk -t12 -d30 http://localhost:8080` and record results

## Results

| Containerization Level | Implementation | Requests/s       | Latency             | Loss wrt native |
| ---------------------- | -------------- | ---------------- | ------------------- | --------------- |
| Native                 | http-pkg       | 8.78k +/- 4.38k  | 5.47ms +/- 7.39ms   | 00.00%          |
| Docker                 | http-pkg       | 3.00k +/- 397.64 | 11.16ms +/- 4.78ms  | 65.83%          |
| Kubernetes             | http-pkg       | 1.44k +/- 253.03 | 24.50ms +/- 18.33ms | 83.59%          |
| Native                 | go-gin         | 9.01k +/- 4.62k  | 5.65ms +/- 7.64ms   | 00.00%          |
| Docker                 | go-gin         | 2.95k +/- 367.79 | 11.37ms +/- 4.94ms  | 66.40%          |
| Kubernetes             | go-gin         | 1.47k +/- 242.75 | 23.76ms +/- 16.65ms | 83.68%          |

## Conclusion

Containerization reduces the performance of http-traffic significantly. I can't tell why that is the case though. I heard, that Docker applies a lot of security processing on top of just forwarding the requests. If this is switched off or tuned, then performance might be better.

## Future work

- Deploy app in the cloud and test again (native vs. docker vs. K8s)
- Add more real-world processing on backend (I/O, calculations, returning json instead of string literal)
- Add nginx-ingress as a fourth level

## Appendix - Full wrk Output

### native - http-pkg

```
Running 30s test @ http://localhost:8080
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.47ms    7.39ms 126.27ms   88.70%
    Req/Sec     8.78k     4.38k   63.92k    82.47%
  3142392 requests in 30.09s, 392.58MB read
  Socket errors: connect 0, read 376, write 0, timeout 0
Requests/sec: 104418.14
Transfer/sec:     13.05MB
```

### native - go-gin

```
Running 30s test @ http://localhost:8080
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.65ms    7.64ms 107.80ms   87.77%
    Req/Sec     9.01k     4.62k   93.69k    82.20%
  3222179 requests in 30.08s, 402.55MB read
  Socket errors: connect 0, read 378, write 0, timeout 0
Requests/sec: 107115.65
Transfer/sec:     13.38MB
```

### Docker - http-pkg

```
Running 30s test @ http://localhost:8080
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    11.16ms    4.78ms  95.97ms   78.89%
    Req/Sec     3.00k   397.64     4.15k    71.00%
  1076704 requests in 30.04s, 134.51MB read
  Socket errors: connect 0, read 261, write 0, timeout 0
Requests/sec:  35839.72
Transfer/sec:      4.48MB
```

### Docker - go-gin

```
Running 30s test @ http://localhost:8080
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    11.37ms    4.94ms  98.28ms   77.62%
    Req/Sec     2.95k   369.79     4.51k    71.00%
  1057120 requests in 30.04s, 132.07MB read
  Socket errors: connect 0, read 255, write 0, timeout 0
Requests/sec:  35188.61
Transfer/sec:      4.40MB
```

### Kubernetes - http-pkg

```
Running 30s test @ http://localhost:8080
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    24.50ms   18.33ms 333.77ms   96.43%
    Req/Sec     1.44k   253.03     2.00k    69.28%
  512572 requests in 30.07s, 64.04MB read
  Socket errors: connect 0, read 383, write 0, timeout 0
Requests/sec:  17044.39
Transfer/sec:      2.13MB
```

### Kubernetes - go-gin

```
Running 30s test @ http://localhost:8080
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    23.76ms   16.65ms 316.21ms   96.12%
    Req/Sec     1.47k   242.75     2.03k    70.18%
  524300 requests in 30.08s, 65.50MB read
  Socket errors: connect 0, read 384, write 0, timeout 0
Requests/sec:  17427.49
Transfer/sec:      2.18MB
```

### Native - actix

```
Running 30s test @ http://localhost:8080
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     9.25ms   22.14ms 291.41ms   91.73%
    Req/Sec    17.10k     6.02k  120.98k    85.85%
  6112489 requests in 30.09s, 512.98MB read
  Socket errors: connect 0, read 381, write 0, timeout 0
Requests/sec: 203133.20
Transfer/sec:     17.05MB
```

### Native - rocket

```
Running 30s test @ http://localhost:8000
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.23ms    6.41ms 254.09ms   96.84%
    Req/Sec    13.38k     1.45k   47.20k    87.34%
  4791067 requests in 30.10s, 1.11GB read
  Socket errors: connect 0, read 373, write 0, timeout 0
Requests/sec: 159194.28
Transfer/sec:     37.65MB
```

