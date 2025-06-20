# 🧭 Velum – Backend

**Velum** est une application mobile d’exploration de **cartes historiques géoréférencées**.  
Ce dépôt contient le **backend en Go** structuré selon une architecture modulaire et évolutive.

---

## 📌 Objectif de la Phase 2/3

✅ Créer un serveur Go avec le framework Gin

✅ Structurer le projet de manière modulaire

✅ Exposer une route GET /maps renvoyant plusieurs cartes (simulées ou via SQLite)

✅ Ajouter une route GET /maps/:id pour les détails d’une carte

🔄 En cours : intégration d’un microservice FastAPI pour la recherche sémantique IA

🔜 Objectif MVP IA : GET /semantic-search?q=... → appelle /search (FastAPI)

---

## 🛠️ Stack technique

- **Langage** : Go (Golang)
- **Framework Web** : [Gin](https://github.com/gin-gonic/gin)
- **Architecture** : Clean Architecture (modulaire et extensible)
- **Base de données** : SQLite (évolutif vers PostgreSQL + PostGIS)
- **Interopérabilité** : microservice Python (FastAPI)

---

## 📂 Arborescence
```bash
backend-velum/
├── cmd/
│   └── server/
│       └── main.go              # Point d’entrée du serveur
├── internal/
│   ├── maps/
│   │   ├── handler.go           # Routes /maps et /maps/:id
│   │   ├── model.go             # Struct MapMetadata
│   │   └── service.go           # Logique métier
│   └── semanticsearch/
│       └── handler.go           # Route /semantic-search → appelle FastAPI
├── database/
│   ├── maps.db                  # Base SQLite (si utilisée)
│   └── init.sql                 # Script de création initiale
└── go.mod                       # Déclaration du module Go

```
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
📍 GET /maps
Renvoie la liste des cartes disponibles.

📍 GET /maps/:id
Renvoie une carte par son ID (titre, image, bbox, etc.)

🧠 GET /semantic-search?q=... (en cours)
Fait une recherche sémantique via un microservice Python et retourne les lieux similaires.
```
## 📦 Base de données

Le fichier `database/maps.db` contient les cartes historiques utilisées pour l’API.

### Intégration IA – FastAPI (Phase 4)
⚡ En développement

Le backend Go délègue les recherches sémantiques à un microservice FastAPI basé sur HuggingFace MiniLM.

Endpoint FastAPI : POST /search

Backend Go appelle : GET /semantic-search?q=... → interne à handler.go

Le microservice lit une base JSON de lieux historiques (nom, description, coordonnées)

## 🗃️ Base de données 

Le fichier database/maps.db contient des données de cartes si tu veux utiliser SQLite.

Recréation :
```bash
sqlite3 database/maps.db < database/init.sql

```
## 📅 Roadmap Velum

🟢 Phase 1 – MVP statique
🟡 Phase 2 – Backend dynamique (/maps, /maps/:id)
🔵 Phase 3 – UI dynamique & superposition
🧠 Phase 4 – IA : recherche sémantique (FastAPI)
🚀 Phase 5 – Cartographie avancée (AR, 3D)
📦 Phase 6 – Finalisation, test & publication

## 👩‍💻 Auteure

Amalia Maturana – https://www.amaliamaturana.com
Passionnée de cartographie, d’archéologie et de clean code ❤️

## 📜 Licence

Ce projet est open-source sous licence MIT.