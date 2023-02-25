import { Component, OnInit } from '@angular/core';
import { UsersService } from './services/user.service';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

//implements OnInit
export class AppComponent  {
  title = 'project1';
  // username='';
  // password='';
  
  // users:any;
  // constructor(private us:UsersService) {

  //  }

  //  ngOnInit(): void {
  //    //this.users=this.us.giveUserDetails()
     
  //  }

  //  UserLogIn(){
  //   console.log("Welcome " + (this.username));
  //  }


}