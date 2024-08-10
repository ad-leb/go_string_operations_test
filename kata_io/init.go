package kata_io
import (
	"bufio"
	"os"
)



func init () {
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
}
