/*
 * File: /internals/model/model_test.go                                        *
 * Project: blog-service                                                       *
 * Created At: Friday, 2022/06/3 , 15:59:07                                    *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Tuesday, 2022/06/21 , 15:08:40                               *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package model

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestFieldTag(t *testing.T) {
	tests := []struct {
		jsonString string
		res        Model
	}{
		{
			`{
				"updated_by": "xAtzU",
				"state": 0,
				"name": "Ukqcjl Ljcwlcxso Wjkwa Hmwcdlf Svfxqjcx Huwqpu Cqm"
			}`,
			Model{
				UpdatedBy: "xAtzU",
				// State:      0,
				// Name:       "Ukqcjl Ljcwlcxso Wjkwa Hmwcdlf Svfxqjcx Huwqpu Cqm",
			},
		},
		{
			`{
				"updated_by": "M%a",
				"state": 0,
				"name": "Ukqcjl Ljcwlcxso Wjkwa Hmwcdlf Svfxqjcx Huwqpu Cqm"
			}`,
			Model{
				UpdatedBy: "M%a",
				// State:      0,
				// Name:       "Ukqcjl Ljcwlcxso Wjkwa Hmwcdlf Svfxqjcx Huwqpu Cqm",
			},
		},
	}
	for _, test := range tests {
		r := strings.NewReader(test.jsonString)
		var m Model
		err := json.NewDecoder(r).Decode(&m)
		if err != nil {
			t.Errorf("invalid json string")
		}

		if !reflect.DeepEqual(m, test.res) {
			t.Logf("%#v", m)
			t.Logf("%#v", test.res)
			t.Errorf("json decode failed")
		}
	}

}
