package main

//j'ai separé mon code en 3 fichiers
// 1 pour le main
//1 autre pour tout ce qui crud sur le l'employé
// et 1 dernié pour la connexio a la bd

import (
	exercice2 "exercice2/exercic2"
	"fmt"
	"strings"
)

func main() {
	exercice2.ConnectDB()
	var reponse int 
	fmt.Println("Menu:")
	fmt.Println("Que voulez vous faire: choisissez un numero et apuyer sur la touche entrer")
	fmt.Println("1 - Ajouter un employé")
	fmt.Println("2 - Afficher la liste des employés")
	fmt.Println("3 - Modifier un employés")
	fmt.Println("4- suprimer un employés")
	fmt.Println("5 - Exporté  la liste d'employé dans un fichier excel")
	fmt.Println("6- Quitter sur l'application")

	fmt.Scan("%d",reponse)
	fmt.Scan(&reponse)

switch reponse {
	case 1:

		   fmt.Println(strings.Repeat("-", 25))
			println("vous avez choisi 1 - Ajouter un employé")
			fmt.Println(strings.Repeat("-", 25))

		    exercice2.AddEmployee()

 case 2:
fmt.Println(strings.Repeat("-", 25))	
println("vous avez choisi 2 - Afficher la liste des employés")
	fmt.Println(strings.Repeat("-", 25))

	exercice2.ListAllUser()

 case 3:   
 fmt.Println(strings.Repeat("-", 35))
 println("vous avez choisi 3-Modifier un employés")
 fmt.Println(strings.Repeat("-", 35))

	exercice2.UpdatEmployee()
 case 4 :
	fmt.Println(strings.Repeat("-", 35))
	println("vous avez choisi 4- suprimer un employés")
	fmt.Println(strings.Repeat("-", 35))
	exercice2.DeleteEmployee()

 case 5:
	fmt.Println(strings.Repeat("-", 35))
	println("vous avez choisi5 - Exporté  la liste d'employé dans un fichier excel")
	fmt.Println(strings.Repeat("-", 25))
	exercice2.ExporteExcel()

 case 6:
	fmt.Println(strings.Repeat("-", 35))
	println("vous avez choisi 6- Quitter sur l'application")
	fmt.Println(strings.Repeat("-", 35))
  fmt.Printf(" Vous avez quittez lapplication")
return
 default:
	fmt.Println("Choix invalide veullez ressayer svp!")
}
	//exercice2.Server()
}
