package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bdomars/gortar/grid"
)

func main() {

	var gridsize float64
	flag.Float64Var(&gridsize, "grid", 300.0, "Define a custom grid size")
	flag.Parse()

	base_str := flag.Arg(0)
	target_str := flag.Arg(1)

	g := grid.NewGrid(gridsize)

	base, err := g.Parse(base_str)
	if err != nil {
		fmt.Printf("Failed to parse base coordinate: %v\n", err)
		os.Exit(1)
	}

	target, err := g.Parse(target_str)
	if err != nil {
		fmt.Printf("Failed to parse target coordinate: %v\n", err)
		os.Exit(1)
	}

	base_pos := base.Position()
	target_pos := target.Position()

	flight_path := target_pos.Sub(base_pos)

	distance := grid.GetDistance(flight_path)
	angle := grid.GetAngleDegrees(flight_path)

	fmt.Printf("Base is %s\n", base)
	fmt.Printf("Target is %s\n\n", target)

	fmt.Printf("Distance\t%6.1f m\n", distance)

	if distance > 1250 {
		fmt.Printf("Target is out of range.\n")
	}

	if distance < 50 {
		fmt.Printf("Target is too close.\n")
	}

	fmt.Printf("Angle\t\t%6.1f °\n", angle)
	fmt.Printf("Elevation\t%6.1f mil\n\n", grid.GetMils(distance))

}
