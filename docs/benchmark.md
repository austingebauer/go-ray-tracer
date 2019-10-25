# Benchmark Results Tracking

This document serves to track benchmark results as I add more features 
to go-ray-tracer.

# Benchmark Test Results

```bash
# After using pointer arguments for matrix operations instead of copy values
> $ make bench
go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/austingebauer/go-ray-tracer
BenchmarkRenderRayTracedWorld3D-4              1        14655736669 ns/op
BenchmarkRenderRayTracedSphere3D-4             1        3184884970 ns/op
BenchmarkRenderRayTracedSphere2D-4             1        2220374784 ns/op
BenchmarkRenderClock-4                      1212            887304 ns/op
BenchmarkRenderProjectile-4                  681           1649569 ns/op
PASS
ok      github.com/austingebauer/go-ray-tracer  22.859s

# After caching inverse transform matrix in camera
> $ make bench
go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/austingebauer/go-ray-tracer
BenchmarkRenderRayTracedWorld3D-4              1        17952116969 ns/op
BenchmarkRenderRayTracedSphere3D-4             1        3865791821 ns/op
BenchmarkRenderRayTracedSphere2D-4             1        2623375120 ns/op
BenchmarkRenderClock-4                      1129           1052287 ns/op
BenchmarkRenderProjectile-4                  675           1746618 ns/op
PASS
ok      github.com/austingebauer/go-ray-tracer  27.465s

> $ make bench
go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/austingebauer/go-ray-tracer
BenchmarkRenderRayTracedWorld3D-4              1        23918012287 ns/op
BenchmarkRenderRayTracedSphere3D-4             1        4553134003 ns/op
BenchmarkRenderRayTracedSphere2D-4             1        3081435354 ns/op
BenchmarkRenderClock-4                       944           1170333 ns/op
BenchmarkRenderProjectile-4                  564           2159013 ns/op
PASS
ok      github.com/austingebauer/go-ray-tracer  35.491s
```

# Rendering Image Results

```bash
# After using pointer arguments for matrix operations instead of copy values
> $ make run
go run main.go
Wrote ppm rendering to: docs/renderings/clock/clock.ppm

Wrote ppm rendering to: docs/renderings/projectile/projectile.ppm

Wrote ppm rendering to: docs/renderings/sphere_2d/sphere_2d.ppm

Wrote ppm rendering to: docs/renderings/sphere_3d/sphere_3d.ppm

Wrote ppm rendering to: docs/renderings/world_3d/world_3d.ppm

Render time: 16.923434931 seconds

# After caching inverse transform matrix in camera
> $ make run
go run main.go
Wrote ppm rendering to: docs/renderings/clock/clock.ppm

Wrote ppm rendering to: docs/renderings/projectile/projectile.ppm

Wrote ppm rendering to: docs/renderings/sphere_2d/sphere_2d.ppm

Wrote ppm rendering to: docs/renderings/sphere_3d/sphere_3d.ppm

Wrote ppm rendering to: docs/renderings/world_3d/world_3d.ppm

Render time: 21.839597926 seconds

> $ make run
go run main.go
Wrote ppm rendering to: docs/renderings/clock/clock.ppm

Wrote ppm rendering to: docs/renderings/projectile/projectile.ppm

Wrote ppm rendering to: docs/renderings/sphere_2d/sphere_2d.ppm

Wrote ppm rendering to: docs/renderings/sphere_3d/sphere_3d.ppm

Wrote ppm rendering to: docs/renderings/world_3d/world_3d.ppm

Render time: 25.566676662 seconds
```
