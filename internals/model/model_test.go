/*
 * File: /internals/model/model_test.go                                        *
 * Project: blog-service                                                       *
 * Created At: Friday, 2022/06/3 , 15:59:07                                    *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/3 , 16:40:23                                 *
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
				"modified_by": "xAtzU",
				"state": 0,
				"name": "Ukqcjl Ljcwlcxso Wjkwa Hmwcdlf Svfxqjcx Huwqpu Cqm"
			}`,
			Model{
				ModifiedBy: "xAtzU",
				// State:      0,
				// Name:       "Ukqcjl Ljcwlcxso Wjkwa Hmwcdlf Svfxqjcx Huwqpu Cqm",
			},
		},
		{
			`{
				"modified_by": "M%a",
				"state": 0,
				"name": "Ukqcjl Ljcwlcxso Wjkwa Hmwcdlf Svfxqjcx Huwqpu Cqm"
			}`,
			Model{
				ModifiedBy: "M%a",
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
