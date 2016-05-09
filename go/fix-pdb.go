package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func LoadFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := make([]string, 0, 1000)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func FixPDB(lines []string) {
	index := 1

	mustReplace := regexp.MustCompile("(?i)(H[efgs]|B[eraihk]|C[laroudsnmf]|N[eaibdpo]|Os|F[erlm]|P[dtbormau]|S[icernbgm]|Kr|I[nr])")

	for l, line := range lines {
		if len(line) < 6 {
			continue
		}

		section := line[:6]

		if section == "ATOM  " || section == "HETATM" {
			line = fmt.Sprintf("%v%5d%v", section, index, line[11:])
			lines[l] = line
			index += 1
		}

		if section == "HETATM" {
			element := strings.TrimSpace(line[76:78])
			name := line[12:16]

			if len(element) == 1 && mustReplace.MatchString(name) {
				beginEnd := mustReplace.FindIndex([]byte(name))
				begin := int(beginEnd[0]) + 1
				line = fmt.Sprintf("%vX%v renamed", line[:12+begin], line[12+begin+1:])
				lines[l] = line
			}
		}
	}
}

func main() {
	lines, err := LoadFile(os.Getenv("INPUT_PDB"))
	if err != nil {
		panic(err)
	}

	FixPDB(lines)

	out, err := os.Create(os.Getenv("OUTPUT_PDB"))
	if err != nil {
		panic(err)
	}
	defer out.Close()

	writer := bufio.NewWriter(out)

	for _, line := range lines {
		writer.WriteString(line)
		writer.WriteRune('\n')
	}

	writer.Flush()
}
