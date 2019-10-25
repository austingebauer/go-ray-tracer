package main

import "testing"

func BenchmarkRenderRayTracedWorld3D(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RenderRayTracedWorld3D()
	}
}

func BenchmarkRenderRayTracedSphere3D(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RenderRayTracedSphere3D()
	}
}

func BenchmarkRenderRayTracedSphere2D(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RenderRayTracedSphere2D()
	}
}

func BenchmarkRenderClock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RenderClock()
	}
}

func BenchmarkRenderProjectile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RenderProjectile()
	}
}
