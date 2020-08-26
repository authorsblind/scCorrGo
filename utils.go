package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type dfRow struct {
	a float64
	b float64
	c float64
	d float64
}

type CorrTable map[string][]int

func readCsv(filepth string) ([][]string, error) {
	csvfile, err := os.Open(filepth)

	if err != nil {
		return nil, err
	}

	defer csvfile.Close()

	reader := csv.NewReader(csvfile)
	fields, err := reader.ReadAll()

	return fields, nil
}

func parseFloat64(s string) float64 {
	t, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatal("Error parsing ", s, "exiting")
	}

	return t
}

func parseCsv(pth string) ([]dfRow, error) {
	records, err := readCsv(pth)

	if err != nil {
		return nil, err
	}

	arr := make([]dfRow, len(records))

	for i := 0; i < len(records); i++ {
		arr[i].a = parseFloat64(records[i][0])
		arr[i].b = parseFloat64(records[i][1])
		arr[i].c = parseFloat64(records[i][2])
		arr[i].d = parseFloat64(records[i][3])
	}

	return arr, nil
}

func (table CorrTable) updateTable(key float64, val int) {
	roundedKey := fmt.Sprintf("%.5f", key)
	table[roundedKey] = append(table[roundedKey], val)
}

func toSeqString(idxArr []int, sep string) string {
	var ret []string
	for _, i := range idxArr {
		ret = append(ret, strconv.Itoa(i))
	}
	return strings.Join(ret, sep)
}

func (table CorrTable) writeTable(name string, n int) {
	pth := fmt.Sprintf("n%v/go_%v_%v_rounded_freqs.txt", n, name, n)
	f, err := os.Create(pth) // Will truncate the file if already exists
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var tmpStr, tmpSeq string

	for k, v := range table {
		tmpSeq = toSeqString(v, "|")
		tmpStr = fmt.Sprintf("%v::%v\n", k, tmpSeq)
		f.WriteString(tmpStr)
	}
}
