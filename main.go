package main
 
import "github.com/emmaanuel/transfer.sh/cmd"

func main() {
	app := cmd.New()
	app.RunAndExitOnError()
}
