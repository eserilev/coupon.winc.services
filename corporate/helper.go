//
// Corporate Order Helpers
//

package corporate

import (
	"encoding/csv"
	"log"
	"os"
)

func readCsvFile(file os.FileInfo, path string) [][]string {
	f, err := os.Open(path + file.Name())
	if err != nil {
		log.Fatal("Unable to read input file "+path+file.Name(), err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+path+file.Name(), err)
	}

	return records
}
