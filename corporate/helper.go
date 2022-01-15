//
// Corporate Order Helpers
//

package corporate

const columns int = 18

func readCsvFile(content []byte) []string {
	c := string(content[:])
	contentLength := len(c)
	rows := contentLength / columns
	stringContent := make([]string, rows)
	for i := 0; i < rows; i++ {
		start := columns * i
		end := columns * (i + 1)
		stringContent[i] = c[start:end]
	}
	return stringContent
}
