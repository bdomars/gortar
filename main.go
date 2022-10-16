package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bdomars/gortar/grid"
)

func main() {

	gridsize := flag.Float64("gridsize", 300.0, "Define a custom grid size")
	flag.Parse()
	base_str := flag.Arg(0)
	target_str := flag.Arg(1)

	g := grid.NewGrid(*gridsize)

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

	flight_path := target_pos
	flight_path.Sub(base_pos)
	distance := grid.GetDistance(flight_path)
	angle := grid.GetAngleDegrees(flight_path)

	fmt.Printf("Base is %s at %+v\n", base, base_pos)
	fmt.Printf("Target is %s at %+v\n", target, target_pos)

	fmt.Printf("Ditance %.1fm\n", distance)
	if distance > 1250 {
		fmt.Printf("Target is out of range.\n")
	}

	fmt.Printf("Angle %.1f\n", angle)
	fmt.Printf("Mils %.1f\n\n", grid.GetMils(distance))

}
