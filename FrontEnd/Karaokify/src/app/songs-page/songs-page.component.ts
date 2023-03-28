import { Component, OnInit } from '@angular/core';
import { UsersService } from '../services/user.service';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';
import {
  HttpEvent, HttpInterceptor, HttpHandler, HttpRequest
} from '@angular/common/http';
import { Injectable } from '@angular/core';
import { NoopInterceptorService } from '../services/intercept.service';
import { ActivatedRoute, Router } from '@angular/router';
import { param } from 'cypress/types/jquery';

@Component({
  selector: 'app-songs-page',
  templateUrl: './songs-page.component.html',
  styleUrls: ['./songs-page.component.css']
})
export class SongsPageComponent implements OnInit {

  str1:string="test";
  str2:string="test";
  httpData : any[] = [];
  constructor(private http: HttpClient, private route: ActivatedRoute, private us:UsersService, private router: Router){}

  ngOnInit(){
    const url = `http://localhost:8080/newSpotifySession`;

    this.us.currentMessage.subscribe(
      (message) => (this.str1 = message)
    );

    this.route.queryParamMap
    .subscribe(params => {
      this.http.post(url, JSON.stringify({playlistID:this.str1})).subscribe(data =>{
        const serverResult = JSON.parse(JSON.stringify(data))
        console.log(serverResult)
        const mapResult = Object.entries(serverResult)
        console.log(mapResult)
        for (let i = 0; i < mapResult.length; i++) {
          this.httpData.push(mapResult[i][1])
          console.log(mapResult[i][1])
        } 
        
      })
      console.log(JSON.stringify({playlistID:this.str1}))
      //console.log(JSON.stringify({o_code:params.get("code")}))
      //console.log(params.get("code"))
    })

  }


  selectSongs(songID:string){
    this.us.ChangeMessage(songID);
    this.router.navigateByUrl('/lyrics');
    console.log(songID);

  }
}

