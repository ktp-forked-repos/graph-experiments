package layout

import (
	"math"
)

// integrate performs forces integration using Euler numerical
// integration method.
//
// F = d(m * v) / dt
//  (mass is constant in our case)
// v = d{x,y,z}/dt
//
// dv = dt * F / m
//
// d{x,y,z} = v * dt
func (l *Layout3D) integrate(forces []*force) {
	const dt = float64(3) // FIXME: 20 what?
	for i := 0; i < len(l.nodes); i++ {
		body := l.nodes[i]
		coeff := dt / float64(body.Mass)

		dvx := coeff * forces[i].dx
		dvy := coeff * forces[i].dy
		dvz := coeff * forces[i].dz
		v := math.Sqrt(dvx*dvx + dvy*dvy + dvz*dvz)

		if v > 1 {
			dvx = dvx / v
			dvy = dvy / v
			dvz = dvz / v
		}

		dx := dt * dvx
		dy := dt * dvy
		dz := dt * dvz

		l.nodes[i].X += int32(dx)
		l.nodes[i].Y += int32(dy)
		l.nodes[i].Z += int32(dz)
	}
}
