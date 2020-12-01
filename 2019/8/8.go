package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	pictureX := 25
	pictureY := 6
	layerSize := pictureX * pictureY
	var pixelSlice []int
	for _, char := range line {

		pixel, _ := strconv.Atoi(string(char))
		pixelSlice = append(pixelSlice, pixel)
	}

	picture := chunkPixels(pixelSlice, layerSize)
	//taskA
	getMaxZeroLayer(picture, layerSize)
	//taskB
	finalPicture := getFinalPicture(picture, layerSize)
	printPicture(finalPicture, layerSize, pictureX)
}

func getFinalPicture(picture [][]int, layerSize int) (finalLayer []int) {
	layers := len(picture)
	for i := 0; i < layerSize; i++ {
		for l := 0; l < layers; l++ {
			done := false
			switch picture[l][i] {
			case 0:
				finalLayer = append(finalLayer, 0)
				done = true
				break
			case 1:
				finalLayer = append(finalLayer, 1)
				done = true
				break
			case 2:
			}
			if done {
				break
			}
			if l+1 == layers {
				finalLayer[i] = 2

			}
		}
	}
	return
}

func printPicture(finalLayer []int, layerSize int, bufferSize int) {
	for i := 0; i < layerSize; i++ {
		if finalLayer[i] == 1 {
			fmt.Print("#")
		} else {
			fmt.Print(" ")
		}
		if (i+1)%bufferSize == 0 {
			fmt.Println(" ")
		}

	}
}

func chunkPixels(pixels []int, layerSize int) (picture [][]int) {
	for i := 0; i < len(pixels); i += layerSize {
		end := i + layerSize
		if end > len(pixels) {
			end = len(pixels)
		}
		picture = append(picture, pixels[i:end])
	}
	return
}

func getMaxZeroLayer(picture [][]int, layerSize int) {
	fewestZeros := layerSize
	for layNo, layer := range picture {
		zeroCount := 0
		oneCount := 0
		twoCount := 0
		for _, pixel := range layer {
			switch pixel {
			case 0:
				zeroCount = zeroCount + 1
			case 1:
				oneCount = oneCount + 1
			case 2:
				twoCount = twoCount + 1
			}
		}
		if zeroCount < fewestZeros {
			fmt.Println("new boss", layNo, zeroCount, oneCount*twoCount)
			fewestZeros = zeroCount
		}

	}
}
