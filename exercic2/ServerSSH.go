package exercice2

// pour la realisation de ce code j'ai  tout simplement regardé la documentation(suivre le lien golang.org/x/crypto/ssh pour voir) et appliqué  ce qui etait dessus en modifiant ce qu'il avait à modifier  golang.org/x/crypto/ssh
import (
	"context"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/ssh"

	issh "github.com/jlandowner/go-interactive-ssh"
)

// defimition de notre struture
type configuration struct {
	host     string
	user     string
	password string
}

func ServerSSh() {
	//ici je demane a l'utilsateur d'entré les information de sa Vm à savoir son adresse ip, son usename et son password
	var c configuration
	ctx := context.Background()
	fmt.Println("Entrer l'Adresse IP de la machine virtuelle Ubuntu:")
	fmt.Scan(&c.host)
	fmt.Println("Entrer le  Username de la machine virtuelle:")
	fmt.Scan(&c.user)
	fmt.Println("Entrer le Password de la machine virtuelle:")
	fmt.Scan(&c.password)

	// une fois cela recuperé  je teste la connection

	// Informations d'authentification
	config := &ssh.ClientConfig{
		User: c.user,

		Auth: []ssh.AuthMethod{
			ssh.Password(c.password),
		},
		// Ignorer les vérifications de clé de serveur SSH (à utiliser avec prudence)
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Établissement de la connexion SSH
	client := issh.NewClient(config, c.host, "22", []issh.Prompt{issh.DefaultPrompt})
	// pour monter que cela passe vraiment j'execute quelque commande avant d'affiché  le "connexion à la VM reusssi si tout ce passe bien biensur"
	err := client.Run(ctx, []*issh.Command{
		issh.CheckUser(c.user),
		issh.ChangeDirectory("/tmp"),
		issh.NewCommand("sleep 2"),
		issh.NewCommand("cat ~/.ssh/id_rsa.pub", issh.WithOutputLevelOption(issh.Output)),
		issh.NewCommand("ls -l", issh.WithOutputLevelOption(issh.Output)),
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(strings.Repeat("-", 75))
	fmt.Println("Connexion à la machine virtuelle reussie")
	fmt.Println(strings.Repeat("-", 75))

}
