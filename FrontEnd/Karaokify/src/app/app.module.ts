import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import {MatButtonModule} from '@angular/material/button';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {HttpClientModule} from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { RouterModule, Routes } from '@angular/router';
import { UsersService } from './services/user.service';
import { AppRoutingModule, routingComponents } from './app-routing.module';
import { SearchScreenItIsComponent } from './search-screen-it-is/search-screen-it-is.component';
import { HomeScreenComponent } from './home-screen/home-screen.component';
import { ItunesService } from './services/itunes.service';
import { HTTP_INTERCEPTORS } from '@angular/common/http';
import {NoopInterceptorService} from './services/intercept.service';
export const HTTPINTProviders = [
  { provide: HTTP_INTERCEPTORS, useClass: NoopInterceptorService, multi: true },
];

const appRoutes: Routes = [
  { path: '', redirectTo: '/login', pathMatch: 'full' },
];

@NgModule({
  declarations: [
    AppComponent,
    routingComponents,
    SearchScreenItIsComponent,
    HomeScreenComponent
  ],
  imports: [
    AppRoutingModule,
    MatButtonModule,
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    FormsModule,
    RouterModule.forRoot([])
  ],
  providers: [UsersService, ItunesService, HTTPINTProviders],
  bootstrap: [AppComponent, SearchScreenItIsComponent, HomeScreenComponent]
})
export class AppModule { }
