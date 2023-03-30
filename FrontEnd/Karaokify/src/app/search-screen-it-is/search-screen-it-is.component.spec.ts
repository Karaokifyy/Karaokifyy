import { ComponentFixture, TestBed } from '@angular/core/testing';
import { SearchScreenItIsComponent } from './search-screen-it-is.component';
import { UsersService } from '../services/user.service';
import { convertToParamMap, Router } from '@angular/router';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ActivatedRoute } from '@angular/router';
import { of } from 'rxjs';


describe('SearchScreenItIsComponent', () => {
 let component: SearchScreenItIsComponent;
 let fixture: ComponentFixture<SearchScreenItIsComponent>;


 beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ SearchScreenItIsComponent ],
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
    fixture = TestBed.createComponent(SearchScreenItIsComponent);
    component = fixture.componentInstance;
  });


 beforeEach(() => {
   fixture = TestBed.createComponent(SearchScreenItIsComponent);
   component = fixture.componentInstance;
   fixture.detectChanges();
 });


 it('should always be true', () => {
   (expect as any)(true).toBeTrue();
 });
 
 it('should create the component', () => {
   (expect as any)(component).toBeTruthy();
 });


 it('should have a button', () => {
    const buttonElement = fixture.nativeElement.querySelector('button');
    (expect as any)(buttonElement).toBeTruthy();
  });




});


