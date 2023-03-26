# NI strees utility for servers

Use to periodically strees CPU and to load memory and network bandwidth.

## Usage


```shell
./NI -cpu 2h5m30s -cpu-d 15m -cpu-p 50 -mem 2.5 -net 4h
```

```
  -cpu duration
        Interval of CPU streess (enables CPU stress)
  -cpu-d duration
        Min. duration for each CPU stress (default 2s)
  -cpu-m float
        Max limit of system's total CPU load perceent (default 100)
  -cpu-n int
        Number of CPU cores to stress (default AllCores)
  -cpu-p float
        Each CPU's load percentage (default 100)
  -mem float
        GiB of memory to use
  -net duration
        Interval for network speed test
  -net-c int
        Set concurrent connections for network speed test (default 10)
```
