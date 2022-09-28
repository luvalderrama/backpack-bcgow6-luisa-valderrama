package main

import (
    "os"
    "encoding/csv"
    "log"
)

func main() {

    rows := [][]string {
        {"ID", "Precio", "Cantidad"},
        {"1", "273888", "4"},
    }
	csvfile, err := os.Create("productos.csv")
 
    if err != nil {
        log.Fatalf("failed creating file: %s", err)
    }
 
    csvwriter := csv.NewWriter(csvfile)
 
    for _, row := range rows {
        _ = csvwriter.Write(row)
    }
    csvwriter.Flush()
	
    csvfile.Close()
}
