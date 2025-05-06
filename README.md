# 🍇 FruitStore - Application Web en Go

Bienvenue sur **FruitStore**, une application web simple codée en Go.  
Elle vous permet de gérer un petit stock de fruits 🍎🍌🍓 dans une interface web légère basée sur des templates HTML.

---

## 🚀 Fonctionnalités prévues

- [x] Affichage d'une page d'accueil
- [ ] Ajout de produits (fruit, origine, poids)
- [ ] Visualisation du stock
- [ ] Calcul du prix avec TVA
- [ ] Interface claire et responsive (HTML/CSS)

---

## 🧱 Structure du projet

Fruit-store/
├── main.go
├── templates/
│   └── index.html
├── static/
│   └── css/
│       └── style.css
├── .gitignore
├── README.md

---

## ▶️ Lancer l'application

1. Clonez le dépôt :

```bash
git clone https://github.com/votre-utilisateur/Fruit-store.git
cd Fruit-store
```

2. Lancer le server Go

```bash
go run main.go
```

3. Ouvrez votre navigateur à l’adresse :

```bash
http://localhost:8080
```


📚 Technologies utilisées

- Go pour le serveur

- net/http et html/template pour le routage et les pages HTML