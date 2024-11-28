package main

import "photomosaics/utils"

func main() {
	filename, pixels := "dummy.png", 22
	utils.Rotate(filename, pixels)
}
