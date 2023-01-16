package models

type UserViewModel struct {
	Page  Page
	Users []User
}

//UserViewModel tipi üzerinden hem sayfa bilgilerini hem de kullanıcı bilgilerini tutuyoruz.
//Bu sayede kullanıcı bilgilerini ve sayfa bilgilerini tek bir struct üzerinden yönetebiliyoruz.
