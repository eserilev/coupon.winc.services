//
// Corporate Order Helpers
//

package corporate

import (
	"encoding/csv"
	"io"
	"net/http"
)

func readCsvFile(r *http.Request) [][]string {
	reader := csv.NewReader(r.Body)
	var results [][]string
	for {
		// read one row from csv
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil
		}

		// add record to result set
		results = append(results, record)
	}
	return results
}
