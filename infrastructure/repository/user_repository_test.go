package repository

import (
	"reflect"
	"testing"

	"github.com/taisa831/go-ddd/domain/model"
	"gorm.io/gorm"
)

func Test_dbRepository_FindUsers(t *testing.T) {
	type fields struct {
		db *gorm.DB
	}
	tests := []struct {
		name          string
		fields        fields
		want          []*model.User
		wantErr       bool
		insertFixture func(db *gorm.DB)
	}{
		{
			name: "FindUsers",
			fields: fields{
				db: rdb,
			},
			want: []*model.User{
				{
					ID:   "u-1",
					Name: "name-1",
				},
			},
			wantErr: false,
			insertFixture: func(db *gorm.DB) {
				u := model.User{
					ID:   "u-1",
					Name: "name-1",
				}
				if err := db.Create(&u).Error; err != nil {
					t.Fatal(err)
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer truncate(tt.fields.db)
			tt.insertFixture(tt.fields.db)
			r := &dbRepository{
				db: tt.fields.db,
			}
			got, err := r.FindUsers()
			if (err != nil) != tt.wantErr {
				t.Errorf("dbRepository.FindUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dbRepository.FindUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}
