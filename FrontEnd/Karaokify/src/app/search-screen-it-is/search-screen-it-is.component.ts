import { Component, OnInit } from '@angular/core';
import { SpotifyService } from '../services/spotify.service';
import { ItunesService } from '../services/itunes.service';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import {
  HttpEvent, HttpInterceptor, HttpHandler, HttpRequest
} from '@angular/common/http';
import { Injectable } from '@angular/core';
import { NoopInterceptorService } from '../services/intercept.service';
import { ActivatedRoute } from '@angular/router';
import { param } from 'cypress/types/jquery';

@Component({
  selector: 'app-search-screen-it-is',
  templateUrl: './search-screen-it-is.component.html',
  styleUrls: ['./search-screen-it-is.component.css']
  //providers:[SpotifyService]
})
export class SearchScreenItIsComponent implements OnInit{

  searchStr:String="";
  songs:any;
  constructor(private http: HttpClient, private route: ActivatedRoute){}

  // ngOnInit(): void {}
    
  
  ngOnInit(){
    const url = `http://localhost:8080/newSpotifySession`;
    this.route.queryParamMap
    .subscribe(params => {
      this.http.post(url, JSON.stringify({o_code:params.get("code")})).subscribe()
      console.log(JSON.stringify({o_code:params.get("code")}))
      console.log(params.get("code"))})
      
  }

  httpData = [];

  httpResult = ($event: any) => {
    const url = `https://itunes.apple.com/search?term=${$event}&limit=500`;
    console.log(url);
    this.http.get<any>(url).subscribe(data => {
      this.httpData = data.results;
      console.log(data.results);
    });
  };
}

