package main

import (
	"image/gif"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", ulam)
	http.ListenAndServe(":3000", nil)
}

// Memoize ulamMem to save data for the next call
var ulamMem = memoize(createUlam)

func ulam(w http.ResponseWriter, r *http.Request) {

	var p []position //positions are the positions for each line in the ulam spiral

	ulam := New()

	//create a number of ulam positions to send to writeUlam
	for i := 0; i < 2000; i++ {
		p = append(p, ulam.next())
	}

	if err := gif.EncodeAll(w, ulamMem(p)); err != nil {
		io.WriteString(w, err.Error())
	}
}

// func main() {
// 	var p []position //positions are the positions for each line in the ulam spiral

// 	ulam := New()

// 	//create a number of ulam positions to send to writeUlam
// 	for i := 0; i < 10; i++ {
// 		p = append(p, ulam.next())
// 	}

// 	ulamMem := memoize(createUlam)

// 	encodeUlam(ulamMem(p), os.Stdout)
// }
