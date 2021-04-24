package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
)

func main() {
	inputArgs := os.Args[1:]

	if len(inputArgs) != 3 {
		printUsage()
		return
	}
	sourceFolder := inputArgs[0]
	pictureDest := inputArgs[1]
	videoDest := inputArgs[2]

	movePictures(sourceFolder, pictureDest)
	moveVideos(sourceFolder, videoDest)
}

func movePictures(source string, destination string) {
	err := filepath.Walk(source, func(path string, file os.FileInfo, err error) error {
		match, err := filepath.Match("*.[j|p][p|n]g", file.Name())
		if err != nil {
			return err
		}
		if match {
			isDupe, err := filepath.Match("*-[1-9].[j|p][p|n]g", file.Name())
			if err != nil {
				return err
			}
			if isDupe {
				err := os.Rename(path, fmt.Sprintf("C:/Temp/DupePics/%s", file.Name()))
				if err != nil {
					return err
				}
			} else {
				img, err := os.Open(path)
				if err != nil {
					return err
				}

				exifData, err := exif.Decode(img)
				if err != nil {
					return err
				}
				img.Close()

				taken, err := exifData.DateTime()
				if err != nil {
					fmt.Printf("DateTime data not found on %s", file.Name())
					err := os.Rename(path, fmt.Sprintf("C:/Temp/DupePics/%s", file.Name()))
					if err != nil {
						return err
					}
					return nil
				}
				err = os.Rename(path, fmt.Sprintf("%s/%d/%s", destination, taken.Year(), file.Name()))
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(fmt.Errorf("Error moving pictures: %s", err))
	}
}

func moveVideos(source string, destination string) {
	err := filepath.Walk(source, func(path string, file os.FileInfo, err error) error {
		match, err := filepath.Match("*.[m|3|g][p|g|i|o][4|p|f|v]", file.Name())
		if err != nil {
			return err
		}
		if match {
			takenRaw := strings.Split(file.Name(), "-")[0]
			taken, err := strconv.Atoi(takenRaw)
			if err != nil {
				fmt.Printf("Year data not found on %s", file.Name())
				return nil
			}
			err = os.Rename(path, fmt.Sprintf("%s/%d/%s", destination, taken, file.Name()))
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(fmt.Errorf("Error moving videos: %s", err))
	}
}

func printUsage() {
	fmt.Println("---------------")
	fmt.Println("--Media Mover--")
	fmt.Println("---------------")
	fmt.Println("Usage: mediamover [source folder] [picture destination] [video destination]")
	fmt.Println("---------------")
}