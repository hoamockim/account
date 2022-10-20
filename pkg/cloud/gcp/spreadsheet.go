package gcp

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

//TODO: calculate range

type ISpreadSheetService interface {
	New(title, sheetTile string) (string, string, error)
	AddSheet(spreadSheetId, sheetTitle string) error
	AddRow(spreadSheetId, rowRange string, data [][]interface{}, isOverWrite bool) ([]byte, error)
	UpdateRow(spreadSheetId, rowRange string, data [][]interface{}) ([]byte, error)
	GetContent(spreadSheetId string, dataRange string) ([]byte, error)
	Download(spreadId string, dataRange string) error
}

type SpreadSheetService struct {
	service *sheets.Service
	ctx     context.Context
}

type SheetDataRage struct {
	Title string
	From  string
	To    string
}

func NewGoogleSheetClient(secret string) ISpreadSheetService {
	data, err := ioutil.ReadFile(secret)
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	conf, err := google.JWTConfigFromJSON(data, sheets.SpreadsheetsScope)
	ctx := context.Background()
	client := conf.Client(ctx)
	sheet, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {

	}
	return &SpreadSheetService{sheet, ctx}
}

func (srv *SpreadSheetService) New(title, sheetTile string) (string, string, error) {
	spr := sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: title,
		},
	}

	if strings.Trim(sheetTile, " ") != "" {
		newSheet := sheets.Sheet{
			Properties: &sheets.SheetProperties{
				Title: sheetTile,
			},
		}
		spr.Sheets = append(spr.Sheets, &newSheet)
	}

	resp, err := srv.service.Spreadsheets.Create(&spr).Context(srv.ctx).Do()
	if err != nil {
		return "", "", err
	}
	return resp.SpreadsheetId, resp.SpreadsheetUrl, nil
}
func (srv *SpreadSheetService) AddRow(spreadSheetId, rowRange string, data [][]interface{}, isOverWrite bool) ([]byte, error) {
	valueInputOption := "RAW"
	insertDataOption := "INSERT_ROWS"
	if isOverWrite {
		insertDataOption = "OVERWRITE"
	}
	valueRow := &sheets.ValueRange{
		Range:  rowRange,
		Values: data,
	}
	resp, err := srv.service.Spreadsheets.Values.Append(spreadSheetId, rowRange, valueRow).
		ValueInputOption(valueInputOption).
		InsertDataOption(insertDataOption).Context(srv.ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp.MarshalJSON()
}
func (srv *SpreadSheetService) UpdateRow(spreadSheetId, rowRange string, data [][]interface{}) ([]byte, error) {
	valueInputOption := "RAW"
	valueRow := &sheets.ValueRange{
		Range:  rowRange,
		Values: data,
	}
	resp, err := srv.service.Spreadsheets.Values.Update(spreadSheetId, rowRange, valueRow).
		ValueInputOption(valueInputOption).
		Context(srv.ctx).Do()
	if err != nil {
		return nil, err
	}
	return resp.MarshalJSON()
}

func (srv *SpreadSheetService) AddSheet(spreadSheetId, sheetTitle string) error {
	updateRequest := sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			{
				AddSheet: &sheets.AddSheetRequest{
					Properties: &sheets.SheetProperties{
						Title: sheetTitle,
					},
				},
			},
		},
	}
	_, err := srv.service.Spreadsheets.BatchUpdate(spreadSheetId, &updateRequest).Context(srv.ctx).Do()
	return err
}

func (srv *SpreadSheetService) GetContent(spreadSheetId string, dataRange string) ([]byte, error) {
	resp, err := srv.service.Spreadsheets.Values.Get(spreadSheetId, dataRange).Context(srv.ctx).Do()
	if err != nil {
		return nil, err
	}
	return json.Marshal(resp.Values)
}

func (srv *SpreadSheetService) Download(spreadId string, dataRange string) error {
	return nil
}
