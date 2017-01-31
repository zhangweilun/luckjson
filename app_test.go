package luckjson

import (


	"testing"
	"log"
)

/**
*
* @author willian
* @created 2017-01-31 20:21
* @email 18702515157@163.com
**/

func TestJs_Array(t *testing.T) {
	//imitate data, it can be download from other website.
	data := `{"name":"light","weigth":"maybe65kg","result":["light","fish","dylan"]}`
	json := NewJson(data)

	// mq is data structure map[string]interface{} so if you want to get value assert as need.
	get := json.Get("name")
	if m, ok := (get.data).(string); ok {
		log.Println(m)
	}

}