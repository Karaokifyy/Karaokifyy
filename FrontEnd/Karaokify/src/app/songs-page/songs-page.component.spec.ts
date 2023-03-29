import { ComponentFixture, TestBed } from '@angular/core/testing';
import { SongsPageComponent } from './songs-page.component';
import { UsersService } from '../services/user.service';
import { Router } from '@angular/router';

describe('SongsPageComponent', () => {
  let component: SongsPageComponent;
  let fixture: ComponentFixture<SongsPageComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ SongsPageComponent ],
      providers: [
        { provide: UsersService, useValue: {} },
        { provide: Router, useValue: {} }
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

  it('should pass if button is clicked', () => {
    const button = fixture.debugElement.nativeElement.querySelector('button');
    button.click();
    (expect as any)(component.buttonClicked).toBeTrue();
  });


});
