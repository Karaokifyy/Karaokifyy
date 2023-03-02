import { Injectable } from '@angular/core';
import { HttpClient, HttpClientModule, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';
import { ItunesService } from './itunes.service';

@Injectable({
  providedIn: 'root'
})
export class SpotifyService {
    private searchUrl: string="";
    private key = 'Bearer BQDvjnM27yu4X2LO6OsQnoglwqYwo6DHrF-mMc68O7njltPGHbay1_0bzYg3b7Hne-3VIZXhkb-5DfNXw8MdV13VZjowam00Witmvn6xINC8LBtKFHU8hXlVbCkLE8Wd9irwMVomNAqKiKYFUSIiBVQYTgQSUSeMqBeAkcvSB4c9Pn_nu2ei';
    
    private httpOptions = {
        headers: new HttpHeaders({
            //'Accept': 'application/json',
            'Content-Type' : 'application/json',
            'Authorization' : this.key
        })

    }
  constructor(private _http:HttpClient) { }

  SearchMusic(str:String, type='artist'):Observable<any>{

    this.searchUrl= 'https:api.spotify.com/v1/search?query='+str+'&offset=0&limit=20&type='+type+'&market=US';
    //'https://api.spotify.com/v1/search?type=artist&limit=10&client_id='+this.key
    //'https://api.spotify.com/v1/search?q=' + String + '&type=' + type;
    //https://api.spotify.com/v1/search?query='+str+'&offset=0&limit=20&type='+type+'&market=US';
    return this._http.get(this.searchUrl);
    //, this.httpOptions


    //return this.http.post(("https://localhost:8080/users",myinputjsonvalue);


  }
  UserData(data:any){
    console.warn(data);
  }
}