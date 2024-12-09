package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"advent-of-code-2024.com/internal/shared"
)

func main() {
	t := &today{}
	_, _, err := shared.Run(t, "testdata/input.txt")
	if err != nil {
		log.Fatal(err)
	}
}

type today struct {
}

func (t *today) ReadInput(filename string) (any, error) {
	return os.ReadFile(filename)
}

func (t *today) Part1(input any) (int, error) {
	return implementation(input /*completeFiles=*/, false)
}

func (t *today) Part2(input any) (int, error) {
	return implementation(input /*completeFiles=*/, true)
}

type block struct {
	id   int // order for free blocks
	size uint
	typ  BlockType
}

type BlockType int

const (
	Free BlockType = iota
	File
)

func implementation(input any, completeFiles bool) (int, error) {
	b, ok := input.([]byte)
	if !ok {
		return -1, fmt.Errorf("unexpected input format %T", b)
	}
	text := []rune(string(b))
	minified, err := shared.RuneSliceToInt(text)
	if err != nil {
		return -1, err
	}
	var (
		id       int
		expanded []string
		blocks   []*block
	)
	index := 0
	for i, size := range minified {
		if i%2 == 0 {
			// file
			for range size {
				expanded = append(expanded, fmt.Sprintf("%d", id))
			}
			blocks = append(blocks, &block{
				id:   id,
				size: uint(size),
				typ:  File,
			})
			id++
		} else {
			// free space
			for range size {
				expanded = append(expanded, ".")
			}
			blocks = append(blocks, &block{
				size: uint(size),
				typ:  Free,
			})
		}
		index += size
	}
	var shifted []string
	if completeFiles {
		shifted, err = shiftBlocks(blocks)
	} else {
		shifted, err = shiftFiles(expanded)
	}
	if err != nil {
		return -1, err
	}
	return checksum(shifted)
}

func shiftFiles(text []string) ([]string, error) {
	// find the last index of a file
	lastFileIndex := -1
	for i := len(text) - 1; i >= 0; i-- {
		if text[i] == "." {
			continue
		}
		lastFileIndex = i
		break
	}
	for i, value := range text {
		if value != "." {
			continue
		}
		if lastFileIndex <= i {
			break
		}
		text[i] = text[lastFileIndex]
		text[lastFileIndex] = "."
		lastFileIndex--
		for text[lastFileIndex] == "." {
			lastFileIndex--
		}
	}
	return text, nil
}

func shiftBlocks(blocks []*block) ([]string, error) {
	// you still need to track last file index
	findLastFile := func(start int) int {
		for i := start; i >= 0; i-- {
			if blocks[i].typ != File {
				continue
			}
			return i
		}
		return -1
	}
	lastFileIndex := findLastFile(len(blocks) - 1)
	for lastFileIndex >= 0 {
		file := blocks[lastFileIndex]
		spaceIndex := -1
		for i := 0; i < lastFileIndex; i++ {
			space := blocks[i]
			// found a match
			if space.typ == Free && space.size >= file.size {
				spaceIndex = i
				break
			}
		}
		if spaceIndex >= 0 {
			space := blocks[spaceIndex]

			// insert the file here
			newBlocks := append([]*block{}, blocks[0:spaceIndex]...)
			newBlocks = append(newBlocks, file)

			remainderIndex := spaceIndex + 1

			// leftover space?
			if space.size > file.size {
				remainingSpace := &block{
					typ:  Free,
					size: space.size - file.size,
				}
				newBlocks = append(newBlocks, remainingSpace)

				// check if the block after contains free space to merge up? (unlikely...)

				if spaceIndex+1 < len(blocks) {
					if after := blocks[spaceIndex+1]; after.typ == Free {
						remainingSpace.size += after.size
						remainderIndex = spaceIndex + 2
					}
				}
			}

			newBlocks = append(newBlocks, blocks[remainderIndex:lastFileIndex]...)
			// leave free space for the file
			newBlocks = append(newBlocks, &block{
				size: file.size,
				typ:  Free,
			})
			newBlocks = append(newBlocks, blocks[lastFileIndex+1:]...)

			// reset...
			blocks = newBlocks
		}
		// reset last file index
		lastFileIndex = findLastFile(lastFileIndex - 1)
	}
	return blocksToText(blocks), nil
}

func blocksToText(blocks []*block) []string {
	var result []string
	for _, b := range blocks {
		for range b.size {
			switch b.typ {
			case File:
				result = append(result, fmt.Sprintf("%d", b.id))
			case Free:
				result = append(result, ".")
			}
		}
	}
	return result
}

func checksum(expanded []string) (int, error) {
	var result int
	for position, value := range expanded {
		if value == "." {
			continue
		}
		asInt, err := strconv.Atoi(string(value))
		if err != nil {
			return 0, err
		}
		result += position * asInt
	}
	return result, nil
}
