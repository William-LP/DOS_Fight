package bot

import (
	"flag"
	"math/rand"
	"time"

	petname "github.com/dustinkirkland/golang-petname"
)

var (
	words     = flag.Int("words", 1, "The number of words in the pet name")
	separator = flag.String("separator", "-", "The separator between words in the pet name")
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func getName() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	return (petname.Generate(*words, *separator))
}
