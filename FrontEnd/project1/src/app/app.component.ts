import { Component, OnInit } from '@angular/core';
import { UsersService } from './services/users.service';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'project1';
  username='';
  password='';
  
  users:any;
  router: any;
  constructor(private us:UsersService) {

   }

   ngOnInit(): void {
     //this.users=this.us.giveUserDetails()
     
   }

   UserLogIn(){
    console.log("Welcome " + (this.username));
    //this.router.navigate(['/dashboard']);
    const authUrl = 'https://accounts.spotify.com/authorize?client_id=8b82fa6e6c4b4d2fb3ce122fe2e3d6ad&redirect_uri=YOUR_REDIRECT_URI&response_type=code';
    window.location.href = authUrl;
   }
   

}
