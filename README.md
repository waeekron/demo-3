# Demo 3

### piiput-ja-filtterit
- filtterit = funktioita joita putkitetaan eli yhdistetään
- `pipe`-funktio ottaa putkessa kulkevan datan ja filtter-funktiot, joita se kutsuu silmukassa muokaten putkessa kulkevaa dataa

### piiput-ja-filtterit-V2

Sama idea kuin yllä, mutta nyt `pipeV2`-funktio osaa sovittaa filterit yhteen reflektion avulla. Reflektio osuudessa on
käytetty tukiälyä mukavan paljon:) 

- `pipeV2`-funktio tarkistaa edellisen funktion paluuarvojen ja seuraavan funktion parametrilistan yhteensopivuuden
  - panikoi, jos funktiot eivät yhdisty oikein

Toteutuksen järkevyyteen en osaa ottaa kantaa, mutta näinkin voi näköjään tehdä:))


### implisiittinen-kutsu
- kaksi goroutinea, kaksi kanavaa
- ensimmäin goroutine lukee syötettä, kirjoittaa sen kanavalle, jonka jälkeen se sulkee kanavan
- toinen goroutine kuuntelee kanavaa, laskee laskun, jonka jälkeen se sulkee kanavan
  
