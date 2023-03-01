import { Injectable } from '@angular/core';
import { HttpClient, HttpClientJsonpModule, HttpClientModule, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';


@Injectable()
export class ItunesService {

    apiRoot: string = "https://itunes.apple.com/search";
    results: [];
  constructor(private jsonp: HttpClient) {
    this.results = [];
   }

  search(term:String){

    this.results = [];
      let apiURL = `${this.apiRoot}?term=${term}&media=music&limit=10`;
      
      console.log(apiURL);
      this.jsonp.get<any>(apiURL).subscribe(data => {
        this.results=data.results;
        console.log(data.results)
      })
  }
}