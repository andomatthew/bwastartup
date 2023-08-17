package user

import "gorm.io/gorm"

// function yang boleh digunakan oleh instance db (struct repository)
// untuk mengakses database
type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(id int) (User, error)
	Update(user User) (User, error)
}

// struct ini dibutuhkan untuk membuat instance database khusus untuk file repository
type repository struct { // membuat instance db
	db *gorm.DB
}


func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}


// semua function dibawah adalah milik "repository" dengan catatan, 
// nama function, parameter dan return value nya harus sesuai dengan yang di deklarasi di dalam interface Repository
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindById(id int) (User, error) {
	var user User
	err := r.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}