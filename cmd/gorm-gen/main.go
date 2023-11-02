package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"os"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./internal/dao",
		ModelPkgPath:      "./internal/entity",
		WithUnitTest:      false,
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldWithIndexTag: true,
	})
	db, _ := gorm.Open(mysql.Open(os.Getenv("DB_DSN")))
	g.UseDB(db)

	g.ApplyBasic(
		g.GenerateAllTable()...,
	)
	g.Execute()
}