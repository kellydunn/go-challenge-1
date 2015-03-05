// This file provides the ability to encode a pattern into a file.
package drum

import (
	"encoding/binary"
	"os"
        "bytes"
)

var SPLICE_FILE_HEADER = "SPLICE"

func EncodePattern(pattern *Pattern, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	err = binary.Write(file, binary.BigEndian, []byte(SPLICE_FILE_HEADER))
	if err != nil {
		return err
	}

	var size uint64
	size = uint64(VERSION_SIZE + BPM_SIZE)
	for _, track := range pattern.Tracks {
		size += uint64(TRACK_ID_SIZE + 4 + len(track.Name) + STEP_SEQUENCE_SIZE)
	}

	err = binary.Write(file, binary.BigEndian, &size)
	if err != nil {
		return err
	}

	version := make([]byte, VERSION_SIZE)
	tmp := []byte(pattern.Version)
	for i, _ := range tmp {
		version[i] = tmp[i]
	}
	err = binary.Write(file, binary.BigEndian, version)
	if err != nil {
		return err
	}

        err = binary.Write(file, binary.LittleEndian, &pattern.Tempo)
	if err != nil {
		return err
	}

        for _, track := range pattern.Tracks {
            err = binary.Write(file, binary.BigEndian, &track.Id)
            if err != nil {
               return err
            }

            var trackNameLen uint32
            trackNameLen = uint32(len(track.Name))
            err = binary.Write(file, binary.BigEndian, &trackNameLen)
            if err != nil {
                   return err
            }

            trackName := []byte(track.Name)
            err = binary.Write(file, binary.BigEndian, trackName)
            if err != nil {
                  return err
            }

            for _, step := range track.StepSequence.Steps {
                if bytes.Compare([]byte{step}, []byte{0}) == 0 {
                   err = binary.Write(file, binary.BigEndian, byte(0))
                } else if bytes.Compare([]byte{step}, []byte{1}) == 0 {
                   err = binary.Write(file, binary.BigEndian, byte(1))
                }

                if err != nil {
                   return err
                }
            }
             
        }


	return nil
}
