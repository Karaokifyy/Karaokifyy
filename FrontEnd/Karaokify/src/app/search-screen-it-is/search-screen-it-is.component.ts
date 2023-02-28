import { Component } from '@angular/core';
import { SpotifyService } from '../services/spotify.service';
import { ItunesService } from '../services/itunes.service';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import {
  HttpEvent, HttpInterceptor, HttpHandler, HttpRequest
} from '@angular/common/http';
import { Injectable } from '@angular/core';
import { NoopInterceptorService } from '../services/intercept.service';

@Component({
  selector: 'app-search-screen-it-is',
  templateUrl: './search-screen-it-is.component.html',
  styleUrls: ['./search-screen-it-is.component.css']
  //providers:[SpotifyService]
})
export class SearchScreenItIsComponent {

  searchStr:String="";
  songs:any;
  constructor(private http: HttpClient) {}
  httpData = [];

  httpResult = ($event: any) => {
    const url = `https://itunes.apple.com/search?term=${$event}&limit=11`;
    console.log(url);
    this.http.get<any>(url).subscribe(data => {
      this.httpData = data.results;
      console.log(data.results);
    });
  };
}

