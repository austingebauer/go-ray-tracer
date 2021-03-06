package ray

import (
	"testing"

	"github.com/austingebauer/go-ray-tracer/matrix"
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/vector"
	"github.com/stretchr/testify/assert"
)

func TestNewRay(t *testing.T) {
	type args struct {
		origin    *point.Point
		direction *vector.Vector
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "ray has origin and direction",
			args: args{
				origin:    point.NewPoint(1, 2, 3),
				direction: vector.NewVector(4, 5, 6),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRay(*tt.args.origin, *tt.args.direction)
			assert.Equal(t, tt.args.origin, r.Origin)
			assert.Equal(t, tt.args.direction, r.Direction)
		})
	}
}

func TestPosition(t *testing.T) {
	type args struct {
		ray *Ray
		t   float64
	}
	tests := []struct {
		name string
		args args
		want *point.Point
	}{
		{
			name: "compute point that lies distance t along the ray 1",
			args: args{
				ray: NewRay(*point.NewPoint(2, 3, 4), *vector.NewVector(1, 0, 0)),
				t:   0,
			},
			want: point.NewPoint(2, 3, 4),
		},
		{
			name: "compute point that lies distance t along the ray 2",
			args: args{
				ray: NewRay(*point.NewPoint(2, 3, 4), *vector.NewVector(1, 0, 0)),
				t:   1,
			},
			want: point.NewPoint(3, 3, 4),
		},
		{
			name: "compute point that lies distance t along the ray 3",
			args: args{
				ray: NewRay(*point.NewPoint(2, 3, 4), *vector.NewVector(1, 0, 0)),
				t:   -1,
			},
			want: point.NewPoint(1, 3, 4),
		},
		{
			name: "compute point that lies distance t along the ray 4",
			args: args{
				ray: NewRay(*point.NewPoint(2, 3, 4), *vector.NewVector(1, 0, 0)),
				t:   2.5,
			},
			want: point.NewPoint(4.5, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Position(tt.args.ray, tt.args.t))
		})
	}
}

func TestTransform(t *testing.T) {
	type args struct {
		ray *Ray
		m   *matrix.Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    *Ray
		wantErr bool
	}{
		{
			name: "transform ray with translation matrix",
			args: args{
				ray: NewRay(*point.NewPoint(1, 2, 3), *vector.NewVector(0, 1, 0)),
				m:   matrix.NewTranslationMatrix(3, 4, 5),
			},
			want:    NewRay(*point.NewPoint(4, 6, 8), *vector.NewVector(0, 1, 0)),
			wantErr: false,
		},
		{
			name: "transform ray with matrix not of 4x4 order for error",
			args: args{
				ray: NewRay(*point.NewPoint(1, 2, 3), *vector.NewVector(0, 1, 0)),
				m:   matrix.NewMatrix(1, 4),
			},
			want:    NewRay(*point.NewPoint(4, 6, 8), *vector.NewVector(0, 1, 0)),
			wantErr: true,
		},
		{
			name: "transform ray with scaling matrix",
			args: args{
				ray: NewRay(*point.NewPoint(1, 2, 3), *vector.NewVector(0, 1, 0)),
				m:   matrix.NewScalingMatrix(2, 3, 4),
			},
			want:    NewRay(*point.NewPoint(2, 6, 12), *vector.NewVector(0, 3, 0)),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Transform(tt.args.ray, tt.args.m)

			if tt.wantErr {
				assert.Nil(t, got)
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestEquals(t *testing.T) {
	type args struct {
		r1 *Ray
		r2 *Ray
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "rays have equal origin and direction",
			args: args{
				r1: NewRay(*point.NewPoint(0, 0, 0), *vector.NewVector(1, 0, 0)),
				r2: NewRay(*point.NewPoint(0, 0, 0), *vector.NewVector(1, 0, 0)),
			},
			want: true,
		},
		{
			name: "rays do not have equal origin",
			args: args{
				r1: NewRay(*point.NewPoint(1, 0, 0), *vector.NewVector(1, 0, 0)),
				r2: NewRay(*point.NewPoint(0, 0, 0), *vector.NewVector(1, 0, 0)),
			},
			want: false,
		},
		{
			name: "rays do not have equal direction",
			args: args{
				r1: NewRay(*point.NewPoint(0, 0, 0), *vector.NewVector(-1, 0, 0)),
				r2: NewRay(*point.NewPoint(0, 0, 0), *vector.NewVector(1, 0, 0)),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Equals(tt.args.r1, tt.args.r2))
		})
	}
}
