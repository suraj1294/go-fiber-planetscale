package repository

import (
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/suraj1294/go-fiber-planetscale/db"
	"github.com/suraj1294/go-fiber-planetscale/logger"
	"github.com/suraj1294/go-fiber-planetscale/utils"
)

type Product struct {
	Id    int64   `json:"id,omitempty" db:"id"  goqu:"skipinsert,skipupdate"`
	Name  *string `json:"name,omitempty" db:"name" `
	Price *int    `json:"price,omitempty" db:"price"`
}

type ProductRepository struct {
	sqlx  *sqlx.DB
	mysql *goqu.Database
}

func (p ProductRepository) GetAll() (*[]Product, *error) {

	err := p.sqlx.DB.Ping()

	if err != nil {
		logger.Error("failed to connect to DB" + err.Error())
		return nil, &err
	}

	ds, _, err := p.mysql.From("products").ToSQL()

	if err != nil {
		logger.Error("failed to generate query all products" + err.Error())
		return nil, &err
	}

	products := []Product{}
	err = p.sqlx.Select(&products, ds)
	if err != nil {
		logger.Error("failed to get products" + err.Error())
		return nil, &err
	}

	return &products, nil
}

func (p ProductRepository) GetById(id int) (*Product, *error) {

	ds, _, _ := p.mysql.From("products").Select("*").Where(goqu.C("id").Eq(id)).ToSQL()

	product := Product{}
	err := p.sqlx.Get(&product, ds)
	if err != nil {
		logger.Error("failed to get product" + err.Error())
		return nil, &err
	}

	return &product, nil
}

func (p ProductRepository) Add(newProduct *Product) (*Product, *error) {

	ds := p.mysql.Insert("products").Rows(Product{Name: newProduct.Name, Price: newProduct.Price})

	addQuery, _, _ := ds.ToSQL()

	res, err := p.sqlx.Exec(addQuery)
	if err != nil {
		logger.Error("(CreateProduct) db.Exec")
		return nil, &err
	}
	Id, err := res.LastInsertId()

	newProduct.Id = Id
	if err != nil {
		logger.Error("(CreateProduct) res.LastInsertId")
		return nil, &err
	}

	return newProduct, nil
}

func (p ProductRepository) Update(update *Product, id int) (*Product, *error) {

	output, _ := utils.MarshalRequest(update)

	ds := p.mysql.Update("products").Set(output).Where(goqu.Ex{"id": goqu.Op{"eq": id}})

	updateQuery, _, _ := ds.ToSQL()

	fmt.Println(updateQuery)

	if updateQuery != "" {
		_, err := p.sqlx.Exec(updateQuery)
		if err != nil {
			logger.Error("(UpdateProduct) db.Exec")
			return nil, &err
		}
	}

	updateId := int64(id)
	update.Id = updateId

	return update, nil
}

func (p ProductRepository) Delete(id int) error {
	ds := p.mysql.From("products").Delete().Where(goqu.Ex{"id": goqu.Op{"eq": id}})

	deleteQuery, _, _ := ds.ToSQL()

	fmt.Println(deleteQuery)

	_, err := p.sqlx.Exec(deleteQuery)

	return err
}

func GetProductRepository() *ProductRepository {

	dbCon := db.NewDatabaseConnection()

	return &ProductRepository{sqlx: dbCon, mysql: goqu.New("mysql", dbCon)}
}
