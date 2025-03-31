package main

import "slices"

type Role string

const (
	Admin       Role = "Admin"
	Customer    Role = "Customer"
	Seller      Role = "Seller"
	AuthService Role = "AuthSerice"
)

type Permission int

const (
	CreateUser Permission = iota
	CreateProduct
)

var permissionTable = map[Role][]Permission{
	Admin:       {CreateUser, CreateProduct},
	Customer:    {},
	Seller:      {CreateProduct},
	AuthService: {CreateUser},
}

func HasPermission(role Role, perm Permission) bool {
	return slices.Contains(permissionTable[role], perm)
}
