package main

import (
	"fmt"
	"math"
)

type moon struct {
	x    float64
	y    float64
	z    float64
	xvel float64
	yvel float64
	zvel float64
}

func main() {
	io := moon{x: 3, y: 15, z: 8}
	europa := moon{x: 5, y: -1, z: -2}
	ganymede := moon{x: 10, y: 8, z: 2}
	callisto := moon{x: 8, y: 4, z: -5}

	/*
		io := moon{x: -8, y: -10, z: 0}
		europa := moon{x: 5, y: 5, z: 10}
		ganymede := moon{x: 2, y: -7, z: 3}
		callisto := moon{x: 9, y: -8, z: -3}
	*/
	//oldIo := io
	//oldEuropa := europa
	//oldGanymede := ganymede
	//oldCallistio := callisto

	for i := 1; i < 1001; i++ {

		//IO
		tempX, tempY, tempZ := applyNewSpeed(io, europa)
		tempXx, tempYy, tempZz := applyNewSpeed(io, ganymede)
		tempXxx, tempYyy, tempZzz := applyNewSpeed(io, callisto)

		velX := tempX + tempXx + tempXxx
		velY := tempY + tempYy + tempYyy
		velZ := tempZ + tempZz + tempZzz
		newIO := moon{
			x:    io.x + velX + io.xvel,
			y:    io.y + velY + io.yvel,
			z:    io.z + velZ + io.zvel,
			xvel: io.xvel + velX,
			yvel: io.yvel + velY,
			zvel: io.zvel + velZ,
		}

		//EUROPA
		tempX, tempY, tempZ = applyNewSpeed(europa, io)
		tempXx, tempYy, tempZz = applyNewSpeed(europa, ganymede)
		tempXxx, tempYyy, tempZzz = applyNewSpeed(europa, callisto)

		velX = tempX + tempXx + tempXxx
		velY = tempY + tempYy + tempYyy
		velZ = tempZ + tempZz + tempZzz

		newEuropa := moon{
			x:    europa.x + velX + europa.xvel,
			y:    europa.y + velY + europa.yvel,
			z:    europa.z + velZ + europa.zvel,
			xvel: europa.xvel + velX,
			yvel: europa.yvel + velY,
			zvel: europa.zvel + velZ,
		}

		//GANYMEDE
		tempX, tempY, tempZ = applyNewSpeed(ganymede, io)
		tempXx, tempYy, tempZz = applyNewSpeed(ganymede, europa)
		tempXxx, tempYyy, tempZzz = applyNewSpeed(ganymede, callisto)

		velX = tempX + tempXx + tempXxx
		velY = tempY + tempYy + tempYyy
		velZ = tempZ + tempZz + tempZzz

		newGanymede := moon{
			x:    ganymede.x + velX + ganymede.xvel,
			y:    ganymede.y + velY + ganymede.yvel,
			z:    ganymede.z + velZ + ganymede.zvel,
			xvel: ganymede.xvel + velX,
			yvel: ganymede.yvel + velY,
			zvel: ganymede.zvel + velZ,
		}

		//Callisto
		tempX, tempY, tempZ = applyNewSpeed(callisto, io)
		tempXx, tempYy, tempZz = applyNewSpeed(callisto, europa)
		tempXxx, tempYyy, tempZzz = applyNewSpeed(callisto, ganymede)

		velX = tempX + tempXx + tempXxx
		velY = tempY + tempYy + tempYyy
		velZ = tempZ + tempZz + tempZzz

		newCallisto := moon{
			x:    callisto.x + velX + callisto.xvel,
			y:    callisto.y + velY + callisto.yvel,
			z:    callisto.z + velZ + callisto.zvel,
			xvel: callisto.xvel + velX,
			yvel: callisto.yvel + velY,
			zvel: callisto.zvel + velZ,
		}

		//oldIo = io
		io = newIO

		//	oldEuropa = europa
		europa = newEuropa

		//	oldGanymede = ganymede
		ganymede = newGanymede

		//	oldCallistio = callisto
		callisto = newCallisto

		fmt.Println("After ", i, "step:")
		fmt.Println(io)
		fmt.Println(europa)
		fmt.Println(ganymede)
		fmt.Println(callisto)
	}

	pot1 := math.Abs(io.x) + math.Abs(io.y) + math.Abs(io.z)
	kin1 := math.Abs(io.xvel) + math.Abs(io.yvel) + math.Abs(io.zvel)

	pot2 := math.Abs(europa.x) + math.Abs(europa.y) + math.Abs(europa.z)
	kin2 := math.Abs(europa.xvel) + math.Abs(europa.yvel) + math.Abs(europa.zvel)

	pot3 := math.Abs(ganymede.x) + math.Abs(ganymede.y) + math.Abs(ganymede.z)
	kin3 := math.Abs(ganymede.xvel) + math.Abs(ganymede.yvel) + math.Abs(ganymede.zvel)

	pot4 := math.Abs(callisto.x) + math.Abs(callisto.y) + math.Abs(callisto.z)
	kin4 := math.Abs(callisto.xvel) + math.Abs(callisto.yvel) + math.Abs(callisto.zvel)

	tot := (pot1 * kin1) + (pot2 * kin2) + (pot3 * kin3) + (pot4 * kin4)

	fmt.Println(tot)

}

func applyNewSpeed(a moon, b moon) (newX float64, newY float64, newZ float64) {
	newX = 0
	newY = 0
	newZ = 0

	if a.x < b.x {
		newX = newX + 1
	} else if a.x > b.x {
		newX = newX - 1
	}
	if a.y < b.y {
		newY = newY + 1
	} else if a.y > b.y {
		newY = newY - 1
	}
	if a.z < b.z {
		newZ = newZ + 1
	} else if a.z > b.z {
		newZ = newZ - 1
	}
	return
}
