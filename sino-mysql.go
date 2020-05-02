package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"strings"
)

func main() {
	var r, tableName,tableDetail string
	fmt.Println("功能列表: ​1.建表  2.帮助  3.关于 0.退出\n")
	for {
		fmt.Print("请选择:")
		fmt.Scan(&r)
		switch r {
		case "1":
			//fmt.Println( "1.建表\n");
			fmt.Print("\n请输入表名:​");
			fmt.Scan(&tableName);
			fmt.Print("\n请输入表名描述(如:字典表):​");
			fmt.Scan(&tableDetail);
			var sql,column1,column2,column3,column4,column5,column6,key string
			tableName = strings.TrimSpace(tableName)
			tableDetail = strings.TrimSpace(tableDetail)
			sql = "DROP TABLE IF EXISTS " + tableName + ";CREATE TABLE "+tableName+" ("
			xlsx, err := excelize.OpenFile("D:/sql.xlsx")
			if err != nil {
				fmt.Println("请建立excel表格(D:/sql.xlsx)")
				return
			}
			rows := xlsx.GetRows("Sheet1")
			var list []string
			for i, row := range rows {
				// 去掉第一行，第一行是表头
				if i == 0 {
					continue
				}
				for j, colCell := range row {
					// 第1列
					if j == 0 {
						column1 =  strings.ToUpper(strings.TrimSpace(colCell))
					}
					// 第2列
					if j == 1 {
						column2 = strings.ToUpper(strings.TrimSpace(colCell))
					}
					// 第3列
					if j == 2 {
						column3 = strings.ToUpper(strings.TrimSpace(colCell))
					}
					// 第4列
					if j == 3 {
						column4 =  strings.ToUpper(strings.TrimSpace(colCell))
					}
					// 第5列
					if j == 4 {
						column5 =  strings.ToUpper(strings.TrimSpace(colCell))
					}
					// 第6列
					if j == 5 {
						column6 =  strings.ToUpper(strings.TrimSpace(colCell))
					}
					// 第7列
					if j == 6 {
						//column7 = colCell
					}
				}
				sql += "`" + column1 + "` "
				sql += column3
				if column6 == "主键索引" {
					key = strings.TrimSpace(column1)
					sql +=" NOT NULL AUTO_INCREMENT COMMENT '" + column2 + "',"
				} else {
					if column6 == "普通索引" {
						list = append(list, column1)
					}
					if column4 == "是" {
						if column5 == "Y" {
							sql += " NOT NULL DEFAULT 1 COMMENT '" + column2 + "',"
						} else {
							sql += " DEFAULT 1 COMMENT '" + column2 + "',"
						}
					} else {
						if column5 == "Y" {
							sql += " NOT NULL COMMENT '" + column2 + "',"
						} else {
							sql += " COMMENT '" + column2 + "',"
						}
					}
				}
			}
			if len(list) > 0 {
				sql += " PRIMARY KEY (`" + key + "`) USING BTREE,"
				for i := 0; i < len(list); i++ {
					if i == len(list)-1 {
						sql += "INDEX `idx_" + tableName + "_" + strings.ToLower(list[i]) + "`(`" + list[i] + "`) USING BTREE)"
					} else {
						sql += "INDEX `idx_" + tableName + "_" + strings.ToLower(list[i]) + "`(`" + list[i] + "`) USING BTREE,"
					}
				}

				sql += "ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8"
			} else {
				sql += " PRIMARY KEY (`" + key + "`) USING BTREE) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8"
			}
			if tableDetail != ""{
				sql += " COMMENT = '"+tableDetail+"'"
			}
			sql += " ROW_FORMAT = Compact"
			fmt.Println("\n请复制下面SQL建表语句:\n\n")
			fmt.Println(sql)
			fmt.Println("\n\n")
			fmt.Println("功能列表: ​1.建表  2.帮助  3.关于 0.退出\n")
		case "2":
			//fmt.Println( "2.帮助\n")
			fmt.Println("\n1.字段及类型 需严格遵守mysql规范，此SQL语句不会进行检验。\n2.针对excel建表，默认文件路径D:/sql.xlsx\n3.使用技巧:将表格从confluence复制到excel中")
			fmt.Println("备注:未添加唯一索引 如有需要建表后自行添加\n")
			fmt.Println("+---------+-----------+------------+--------+----------+---------+--------------+\n| 字段名  |  注 释    |    类型    | 默认值 | 是否必填 |  索引   |  备注        |\n+---------+-----------+------------+--------+----------+---------+--------------+\n| id      |  主 键    |  int(11)   |        |     Y    | 主键索引|              |\n+---------+-----------+------------+--------+----------+---------+--------------+\n| name    |  姓 名    | varchar(8) |        |     Y    | 普通索引|   姓名       |\n+---------+-----------+------------+--------+----------+---------+--------------+\n|is_effect|  是否有效 |  int(3)    |   是   |          |         |              |\n+---------+-----------+------------+--------+----------+---------+--------------+")
			fmt.Println("\n功能列表: ​1.建表  2.帮助  3.关于 0.退出\n")
		case "3":
			//fmt.Println("3.关于\n")
			fmt.Println("\n##作者:李爱朋\n\n##邮箱:228171835@qq.com\n\n##微信:godream_cn\n")
			fmt.Println("功能列表: ​1.建表  2.帮助  3.关于 0.退出\n")
		case "0":
			fmt.Println("0.退出")
			break
		default:
			fmt.Println("\n输入命令不正确。。。\n")
			fmt.Println("功能列表: ​1.建表  2.帮助  3.关于 0.退出\n")
		}
	}
}
