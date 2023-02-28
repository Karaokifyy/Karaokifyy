import { Injectable } from '@angular/core';
import { HttpClient, HttpClientModule, HttpHeaders } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable()
export class ItunesService {

  constructor(private _http: HttpClient) { }

  search(str:String) {

    return this._http.get(`https://itunes.apple.com/search?term=${String}&limit=${"10"}`);
    //this._http.request('GET', )
    //httpClient.request('GET', this.heroesUrl + '?' + 'name=term', {responseType:'json'});

  }
}