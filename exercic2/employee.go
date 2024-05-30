package exercice2

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

//pour eviter de devoir définir a chaque fois les variables dans chaque fonction
//j'ai creé un objet qui contiendra les variables nécéssaire pour notre entité employé

type employee struct {
	id       int
	name     string
	prenom   string
	fonction string
	email    string
}

// fonction pour un  ajouté employés
func AddEmployee() {
	var e employee
	fmt.Println("--Veuillez entrer les informations de l'employé")
	// Lire les entrées de l'utilisateur
	fmt.Print("Veuillez entrer un nom : ")
	fmt.Scan(&e.name)
	fmt.Print("Veuillez entrer un prénom : ")
	fmt.Scan(&e.prenom)
	fmt.Print("Veuillez entrer une fonction : ")
	fmt.Scan(&e.fonction)
	fmt.Print("Veuillez entrer une email : ")
	fmt.Scan(&e.email)
	// Préparer l'insertion des données
	insert, err := DB.Query("INSERT INTO employee (nom, prenom, fonction,email) VALUES (?, ?, ?,?)", e.name, e.prenom, e.fonction, e.email)
	if err != nil {
		fmt.Println("Une erreur est survenue :", err)
		return
	}
	defer insert.Close()
	fmt.Println(strings.Repeat("-", 75))
	fmt.Println("L'employé a été ajouté avec succès à la base de données.")
	fmt.Println(strings.Repeat("-", 75))

}

///***********************(- -)*************************
// fonction pour lister les  employés

func ListAllUser() {

	//filename := "C:/Project/Tp_Hermes_nguefack/exercic2/employee.xlsx"
	var e employee
	fmt.Println("--La liste des utilisateurs est :")

	results, err := DB.Query("SELECT * FROM employee")
	if err != nil {
		fmt.Println("Une erreur est survenue lors de la requête :", err)
		return
	}
	defer results.Close()

	idd := 0 // Initialise l'identifiant pour eviter que lorsque suprime un element dans la table les id soit desordonne
	fmt.Println(strings.Repeat("-", 126))
	fmt.Printf("%-10s | %-30s | %-40s | %-40s| %-20s\n", "ID", "Nom", "Prénom", "Fonction", "Email")
	fmt.Println(strings.Repeat("-", 146))
	for results.Next() {
		idd++
		err = results.Scan(&e.id, &e.name, &e.prenom, &e.fonction, &e.email)
		if err != nil {
			fmt.Println("Une erreur est survenue lors de la lecture des résultats :", err)
			continue
		}

		fmt.Printf("%-10d | %-30s | %-40s | %-40s| %-20s\n", idd, e.name, e.prenom, e.fonction, e.email)

		// Crée les références des cellules pour chaque colonne

	}

}

// /************************************************
// fonction pour mofifier un employé
func UpdatEmployee() {
	var e employee
	var email_entre string

	// Récupération du nom, prénom et de la fonction
	fmt.Println("--Pour modifier, il faudra entrer l'email de l'employé à modifier afin de vérifier s'il existe")

	fmt.Println("--Entrer un email de l'employé à modifier :")
	fmt.Scan(&email_entre)

	// Vérification si l'employé existe
	res, err := DB.Query("SELECT id, Nom, Prenom,Fonction,email FROM employee WHERE email = ?", email_entre)
	if err != nil {
		fmt.Println("Une erreur est survenue :", err)
		return
	}
	defer res.Close()

	// Parcours de la table afin de retrouver un employé qui vérifie la requête
	rowFound := false
	for res.Next() {
		err = res.Scan(&e.id, &e.name, &e.prenom, &e.fonction, &e.email)
		if err != nil {
			fmt.Println("Une erreur est survenue lors de la récupération des données :", err)
			continue
		}
		// renvoi tru s'il la requete sexecute bien
		rowFound = true
	}

	if !rowFound {
		fmt.Println(strings.Repeat("-", 75))
		fmt.Println("Aucun employé retrouvé avec ces informations")
		fmt.Println(strings.Repeat("-", 75))

	} else {
		///fmt.Printf("l'id retrouver pour est %d\n",e.id)
		fmt.Println(strings.Repeat("-", 25))
		println("verification ok")
		fmt.Println(strings.Repeat("-", 25))
		var rep int
		fmt.Println("Quelle informations de l'emplyé voulez vous modifier?")
		fmt.Println("1-Nom")
		fmt.Println("2-Prenom")
		fmt.Println("3-Fonction")
		fmt.Println("4-Email")
		fmt.Scan(&rep)

		switch rep {

		case 1:
			fmt.Print("Veuillez entrer le nouveux  nom : ")
			fmt.Scan(&e.name)
			upp, err := DB.Query("UPDATE employee SET Nom=? WHERE id=?", e.name, e.id)
			if err != nil {
				fmt.Println("Une erreur est survenue :", err)
				return
			}
			defer upp.Close()
			fmt.Println(strings.Repeat("-", 75))
			fmt.Println("Le nom de l'employé a été mis a jour avec succès .")
			fmt.Println(strings.Repeat("-", 75))

		case 2:
			fmt.Print("Veuillez entrer le nouveux  Prenom : ")
			fmt.Scan(&e.prenom)
			upp, err := DB.Query("UPDATE employee SET Prenom=? WHERE id=?", e.prenom, e.id)
			if err != nil {
				fmt.Println("Une erreur est survenue :", err)
				return
			}
			defer upp.Close()
			fmt.Println(strings.Repeat("-", 75))

			fmt.Println("Le prenom d l'employé a été mi a jour avec succès .")

			fmt.Println(strings.Repeat("-", 75))

		case 3:
			fmt.Print("Veuillez entrer la nouvelle  fonction : ")
			fmt.Scan(&e.name)
			upp, err := DB.Query("UPDATE employee SET Nom=? WHERE id=?", e.fonction, e.id)
			if err != nil {
				fmt.Println("Une erreur est survenue :", err)
				return
			}
			defer upp.Close()

			fmt.Println(strings.Repeat("-", 75))
			fmt.Println("La fonction de l'employé a été mi a jour avec succès .")
			fmt.Println(strings.Repeat("-", 75))

		case 4:
			fmt.Print("Veuillez entrer un nouveux  mail : ")
			fmt.Scan(&e.name)
			upp, err := DB.Query("UPDATE employee SET email=? WHERE id=?", e.email, e.id)
			if err != nil {
				fmt.Println("Une erreur est survenue :", err)
				return
			}
			defer upp.Close()
			fmt.Println(strings.Repeat("-", 75))
			fmt.Println("L'email de l'employé a été mi a jour avec succès .")
			fmt.Println(strings.Repeat("-", 75))

		}

	}

}

// /************************************************
// fonction pour supprimé un employé
func DeleteEmployee() {
	var e employee
	var email_entre string

	// Récupération du nom, prénom et de la fonction
	fmt.Println("Pour modifier, il faudra entrer l'email de l'employé à modifier afin de vérifier s'il existe")

	fmt.Println("Entrer un email de l'employé à modifier :")
	fmt.Scan(&email_entre)

	// Vérification si l'employé existe
	res, err := DB.Query("SELECT id, Nom, Prenom, Fonction FROM employee WHERE email = ?", email_entre)
	if err != nil {
		fmt.Println("Une erreur est survenue :", err)
		return
	}
	defer res.Close()

	// Parcours de la table afin de retrouver un employé qui vérifie la requête
	rowFound := false
	for res.Next() {
		err = res.Scan(&e.id, &e.name, &e.prenom, &e.fonction, &e.email)
		if err != nil {
			fmt.Println("Une erreur est survenue lors de la récupération des données :", err)
			continue
		}
		// renvoi tru s'il la requete sexecute bien
		rowFound = true
	}

	if !rowFound {
		fmt.Println(strings.Repeat("-", 65))
		fmt.Println("Aucun employé retrouvé avec ces informations")
		fmt.Println(strings.Repeat("-", 65))

	} else {
		///fmt.Printf("l'id retrouver pour est %d\n",e.id)
		fmt.Println(strings.Repeat("-", 25))
		println("verification ok")
		fmt.Println(strings.Repeat("-", 25))
		fmt.Println("Veuillez les nouveaux informations de l'employé")

		// on ecrit notre requete de la suppression
		upp, err := DB.Query("DELETE FROM employee WHERE id=?", e.id)
		if err != nil {
			fmt.Println("Une erreur est survenue :", err)
			return
		}
		defer upp.Close()
		fmt.Println(strings.Repeat("-", 45))
		fmt.Println("L'employé a été supprimé.")
		fmt.Println(strings.Repeat("-", 45))

	}

}

func ExporteExcel() {

	filename := "./././exercic2/employee.xlsx"
	var e employee

	results, err := DB.Query("SELECT * FROM employee")
	if err != nil {
		fmt.Println("Une erreur est survenue lors de la requête :", err)
		return
	}
	defer results.Close()

	// Création du fichier Excel et de la feuille
	f := excelize.NewFile()
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println("Erreur lors de la création de la feuille :", err)
		return
	}

	// Ajout des en-têtes dans le fichier Excel
	f.SetCellValue("Sheet1", "A1", "id")
	f.SetCellValue("Sheet1", "B1", "Name")
	f.SetCellValue("Sheet1", "C1", "prenom")
	f.SetCellValue("Sheet1", "D1", "Fonction")

	row := 2 // Initialise la ligne de début d'insertion

	idd := 0 // Initialise l'identifiant pour eviter que lorsque suprime un element dans la table les id soit desordonne

	for results.Next() {
		idd++
		err = results.Scan(&e.id, &e.name, &e.prenom, &e.fonction, &e.email)
		if err != nil {
			fmt.Println("Une erreur est survenue lors de la lecture des résultats :", err)
			continue
		}

		// Crée les références des cellules pour chaque colonne
		cellId := fmt.Sprintf("A%d", row)
		cellName := fmt.Sprintf("B%d", row)
		cellPrenom := fmt.Sprintf("C%d", row)
		cellFonction := fmt.Sprintf("D%d", row)
		cellEmail := fmt.Sprintf("D%d", row)

		// Ajoute les données dans les cellules correspondantes
		f.SetCellValue("Sheet1", cellId, idd)
		f.SetCellValue("Sheet1", cellName, e.name)
		f.SetCellValue("Sheet1", cellPrenom, e.prenom)
		f.SetCellValue("Sheet1", cellFonction, e.fonction)
		f.SetCellValue("Sheet1", cellEmail, e.email)

		row++
		fmt.Println("Exportation effectuée avec succes")

	}

	// Définir la feuille active du classeur
	f.SetActiveSheet(index)

	// Enregistrer le fichier Excel par le chemin donné
	if err := f.SaveAs(filename); err != nil {
		fmt.Println("Erreur lors de l'enregistrement du fichier Excel :", err)
		return
	}
}
