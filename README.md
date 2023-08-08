# gforwarder
A high-performance network forwarder developed in go language.


gforwarder is a high-performance network forwarder developed in go language.

Support tcp protocol.

Adopt the csp concurrency model, the net library in the go language standard library.

The coroutine pool technology is adopted to greatly reduce the number of coroutines and memory usage.


Supported operating systems: windows, linux, macos


## Performance:

qps: more than 40000/s

(Single connection data volume 4KiB, Kbytes)

Throughput: more than 1000MiB/s

(Mbytes per second) (Single connection data volume 1GiB, Gbytes)


### test environment:

ubuntu 22.04.2

Intel i5 1135g7 2.4GHz

4 cores 4 threads without hyperthreading


## New features that will be added:

support udp protocol

Support load balancing

access control (ip)

network speed limit
