import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HomeScreenComponent } from './home-screen.component';
import { UsersService } from '../services/user.service';
import { Router } from '@angular/router';
//import  expect from 'jasmine-expect';

describe('HomeScreenComponent', () => {
  let component: HomeScreenComponent;
  let fixture: ComponentFixture<HomeScreenComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [ HomeScreenComponent ],
      providers: [
        { provide: UsersService, useValue: {} },
        { provide: Router, useValue: {} }
      ]
    });

    fixture = TestBed.createComponent(HomeScreenComponent);
    component = fixture.componentInstance;
  });

  it('should create the component', () => {
    (expect as any)(component).toBeTruthy();
  });

  it('should have empty username and password fields on initialization', () => {
    (expect as any)(component.username).toBe('');
    (expect as any)(component.password).toBe('');
  });
});
