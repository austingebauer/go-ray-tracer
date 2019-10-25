package matrix

import (
	"errors"
	"github.com/austingebauer/go-ray-tracer/point"
	"github.com/austingebauer/go-ray-tracer/vector"
	"math"
)

// NewTranslationMatrix returns a new 4x4 translation Matrix.
//
// The translation Matrix returned has the form:
//   | 1 0 0 x |
//   | 0 1 0 y |
//   | 0 0 1 z |
//   | 0 0 0 1 |
func NewTranslationMatrix(x, y, z float64) *Matrix {
	m := NewIdentityMatrix(4)
	m.setValue(0, 3, x)
	m.setValue(1, 3, y)
	m.setValue(2, 3, z)
	return m
}

// NewScalingMatrix returns a new 4x4 scaling Matrix.
//
// The scaling Matrix returned has the form:
//   | x 0 0 0 |
//   | 0 y 0 0 |
//   | 0 0 z 0 |
//   | 0 0 0 1 |
func NewScalingMatrix(x, y, z float64) *Matrix {
	m := NewIdentityMatrix(4)
	m.setValue(0, 0, x)
	m.setValue(1, 1, y)
	m.setValue(2, 2, z)
	return m
}

// NewXRotationMatrix returns a new 4x4 rotation Matrix that can be used to rotate
// a Point or Vector around the X axis by the passed number of radians.
//
// The rotation Matrix returned has the form:
//   | 1 0      0       0 |
//   | 0 cos(r) -sin(r) 0 |
//   | 0 sin(r) cos(r)  0 |
//   | 0 0      0       1 |
func NewXRotationMatrix(radians float64) *Matrix {
	m := NewIdentityMatrix(4)
	m.setValue(1, 1, math.Cos(radians))
	m.setValue(1, 2, -1*math.Sin(radians))
	m.setValue(2, 1, math.Sin(radians))
	m.setValue(2, 2, math.Cos(radians))
	return m
}

// NewYRotationMatrix returns a new 4x4 rotation Matrix that can be used to rotate
// a Point or Vector around the Y axis by the passed number of radians.
//
// The rotation Matrix returned has the form:
//   | cos(r)  0 sin(r) 0 |
//   | 0       1 0      0 |
//   | -sin(r) 0 cos(r) 0 |
//   | 0       0 0      1 |
func NewYRotationMatrix(radians float64) *Matrix {
	m := NewIdentityMatrix(4)
	m.setValue(0, 0, math.Cos(radians))
	m.setValue(0, 2, math.Sin(radians))
	m.setValue(2, 0, -1*math.Sin(radians))
	m.setValue(2, 2, math.Cos(radians))
	return m
}

// NewZRotationMatrix returns a new 4x4 rotation Matrix that can be used to rotate
// a Point or Vector around the z-axis by the passed number of radians.
//
// The rotation Matrix returned has the form:
//   | cos(r) -sin(r) 0 0 |
//   | sin(r) cos(r)  0 0 |
//   | 0      0       1 0 |
//   | 0      0       0 1 |
func NewZRotationMatrix(radians float64) *Matrix {
	m := NewIdentityMatrix(4)
	m.setValue(0, 0, math.Cos(radians))
	m.setValue(0, 1, -1*math.Sin(radians))
	m.setValue(1, 0, math.Sin(radians))
	m.setValue(1, 1, math.Cos(radians))
	return m
}

// NewShearingMatrix returns a new 4x4 shearing Matrix that can be used for a
// shear transformation of a Point.
//
// The shearing Matrix returned has the form:
//   | 1  xy xz 0 |
//   | yx 1  yz 0 |
//   | zx zy 1  0 |
//   | 0  0  0  1 |
func NewShearingMatrix(xy, xz, yx, yz, zx, zy float64) *Matrix {
	m := NewIdentityMatrix(4)
	m.setValue(0, 1, xy)
	m.setValue(0, 2, xz)
	m.setValue(1, 0, yx)
	m.setValue(1, 2, yz)
	m.setValue(2, 0, zx)
	m.setValue(2, 1, zy)
	return m
}

// Translate translates this Matrix by the passed x, y, and z values.
func (m *Matrix) Translate(x, y, z float64) *Matrix {
	mRes, _ := Multiply(*NewTranslationMatrix(x, y, z), *m)
	m.data = mRes.data
	m.rows = mRes.rows
	m.cols = mRes.cols
	return m
}

// Scale scales this Matrix by the passed x, y, and z values.
func (m *Matrix) Scale(x, y, z float64) *Matrix {
	mRes, _ := Multiply(*NewScalingMatrix(x, y, z), *m)
	m.data = mRes.data
	m.rows = mRes.rows
	m.cols = mRes.cols
	return m
}

// RotateX rotates this Matrix around the x-axis by the passed number of radians.
// Rotation happens clockwise when looking down the positive x-axis towards the negative x-axis.
func (m *Matrix) RotateX(radians float64) *Matrix {
	mRes, _ := Multiply(*NewXRotationMatrix(radians), *m)
	m.data = mRes.data
	m.rows = mRes.rows
	m.cols = mRes.cols
	return m
}

// RotateY rotates this Matrix around the y-axis by the passed number of radians.
// Rotation happens clockwise when looking down the positive y-axis towards the negative y-axis.
func (m *Matrix) RotateY(radians float64) *Matrix {
	mRes, _ := Multiply(*NewYRotationMatrix(radians), *m)
	m.data = mRes.data
	m.rows = mRes.rows
	m.cols = mRes.cols
	return m
}

// RotateZ rotates this Matrix around the z-axis by the passed number of radians.
// Rotation happens clockwise when looking down the positive z-axis towards the negative z-axis.
func (m *Matrix) RotateZ(radians float64) *Matrix {
	mRes, _ := Multiply(*NewZRotationMatrix(radians), *m)
	m.data = mRes.data
	m.rows = mRes.rows
	m.cols = mRes.cols
	return m
}

// Shear shears or skews this Matrix in on coordinate relative to another coordinate.
// For example, the parameter xy represents how much to shear x relative to y.
func (m *Matrix) Shear(xy, xz, yx, yz, zx, zy float64) *Matrix {
	mRes, _ := Multiply(*NewShearingMatrix(xy, xz, yx, yz, zx, zy), *m)
	m.data = mRes.data
	m.rows = mRes.rows
	m.cols = mRes.cols
	return m
}

// ViewTransform returns a transformation matrix that can
// be used to transform a camera view in a scene.
//
// The from parameter specifies where the eye is at in the scene.
// The to parameter specifies the point in the scene at which to look at.
// The up parameter specifies which direction is up.
func ViewTransform(from, to point.Point, up vector.Vector) *Matrix {
	// Compute the forward vector by subtracting from from to
	forwardVec := point.Subtract(to, from).Normalize()

	// Compute the left vector by taking the cross product of forward and normalized up
	leftVec := vector.CrossProduct(*forwardVec, *vector.Normalize(up))

	// Compute the true up vector by taking the cross product of left and forward vectors.
	// This allows the original up vector to be only approximately up.
	trueUpVec := vector.CrossProduct(leftVec, *forwardVec)

	// Construct a matrix that represents the orientation transformation
	//   | leftX     leftY     leftZ      0 |
	//   | trueupX   trueupY   trueupZ    0 |
	//   | -forwardX -forwardY -forwardZ  0 |
	//   | 0         0         0          1 |
	orientationM := NewIdentityMatrix(4)
	orientationM.setValue(0, 0, leftVec.X)
	orientationM.setValue(0, 1, leftVec.Y)
	orientationM.setValue(0, 2, leftVec.Z)
	orientationM.setValue(1, 0, trueUpVec.X)
	orientationM.setValue(1, 1, trueUpVec.Y)
	orientationM.setValue(1, 2, trueUpVec.Z)
	orientationM.setValue(2, 0, -1*forwardVec.X)
	orientationM.setValue(2, 1, -1*forwardVec.Y)
	orientationM.setValue(2, 2, -1*forwardVec.Z)

	// Multiply the orientation transformation by a translation transformation
	// in order to move the scene into place before orienting it.
	translationM := NewTranslationMatrix(-1*from.X, -1*from.Y, -1*from.Z)
	viewTransformM, _ := Multiply(*orientationM, *translationM)
	return viewTransformM
}

// PointToMatrix returns a 4x1 Matrix that represents the passed Point.
// The returned Matrix is known as a 'column vector' in linear algebra.
func PointToMatrix(pt *point.Point) *Matrix {
	m := NewMatrix(4, 1)
	_ = m.SetValue(0, 0, pt.X)
	_ = m.SetValue(1, 0, pt.Y)
	_ = m.SetValue(2, 0, pt.Z)
	_ = m.SetValue(3, 0, 1)

	return m
}

// MatrixToPoint returns a Point representation of the passed Matrix.
// An error is returned if the passed Matrix is not of a 3x1 or 4x1 dimension.
func MatrixToPoint(m *Matrix) (*point.Point, error) {
	if m.GetRows() != 3 && m.GetRows() != 4 {
		return nil, errors.New("matrix m must have 3 or 4 rows to be converted to a point")
	}

	if m.GetCols() != 1 {
		return nil, errors.New("matrix m must have 1 column to be converted to a point")
	}

	x, _ := m.GetValue(0, 0)
	y, _ := m.GetValue(1, 0)
	z, _ := m.GetValue(2, 0)
	return point.NewPoint(x, y, z), nil
}

// VectorToMatrix returns a 4x1 Matrix that represents the passed Vector.
// The returned Matrix is known as a 'column vector' in linear algebra.
func VectorToMatrix(vec *vector.Vector) *Matrix {
	m := NewMatrix(4, 1)
	_ = m.SetValue(0, 0, vec.X)
	_ = m.SetValue(1, 0, vec.Y)
	_ = m.SetValue(2, 0, vec.Z)
	_ = m.SetValue(3, 0, 0)

	return m
}

// MatrixToVector returns a Point representation of the passed Matrix.
// An error is returned if the passed Matrix is not of a 3x1 or 4x1 dimension.
func MatrixToVector(m *Matrix) (*vector.Vector, error) {
	if m.GetRows() != 3 && m.GetRows() != 4 {
		return nil, errors.New("matrix m must have 3 or 4 rows to be converted to a vector")
	}

	if m.GetCols() != 1 {
		return nil, errors.New("matrix m must have 1 column to be converted to a vector")
	}

	x, _ := m.GetValue(0, 0)
	y, _ := m.GetValue(1, 0)
	z, _ := m.GetValue(2, 0)
	return vector.NewVector(x, y, z), nil
}
