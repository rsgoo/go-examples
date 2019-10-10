package unit

import (
	"testing"
	"fmt"
)

func TestEncodePersonToJsonFile(t *testing.T) {
	t.Log("start to test func TestEncodePersonToJsonFile....")
	fileName := "TestEncodePersonToJsonFile.json"
	p := Person{
		Name:   "Tom",
		Age:    2,
		Rmb:    268000,
		Gender: "Male",
		Hobby:  []string{"eat fish", "play with tom"},
	}
	ok := EncodePersonToJsonFile(&p, fileName)

	if ok {
		dPerson, err := DecodeJsonFileToPerson(fileName)
		if err != nil {
			fmt.Println("decode json err ", err)
			return
		}

		if dPerson.Name == p.Name && dPerson.Age == p.Age && dPerson.Rmb == p.Rmb && dPerson.Gender == p.Gender {
			t.Log("TestEncodePersonToJsonFile test success")
		} else {
			t.Fatal("编码数据不正确")
		}
	} else {
		t.Fatal("编码失败")
	}

}
