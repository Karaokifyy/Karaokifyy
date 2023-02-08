import { Component, OnInit } from '@angular/core';
import { UsersService } from './services/users.service';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'project1';
  users:any;
  constructor(private us:UsersService) {

   }

   ngOnInit(): void {
     this.users=this.us.giveUserDetails()
     //
   }


}
