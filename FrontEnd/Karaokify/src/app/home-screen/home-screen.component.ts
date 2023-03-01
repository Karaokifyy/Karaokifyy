import { Component, OnInit } from '@angular/core';
import { UsersService } from '../services/user.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-home-screen',
  templateUrl: './home-screen.component.html',
  styleUrls: ['./home-screen.component.css']
})
export class HomeScreenComponent implements OnInit{
  username='';
  password='';
  constructor(private us:UsersService, private router: Router) {

  }

  ngOnInit(): void {
    //this.users=this.us.giveUserDetails()
    
  }

  UserLogIn(){
   console.log("Welcome " + (this.username));
   //this.router.navigateByUrl('/screen-search');
   //http://localhost:4200/screen-search
   const scopes = '&scope=playlist-read-private'
   const authUrl = 'https://accounts.spotify.com/authorize?client_id=8b82fa6e6c4b4d2fb3ce122fe2e3d6ad&redirect_uri=http://localhost:4200/screen-search&response_type=code'+scopes;
   window.location.href = authUrl;
  }



}
