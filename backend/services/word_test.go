package services

import (
	"reflect"
	"testing"

	"sinistra/deltatre/backend/models"
)

func TestWordService_AddWord(t *testing.T) {
	type args struct {
		newWord string
	}
	tests := []struct {
		name    string
		args    args
		want    models.Word
		wantErr bool
	}{
		{
			name: "AddTest",
			args: args{newWord: "orange"},
			want: models.Word{
				Text:  "orange",
				Count: 0,
			},
			wantErr: false,
		}, {
			name: "AddTest2",
			args: args{newWord: "yes"},
			want: models.Word{
				Text:  "",
				Count: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := WordService{}
			got, err := w.AddWord(tt.args.newWord)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddWord() got = %v, want %v", got, tt.want)
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
		wantErr bool
	}{
		{
			name: "GetAWord",
			args: args{search: "yes"},
			want: models.Word{
				Text:  "yes",
				Count: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := WordService{}
			got, err := w.GetWord(tt.args.search)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWord() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordService_GetWords(t *testing.T) {
	tests := []struct {
		name    string
		want    []models.Word
		wantErr bool
	}{
		{
			name: "GetWords",
			want: []models.Word{
				{"hello", 0},
				{"goodbye", 0},
				{"simple", 0},
				{"list", 0},
				{"search", 0},
				{"filter", 0},
				{"yes", 0},
				{"no", 0},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := WordService{}
			got, err := w.GetWords()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWords() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWordService_TopWords(t *testing.T) {
	tests := []struct {
		name    string
		want    []models.Word
		wantErr bool
	}{
		{
			name: "TopWords",
			want: []models.Word{
				{"hello", 0},
				{"goodbye", 0},
				{"simple", 0},
				{"list", 0},
				{"search", 0},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := WordService{}
			got, err := w.TopWords()
			if (err != nil) != tt.wantErr {
				t.Errorf("TopWords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopWords() got = %v, want %v", got, tt.want)
			}
		})
	}
}
