package repository

import (
	"sync"
	"testing"

	"gorm.io/gorm"
)

func TestOpenDB(t *testing.T) {
	type args struct {
		count int
	}
	tests := []struct {
		name    string
		want    *gorm.DB
		wantErr bool
		args args
	}{
		{
			name: "接続上限の時",
			args: args{
				count: 151,
			},
		},
		{
			name: "接続上限を超えた時",
			args: args{
				count: 152,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := OpenDB()
			if err != nil {
				panic(err)
			}

			sqlDB, err := db.DB()
			sqlDB.SetMaxOpenConns(151)
			if err != nil {
				panic(err)
			}

			wg := &sync.WaitGroup{}
            for index := 0; index < tt.args.count; index++ {
                go func() {
                    wg.Add(1)
                    defer wg.Done()
                    if err := db.Exec("select sleep(5)").Error; err != nil {
                        t.Errorf("%v\n", err)
                    }
                }()
            }
            wg.Wait()

			sqlDB.Close()
		})
	}
}