package data

import (
	"context"
	"database/sql"
	"iris-server/models"
	"reflect"
	"strconv"
)

type Product interface {
	Conn() error
	Insert(product *models.Product) (int64, error)
	Delete(int64) error
	Update(product *models.Product) error
	Select(int64) (*models.Product, error)
	SelectAll() ([]*models.Product, error)
}

type ProductManager struct {
	mysql *sql.DB
	table string
}

func NewProductManager() Product {
	return &ProductManager{}
}

func (p *ProductManager) Conn() error {
	_, err := p.mysql.Conn(context.Background())

	return err
}

func (p *ProductManager) Insert(product *models.Product) (int64, error) {
	if err := p.Conn(); err != nil {
		return 0, err
	}

	sqlStr := "INSERT product SET name=?, number=?, image=?, url=?"
	stmt, err := p.mysql.Prepare(sqlStr)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(product.Name, product.Number, product.Image, product.URL)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (p *ProductManager) Delete(id int64) error {
	if err := p.Conn(); err != nil {
		return err
	}

	sqlStr := "DELETE FROM product WHERE ID = ?"
	stmt, err := p.mysql.Prepare(sqlStr)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(id)

	return err
}

func (p *ProductManager) Update(product *models.Product) error {
	if err := p.Conn(); err != nil {
		return err
	}

	sqlStr := "UPDATE FROM product where Name = ?, Number = ?, Image = ?, URL = ? WHERE ID = ?"
	stmt, err := p.mysql.Prepare(sqlStr)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(product.Name, product.Number, product.Image, product.URL, strconv.FormatInt(product.ID, 10))

	return err
}

func (p *ProductManager) Select(id int64) (*models.Product, error) {
	if err := p.Conn(); err != nil {
		return nil, err
	}

	sqlStr := "SELECT * FROM " + p.table + " WHERE ID =" + strconv.FormatInt(id, 10)
	_, err := p.mysql.Query(sqlStr)
	if err != nil {
		return nil, err
	}

	return nil, err
}

func (p *ProductManager) SelectAll() ([]*models.Product, error) {
	//TODO implement me
	panic("implement me")
}

func mappingToStructByTagSql(mapping map[string]string, obj interface{}) (*models.Product, error) {
	objValue := reflect.ValueOf(obj).Elem()
	for i := 0; i < objValue.NumField(); i++ {
		// 获取实体上每个字段的sql标签值，其对应了数据表返回的数据
		tag := objValue.Type().Field(i).Tag.Get("sql")
		// 根据标签获取数据表中的具体数据
		val := mapping[tag]
		// 获取实体上的字段名
		field := objValue.Type().Field(i).Name
		// 创建实例
		structType := objValue.Field(i).Type()
		if structType == reflect.Struct {

		}
	}
}
