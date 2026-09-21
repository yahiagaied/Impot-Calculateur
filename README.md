# Calculateur de Retenue à la Source (Tunisie)

Application web backend développée en **Go** conçue pour les professionnels de la finance et de la comptabilité. Elle automatise le calcul des **retenues à la source (RAS)** et de la **TVA** conformément à la législation fiscale tunisienne.

---

## 📋 Barème et Taux Appliqués (Législation Tunisienne)

| Type de Tiers / Prestataire | Taux de Retenue à la Source (RAS) | Retenue sur TVA |
| :--- | :--- | :--- |
| **Personnes Morales** | **1 %** | 25 % de la TVA (soit 19 %) |
| **Personnes Physiques** | **1,5 %** | 25 % de la TVA (soit 19 %) |
| **Bureaux d'Études** | **3 %** | 25 % de la TVA (soit 19 %) |

> **Note sur la TVA :** Le taux général de la TVA en Tunisie est de **19%**. La retenue à la source sur la TVA représente **25%** du montant de cette TVA.

---

## 💻 Utilisation

### Option 1 : Lancer directement l'exécutable (Recommandé)
Si vous disposez du fichier exécutable compilé de l'application :
1. Double-cliquez sur le fichier exécutable (ou lancez-le depuis votre terminal).
2. Ouvrez votre navigateur web et accédez à l'adresse suivante :
   ```text
   http://localhost:8080


## ✨ Fonctionnalités

- Calcul précis de la RAS selon la nature juridique du fournisseur (Personne Morale, Personne Physique, Bureau d'Études).
- Calcul automatique de la TVA (19%) et de la retenue sur TVA (25% de la TVA).
- Calcul de montant net a payer par les services de l'etat.
- Interface web locale gérée par un serveur HTTP en **Go**.
---

## 🛠️ Technologies utilisées

- Golang
- HTML5
- Bootstrap 5

---

## 🚀 Installation

### Cloner le projet

```bash
git clone https://github.com/yahiagaied/Impot-Calculateur.git
```

### Accéder au dossier

```bash
cd calculimpot
```

### Lancer l'application

```bash
go run main.go
```

Ouvrir ensuite :

```text
http://localhost:8080
```
---

## 🎯 Objectifs du projet

Ce projet a été réalisé afin de :

- pratiquer le développement web avec Go ;
- utiliser les templates HTML ;
- améliorer mes compétences en Bootstrap ;
- créer une interface utilisateur moderne.

---

## 👨‍💻 Auteur

Développé avec ❤️ par **Yahia Gaied**

Passionné par :

- Golang 💻
- Développement Web 🌐
- Musculation 🏋️

---

## ⭐ Si ce projet vous plaît

N'hésitez pas à laisser une étoile ⭐ sur le dépôt GitHub.
