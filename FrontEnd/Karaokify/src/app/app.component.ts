import { Component, OnInit } from '@angular/core';
import { UsersService } from './services/user.service';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})

//implements OnInit
export class AppComponent  {

  constructor(private title: Title) {}
  public ngOnInit(): void {
    this.title.setTitle('Karaokifyy');
}

public getTitle(): string {
    this.title.setTitle('Karaokifyy');
    return this.title.getTitle();
}



}