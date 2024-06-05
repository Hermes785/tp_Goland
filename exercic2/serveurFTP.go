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

const ftpServer = "localhost:21"

// definition de notre structure
type utiliateurConfig struct {
	ftpuser string
	ftppass string
}

func ServerFTP() {
	var u utiliateurConfig

	fmt.Println("veuillez entrer votre nom utiliateur")
	fmt.Scan(&u.ftpuser)

	fmt.Println("veuillez entrer votre password")
	fmt.Scan(&u.ftppass)

	// Connexion au serveur FTP en utilisant EPSV par défaut
	c, err := ftp.Dial(ftpServer, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer c.Quit() // Déconnexion du serveur FTP à la fin de la fonction

	// Authentification auprès du serveur FTP
	err = c.Login(u.ftpuser, u.ftppass)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Repeat("-", 35))
	fmt.Println("Connexion au serveur FTP réussie !")
	fmt.Println(strings.Repeat("-", 35))

}
