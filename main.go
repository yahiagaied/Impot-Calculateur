package main

import (
	// fmt pour l'affichage des messages dans la console
	"fmt"
	// template pour le rendu des templates HTML
	"html/template"
	// log pour la journalisation des erreurs
	"log"
	// net/http pour la gestion des requêtes HTTP
	"net/http"
	// strconv pour la conversion de chaînes en nombres
	"strconv"
	// time pour l'affichage de la date actuelle
	"time"
	// strings pour la manipulation de chaînes de caractères
	"strings"
	
)
// Page d'accueil
    func accueil(w http.ResponseWriter, r *http.Request) {
	// Charger le template HTML pour la page d'accueil
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Exécuter le template sans données
	tmpl.Execute(w, nil)
}
//page option1:
func option1(w http.ResponseWriter, r *http.Request) {
	// Charger le template HTML pour la page d'accueil
	tmpl, err := template.ParseFiles("templates/option1.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Exécuter le template sans données
	tmpl.Execute(w, nil)
}
//page option2:
func option2(w http.ResponseWriter, r *http.Request) {
	// Charger le template HTML pour la page d'accueil
	tmpl, err := template.ParseFiles("templates/option2.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Exécuter le template sans données
	tmpl.Execute(w, nil)
}
//page option3:
func option3(w http.ResponseWriter, r *http.Request) {
	// Charger le template HTML pour la page d'accueil
	tmpl, err := template.ParseFiles("templates/option3.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Exécuter le template sans données
	tmpl.Execute(w, nil)
}

// Page de traitement des données du formulaire
func traitement(w http.ResponseWriter, r *http.Request) {
	// Vérifier que la méthode est POST
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Récupérer les valeurs du formulaire et remplacer les virgules par des points (pour les décimaux)
	vhStr := strings.ReplaceAll(r.FormValue("vh"), ",", ".")
	vttcStr := strings.ReplaceAll(r.FormValue("vttc"), ",", ".")
	tvaStr := strings.ReplaceAll(r.FormValue("tva"), ",", ".")

	// Convertir les valeurs en float64
	vh, err1 := strconv.ParseFloat(vhStr, 64)
	vttc, err2 := strconv.ParseFloat(vttcStr, 64)
	tva, err3 := strconv.ParseFloat(tvaStr, 64)

	// Validation
	if err1 != nil || err2 != nil || err3 != nil || vh <= 0 || vttc <= 0 || tva <= 0 {
		http.Error(w, " 🥸🥸 Veuillez entrer des valeurs valides pour les montants et le taux de TVA !!!", http.StatusBadRequest)
		return
	}

	// Calculer retenu a la source et tva :
	retenu1 := vttc * 1 / 100
	retenu2 := tva * 25 / 100

	// Calculer le montant de retenue après retenue à la source et TVA
	var MontantNetPaye float64
	MontantNetPaye = vttc - (retenu1 + retenu2)

	// Charger le template de traitement
	tmpl, err := template.ParseFiles("templates/traitement.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Préparer les données à passer au template
	data := struct {
		Vh             string
		Tva            string
		Vttc           string
		Retenu1        string
		Retenu2        string
		MontantNetPaye string
	}{
		Vh:             vhStr,
		Tva:            tvaStr,
		Vttc:           vttcStr,
		Retenu1:        fmt.Sprintf("%.2f", retenu1),
		Retenu2:        fmt.Sprintf("%.2f", retenu2),
		MontantNetPaye: fmt.Sprintf("%.2f", MontantNetPaye),
	}

	// Exécuter le template avec les résultats
	tmpl.Execute(w, data)
}

// Page de traitement2 des données du formulaire
func traitement2(w http.ResponseWriter, r *http.Request) {
	// Vérifier que la méthode est POST
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	// Récupérer les valeurs du formulaire
	vhStr := strings.ReplaceAll(r.FormValue("vh"), ",", ".")
	vttcStr := strings.ReplaceAll(r.FormValue("vttc"), ",", ".")
	tvaStr := strings.ReplaceAll(r.FormValue("tva"), ",", ".")
	
	// Convertir les valeurs en float64
	vh, err1 := strconv.ParseFloat(vhStr, 64)
	vttc, err2 := strconv.ParseFloat(vttcStr, 64)
	tva, err3 := strconv.ParseFloat(tvaStr, 64)

	if err1 != nil || err2 != nil || err3 != nil || vh <= 0 || vttc <= 0 || tva <= 0 {
		http.Error(w, " 🥸🥸 Veuillez entrer des valeurs valides pour les montants et le taux de TVA !!!", http.StatusBadRequest)
		return
	}
	// Calculer retenu a la source et tva :
	  retenu1 := vttc * 1.5/100 
	  retenu2 := tva * 25/100 
    
   // Calculer le montant de retenue après retenue à la source et TVA
     var MontantNetPaye float64
     MontantNetPaye = vttc - (retenu1 + retenu2)
	
	// Charger le template de traitement
	tmpl, err := template.ParseFiles("templates/traitement.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Préparer les données à passer au template
	data := struct {
		Vh             string
		Tva            string
		Vttc           string
		Retenu1        string
		Retenu2        string
		MontantNetPaye string
	}{
		Vh:             vhStr,
		Tva:            tvaStr,
		Vttc:           vttcStr,
		Retenu1:        fmt.Sprintf("%.2f", retenu1),
		Retenu2:        fmt.Sprintf("%.2f", retenu2),
		MontantNetPaye: fmt.Sprintf("%.2f", MontantNetPaye),
	}

	// Exécuter le template avec les résultats
	tmpl.Execute(w, data)
}

	
// Page de traitement3 des données du formulaire
func traitement3(w http.ResponseWriter, r *http.Request) {
	// Vérifier que la méthode est POST
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Récupérer les valeurs du formulaire
	vhStr := strings.ReplaceAll(r.FormValue("vh"), ",", ".")
	vttcStr := strings.ReplaceAll(r.FormValue("vttc"), ",", ".")
	tvaStr := strings.ReplaceAll(r.FormValue("tva"), ",", ".")

	// Convertir les valeurs en float64
	vh, err1 := strconv.ParseFloat(vhStr, 64)
	vttc, err2 := strconv.ParseFloat(vttcStr, 64)
	tva, err3 := strconv.ParseFloat(tvaStr, 64)

	if err1 != nil || err2 != nil || err3 != nil || vh <= 0 || vttc <= 0 || tva <= 0 {
		http.Error(w, " 🥸🥸 Veuillez entrer des valeurs valides pour les montants et le taux de TVA ...!!!", http.StatusBadRequest)
		return
	}// Calculer retenu a la source et tva :
	  retenu1 := vttc * 3/100 
	  retenu2 := tva * 25/100 
    
   // Calculer le montant de retenue après retenue à la source et TVA
     var MontantNetPaye float64
     MontantNetPaye = vttc - (retenu1 + retenu2)

	// Charger le template de traitement
	tmpl, err := template.ParseFiles("templates/traitement.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Préparer les données à passer au template
	data := struct {
		Vh             string
		Tva            string
		Vttc           string
		Retenu1        string
		Retenu2        string
		MontantNetPaye string
	}{
		Vh:             vhStr,
		Tva:            tvaStr,
		Vttc:           vttcStr,
		Retenu1:        fmt.Sprintf("%.2f", retenu1),
		Retenu2:        fmt.Sprintf("%.2f", retenu2),
		MontantNetPaye: fmt.Sprintf("%.2f", MontantNetPaye),
			
	}
	// Exécuter le template avec les résultats
	tmpl.Execute(w, data)

}

// Fonction principale pour démarrer le serveur HTTP

func main() {
	// Définir les routes pour les différentes pages
	     http.HandleFunc("/", accueil)


    // Définir les routes pour les pages d'option1 et de traitement
		http.HandleFunc("/option1", option1)
		http.HandleFunc("/traitement", traitement)

	// Définir les routes pour les pages d'option2et de traitement2
		http.HandleFunc("/option2", option2)
		http.HandleFunc("/traitement2", traitement2)

	// Définir les routes pour les pages d'option3 et de traitement3
		http.HandleFunc("/option3", option3)
		http.HandleFunc("/traitement3", traitement3)

	
    // Servir les fichiers statiques (CSS, images, etc.)
    fs := http.FileServer(http.Dir("static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Afficher la date actuelle dans la console
	fmt.Println("Date actuelle :", time.Now().Format("02-01-2006"))

	// Démarrer le serveur HTTP sur le port 8080
	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

   
}