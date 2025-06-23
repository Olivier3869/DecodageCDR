# goFonctions

Version : 23

Ce depot héberge les fonctions communes aux outils de décodage de CRAs pour : 
 - **norego** , l'outil de NONREG  
 - **decodeCDR** pour le décodage des CRAs (en vue du remplacement du traiterLot)

## Utilisation 

Pour utiliser les fonctions dans un autre projet `Go` il faut configurer vote projet avec l'utilisation d'un fichier `go.mod`.  
Dans ce fichier, ajouter les lignes :

```mod
replace local/goFonctions => ../goFonctions
require local/goFonctions v0.0.0
```

Pour chaque fichier Go où il est nécessaire d'utiliser une fonction, il faut ajouter l'`import` :

```go
import "local/goFonctions"
```

Et pour utiliser, par exemple la fonction `ByteArrayToHexString` :

```go
valeur := goFonctions.ByteArrayToHexString(val)
```

## Fonctions

### Fonction `ByteArrayToHexString`

Convertit un tableau d'octets en une chaîne hexadécimale.

#### Paramètre
- `byteArray` : tableau d'octets à convertir.

#### Retour
- Chaîne de caractères représentant la valeur hexadécimale de chaque octet, concaténée.

#### Fonctionnement
- Parcourt chaque octet du tableau.
- Convertit chaque octet en sa représentation hexadécimale à deux chiffres avec `fmt.Sprintf("%02x", b)`.
- Concatène le résultat dans une chaîne de caractères.

---

### Fonction `BytesToIntBigEndian`

Convertit un tableau d'octets en un entier `int` en utilisant l'Endianness Big Endian.

#### Paramètre
- `data` : tableau de bytes à convertir (doit contenir entre 1 et 8 octets).

#### Retour
- Pointeur vers l'entier converti si la conversion réussie.
- `nil` si le tableau est vide, trop grand ou non supporté.

#### Fonctionnement
- Vérifie si le tableau est vide ou trop grand.
- Selon la taille du tableau, utilise `binary.BigEndian.Uint16` ou `Uint32` pour convertir.
- Retourne un pointeur vers la valeur convertie.

---

### Fonction `BytesToIntLittleEndian`

Convertit un tableau d'octets en un entier `int` en utilisant l'Endianness Little Endian.

#### Paramètre
- `data` : tableau de bytes à convertir (doit contenir entre 1 et 8 octets).

#### Retour
- Valeur entière convertie si la conversion réussie.
- `0` si le tableau est vide, trop grand ou non supporté.

#### Fonctionnement
- Vérifie si le tableau est vide ou trop grand.
- Selon la taille du tableau, utilise `binary.LittleEndian.Uint16` ou `Uint32` pour convertir.
- Retourne la valeur convertie.

---

### Fonction `ConvertDate`

Convertit une chaîne hexadécimale représentant une date en une chaîne formatée `"dd/mm/yyyy"`.

#### Paramètre
- `dateHex` : chaîne hexadécimale de la date.

#### Retour
- Chaîne formatée `"dd/mm/yyyy"` si la conversion réussie.
- Chaîne vide en cas d'erreur.

#### Fonctionnement
- Convertit la chaîne hexadécimale en entier avec `strconv.ParseInt`.
- Extrait le jour, le mois et l'année en utilisant des opérations bit à bit.
- Ajoute 2000 à l'année extraite.
- Retourne la date formatée avec `fmt.Sprintf`.

---

### Fonction `ConvertTime`

Convertit une chaîne hexadécimale représentant une heure en une chaîne formatée `"HH:MM:SS"`.

#### Paramètre
- `timeHex` : chaîne hexadécimale de l'heure.

#### Retour
- Chaîne formatée `"HH:MM:SS"` si la conversion réussie.
- Chaîne vide en cas d'erreur.

#### Fonctionnement
- Convertit la chaîne hexadécimale en entier avec `strconv.ParseInt`.
- Extrait les heures, minutes et secondes en utilisant des opérations bit à bit.
- Retourne l'heure formatée avec `fmt.Sprintf`.

---

### Fonction `ConvertToStringSlice`

Convertit un tableau d'interfaces en un tableau de chaînes de caractères.

#### Paramètre
- `input` : slice de `interface{}` à convertir.

#### Retour
- Slice de chaînes de caractères si la conversion réussie.
- Erreur si un élément n'est pas une chaîne.

#### Fonctionnement
- Crée un nouveau slice de chaînes avec la même longueur que l'entrée.
- Parcourt chaque élément de l'entrée.
- Vérifie si l'élément est une chaîne, sinon retourne une erreur.
- Ajoute la chaîne au slice de sortie.

---

### Fonction `CopyFile`

Copie le contenu d'un fichier source vers un fichier de destination.

#### Paramètres
- `src` : chemin du fichier source.
- `dst` : chemin du fichier de destination.

#### Retour
- `nil` si la copie réussie.
- Erreur si une opération échoue.

#### Fonctionnement
- Ouvre le fichier source en lecture.
- Crée le fichier de destination.
- Copie le contenu du fichier source dans le fichier destination avec `io.Copy`.
- Gère et retourne toute erreur rencontrée.

---

### Fonction `EbcdicToString`

Convertit une chaîne hexadécimale EBCDIC en une chaîne ASCII.

#### Paramètre
- `ebcdicHex` : chaîne hexadécimale représentant des données EBCDIC.

#### Retour
- Chaîne ASCII correspondante, avec `?` pour les octets inconnus.

#### Fonctionnement
- Décode la chaîne hexadécimale en tableau de bytes.
- Parcourt chaque byte EBCDIC.
- Convertit chaque byte en caractère ASCII à l'aide d'une table de correspondance `ebcdicToAscii`.
- Ajoute `?` si la correspondance n'existe pas.
- Retourne la chaîne ASCII construite.

---

### Fonction `EstEntete`

Détecte si un tableau de bytes correspond à un entête spécifique et retourne un booléen ainsi qu'une valeur associée.

#### Paramètre
- `data` : tableau de bytes à analyser.

#### Retour
- `true` et une valeur entière si l'entête est reconnue.
- `false` et `0` sinon.

#### Fonctionnement
- Convertit la donnée en un entier 64 bits en utilisant `binary.BigEndian.Uint64`.
- Ajoute des masques pour vérifier différentes signatures d'entêtes.
- Retourne `true` avec la valeur associée si une correspondance est trouvée.
- Sinon, retourne `false` et `0`.

---

### Fonction `FormatDateTime`

Formate une chaîne de date et heure de 14 caractères en une représentation `"dd/mm/yyyy hh:mm:ss"`.

#### Paramètre
- `dateTime` : chaîne de 14 caractères représentant la date et l'heure.

#### Retour
- Chaîne formatée si la chaîne est valide.
- Chaîne vide si la longueur est incorrect ou si les valeurs ne sont pas numériques.

#### Fonctionnement
- Vérifie que la longueur de la chaîne est de 14 caractères.
- Extrait les composants année, mois, jour, heure, minute, seconde.
- Vérifie que chaque composant est numérique avec `IsNumeric`.
- Construit la chaîne formatée avec `fmt.Sprintf`.

---

### Fonction `FormatDateTime8`

Formate une chaîne de date et heure de 12 caractères en une représentation `"dd/mm/yyyy hh:mm:ss"`.

#### Paramètre
- `dateTime` : chaîne de 12 caractères représentant la date et l'heure.

#### Retour
- Chaîne formatée si la chaîne est valide.
- Chaîne vide si la longueur est incorrect ou si les valeurs ne sont pas numériques.

#### Fonctionnement
- Vérifie que la longueur de la chaîne est de 12 caractères.
- Extrait les composants année, mois, jour, heure, minute, seconde.
- Ajoute `"20"` devant l'année pour former l'année complète.
- Vérifie que chaque composant est numérique avec `IsNumeric`.
- Construit la chaîne formatée avec `fmt.Sprintf`.

---

### Fonction `FormatDateTime14`

Convertit une chaîne de caractères représentant une date et une heure au format `"yyyyMMddHHmmss"` en une chaîne formatée `"dd/MM/yyyy HH:MM:SS"`.

#### Paramètre
- `dateTime` : chaîne de caractères au format `"yyyyMMddHHmmss"`.

#### Retour
- Chaîne formatée `"dd/MM/yyyy HH:MM:SS"` si la conversion réussie.
- Chaîne vide si la chaîne d'entrée ne correspond pas au format attendu.

#### Fonctionnement
- Analyse la chaîne d'entrée avec `time.Parse` en utilisant le format `"20060102150405"`.
- Si l'analyse échoue, retourne une chaîne vide.
- Sinon, reformate la date et l'heure dans le format `"02/01/2006 15:04:05"` et retourne le résultat.

---

### Fonction `FormatDateTimeGMT`

Formate une chaîne de date et heure de 18 caractères en une représentation `"dd/mm/yyyy hh:mm:ss +offset"`.

#### Paramètre
- `input` : chaîne de 18 caractères représentant la date, l'heure et le décalage GMT.

#### Retour
- Chaîne formatée si la chaîne est valide.
- La chaîne d'entrée inchangée si la longueur n'est pas 18 ou si une valeur n'est pas numérique.

#### Fonctionnement
- Vérifie que la longueur de la chaîne est de 18 caractères.
- Extrait les composants année, mois, jour, heure, minute, seconde, offset.
- Ajoute `"20"` devant l'année pour former l'année complète.
- Vérifie que chaque composant est numérique avec `IsNumeric`.
- Construit la chaîne formatée avec `fmt.Sprintf`.

---

### Fonction `FormatDateTimeInv`

Inverse l'ordre des composants d'une chaîne de date et heure de 14 caractères pour produire une date au format `"dd/mm/yyyy hh:mm:ss"`.

#### Paramètre
- `dateTime` : chaîne de 14 caractères représentant la date et l'heure dans un ordre inversé.

#### Retour
- Chaîne formatée si la chaîne est valide.
- Chaîne vide si la longueur n'est pas 14 ou si une valeur n'est pas numérique.

#### Fonctionnement
- Vérifie que la longueur de la chaîne est de 14 caractères.
- Extrait les composants année, mois, jour, heure, minute, seconde en inversant l'ordre.
- Vérifie que chaque composant est numérique avec `IsNumeric`.
- Construit la chaîne formatée avec `fmt.Sprintf`.

---

### Fonction `HexTimestampToDate`

Convertit une chaîne hexadécimale représentant un timestamp Unix en une date formatée `"dd/mm/yyyy hh:mm:ss"`.

#### Paramètre
- `hexStr` : chaîne hexadécimale du timestamp.

#### Retour
- Chaîne formatée si la conversion réussie.
- Chaîne vide en cas d'erreur.

#### Fonctionnement
- Convertit la chaîne hexadécimale en entier 64 bits avec `strconv.ParseInt`.
- Convertit le timestamp en temps avec `time.Unix`.
- Formate la date avec `t.Format` selon le modèle `"02/01/2006 15:04:05"`.

---

### Fonction `HexToDecimal`

Convertit une chaîne hexadécimale en sa représentation décimale sous forme de chaîne.

#### Paramètre
- `hexStr` : chaîne hexadécimale à convertir.

#### Retour
- Chaîne décimale si la conversion réussie.
- Chaîne vide en cas d'erreur.

#### Fonctionnement
- Convertit la chaîne hexadécimale en entier 64 bits avec `strconv.ParseInt`.
- Convertit l'entier en chaîne décimale avec `strconv.FormatInt`.

---

### Fonction `HexToIntLittleEndian`

Convertit une chaîne hexadécimale en un entier en respectant l'ordre Little-endian.

#### Paramètre
- `hexString` : chaîne hexadécimale à convertir.

#### Retour
- Entier correspondant à la valeur Little-endian.
- `0` en cas d'erreur de décodage.

#### Fonctionnement
- Décode la chaîne hexadécimale en slice de bytes.
- Inverse l'ordre des octets pour respecter le Little-endian.
- Assemble les octets en un entier en décalant chaque octet.

---

### Fonction `HexToIPv4`

Convertit une chaîne hexadécimale de 8 caractères en une adresse IPv4.

#### Paramètre
- `hex` : chaîne hexadécimale représentant l'adresse IP.

#### Retour
- Adresse IP au format `"x.x.x.x"` si la chaîne est valide.
- Chaîne vide si la longueur n'est pas 8 ou si une conversion échoue.

#### Fonctionnement
- Vérifie que la longueur de la chaîne est de 8 caractères.
- Extrait chaque octet en utilisant `fmt.Sscanf`.
- Vérifie que chaque octet est dans la plage 0-255.
- Formate l'adresse IP avec `fmt.Sprintf`.

---

### Fonction `HexToString`

Convertit une chaîne hexadécimale en une chaîne ASCII.

#### Paramètre
- `hexString` : chaîne hexadécimale à convertir.

#### Retour
- Chaîne ASCII correspondant à la conversion hexadécimale.
- Chaîne vide si une erreur survient lors de la conversion.

#### Fonctionnement
- Parcourt la chaîne hexadécimale par segments de 2 caractères.
- Convertit chaque segment en entier avec `strconv.ParseInt`.
- Ajoute le caractère correspondant à la chaîne de sortie.

---

### Fonction `InverserDeuxParDeux`

Inverse chaque paire de caractères dans une chaîne.

#### Paramètre
- `s` : chaîne de caractères à traiter.

#### Retour
- Chaîne avec chaque paire de caractères inversée.

#### Fonctionnement
- Convertit la chaîne en slice de runes.
- Parcourt le slice par pas de 2, échangeant chaque paire.
- Reconstruit la chaîne à partir du slice modifié.

---

### Fonction `IsNumeric`

Vérifie si une chaîne ne contient que des caractères numériques.

#### Paramètre
- `s` : chaîne à vérifier.

#### Retour
- `true` si la chaîne ne contient que des chiffres.
- `false` sinon.

#### Fonctionnement
- Parcourt chaque caractère de la chaîne.
- Vérifie si le caractère est compris entre `'0'` et `'9'`.

---

### Fonction `Logger`

#### Variables et constantes
- `timerReference` : variable globale pour stocker la dernière heure affichée.
- Constantes ANSI pour la coloration du texte.

#### Fonction `formatTime()`
- Retourne l'heure actuelle au format `HH:MM:SS` en fuseau horaire "Europe/Paris".
- Si l'heure est identique à la dernière appelée, retourne une chaîne de 10 espaces.
- Sinon, met à jour `timerReference` et retourne l'heure entre crochets `[HH:MM:SS]`.

#### Fonctions de logging
- `Log(message string)` : affiche le message avec l'heure en turquoise sombre, sans couleur pour le message.
- `LogAlerte(message string)` : affiche le message avec l'heure en turquoise sombre et le texte en rouge.
- `LogInfo(message string)` : affiche le message avec l'heure en turquoise sombre et le texte en jaune.
- `LogSuccess(message string)` : affiche le message avec l'heure en turquoise sombre et le texte en vert gras.
- `LogRule(message string)` : affiche un message encadré par des traits horizontaux, en turquoise, avec une largeur totale de 150 caractères.

---

### Fonction `ReadBinaryFile`

Lit le contenu binaire d'un fichier et retourne un tableau d'octets.

#### Paramètre
- `filename` : nom du fichier à lire.

#### Retour
- `[]byte` : contenu du fichier.
- `error` : erreur éventuelle lors de la lecture.

#### Fonctionnement
- Ouvre le fichier en mode lecture avec `os.Open`.
- Assure la fermeture du fichier avec `defer`.
- Récupère la taille du fichier avec `file.Stat`.
- Alloue un tableau d'octets de la taille du fichier.
- Lit le contenu du fichier dans le tableau avec `binary.Read` en utilisant l'Endianness LittleEndian.

---

### Map `ebcdicToAscii`

Dictionnaire de correspondance entre les octets EBCDIC et ASCII.

#### Description
- Clé : byte en EBCDIC.
- Valeur : byte en ASCII correspondant.

#### Utilisation
- Permet la conversion de données encodées en EBCDIC vers ASCII pour traitement ou affichage.
