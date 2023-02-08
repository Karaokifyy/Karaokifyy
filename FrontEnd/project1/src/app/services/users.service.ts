import { Injectable } from '@angular/core';
import { HttpClient, HttpClientModule } from '@angular/common/http';


@Injectable({
  providedIn: 'root'
})
export class UsersService {

  constructor(private http:HttpClient) { }

  giveUserDetails(){
    return {id:'test', password:'test1'}
    //<!-- return this.http.get-->
  }
}
