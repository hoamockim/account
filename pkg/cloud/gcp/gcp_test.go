package gcp

import (
	"fmt"
	"testing"
)

func Test_NewSpreadSheet(t *testing.T) {
	srv := NewGoogleSheetClient("tipee-cart-331216-bcb677088ed7.json")
	spreadId, url, err := srv.New("test ", "test 1")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("spreadId: ", spreadId)
	fmt.Println("url: ", url)
}

func Test_GetContent(t *testing.T) {
	srv := NewGoogleSheetClient("tipee-cart-331216-bcb677088ed7.json")
	datas, err := srv.GetContent("11UVuSs7TOtjCRqOyrdeP2YriF7S-tIqNVbaBpnku4uo", "abb!A1:B3")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("data: ", string(datas))
}

func Test_AddSheet(t *testing.T) {
	srv := NewGoogleSheetClient("tipee-cart-331216-bcb677088ed7.json")
	err := srv.AddSheet("11UVuSs7TOtjCRqOyrdeP2YriF7S-tIqNVbaBpnku4uo", "ddd")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_Addrow(t *testing.T) {
	var data [][]interface{}
	var row1 = []interface{}{"f9", "5"}
	var row2 = []interface{}{"f10", "11"}
	data = append(data, row1)
	data = append(data, row2)
	srv := NewGoogleSheetClient("tipee-cart-331216-bcb677088ed7.json")
	resp, err := srv.AddRow("11UVuSs7TOtjCRqOyrdeP2YriF7S-tIqNVbaBpnku4uo", "abb!A1:B1", data, false)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Data: ", string(resp))
}

func Test_UpdateRow(t *testing.T) {
	var data [][]interface{}
	var row1 = []interface{}{"f22", "3"}
	data = append(data, row1)
	srv := NewGoogleSheetClient("tipee-cart-331216-bcb677088ed7.json")
	resp, err := srv.UpdateRow("11UVuSs7TOtjCRqOyrdeP2YriF7S-tIqNVbaBpnku4uo", "abb!A10:B10", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("Data: ", string(resp))
}
