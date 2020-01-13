package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	tr, trReplay, zond, zondReplay, fam, alk, ln, allSacses, all, reBuild, rox, hex, joe, cy5, tmr, r6g int
)

func main() {
	counts := make(map[string]int)
	countsDubl := make(map[string]int)
	files := filesNames(os.Stdin)
	for _, arg := range files {
		if strings.Contains(arg, "доп") {
			continue
		}
		if strings.Contains(arg, "ljg") {
			continue
		}
		f, err := os.Open(arg)
		if err != nil {
			fmt.Println(arg)
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}
		countLines(f, counts, countsDubl, arg)
		f.Close()
	}
	statistic(counts, countsDubl)
	fmt.Printf(" Всего синтезировано %d\n Всего синтезированo с первого раза %d\n Синтезировано повторно %d\n", all, allSacses, reBuild)
	fmt.Printf(" Fam %d\n Другие красители %d\n Зонды %d\n Зонд повторно %d\n LNA %d\n Очистка на картриджах %d\n Очистка на картриджах повторно %d\n", fam, alk, zond, zondReplay, ln, tr, trReplay)
	fmt.Printf("  ROX d\n HEX %d\n JOE %d\n, Cy5 %d\n, TMR %d\n, R6G %d\n, Другие красители %d\n", rox, hex, joe, cy5, tmr, r6g, alk-rox-hex-joe-cy5-tmr-r6g)
}

func filesNames(in *os.File) []string {
	var (
		folder string
		files  []string
	)
	i := bufio.NewScanner(in)
	i.Scan()
	folder = i.Text()
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}
func countLines(f *os.File, counts map[string]int, countsDubl map[string]int, arg string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		if strings.Contains(input.Text(), ",") {
			if strings.Contains(input.Text(), "#") {
				if strings.Contains(input.Text(), "+") {
					oligName := strings.Split(strings.Split(input.Text(), ",")[0], "___")[1] + "tr," + strings.Split(input.Text(), ",")[1]
					counts[oligName]++
					countsDubl[arg+oligName]++
				} else {
					oligName := strings.Split(strings.Split(input.Text(), ",")[0], "___")[1] + "," + strings.Split(input.Text(), ",")[1]
					counts[oligName]++
					countsDubl[arg+oligName]++
				}

			} else {
				oligName := strings.Split(input.Text(), ",")[1] + "," + strings.Split(input.Text(), ",")[2]
				counts[oligName]++
				countsDubl[arg+oligName]++
			}
		} else {
			continue
		}

	}

}
func statistic(counts map[string]int, countsDubl map[string]int) {
	var dubl int
	for line, n := range counts {
		line = strings.ToUpper(line)
		if strings.Contains(line, "[5]") {
			fam++
		}
		if strings.Contains(line, "[7]") {
			alk++
			if strings.Contains(line, "ROX") {
				rox++
			}
			if strings.Contains(line, "HEX") {
				hex++
			}
			if strings.Contains(line, "JOE") {
				joe++
			}
			if strings.Contains(line, "CY5") {
				cy5++
			}
			if strings.Contains(line, "TMR") {
				tmr++
			}
			if strings.Contains(line, "R6G") {
				r6g++
			}
		}
		if (strings.Contains(line, "-S") || strings.Contains(line, "-D")) && (strings.Contains(line, "[7]") || strings.Contains(line, "[5]")) {
			zond++
			if n != 1 {
				zondReplay += (n - 1)
			}
		}
		if strings.Contains(line, "tr") {
			tr++
			if n != 1 {
				trReplay += (n - 1)
			}
		}
		if strings.Contains(line, "[6]") || strings.Contains(line, "[8]") {
			ln++
		}
		if n == 1 {
			allSacses++
			all++
		} else {
			allSacses++
			reBuild += (n - 1)
			all += n
		}

	}
	for _, n := range countsDubl {
		if n > 1 {
			dubl = dubl + (n - 1)
		} else {
			continue
		}
	}
	reBuild = reBuild - dubl
	allSacses = allSacses + dubl
}
