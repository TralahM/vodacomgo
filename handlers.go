package vodacomgo

import (
	"io"
	"text/template"
)

func GenLogin(wr io.Writer, data Login) {
	t, err := template.New("login").Parse(LoginT)
	if err != nil {
		panic(err)
	}
	err = t.Execute(wr, data)
	if err != nil {
		panic(err)
	}
}

func GenB2C(wr io.Writer, data B2C) {
	t, err := template.New("b2c").Parse(B2CT)
	if err != nil {
		panic(err)
	}
	err = t.Execute(wr, data)
	if err != nil {
		panic(err)
	}
}

func GenC2B(wr io.Writer, data C2B) {
	t, err := template.New("c2b").Parse(C2BT)
	if err != nil {
		panic(err)
	}
	err = t.Execute(wr, data)
	if err != nil {
		panic(err)
	}
}
