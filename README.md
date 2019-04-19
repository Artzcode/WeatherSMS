# WeatherSMS

WeatherSMS est un petit programme en **GoLang**. Il permet d'envoyer un SMS avec la température minimale et maximale de la journée. 
**Ce programme fonctionne :** 
- avec **l'API SMS de Free Mobile**
- avec l'API Météo **https://www.apixu.com/**
- avec https://github.com/tidwall/gjson pour parser les fichiers .json

## Configuration :

Renommer `config.json.dist` en `config.json` et remplacer les valeurs de `user` et `pass` par vos identifiants de l'API Free Mobile. 
Remplacer la valeur `apiKey` par votre clé Api de https://www.apixu.com/.

```json
{
    "user": "identifiant Free Mobile",
    "pass": "Mot de passe de l'API Free Mobile",
    "apiKey": "Votre api key Apixu"
}
```

## Utilisation :

Lancer simplement le programme avec `go run main.go`. Aucun argument n'est nécessaire.