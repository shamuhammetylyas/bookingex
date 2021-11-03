package main

import (
	"net/http"
	"os"
	"testing"
)

//setup_test.go => kabir yagdaylara bizin testlerimiz ishlemaka bir zatlar
//set etmeli bolyarys. Meselem NoSurf middleware-ni test etmek ucin ona bir handler ugratmaly bolyarys.
//sho handleri ugratmak ucin hem ony birinji doretmeli. Sholar yaly test ishlemaka bir zatlar set etmeli
//ya doretmeli zatlarymyzy setup_test.go-da edip bilyaris. File name hokman setup_test.go bolmaly
//setup_test.go bizin testlerimizden on ishleyar

//setup_test.go-da hokman funksiyanyn ady TestMain bolmaly we parametri *testing.M bolmaly
func TestMain(m *testing.M) {

	os.Exit(m.Run())
	//yokardaky setire gelinmaka onunden bir zatlar edilyar
	//shondan son testimiz ishlap bashlasyn diyildigi bolyar
}

type myHandler struct{}

func (mh myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

}
