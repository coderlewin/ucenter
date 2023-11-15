package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

const DSN = "root:12345678@tcp(127.0.0.1:3306)/coderlu_uc?charset=utf8mb4&parseTime=True"

func main() {
	autoGen()
}

func autoGen() {
	db, err := gorm.Open(mysql.Open(DSN))
	if err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		ModelPkgPath: "../../internal/infrastructure/entity",
	})

	g.UseDB(db)

	g.ApplyBasic(
		// 逻辑删除
		g.GenerateAllTable(
			gen.FieldGORMTag("is_delete", func(tag field.GormTag) field.GormTag {
				tag.Append("softDelete", "flag")
				return tag
			}),
			gen.FieldType("is_delete", "soft_delete.DeletedAt"),
		)...,
	)

	g.Execute()
}
