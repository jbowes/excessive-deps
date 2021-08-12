package main

import (
	"github.com/sanity-io/litter"
	_ "gopkg.in/yaml.v2"
)

func Main() {
	litter.Dumps("foo")
}
