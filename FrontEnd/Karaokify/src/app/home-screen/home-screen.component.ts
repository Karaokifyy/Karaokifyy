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
   this.router.navigateByUrl('/screen-search');
   //const authUrl = 'https://accounts.spotify.com/authorize?client_id=YOUR_CLIENT_ID&redirect_uri=&response_type=code';
   //window.location.href = authUrl;
  }



}
