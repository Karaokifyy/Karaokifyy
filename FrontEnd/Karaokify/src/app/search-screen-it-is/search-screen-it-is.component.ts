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
  router: any;
  constructor(private http: HttpClient, private route: ActivatedRoute){}

  // ngOnInit(): void {}
    
  httpData : any[] = [];
  ngOnInit(){
    const url = `http://localhost:8080/newSpotifySession`;
    
    this.route.queryParamMap
    .subscribe(params => {
      this.http.post(url, JSON.stringify({o_code:params.get("code")})).subscribe(data =>{
        const serverResult = JSON.parse(JSON.stringify(data))
        console.log(serverResult)
        const mapResult = Object.entries(serverResult)
        console.log(mapResult)
        for (let i = 0; i < mapResult.length; i++) {
          this.httpData.push(mapResult[i][1])
          console.log(mapResult[i][1])
        } 
        
      })
      console.log(JSON.stringify({o_code:params.get("code")}))
      console.log(params.get("code"))})

  }

  selectPlaylist(playlistID:String){

    const url = `http://localhost:8080/newSpotifySession`;
    console.log(playlistID);


    // this.route.queryParamMap
    // .subscribe(params => {
    //   this.http.post(url, JSON.stringify({playlistID})).subscribe(data =>{
    //     const serverResult = JSON.parse(JSON.stringify(data))
    //     console.log(serverResult)
    //     const mapResult = Object.entries(serverResult)
    //     console.log(mapResult)
    //     for (let i = 0; i < mapResult.length; i++) {
    //       this.httpData.push(mapResult[i][1])
    //       console.log(mapResult[i][1])
    //     } 
  }

  httpResult = ($event: any) => {
    const url = `https://itunes.apple.com/search?term=${$event}&limit=500`;
    console.log(url);
    this.http.get<any>(url).subscribe(data => {
      this.httpData = data.results;
      console.log(data.results);
    });
  };

  UserLyricsPage(){
    console.log("Welcome to lyrics page" );
    //this.router.navigateByUrl('/lyrics-page');
    window.location.href = 'http://localhost:4200/lyrics-page';
  }
}

