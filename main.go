package main

/*
j'ai separé mon code en 3 fichiers
 1 pour le main
1 autre pour tout ce qui crud sur le l'employé
 et 1 dernié pour la connexio a la bd
*/

import (
	Employes "employes/Gestion-Employes"
	"fmt"
	"strings"
)

func main() {
	Employes.ConnectDB()
	for {
		var reponse int
		fmt.Println("Menu:")
		fmt.Println("--> Que voulez vous faire ? Choisissez un numero et appuyer sur la touche entrer")
		fmt.Println("1- Ajouter un employé")
		fmt.Println("2- Modifier un employé")
		fmt.Println("3- suprimer un employé")
		fmt.Println("4- Afficher la liste des employés")
		fmt.Println("5- Exporter la liste des employés sous format Excel")
		fmt.Println("6- Se connecter à un serveur ftp")
		fmt.Println("7- Lancer un serveur http (avec une page web simple connecté à une route)")
		fmt.Println("8- Se connecter à l'une de vos VM en ssh")
		fmt.Println("9- Quitter sur l'application")

		fmt.Scan("%d", reponse)
		fmt.Scan(&reponse)

		switch reponse {
		case 1:

			fmt.Println(strings.Repeat("-", 45))
			println("vous avez choisi 1 - Ajouter un employé")
			fmt.Println(strings.Repeat("-", 45))

			Employes.AddEmployee()

		case 2:
			fmt.Println(strings.Repeat("-", 45))
			println("vous avez choisi 2-Modifier un employés")
			fmt.Println(strings.Repeat("-", 45))

			Employes.UpdatEmployee()

		case 3:
			fmt.Println(strings.Repeat("-", 45))
			println("vous avez choisi 3- suprimer un employés")
			fmt.Println(strings.Repeat("-", 45))

			Employes.DeleteEmployee()

		case 4:
			fmt.Println(strings.Repeat("-", 55))
			println("vous avez choisi 4 - Afficher la liste des employés")
			fmt.Println(strings.Repeat("-", 55))

			Employes.ListAllEmployees()

		case 5:
			fmt.Println(strings.Repeat("-", 75))
			println("vous avez choisi 5 - Exporté  la liste d'employé dans un fichier excel")
			fmt.Println(strings.Repeat("-", 75))

			Employes.ExporteExcel()

		case 6:
			fmt.Println(strings.Repeat("-", 55))
			println("vous avez choisi 6- Se connecter à un serveur ftp")
			fmt.Println(strings.Repeat("-", 55))

			Employes.ServerFTP()

		case 7:
			fmt.Println(strings.Repeat("-", 85))
			println("vous avez choisit 7- Lancer un serveur http (avec une page web simple connecté à une route")
			fmt.Println(strings.Repeat("-", 85))

			Employes.ServerHTTP()

		case 8:
			fmt.Println(strings.Repeat("-", 55))
			println("vous avez choisi 8-  Se connecter à l'une de vos VM en ssh")
			fmt.Println(strings.Repeat("-", 55))

			Employes.ServerSSh()

			/*
				case 9:
							fmt.Println(strings.Repeat("-", 55))
							println("vous avez choisi 9-  Ouvrir l'interface graphique")
							fmt.Println(strings.Repeat("-", 55))

							exercice2.Intertface()
			*/
		case 9:

			fmt.Println(strings.Repeat("-", 55))
			println("vous avez choisi 10- Quitter sur l'application")
			fmt.Println(strings.Repeat("-", 55))

			var rep3 int
			fmt.Println("  --> Voulez vous vraiment quitter  l'application ?")

			fmt.Println("1- Oui")
			fmt.Println("2- Non")

			fmt.Scan(&rep3)

			switch rep3 {

			case 1:
				fmt.Println(" ---> Vous avez choisit 1-Oui")
				fmt.Println(strings.Repeat("-", 10))
				fmt.Println("Good bye")
				fmt.Println(strings.Repeat("-", 10))
				return

			case 2:
				fmt.Println(" ---> Vous avez choisit 2-Non")
				main()

			}

		default:
			fmt.Println(" --->Choix invalide veullez ressayer svp!")
		}

	}
}
