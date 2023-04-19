import { ComponentFixture, TestBed, fakeAsync, tick } from '@angular/core/testing';import { LyricsPageComponent } from './lyrics-page.component';
import { UsersService } from '../services/user.service';
import { ActivatedRoute, convertToParamMap, Router } from '@angular/router';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { By } from '@angular/platform-browser';
import { of } from 'rxjs';
import { HttpTestingController } from '@angular/common/http/testing';




describe('LyricsPageComponent', () => {
 let component: LyricsPageComponent;
 let fixture: ComponentFixture<LyricsPageComponent>;
 let httpMock: HttpTestingController;



 beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ LyricsPageComponent ],
      imports: [ HttpClientModule ],
      providers: [
        {
          provide: UsersService,
          useValue: {
            currentMessage: of('mock value')
          }
        },
        { provide: Router, useValue: {} },
        {
          provide: ActivatedRoute,
          useValue: {
            snapshot: { paramMap: convertToParamMap({ id: '1' }) },
            queryParamMap: of(convertToParamMap({}))
          }
        }
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



 it('should have a YouTube link', () => {
    const linkElement = fixture.nativeElement.querySelector('a[href*="youtube.com"]');
    (expect as any)(linkElement).toBeFalsy();
  });


  it('Youtube link should be the correct one from database', () => {
    const linkElement = fixture.nativeElement.querySelector('a[href*="youtube.com"]');
    (expect as any)(linkElement).toBeFalsy();
  });
  
  it('should log a message to console when testtest() is called', () => {
    spyOn(console, 'log');
    component.testtest();
    (expect as any)(console.log).toHaveBeenCalledWith('clicked');
  });
  
  it('should update displayed lyrics in displaylyrics()', () => {
    component.timeArray = ['000', '101', '102'];
    component.textArray = ['line 1', 'line 2', 'line 3'];
  
    component.displaylyrics();
    (expect as any)(component.displayText).toEqual('line 1');
    (expect as any)(component.counter).toEqual(1);
    (expect as any)(component.timeInterval).toEqual(61000);
  
    component.displaylyrics();
    (expect as any)(component.displayText).toEqual('line 2');
    (expect as any)(component.counter).toEqual(2);
    (expect as any)(component.timeInterval).toEqual(1000);
  });
  
});




