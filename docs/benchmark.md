# Benchmark Results Tracking

This document serves to track benchmark results as I add more features 
to go-ray-tracer.

# Results

```bash
> $ make bench
go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/austingebauer/go-ray-tracer
BenchmarkRenderRayTracedSphere3D-4             1        3396119188 ns/op
BenchmarkRenderRayTracedSphere2D-4             1        2320342009 ns/op
BenchmarkRenderClock-4                      1222            903268 ns/op
BenchmarkRenderProjectile-4                  675           1694453 ns/op
PASS
ok      github.com/austingebauer/go-ray-tracer  8.551s
```
