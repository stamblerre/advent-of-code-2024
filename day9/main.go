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
	id         int // order for free blocks
	size       int
	typ        BlockType
	startIndex int
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
				id:         id,
				size:       size,
				typ:        File,
				startIndex: index,
			})
			id++
		} else {
			// free space
			for range size {
				expanded = append(expanded, ".")
			}
			blocks = append(blocks, &block{
				size:       size,
				typ:        Free,
				startIndex: index,
			})
		}
		index += size
	}
	shifted, err := shiftBlocks(expanded, blocks, completeFiles)
	if err != nil {
		return -1, err
	}
	checksum, err := checksum(shifted)
	return int(checksum), err
}

func shiftBlocks(text []string, blocks []*block, completeFiles bool) ([]string, error) {
	if completeFiles {
		// you still need to track last file index
		for i := len(blocks) - 1; i >= 0; i-- {
			if blocks[i].typ != File {
				continue
			}
			file := blocks[i]
			for j := i - 1; j >= 0; j-- {
				space := blocks[j]
				// found a match
				if space.typ == Free && space.size >= file.size {

					// insert the file here
					newBlocks := append(blocks[0:j], file)

					remainderIndex := j + 1

					// leftover space?
					if space.size > file.size {
						remainingSpace := &block{
							typ:        Free,
							size:       file.size - space.size,
							startIndex: space.startIndex + file.size,
						}
						newBlocks = append(newBlocks, remainingSpace)

						// check if the block after contains free space to merge up? (unlikely...)

						if j+1 < len(blocks) {
							if after := blocks[j+1]; after.typ == Free {
								remainingSpace.size += after.size
								remainderIndex = j + 2
							}
						}
					}
					newBlocks = append(newBlocks, blocks[remainderIndex:i]...)
					newBlocks = append(newBlocks, blocks[i+1:]...)

					blocks = newBlocks
					break // we are done
				}
			}
		}
		return blocksToText(blocks), nil
	}
	if !completeFiles {
		// find the last index of a file
		lastFileIndex := -1
		for i := len(blocks) - 1; i >= 0; i-- {
			if blocks[i].typ != File {
				continue
			}
			lastFileIndex = blocks[i].startIndex
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
	}
	return text, nil
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

func checksum(expanded []string) (uint64, error) {
	var result uint64
	for position, value := range expanded {
		if value == "." {
			continue
		}
		asInt, err := strconv.Atoi(string(value))
		if err != nil {
			return 0, err
		}
		result += (uint64(position) * uint64(asInt))
	}
	return result, nil
}
