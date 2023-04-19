import { Component, OnInit, OnDestroy } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { UsersService } from '../services/user.service';
import { DomSanitizer } from '@angular/platform-browser';
import { ActivatedRoute, Router } from '@angular/router';
import { interval } from 'rxjs';

@Component({
  selector: 'app-lyrics-page',
  templateUrl: './lyrics-page.component.html',
  styleUrls: ['./lyrics-page.component.css']
})
export class LyricsPageComponent implements OnInit {

  displayFlag = false;
  counter: number = 0;
  str1:string="Kq8zlXS2bUg";
  strnew:string="";
  strlyric:string="";
  strnewlyrics:string="";
  checker:string="1";
  num1:string="0";
  str2:string="https://www.youtube.com/embed/";
  url:string= `http://localhost:8080/karaokifyy/`;
  httpData : any[] = [];
  safeUrl:any=this._san.bypassSecurityTrustResourceUrl("https://www.youtube.com/embed");
  timeArray: string[] = [];
  textArray: string[] = [];
  displayText: string="";
  timeInterval: number = 5000;
  interval: any;
  constructor(private http: HttpClient, private us:UsersService, private _san:DomSanitizer, private route: ActivatedRoute){}

  ngOnInit(): void {
    var num = this.us.currentMessage.subscribe(
      (message) => (this.str1 = message)
    );
    console.log(this.str1)
    this.url=this.url+this.str1;
    console.log(this.url)
  
   

    this.timeArray.push("000")


 


    this.route.queryParamMap
    .subscribe(params => {
      this.http.get(this.url).subscribe(data =>{
        const serverResult = JSON.parse(JSON.stringify(data))
        const mapResult = Object.entries(serverResult)
        //console.log(mapResult)
        for (let i = 0; i < mapResult.length; i++) {
          this.httpData.push(mapResult[i][1])
          // console.log(i)
          // console.log(mapResult[i][1])
          if(i==0){
            const myArr1:any=mapResult[i][1]
            //console.log(myArr1)
            this.strnewlyrics=myArr1
            console.log(this.strnewlyrics)
            this.str2=this.strnewlyrics
          }
          if(i!=0){
            const myArr:any=mapResult[i][1]
            this.strnewlyrics=myArr
            this.strnewlyrics=this.strnewlyrics.concat(";")
            for(let k = 0; k<500; k++){
            for(let j = 0; j<10; j++){
              if(j==2||j==4||j==5){
                this.strnew=this.strnew.concat(this.strnewlyrics[j])
              }
            }
            for(let j =11; j<1000; j++){
              if(this.strnewlyrics[j].localeCompare("[")==0){
                this.strnewlyrics=this.strnewlyrics.substring(j)
                break;
              }
              else if(this.strnewlyrics[j].localeCompare(";")==0){
                this.checker="0"
                break;
              }
              else{
                this.strlyric=this.strlyric.concat(this.strnewlyrics[j]);
              }
              

            }

           
            if(k!=0)
              this.timeArray.push(this.strnew)
            //console.log(this.timeArray[this.timeArray.length - 1])
            //console.log(this.timeArray.length)
            this.textArray.push(this.strlyric)
            this.strnew="";
            this.strlyric="";
            if(this.strnewlyrics.length==0){
              break;
            }
            if(this.checker.localeCompare("0")==0){
              break;

          }
        }
        }

          
      }
      //console.log(this.timeArray.length)
      this.interval = setInterval(() => {
        if(this.counter==0){
          this.str2=this.str2+"?autoplay=1";
          this.safeUrl=this._san.bypassSecurityTrustResourceUrl(this.str2);
        }
        setTimeout(() => {
          this.displaylyrics();;
        }, 2000);
         //console.log(this.timeInterval)
      }, this.timeInterval);




        
      })
      
    })
   
    



  }
  displaylyrics(){
    if(this.counter+1==this.timeArray.length){
      clearInterval(this.interval);
      return;
    }
    let nummy2 = 0;
    let nummy1 = 0;
    let nummy = "";
    let nummy3= "";
    nummy = (this.timeArray[this.counter+1])
    nummy3 = (this.timeArray[this.counter])
    nummy1 = (parseInt(nummy3[0])*60)+parseInt(nummy3.substring(1))
    nummy2 = (parseInt(nummy[0])*60)+parseInt(nummy.substring(1))
    nummy2 = (nummy2-nummy1)*1000
    this.displayText = this.textArray[this.counter]
    this.counter++;
    this.timeInterval=(nummy2-50)

    console.log(this.timeInterval)
    clearInterval(this.interval);
    this.interval = setInterval(() => {
      this.displaylyrics();
    }, this.timeInterval);
  }
  testtest(){
    console.log("clicked")
  }



}