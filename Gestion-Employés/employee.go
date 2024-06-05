package exercice2

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/xuri/excelize/v2"
)

/*
pour éviter de devoir définir à chaque fois les variables dans chaques fonctions
j'ai creé une structure qui contiendra les variables nécéssaire pour notre entité employé

*/

type employee struct {
	id       int
	name     string
	prenom   string
	fonction string
	email    string
}

// cette fuction permet de valider si l'email que l'utilsateur entre est correct
func isValidEmail(email string) bool {
	// Expression régulière pour vérifier la présence de '.' et '@'
	pattern := `.*@.*\..*`
	match, _ := regexp.MatchString(pattern, email)
	return match
}

// celle ci valide si tout ce qui est nom,prenom et fonction entré par l'utilsateur est valide cete a dire
// ne contient pas de chiffr ou d'autres caracteres
func isValidString(champs string) bool {
	// Expression régulière pour valider l'adresse email
	pattern := `^[a-zA-Z]+$`
	match, _ := regexp.MatchString(pattern, champs)
	return match
}

var e employee

// fonction pour un  ajouté employés
func AddEmployee() {

	fmt.Println("-->Veuillez entrer les informations de l'employé")
	// Lecture et Validation des entrées utilisateurs pour linsertion
	fmt.Print("-->Entrer le nom  de l'employé: ")
	fmt.Scan(&e.name)
	for !isValidString(e.name) {
		fmt.Printf("le nom:%s n'est pas valide \n", e.name)
		fmt.Println(strings.Repeat("-", 75))
		fmt.Print("-->Entrer le nom  de l'employé: ")
		fmt.Scan(&e.name)

	}

	fmt.Print("-->Entrer le prénom de l'employé: ")
	fmt.Scan(&e.prenom)
	for !isValidString(e.prenom) {
		fmt.Printf("le prenom :%s n'est pas valide \n", e.prenom)
		fmt.Println(strings.Repeat("-", 75))
		fmt.Print("-->Entrer le prénom de l'employé: ")
		fmt.Scan(&e.prenom)

	}

	fmt.Print("-->Entrer la fonction de l'employé : ")
	fmt.Scan(&e.fonction)
	for !isValidString(e.fonction) {
		fmt.Printf("--la fonction:%s n'est pas valide \n", e.fonction)
		fmt.Println(strings.Repeat("-", 75))
		fmt.Print("-->Entrer la fonction de l'employé : ")
		fmt.Scan(&e.fonction)
	}

	fmt.Print("-->Entrer l'adresse email de l'employé: ")
	fmt.Scan(&e.email)
	for !isValidEmail(e.email) {
		fmt.Printf("l'adresse email: %s n'est pas valide \n", e.email)
		fmt.Println(strings.Repeat("-", 55))
		fmt.Print("-->Entrer l'adresse email de l'employé: ")
		fmt.Scan(&e.email)

	}

	// Préparer l'insertion des données

	insert, err := DB.Query("INSERT INTO employee (nom, prenom, fonction,email) VALUES (?, ?, ?,?)", e.name, e.prenom, e.fonction, e.email)
	if err != nil {
		fmt.Println("--> Une erreur est survenue :", err)
		return
	}
	defer insert.Close()
	fmt.Println(strings.Repeat("-", 75))
	fmt.Println("L'employé a été ajouté avec succès à la base de données.")
	fmt.Println(strings.Repeat("-", 75))

}

///***********************(- -)*************************
// fonction pour lister les  employés

func ListAllEmployees() {

	//var e employee
	fmt.Println("--> La liste des utilisateurs est :")

	results, err := DB.Query("SELECT * FROM employee")
	if err != nil {
		fmt.Println("-->Une erreur est survenue lors de la requête :", err)
		return
	}
	defer results.Close()

	idd := 0 // Initialise l'identifiant pour eviter que lorsque suprime un élément dans la table les id soit désordonne lors de l'affichage
	fmt.Println(strings.Repeat("-", 146))
	fmt.Printf("%-10s | %-30s | %-40s | %-40s| %-20s\n", "ID", "Nom", "Prénom", "Fonction", "Email")
	fmt.Println(strings.Repeat("-", 146))
	for results.Next() {
		idd++
		err = results.Scan(&e.id, &e.name, &e.prenom, &e.fonction, &e.email)
		if err != nil {
			fmt.Println("-->Une erreur est survenue lors de la lecture des résultats :", err)
			continue
		}

		fmt.Printf("%-10d | %-30s | %-40s | %-40s| %-20s\n", idd, e.name, e.prenom, e.fonction, e.email)

		// Crée les références des cellules pour chaque colonne

	}

}

// /***********************(- -)*************************
// fonction pour mofifier un employé
func UpdatEmployee() {
	//var e employee
	var email_entre string

	// Récupération de lemail
	fmt.Println("--> Pour modifier, il faudra entrer l'email de l'employé à modifier afin de vérifier s'il existe")

	fmt.Println("     --> Entrer un email de l'employé à modifier :")
	fmt.Scan(&email_entre)

	// Vérification de l'existance de  l'employé
	res, err := DB.Query("SELECT id, Nom, Prenom,Fonction,email FROM employee WHERE email = ?", email_entre)
	if err != nil {
		fmt.Println("-->Une erreur est survenue :", err)
		return
	}
	defer res.Close()

	// Parcours de la table afin de retrouver un employé qui vérifie la requête
	rowFound := false
	for res.Next() {
		err = res.Scan(&e.id, &e.name, &e.prenom, &e.fonction, &e.email)
		if err != nil {
			fmt.Println("-->Une erreur est survenue lors de la récupération des données :", err)
			continue
		}
		// change la valeur de rowFoud a  tru s'il la requete sexecute bien
		rowFound = true
	}

	if !rowFound {
		fmt.Println(strings.Repeat("-", 75))
		fmt.Println("-->Aucun employé retrouvé avec cette adresse email")
		fmt.Println(strings.Repeat("-", 75))

	} else {

		fmt.Println(strings.Repeat("-", 25))
		println("verification ok")
		fmt.Println(strings.Repeat("-", 25))

		var rep int

		fmt.Println("-->Quelle informations de l'emplyé voulez vous modifier?")
		fmt.Println("1-Nom")
		fmt.Println("2-Prenom")
		fmt.Println("3-Fonction")
		fmt.Println("4-Email")

		fmt.Scan(&rep)

		switch rep {

		case 1:
			fmt.Print("-->Entrer le nouveau nom  de l'employé:  ")
			fmt.Scan(&e.name)
			for !isValidString(e.name) {
				fmt.Printf("  -->le nom: %s n'est pas valide \n ", e.name)
				fmt.Println(strings.Repeat("-", 75))
				fmt.Print("-->Entrer le nouveux  nom  de l'employé:  ")
				fmt.Scan(&e.name)

			}
			upp, err := DB.Query("UPDATE employee SET Nom=? WHERE id=?", e.name, e.id)
			if err != nil {
				fmt.Println("-->Une erreur est survenue :", err)
				return
			}
			defer upp.Close()
			fmt.Println(strings.Repeat("-", 75))
			fmt.Println("Le nom de l'employé a été mis à jour avec succès .")
			fmt.Println(strings.Repeat("-", 75))

		case 2:

			fmt.Print("-->Entrer le nouveau  Prenom  de l'employé: ")
			fmt.Scan(&e.prenom)
			for !isValidString(e.prenom) {
				fmt.Printf("  -->le prenom: %s n'est pas valide \n ", e.prenom)
				fmt.Println(strings.Repeat("-", 75))
				fmt.Print("-->Entrer le nouveau  Prenom  de l'employé: ")

				fmt.Scan(&e.prenom)

			}
			upp, err := DB.Query("UPDATE employee SET Prenom=? WHERE id=?", e.prenom, e.id)
			if err != nil {
				fmt.Println("-->Une erreur est survenue :", err)
				return
			}
			defer upp.Close()
			fmt.Println(strings.Repeat("-", 75))

			fmt.Println("Le prenom de l'employé a été mis à jour avec succès .")

			fmt.Println(strings.Repeat("-", 75))

		case 3:
			fmt.Print("-->Entrer la nouvelle  fonction  de l'employé: ")
			fmt.Scan(&e.fonction)
			for !isValidString(e.fonction) {
				fmt.Printf(" -->la fonction:%s n'est pas valide \n ", e.fonction)
				fmt.Println(strings.Repeat("-", 75))
				fmt.Print("-->Entrer la nouvelle  fonction  de l'employé: ")
				fmt.Scan(&e.fonction)

			}
			upp, err := DB.Query("UPDATE employee SET Fonction=? WHERE id=?", e.fonction, e.id)
			if err != nil {
				fmt.Println("Une erreur est survenue :", err)
				return
			}
			defer upp.Close()

			fmt.Println(strings.Repeat("-", 75))
			fmt.Println("La fonction de l'employé a été mise à jour avec succès .")
			fmt.Println(strings.Repeat("-", 75))

		case 4:

			fmt.Print("-->Entrer la nouvelle email : ")
			fmt.Scan(&e.email)
			for !isValidEmail(e.email) {
				fmt.Printf("l'adresse email %s n'est pas valide \n", e.email)
				fmt.Println(strings.Repeat("-", 75))
				fmt.Print("-->Entrer un nouveau  mail  de l'employé: ")
				fmt.Scan(&e.email)
			}
			upp, err := DB.Query("UPDATE employee SET email=? WHERE id=?", e.email, e.id)
			if err != nil {
				fmt.Println("-->Une erreur est survenue :", err)
				return
			}
			defer upp.Close()
			fmt.Println(strings.Repeat("-", 75))
			fmt.Println("Ladress email de l'employé a été mise à jour avec succès .")
			fmt.Println(strings.Repeat("-", 75))

		default:
			fmt.Println("Choix invalide. Veuillez saisir 1,2,3 ou 4.")

		}

	}

}

// /***********************(- -)*************************
// fonction pour supprimé un employé
func DeleteEmployee() {
	//var e employee
	var email_entre string

	// Récupération du nom, prénom et de la fonction
	fmt.Println("-->Pour supprimer, il faudra entrer l'email de l'employé à supprimer afin de vérifier s'il existe")

	fmt.Println("  -->Entrer un email de l'employé à modifier :")
	fmt.Scan(&email_entre)

	// Vérification de  l'existance de l'employé
	res, err := DB.Query("SELECT id, Nom, Prenom, Fonction,email FROM employee WHERE email = ?", email_entre)
	if err != nil {
		fmt.Println("-->Une erreur est survenue :", err)
		return
	}
	defer res.Close()

	// Parcours de la table afin de récuper les informations de l'employé qui vérifi la requête
	rowFound := false
	for res.Next() {
		err = res.Scan(&e.id, &e.name, &e.prenom, &e.fonction, &e.email)
		if err != nil {
			fmt.Println("-->Une erreur est survenue lors de la récupération des données :", err)
			continue
		}
		// change la valeur de rowFoud à  true s'il la requete sexecute bien
		rowFound = true
	}

	if !rowFound {
		fmt.Println(strings.Repeat("-", 65))
		fmt.Println("Aucun employé retrouvé avec ces informations")
		fmt.Println(strings.Repeat("-", 65))

	} else {

		fmt.Println(strings.Repeat("-", 25))

		fmt.Println("verification ok")

		fmt.Println(strings.Repeat("-", 25))

		var rep int
		fmt.Println(" --> Voulez vous vraiment supprimer cet émployé ?")
		fmt.Println("1- Oui")
		fmt.Println("2- Non")
		fmt.Scan(&rep)

		switch rep {
		case 1:
			upp, err := DB.Query("DELETE FROM employee WHERE id=?", e.id)
			if err != nil {
				fmt.Println("-->Une erreur est survenue :", err)
				return
			}
			defer upp.Close()
			fmt.Println(strings.Repeat("-", 45))
			fmt.Println("L'employé à été supprimé.")
			fmt.Println(strings.Repeat("-", 45))

		case 2:
			fmt.Println(strings.Repeat("-", 45))
			fmt.Println("Suppression annulée.")
			fmt.Println(strings.Repeat("-", 45))
		default:
			fmt.Println("Choix invalide. Veuillez saisir 1 ou 2.")
		}

	}

}

// /***********************(- -)*************************
// / function pour exporter les données dans le fichié excel
func ExporteExcel() {
	//chemin vers le fichié excel
	filename := "./././exercic2/employee.xlsx"
	results, err := DB.Query("SELECT * FROM employee")
	if err != nil {
		fmt.Println(" --> Une erreur est survenue lors de l'execution de la requête :", err)
		return
	}
	defer results.Close()

	// Création du fichié Excel et de la feuille
	f := excelize.NewFile()
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println("-->Erreur lors de la création de la feuille :", err)
		return
	}
	// Ajout d'un simple style a notre fichié excel
	style := &excelize.Style{
		Fill: excelize.Fill{
			Type:    "pattern",
			Color:   []string{"#FFFF11"},
			Pattern: 1,
		},
		Font: &excelize.Font{
			Bold: true,
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
		},
	}

	styleID, err := f.NewStyle(style)

	if err != nil {
		panic(err)
	}
	/*
		Ajout des en-têtes dans le fichié Excel
		pour ce faire, j'ai creé deux tableau un pour les valeurs des cellules du header et un autre
		 pour les cellule puis j'ai fait une boucle for pour le remplir
	*/
	const rowheader = 1
	headerCellValu := [5]string{"ID", "Nom", "Prenom", "Fonction", "Email"}
	cell := [5]string{"A", "B", "C", "D", "E"}
	for i := 0; i < len(headerCellValu); i++ {
		f.SetCellValue("Sheet1", fmt.Sprintf("%s"+"%d", cell[i], rowheader), headerCellValu[i])
		f.SetCellStyle("sheet1", fmt.Sprintf("%s"+"%d", cell[i], rowheader), fmt.Sprintf("%s"+"%d", cell[i], rowheader), styleID)
	}
	row := 2 // Initialise la ligne de début d'insertion
	idd := 0 // Initialise l'identifiant pour eviter que lorsqu'on suprime un élément dans la table les id soit en desorde lors de l'affichage

	for results.Next() {
		idd++
		err = results.Scan(&e.id, &e.name, &e.prenom, &e.fonction, &e.email)
		if err != nil {
			fmt.Println("-->Une erreur est survenue lors de la lecture des résultats :", err)
			continue
		}
		/*
			Ajout des données dans les cellules correspondantes. Pour ce faire, j'ai,creé deux tableaux
			un tableau pour les valeurs des cellules
			 et un autre pour les diffentes cellules
			 en suite pour notre tableau de valeurs des cellules, etant donné que les
			les valeurs sont de différent type, en faisant quelque recherche et en ayant vu des tutos sur go
			j'ai vu que en go on pouvait creer un tableau de différents type avec le
			 type "interface" donc j'ai utilisé le type interface qui permet
			de creer des tableux avec des valeurs de differents type
		*/
		CellValuvalue := []interface{}{idd, e.name, e.prenom, e.fonction, e.email}
		cell := [5]string{"A", "B", "C", "D", "E"}
		for i := 0; i < len(CellValuvalue); i++ {
			f.SetCellValue("Sheet1", fmt.Sprintf("%s"+"%d", cell[i], row), CellValuvalue[i])
		}
		row++

	}
	fmt.Println(" -->Exportation effectuée avec succes")
	// Définir la feuille active du classeur
	f.SetActiveSheet(index)

	// Enregistrer le fichié Excel par le chemin donné
	if err := f.SaveAs(filename); err != nil {
		fmt.Println("-->Erreur lors de l'enregistrement du fichier Excel :", err)
		return
	}
}
