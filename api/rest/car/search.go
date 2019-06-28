package car

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/md3756/ustart_tutorial/car/carpb"
)

// Search ...
func (rapi *RESTAPI) Search(w http.ResponseWriter, req *http.Request) {
	regCtx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	//NOTE: this method of retrieving data from a REST request should only be used for GET requests
	//later on you will be shown the difference between what is and when you should use GET or POST
	req.ParseForm()
	fn := req.Form.Get("FirstName")
	ln := req.Form.Get("LastName")
	dob := req.Form.Get("DOB")
	scroll := req.Form.Get("Scroll")

	lookReq := &carpb.SearchRequest{
		FirstName: fn,
		LastName:  ln,
		DOB:       dob,
		Scroll:    scroll,
	}

	ret := make(map[string]interface{})
	//key:type string; value: interface- meaning you can put anything into it

	resp, err := rapi.car.Search(regCtx, lookReq)
	if resp != nil {
		ret["response"] = resp
	} else {
		ret["response"] = ""
	}
	if err != nil {
		ret["error"] = err.Error()
	} else {
		ret["error"] = ""
	}

	data, err := json.Marshal(ret)
	if err != nil {
		logger.Panic(err)
	}

	fmt.Fprintln(w, string(data))
}
