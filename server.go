package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println()
	fmt.Println("======================================")
	fmt.Println("SWAPI by Luiz Janela")
	fmt.Println("======================================")
	fmt.Println("Starting SWAPI Server @ ", ApiEndpoint)
	fmt.Println("May the force be with us!")
	fmt.Println()
	fmt.Println(`
                       .-.
                      |_:_|
                     /(_Y_)\
.                   ( \/M\/ )
 '.               _.'-/'-'\-'._
   ':           _/.--'[[[[]'--.\_
     ':        /_'  : |::"| :  '.\
       ':     //   ./ |oUU| \.'  :\
         ':  _:'..' \_|___|_/ :   :|
           ':.  .'  |_[___]_|  :.':\
            [::\ |  :  | |  :   ; : \
             '-'   \/'.| |.' \  .;.' |
             |\_    \  '-'   :       |
             |  \    \ .:    :   |   |
             |   \    | '.   :    \  |
             /       \   :. .;       |
            /     |   |  :__/     :  \\
           |  |   |    \:   | \   |   ||
          /    \  : :  |:   /  |__|   /|
      snd |     : : :_/_|  /'._\  '--|_\
          /___.-/_|-'   \  \
                         '-'
Art by Shanaka Dias`)

	fmt.Println("======================================")

	r := mux.NewRouter()

	// Routes definitions
	r.HandleFunc("/api/planets", planetCreateHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/planets", planetSearchHandler).Queries("name", "{name}").Methods(http.MethodGet)
	r.HandleFunc("/api/planets", planetListHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/planets/{id}", planetReadHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/planets/{id}", planetDeleteHandler).Methods(http.MethodDelete)

	http.Handle("/", r)

	if err := http.ListenAndServe(ApiEndpoint, nil); err != nil {
		fmt.Println(err)
	}
}
