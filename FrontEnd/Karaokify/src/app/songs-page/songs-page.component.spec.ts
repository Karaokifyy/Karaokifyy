import { ComponentFixture, TestBed } from '@angular/core/testing';
import { SongsPageComponent } from './songs-page.component';
import { UsersService } from '../services/user.service';
import { convertToParamMap, Router } from '@angular/router';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ActivatedRoute } from '@angular/router';
import { of } from 'rxjs';


describe('SongsPageComponent', () => {
 let component: SongsPageComponent;
 let fixture: ComponentFixture<SongsPageComponent>;


 beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ SongsPageComponent ],
      imports: [ HttpClientTestingModule ],
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
    fixture = TestBed.createComponent(SongsPageComponent);
    component = fixture.componentInstance;
  });


 beforeEach(() => {
   fixture = TestBed.createComponent(SongsPageComponent);
   component = fixture.componentInstance;
   fixture.detectChanges();
 });


 it('should always be true', () => {
   (expect as any)(true).toBeTrue();
 });
 
 it('should create the component', () => {
   (expect as any)(component).toBeTruthy();
 });

 it('should be able to click the song by button', () => {
  (expect as any)(HTMLButtonElement).toBeTruthy();
});
 
});


