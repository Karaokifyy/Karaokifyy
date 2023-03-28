import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { UsersService } from '../services/user.service';
import { DomSanitizer } from '@angular/platform-browser';

@Component({
  selector: 'app-lyrics-page',
  templateUrl: './lyrics-page.component.html',
  styleUrls: ['./lyrics-page.component.css']
})
export class LyricsPageComponent implements OnInit {

  str1:string="";
  httpData : any[] = [];
  safeUrl:any;
  constructor(private http: HttpClient, private us:UsersService, private _san:DomSanitizer){}

  ngOnInit(): void {
    this.us.currentMessage.subscribe(
      (message) => (this.str1 = message)
    );

    this.safeUrl=this._san.bypassSecurityTrustResourceUrl("https://www.youtube.com/embed/Kq8zlXS2bUg");



  }



}
