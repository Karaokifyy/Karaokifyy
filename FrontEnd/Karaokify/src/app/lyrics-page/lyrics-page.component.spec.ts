import { ComponentFixture, TestBed } from '@angular/core/testing';
import { LyricsPageComponent } from './lyrics-page.component';
import { UsersService } from '../services/user.service';
import { Router } from '@angular/router';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { By } from '@angular/platform-browser';


describe('LyricsPageComponent', () => {
 let component: LyricsPageComponent;
 let fixture: ComponentFixture<LyricsPageComponent>;


 beforeEach(() => {
   TestBed.configureTestingModule({
     declarations: [ LyricsPageComponent ],
     imports: [HttpClientModule],
     providers: [
       { provide: UsersService, useValue: {} },
       { provide: Router, useValue: {} }
     ]
   });
   fixture = TestBed.createComponent(LyricsPageComponent);
   component = fixture.componentInstance;
 });


 beforeEach(() => {
   fixture = TestBed.createComponent(LyricsPageComponent);
   component = fixture.componentInstance;
   fixture.detectChanges();
 });


 it('should always be true', () => {
   (expect as any)(true).toBeTrue();
 });
 
 it('should create the component', () => {
   (expect as any)(component).toBeTruthy();
 });


 it('should contain a YouTube link', () => {
   const link = fixture.debugElement.query(By.css('a')).nativeElement;
   (expect as any)(link.href).toContain('youtube.com');
 });
});




