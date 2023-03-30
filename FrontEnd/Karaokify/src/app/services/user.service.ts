import { Injectable } from '@angular/core';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { BehaviorSubject } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class UsersService {

  constructor(private http:HttpClient) { }

  messageSource = new BehaviorSubject("Test");
  currentMessage = this.messageSource.asObservable();



  ChangeMessage(message:string): void{
    this.messageSource.next(message);


    //return {id:'test', password:'test1'}
    //return this.http.get("https://jsonplaceholder.typicode.com/users");
    //return this.http.post(("https://localhost:8080/users",myinputjsonvalue);


  }
  UserData(data:any){
    console.warn(data);
  }
}
