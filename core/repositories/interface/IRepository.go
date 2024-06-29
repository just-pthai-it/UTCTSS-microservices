package _interface

import _interface "TSS-microservices/core/models/interface"

type IRepository interface {
	Create(model *_interface.IModel) error
	GetById(model *_interface.IModel, id any) error
	//GetByIds() error
	UpdateById(model *_interface.IModel, id any) error
	DeleteById(model *_interface.IModel, id any) error
}
