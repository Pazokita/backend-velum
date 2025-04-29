# 🧭 Velum – Backend

**Velum** est une application mobile d’exploration de **cartes historiques géoréférencées**.  
Ce dépôt contient le **backend en Go** structuré selon une architecture modulaire et évolutive.

---

## 📌 Objectif de la Phase 1

- ✅ Créer un serveur Go avec le framework **Gin**
- ✅ Structurer le backend selon une Clean Architecture légère
- ✅ Exposer une première route statique `GET /maps` qui renvoie des métadonnées de cartes
- ✅ Préparer le terrain pour l’ajout de base de données et d’upload de cartes dans les prochaines phases

---

## 🛠️ Stack technique

- **Langage** : Go (Golang)
- **Framework Web** : [Gin](https://github.com/gin-gonic/gin)
- **Architecture** : Clean Architecture (modulaire et extensible)
- **API** : RESTful

---

## 📂 Arborescence
backend-velum/
├── cmd/
│   └── server/
│       └── main.go            # Point d’entrée du serveur
├── internal/
│   └── maps/
│       ├── handler.go         # Logique HTTP (GET /maps)
│       ├── model.go           # Définition de la struct MapMetadata
│       └── service.go         # (Réservé pour la logique métier future)
└── go.mod                     # Déclaration du module Go

---

## 🚀 Lancer le serveur localement

### 1. Cloner le repo :

```bash
git clone https://github.com/Pazokita/backend-velum.git
cd backend-velum

```

### 2. Installer les dépendances :

```bash
go mod tidy

```

### 3. Lancer le serveur :
```bash
go run cmd/server/main.go

```
### Le serveur démarre sur : http://localhost:8080

## 🔍 Endpoint disponible
```bash
GET /maps

Renvoie une liste statique de métadonnées de cartes.

Exemple de réponse :
[
  {
    "id": 1,
    "title": "Paris 1750",
    "imageUrl": "https://example.com/paris-1750.png",
    "bbox": [2.3319, 48.8566, 2.3519, 48.8666],
    "opacity": 0.8
  }
]
```

## 📅 Roadmap Velum

🟢 Phase 1 – MVP statique (en cours)
	•	Backend Go
	•	Route statique /maps
	•	Frontend React Native (carte + overlay + slider)

🟡 Phase 2 – Backend dynamique
	•	Connexion à SQLite
	•	Endpoint GET /maps/:id
	•	Serveur REST complet

🔵 Phases suivantes
	•	Upload de cartes
	•	Authentification
	•	PostGIS & requêtes spatiales
	•	Tests & déploiement

## 👩‍💻 Auteure

Amalia Maturana – https://www.amaliamaturana.com
Passionnée de cartographie, d’archéologie et de clean code ❤️

## 📜 Licence

Ce projet est open-source sous licence MIT.