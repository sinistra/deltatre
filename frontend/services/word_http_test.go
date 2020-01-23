package services

import (
	"reflect"
	"testing"

	"github.com/sinistra/deltatre/frontend/models"
)

func TestWordService_AddWord(t *testing.T) {
	type args struct {
		newWord string
	}
	tests := []struct {
		name    string
		args    args
		want    models.Word
		want1   int
		wantErr bool
	}{
		{
			name: "AddTest1",
			args: args{newWord: "orange"},
			want: models.Word{
				Text:  "orange",
				Count: 0,
			},
			want1:   200,
			wantErr: false,
		}, {
			name: "AddTest2",
			args: args{newWord: "yes"},
			want: models.Word{
				Text:  "",
				Count: 0,
			},
			want1:   422,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := WordService{}
			got, got1, err := w.AddWord(tt.args.newWord)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddWord() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("AddWord() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWordService_GetWord(t *testing.T) {
	type args struct {
		search string
	}
	tests := []struct {
		name    string
		args    args
		want    models.Word
		want1   int
		wantErr bool
	}{
		{
			name: "GetAWord",
			args: args{search: "yes"},
			want: models.Word{
				Text:  "yes",
				Count: 1,
			},
			want1:   200,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := WordService{}
			got, got1, err := w.GetWord(tt.args.search)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWord() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetWord() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestWordService_TopWords(t *testing.T) {
	tests := []struct {
		name    string
		want    []models.Word
		want1   int
		wantErr bool
	}{
		{
			name: "GetWords",
			want: []models.Word{
				{"yes", 1},
				{"hello", 0},
				{"goodbye", 0},
				{"simple", 0},
				{"list", 0},
			},
			want1:   200,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := WordService{}
			got, got1, err := w.TopWords()
			if (err != nil) != tt.wantErr {
				t.Errorf("TopWords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopWords() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("TopWords() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_httpRequest(t *testing.T) {
	type args struct {
		method string
		url    string
		body   []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		want1   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := httpRequest(tt.args.method, tt.args.url, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("httpRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("httpRequest() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("httpRequest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
