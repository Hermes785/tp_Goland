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
	for {
		var reponse int
		fmt.Println("Menu:")
		fmt.Println("--> Que voulez vous faire ? choisissez un numero et apuyer sur la touche entrer")
		fmt.Println("1- Ajouter un employé")
		fmt.Println("2- Afficher la liste des employé")
		fmt.Println("3- Modifier un employé")
		fmt.Println("4- suprimer un employé")
		fmt.Println("5- Exporté  la liste d'employé dans un fichier excel")
		fmt.Println("6- Se connecter à un serveur ftp")
		fmt.Println("7- Lancer un serveur http (avec une page web simple connecté à une route")
		fmt.Println("8- Se connecter à l'une de vos VM en ssh")
		fmt.Println("9- Ouvrir l'interface graphique")
		fmt.Println("10- Quitter sur l'application")

		fmt.Scan("%d", reponse)
		fmt.Scan(&reponse)

		switch reponse {
		case 1:

			fmt.Println(strings.Repeat("-", 45))
			println("vous avez choisi 1 - Ajouter un employé")
			fmt.Println(strings.Repeat("-", 45))

			exercice2.AddEmployee()

		case 2:
			fmt.Println(strings.Repeat("-", 55))
			println("vous avez choisi 2 - Afficher la liste des employés")
			fmt.Println(strings.Repeat("-", 55))

			exercice2.ListAllEmployees()

		case 3:
			fmt.Println(strings.Repeat("-", 45))
			println("vous avez choisi 3-Modifier un employés")
			fmt.Println(strings.Repeat("-", 45))

			exercice2.UpdatEmployee()
		case 4:
			fmt.Println(strings.Repeat("-", 45))
			println("vous avez choisi 4- suprimer un employés")
			fmt.Println(strings.Repeat("-", 45))

			exercice2.DeleteEmployee()

		case 5:
			fmt.Println(strings.Repeat("-", 75))
			println("vous avez choisi 5 - Exporté  la liste d'employé dans un fichier excel")
			fmt.Println(strings.Repeat("-", 75))

			exercice2.ExporteExcel()

		case 6:
			fmt.Println(strings.Repeat("-", 55))
			println("vous avez choisi 6- Se connecter à un serveur ftp")
			fmt.Println(strings.Repeat("-", 55))

			exercice2.ServerFTP()

		case 7:
			fmt.Println(strings.Repeat("-", 85))
			println("vous avez choisi Lancer un serveur http (avec une page web simple connecté à une route")
			fmt.Println(strings.Repeat("-", 85))

			exercice2.ServerHTTP()

		case 8:
			fmt.Println(strings.Repeat("-", 55))
			println("vous avez choisi 8-  Se connecter à l'une de vos VM en ssh")
			fmt.Println(strings.Repeat("-", 55))

			exercice2.ServerSSh()

		case 9:
			fmt.Println(strings.Repeat("-", 55))
			println("vous avez choisi 9-  Ouvrir l'interface graphique")
			fmt.Println(strings.Repeat("-", 55))

			exercice2.Intertface()

		case 10:

			fmt.Println(strings.Repeat("-", 55))
			println("vous avez choisi 10- Quitter sur l'application")
			fmt.Println(strings.Repeat("-", 55))

			var rep2 int
			fmt.Println("  --> Voulez vous vraiment quitter  l'application ?")
			fmt.Println("1- Oui")
			fmt.Println("2- Non")

			fmt.Scan(&rep2)

			switch rep2 {

			case 1:
				fmt.Println(strings.Repeat("-", 10))
				fmt.Println("Good bye")
				fmt.Println(strings.Repeat("-", 10))
				return

			case 2:
				fmt.Println(" ---> Vous avez chois 2")
				main()

			}

		default:
			fmt.Println(" --->Choix invalide veullez ressayer svp!")
		}

	}
}
