import { Component, OnInit } from '@angular/core';
//import { UsersService } from './services/users.service';
import { UsersService } from '../services/user.service';

@Component({
  selector: 'app-home-screen',
  templateUrl: './home-screen.component.html',
  styleUrls: ['./home-screen.component.css']
})
export class HomeScreenComponent implements OnInit{
  username='';
  password='';
  constructor(private us:UsersService) {

  }

  ngOnInit(): void {
    //this.users=this.us.giveUserDetails()
    
  }

  UserLogIn(){
   console.log("Welcome " + (this.username));
  }



}
