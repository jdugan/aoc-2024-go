package day09

import (
	"slices"

	"github.com/elliotchance/pie/v2"
)

// ========== DEFINITION ==================================

type FileSystem struct {
	files []File
}

// ========== RECEIVERS ===================================

// func (fs FileSystem) CheckSum() int {
// 	return 0
// }

func (fs *FileSystem) Compress() int {
	disk := make([]int, 0)
	values := make([]int, 0)
	spaces := make([]int, 0)
	idx := 0
	for _, file := range fs.files {
		for i := 0; i < file.blocks; i++ {
			disk = append(disk, file.id)
			values = append(values, idx)
			idx += 1
		}
		for i := 0; i < file.space; i++ {
			disk = append(disk, -1)
			spaces = append(spaces, idx)
			idx += 1
		}
	}
	slices.Reverse(values)
	curr_value, rem_values := pie.Shift(values)
	curr_space, rem_spaces := pie.Shift(spaces)
	for curr_value > curr_space {
		disk[curr_space] = disk[curr_value]
		disk[curr_value] = -1
		curr_value, rem_values = pie.Shift(rem_values)
		curr_space, rem_spaces = pie.Shift(rem_spaces)
	}
	checksum := 0
	for i, v := range disk {
		if v > -1 {
			checksum += i * v
		}
	}
	return checksum
}
