package exercice2

import (
	"fmt"
	"log"

	//"os"
	//"path/filepath"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
)

const ftpServer = "localhost:21" // Adresse du serveur FTP
const ftpUser = "hermes"         // Nom d'utilisateur du serveur FTP
const ftpPass = "golang"         // Mot de passe du serveur FTP

func ServerFTP() {
	// Connexion au serveur FTP en utilisant EPSV par défaut
	c, err := ftp.Dial(ftpServer, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Quit() // Déconnexion du serveur FTP à la fin de la fonction

	// Authentification auprès du serveur FTP
	err = c.Login(ftpUser, ftpPass)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Repeat("-", 35))
	fmt.Println("Connexion au serveur FTP réussie !")
	fmt.Println(strings.Repeat("-", 35))

}
