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
  str2:string="https://www.youtube.com/embed/";
  httpData : any[] = [];
  safeUrl:any;
  constructor(private http: HttpClient, private us:UsersService, private _san:DomSanitizer){}

  ngOnInit(): void {
    this.us.currentMessage.subscribe(
      (message) => (this.str1 = message)
    );

    //this.str2=this.str2+this.str1;
    this.str2=this.str2+"Kq8zlXS2bUg";

    this.safeUrl=this._san.bypassSecurityTrustResourceUrl(this.str2);
    


  }



}
