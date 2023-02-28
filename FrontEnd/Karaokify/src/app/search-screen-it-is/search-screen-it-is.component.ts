import { Component } from '@angular/core';
import { SpotifyService } from '../services/spotify.service';
import { ItunesService } from '../services/itunes.service';


@Component({
  selector: 'app-search-screen-it-is',
  templateUrl: './search-screen-it-is.component.html',
  styleUrls: ['./search-screen-it-is.component.css']
  //providers:[SpotifyService]
})
export class SearchScreenItIsComponent {

  searchStr:String="";
  songs:any;

  constructor(private _s:ItunesService){

  }

  searchMusic(){
    console.log(this.searchStr);
    //console.log(this._s.search(this.searchStr));
    this._s.search(this.searchStr)
    // .subscribe((res) => {
    //   this.songs = res.json().results;
    // });



    
    // .subscribe((data)=>{
    //   this.artists=data
    // })
      //res => {
      //console.log(res);
    //})
  }

}
