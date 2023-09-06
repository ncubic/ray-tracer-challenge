package main

import (
	o "rtc/outputs"
	t "rtc/types"
)

func main() {
	start := t.CreatePoint(0, 1, 0)
	velocity := t.CreateVector(1, 1.8, 0).Normalize().Multiply(11.25)
	p := Projectile{start, velocity}

	gravity := t.CreateVector(0, -0.1, 0)
	wind := t.CreateVector(-0.01, 0, 0)
	e := Environment{gravity, wind}

	c := t.CreateCanvas(900, 550)

	for {
		p = tick(e, p)

		if p.position.Y <= 0 {
			break
		}

		c.WritePixel(int(p.position.X), c.Height-int(p.position.Y), t.CreateColor(50, 50, 50))
	}

	o.ToFile(c, "cannon")
}

type Projectile struct {
	position t.Tupple
	velocity t.Tupple
}

type Environment struct {
	gravity t.Tupple
	wind    t.Tupple
}

func tick(env Environment, projectile Projectile) Projectile {
	return Projectile{
		position: projectile.position.Add(projectile.velocity),
		velocity: projectile.velocity.Add(env.gravity).Add(env.wind),
	}
}
