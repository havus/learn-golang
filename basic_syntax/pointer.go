package main

import "fmt"

type Identity struct {
	FirstName, LastName, Address string
	Age int
	IsMarried bool
}

func (identity *Identity) changeFirstName(newFirstName string) {
	identity.FirstName = newFirstName
}

func main() {
	// https://app.diagrams.net/#G17YWt5h3tSluXloSy9W0x_VLB3VN243jF

	john := Identity{
		FirstName: 	"John",
		LastName: 	"Doe",
		Address:		"Jakarta",
		Age: 				20,
		IsMarried: 	false,
	}

	stanger := new(Identity)

	budi 						:= john
	budi.FirstName 	= "Budi"
	budi.LastName 	= "Slamet"

	fmt.Println(john)
	fmt.Println(budi)
	fmt.Println("\n=== setelah 2 taun john menikah duluan dan budi pindah tempat tinggal ===\n")

	john.Age = 22
	john.IsMarried = true
	budi.Address = "Kota banget euy!"

	fmt.Println("john", john)
	fmt.Println("budi", budi)

	
	

	fmt.Println("\n=== john punya anak 2, ngaku2 marga Hapean ===\n")

	santoso := &john // var santoso *Identity = &Identity
	valarie := &john // var santoso *Identity = &Identity
	john.LastName = "Hapean"
	santoso.Age = 5

	fmt.Println("john", john)
	fmt.Println("budi", budi)
	fmt.Println("santoso", santoso)
	fmt.Println("valarie", valarie)




	fmt.Println("\n=== john ganti nama anaknya santoso ===\n")

	santoso = &Identity{ "Santoso", "Hapean", "Desa sebelah", 5, false }

	fmt.Println("john", john)
	fmt.Println("budi", budi)
	fmt.Println("santoso", santoso)
	fmt.Println("valarie", valarie)



	
	fmt.Println("\n=== valarie cakep, akhirnya orang2 iktan manggil bapaknya dengan identity valarie ===\n")
	*valarie = Identity{ "Valarie", "Hapean", "Desa sebelah", 5, false }

	fmt.Println("john", john)
	fmt.Println("budi", budi)
	fmt.Println("santoso", santoso)
	fmt.Println("valarie", valarie)




	fmt.Println("\n=== santoso pinter TAPI bapaknya ga dipanggil dengan identity santoso ===\n")

	*santoso = Identity{ "Santoso", "Hapean", "Desa sebelah", 10, false }

	fmt.Println("stanger", stanger)
	fmt.Println("john", john)
	fmt.Println("budi", budi)
	fmt.Println("santoso", santoso)
	fmt.Println("valarie", valarie)

	// Pointer Method
	simbada := new(Identity)
	simbada.FirstName = "Simbada"
	simbada.changeFirstName("Simbada Changed")

	fmt.Println(simbada.FirstName)
}
