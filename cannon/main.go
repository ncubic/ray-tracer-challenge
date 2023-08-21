package main

import (
	"fmt"
	t "rtc/types"
	"time"
)

func main() {
	p := Projectile{t.CreatePoint(0, 1, 0), t.CreateVector(1, 1, 0).Normalize()}
	e := Environment{t.CreateVector(0, -0.1, 0), t.CreateVector(-0.01, 0, 0)}

	for {
		p = tick(e, p)

		if p.position.Y <= 0 {
			fmt.Printf("Huston, we have a problem. Impact at %+v with velocity %+v\n", p.position, p.velocity)
			break
		} else {
			fmt.Printf("Huston, problem is closer. Currently at %+v with velocity %+v\n", p.position, p.velocity)
		}

		time.Sleep(1 * time.Second)
	}
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
