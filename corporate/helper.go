//
// Corporate Order Helpers
//

package corporate

import (
	"encoding/csv"
	"net/http"
)

func readCsvFile(r *http.Request) [][]string {
	reader := csv.NewReader(r.Body)
	records, err := reader.ReadAll()
	if err != nil {
		return nil
	}
	return records
}
