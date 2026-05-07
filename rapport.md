# Introduction

Dans le cadre de notre dernière année d’études à la HEIG-VD, 
nous sommes amenés à choisir des cours à option afin d'approfondir nos connaissances dans des domaines spécifiques. 
Ce rapport s’inscrit dans le cadre du cours PLM (Programming Language Multiverse), dont l'objectif est d'explorer et de s'approprier un paradigme de programmation non abordé durant le cursus obligatoire.

Pour ce projet, nous avons choisi d'étudier le paradigme déclaratif. 
La problématique traitée concerne la création d’une bibliothèque permettant de définir une interface graphique de manière déclarative, 
par le biais d’un DSL (Domain Specific Language).

Lors de la présentation des objectifs du cours, deux approches étaient possibles : 
la première consistait à utiliser un paradigme existant, tandis que la seconde que nous avons privilégiée porte sur l'implémentation même du paradigme au sein du langage de programmation de notre choix.

# Motivation

La motivation première de ce projet est née du visionnage d'une vidéo YouTube vulgarisant l'algorithme de résolution de mise en page (layout) au sein d'une interface graphique. Cette découverte nous a convaincus que le cours PLM représentait l'opportunité idéale pour implémenter notre propre bibliothèque.

Au regard des différents paradigmes à disposition et de notre expérience acquise durant le cursus (notamment en WEB et DAA), le paradigme déclaratif est apparu comme le plus adapté pour résoudre cette problématique. Ce choix est d'autant plus pertinent qu'il reflète les standards actuels des bibliothèques d'interface utilisées dans l'industrie.

L'approche déclarative permet de découpler la complexité algorithmique de la définition structurelle de l'interface. 
Dans une implémentation impérative, la responsabilité du placement des éléments incombe à l'utilisateur, qui doit concevoir sa propre logique pour obtenir le rendu souhaité.
À l'inverse, en mode déclaratif, l'utilisateur se contente de décrire la structure et les contraintes de placement. Il revient alors à notre bibliothèque de résoudre ces contraintes, grâce aux algorithmes implémentés, pour générer le résultat attendu.

L'utilisation de l'approche déclarative offre un avantage majeur : 
l'utilisateur se concentre exclusivement sur la définition de son interface. 
Cette abstraction simplifie considérablement le développement et permet de focaliser les efforts sur les autres composants de l'application. En contrepartie, 
l'utilisateur est limité aux contraintes et aux comportements proposés par la bibliothèque, sans pouvoir implémenter sa propre logique de placement sans modifier le cœur de notre solution.

C'est dans cette optique que nous justifions le choix de ce paradigme pour notre projet. 
De plus nous avons privilégié la seconde voie qui est l'implémentation du paradigme lui-même car, dans le contexte de notre problématique, l'utilisation du paradigme déjà implémenté aurait présenté un intérêt limité.

# Le paradigme déclaratif

## Go et déclaratif

# Cahier des charges

# Conclusion