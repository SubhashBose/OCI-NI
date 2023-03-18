# NeverIdle

## Usage


```shell
./NeverIdle -c 2h5m30s -d 15m -p 50 -m 2 -n 4h
```

```
  -c duration
        Interval for CPU load
  -d duration
        Min duration for each CPU load (default 2s)
  -p float
        CPU load percentage (default 100)
  -ncpu int
        Number of CPU cores to load (default 2)
  -m float
        GiB of memory to use
  -n duration
        Interval for network speed test
  -t int
        Set concurrent connections for network speed test (default 10)
```
