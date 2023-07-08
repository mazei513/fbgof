# High throughput Fizz Buzz

Based on the question in https://codegolf.stackexchange.com/questions/215216/high-throughput-fizz-buzz

## Details

Done in pure Go. The final version has concurrency alongside tricks such as unrolling loops and reusing buffers, and at commit 6c3e2b0 is a serial version.

When run on my laptop (AMD Ryzen 7 PRO 6850U), the concurrent version achieves about 2.8-2.9 GiB/s. The serial one reaches 1.9 GiB/s.

## Running the code

```bash
# Requires Go and pv
# pv used to measure throughput, may need to install if not available on system
# Pipe out to /dev/null to avoid getting spammed
go run main.go | pv > /dev/null
```
