package repository

import (
	"errors"
	"go-gorm/model"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *model.Customer) error
	Update(customer *model.Customer, by map[string]interface{}) error
	Update2(customer *model.Customer, by model.Customer) error
	Delete(customer *model.Customer) error
	Delete2(id string) error
	FindById(id string) (model.Customer, error)
	FindFirstBy(by map[string]interface{}) (model.Customer, error)   // where column = ? limit 1
	FindAllBy(by map[string]interface{}) ([]model.Customer, error)   //where column = ?
	FindBy(by string, vals ...interface{}) ([]model.Customer, error) //where column like ?
	BaseRepositoryAggregation
	BaseRepositoryPaging
}

type customerRepository struct {
	db *gorm.DB
}

func (c *customerRepository) Create(customer *model.Customer) error {
	// sudah otomatis deteksi gorm nya kalo fungsi di bawah adalah insert
	// SQL Builder
	result := c.db.Create(customer).Error
	return result
}

func (c *customerRepository) Update(customer *model.Customer, by map[string]interface{}) error {
	result := c.db.Model(customer).Updates(by).Error
	return result
}

func (c *customerRepository) Update2(customer *model.Customer, by model.Customer) error {
	result := c.db.Model(customer).Updates(by).Error
	return result
}

func (c *customerRepository) Delete(customer *model.Customer) error {
	result := c.db.Delete(&model.Customer{}, customer).Error
	return result
}

func (c *customerRepository) Delete2(id string) error {
	result := c.db.Delete(&model.Customer{}, id).Error
	return result
}

func (c *customerRepository) FindById(id string) (model.Customer, error) {
	var customer model.Customer
	result := c.db.Unscoped().First(&customer, "id = ?", id)
	// result := c.db.Unscoped().First(&customer, "id = ?",id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}

func (c *customerRepository) FindFirstBy(by map[string]interface{}) (model.Customer, error) {
	var customer model.Customer
	result := c.db.Where(by).First(&customer)
	// result := c.db.Unscoped().First(&customer, "id = ?",id)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customer, nil
		} else {
			return customer, err
		}
	}
	return customer, nil
}
func (c *customerRepository) FindAllBy(by map[string]interface{}) ([]model.Customer, error) {
	var customers []model.Customer
	result := c.db.Where(by).Find(&customers)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customers, nil
		} else {
			return customers, err
		}
	}
	return customers, nil
}

func (c *customerRepository) FindBy(by string, vals ...interface{}) ([]model.Customer, error) {
	var customers []model.Customer
	result := c.db.Unscoped().Where(by, vals...).Find(&customers)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customers, nil
		} else {
			return customers, err
		}
	}
	return customers, nil
}


func (c *customerRepository) Count(groupBy string) (int, error) {
	var total int
	result := c.db.Model(&model.Customer{}).Select("count(*)").Group(groupBy).First(&total)
	if err := result.Error; err != nil {
		return 0, nil
	}
	return total, nil
}
func (c *customerRepository) GroupBy(result interface{}, selectedBy string, whereBy map[string]interface{}, groupBy string) error {
	res := c.db.Model(&model.Customer{}).Select("count(*)").Where(whereBy).Group(groupBy).Find(result)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		} else {
			return err
		}
	}
	return nil
}
func (c *customerRepository) Paging(page int, itemPerPage int) (interface{}, error) {
	var customers []model.Customer
	offset := itemPerPage * (page - 1)
	result := c.db.Order("created_at").Limit(itemPerPage).Offset(offset).Find(&customers)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return customers, nil
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	repo := new(customerRepository)
	repo.db = db
	return repo
}
