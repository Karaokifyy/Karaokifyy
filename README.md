## Project Name
Karaokifyy

## Project Description
The basis for the website/platform centers around providing users karaoke versions of their favorite songs or Spotify playlists, especially for ones that don't already have an official or fan-made version readily available. Our main feature will link to the user's Spotify and pull the song(s)'s metadata using Spotify's Web API, find the matching instrumental through YouTube's Data API, and lastly play that instrumental via an IFrame while displaying the time-synchronized lyrics pulled from our database. The secondary feature focuses on crowdsourcing the time-synchronized lyrics through an easy-to-use interface involving the volunteer highlighting parts of the lyrics that are currently being played from the original version.

## Instalation
Run FrontEnd Locally
1. Install [Node.js](https://nodejs.org/en/download)
2. Install angular CLI by using "npm install -g @angular/cli" in cmd prompt
3. Download/Clone our project from our [GitHub](https://github.com/Karaokifyy/Karaokifyy)
4. change directory in cmd prompt to ...\Karaokifyy\FrontEnd\Karaokify
5. Run "npm install" in cmd prompt
6. Run "ng serve" in cmd prompt
7. The application will be running on [localhost:4200](http://localhost:4200)

Set up Back-end Locally
1. Run app.exe while in "Karaokiffy/app" folder
OR
1. Ensure CGO is [enabled](https://pkg.go.dev/cmd/cgo) and you have a [working x64 c compiler (e.g.](https://www.msys2.org/)) accesible and in "PATH"
2. open the "app" directory in a terminal and perform either
   1. "go build" and run the generated executable file OR
   2. "go run ."

## Members
Heather Burke - Back End

Marc Guillaume - Back End

Akhil Kancherla - Front End

Skanda Kumaran - Front End
